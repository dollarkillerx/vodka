/**
 * @Author: DollarKiller
 * @Description: etcd 注册中心  (其他注册中心可以参考它)
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 12:51 2019-09-24
 */
package etcd

import (
	"context"
	"fmt"
	"github.com/dollarkillerx/vodka/registry"
	"github.com/dollarkillerx/vodka/utils"
	"github.com/dollarkillerx/vodka/utils/clog"
	"go.etcd.io/etcd/clientv3"
	"log"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxServiceNum       = 8
	//SyncServerCacheTime = 10 * time.Minute // 为防止意外的定时缓存更新  分钟
	SyncServerCacheTime = 10 * time.Second // 为防止意外的定时缓存更新  分钟
)

type EtcdRegistry struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service

	lock               sync.Mutex   // nutex 防止并发超量
	value              atomic.Value // 原子缓存  高效
	registryServiceMap sync.Map
}

var (
	// 单利 饿汉式
	etcdRegistry = &EtcdRegistry{
		serviceCh: make(chan *registry.Service, MaxServiceNum),
	}
)

type RegisterService struct {
	id            clientv3.LeaseID  // 租约id
	service       *registry.Service // 服务体
	registered    bool              // 是否注册
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

// 定义缓存结构体
type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

// 外部初始化
func EtcdInit() {}

func init() {
	serverCache := &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}
	// 初始化缓存
	etcdRegistry.value.Store(serverCache)

	registry.RegisterPlugin(etcdRegistry)
	go etcdRegistry.run()
}

// 实现interface

func (e *EtcdRegistry) Name() string {
	return "etcd"
}

// 初始化配置
func (e *EtcdRegistry) Init(ctx context.Context, opts ...registry.Option) (err error) {
	// 基础设置
	options := registry.Options{}

	for _, k := range opts {
		k(&options)
	}

	e.options = &options

	config := clientv3.Config{
		Endpoints:   options.Address,
		DialTimeout: options.Timeout,
	}

	e.client, err = clientv3.New(config)
	if err != nil {
		return err
	}

	return
}

// 服务注册
func (e *EtcdRegistry) Register(ctx context.Context, service *registry.Service) (err error) {
	select {
	case e.serviceCh <- service:
	default:
		err = fmt.Errorf("register chan is full")
		return
	}
	return
}

// 服务反注册
func (e *EtcdRegistry) Unregister(ctx context.Context, service *registry.Service) (err error) {
	e.registryServiceMap.Delete(service.Name)
	return
}

// 缓存中查询
func (e *EtcdRegistry) getServiceByCache(ctx context.Context, name string) (service *registry.Service, err error) {
	name = e.servicePath(name)
	serverInfo := e.value.Load().(*AllServiceInfo)
	i, ok := serverInfo.serviceMap[name]
	if ok {
		return i, nil
	}
	return nil, fmt.Errorf("cache not data")
}

// etcd中查询
func (e *EtcdRegistry) getServiceByEtcd(ctx context.Context, name string, tag int) (service *registry.Service, err error) {
	if e.options.Debug && tag == 1 {
		log.Println("cache 被穿透 现在 进入 etcd 获取 数据 缓存名称: " + name)
	}
	pathName := e.servicePath(name)
	response, err := e.client.Get(context.TODO(), name, clientv3.WithPrefix())
	if err != nil {
		// 如果不存在 则返回err
		return nil, fmt.Errorf("etcd not data")
	}
	service = &registry.Service{
		Name: name,
	}
	for _, kv := range response.Kvs {
		val := kv.Value
		var item registry.Node
		err := utils.Jsonp.Unmarshal(val, &item)
		if err != nil {
			clog.PrintWa(err)
			return nil, fmt.Errorf("unmarshal error")
		}
		service.Nodes = append(service.Nodes, &item)
	}

	// 已经去除 开始写入缓存

	serverInfo := e.value.Load().(*AllServiceInfo)
	serverInfo.serviceMap[pathName] = service

	e.value.Store(serverInfo)

	return service, nil
}

// 服务发现
func (e *EtcdRegistry) GetService(ctx context.Context, name string) (service *registry.Service, err error) {
	// 服务发现 先想缓存中获取  如果缓存被穿透 就想etcd中获取
	service, err = e.getServiceByCache(ctx, name)
	if err == nil {
		return
	}
	// 如果不存在 就查询  一个查询进行
	e.lock.Lock()
	defer e.lock.Unlock()
	service, err = e.getServiceByCache(ctx, name)
	if err == nil {
		return
	} else {
		// 如果不存在 则 向etcd 中查询
		service, err = e.getServiceByEtcd(ctx, name, 1)
		return
	}
}

func (e *EtcdRegistry) run() {
	ticker := time.NewTicker(SyncServerCacheTime)
	for {
		select {
		case item := <-e.serviceCh:
			// 先查询是否存在  如果存在则续租，反之进行注册
			ser, ok := e.registryServiceMap.Load(item.Name)
			if ok {
				// 存在 添加节点信息
				server := ser.(*RegisterService)
				for _, i := range item.Nodes {
					server.service.Nodes = append(server.service.Nodes, i)
				}
				server.registered = false

				e.registryServiceMap.Store(item.Name, server)
				continue
			} else {
				// 不存在 进行注册
				registryed := &RegisterService{
					service: item,
				}
				e.registryServiceMap.Store(item.Name, registryed)
			}
		case <- ticker.C:
			e.syncUpdateCache()
		default:
			// 续约
			e.registerOrKeepAlive()
			time.Sleep(time.Millisecond * 200)
		}
	}
}

// 定时更新缓存   这个是保护措施
func (e *EtcdRegistry) syncUpdateCache() {
	serverInfo := e.value.Load().(*AllServiceInfo)
	for k, _ := range serverInfo.serviceMap {
		_, err := e.getServiceByEtcd(context.TODO(), k, 2)
		if err != nil {
			clog.PrintWa(err)
		}
	}
}

// 注册 或者 续租
func (e *EtcdRegistry) registerOrKeepAlive() {
	e.registryServiceMap.Range(func(key, value interface{}) bool {
		item := value.(*RegisterService)

		// 如果注册就续约
		if item.registered {
			e.keepAlive(item)
			return true
		}
		// 反之就进行注册
		err := e.registerServer(item)
		if err != nil {
			panic(err)
		}

		return true // 一直往下找
	})
}

// 续约
func (e *EtcdRegistry) keepAlive(server *RegisterService) {
	select {
	case resp := <-server.keepAliveChan:
		if resp == nil {
			// 服务出现问题
			server.registered = false
		}
		if e.options.Debug {
			log.Printf("续租 server: %v, node %v, port: %v", server.service.Name, server.service.Nodes[0].Ip, server.service.Nodes[0].Port)
		}
	}
}

// 注册
func (e *EtcdRegistry) registerServer(server *RegisterService) error {
	// 获取租约id
	response, err := e.client.Grant(context.TODO(), e.options.HeartBeat)
	if err != nil {
		return err
	}

	id := response.ID
	server.id = id

	for _, node := range server.service.Nodes {
		// 为每个节点够着一个server 注入  /dkg/server1/127.0.0.1:8081  /dkg/server1/127.0.0.2:8081
		ser := &registry.Service{
			Name: server.service.Name,
			Nodes: []*registry.Node{
				node,
			},
		}
		path := e.serviceNodePath(ser)
		bytes, err := utils.Jsonp.Marshal(ser)
		if err != nil {
			continue
		}

		_, err = e.client.Put(context.TODO(), path, string(bytes), clientv3.WithLease(id))
		if err != nil {
			continue
		}
	}

	responses, err := e.client.KeepAlive(context.TODO(), id)
	if err != nil {
		return err
	}

	server.keepAliveChan = responses
	server.registered = true
	if e.options.Debug {
		log.Printf("注册完毕 server: %v  host: %v  port: %v", server.service.Name, server.service.Nodes[0].Ip, server.service.Nodes[0].Port)
	}

	// 注册完毕后写入缓存
	name := e.servicePath(server.service.Name)
	serverInfo := e.value.Load().(*AllServiceInfo)
	serverInfo.serviceMap[name] = server.service
	// 缓存写入完毕后存入
	e.value.Store(serverInfo)

	return nil
}

// 获取etcd 配置路径
func (e *EtcdRegistry) serviceNodePath(service *registry.Service) string {
	nodeIP := fmt.Sprintf("%s:%d", service.Nodes[0].Ip, service.Nodes[0].Port)
	return path.Join(e.options.RegistryPath, "/", service.Name, "/", nodeIP)
}

// 获取 service 路径
func (e *EtcdRegistry) servicePath(name string) string {
	return path.Join(e.options.RegistryPath, "/", name, "/")
}
