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
	MaxServiceNum = 8
	//SyncServerCacheTime = 10 * time.Minute // 为防止意外的定时缓存更新  分钟
	SyncServerCacheTime = 20 * time.Second // 为防止意外的定时缓存更新  分钟
)

type EtcdRegistry struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Node

	lock               sync.Mutex   // nutex 防止并发超量
	value              atomic.Value // 原子缓存  高效
	registryServiceMap sync.Map
}

var (
	// 单利 饿汉式
	etcdRegistry = &EtcdRegistry{
		serviceCh: make(chan *registry.Node, MaxServiceNum),
	}
)

type RegisterService struct {
	name          string           // 服务名称
	path          string           // 服务path
	id            clientv3.LeaseID // 租约id
	node          *registry.Node   // 服务体
	registered    bool             // 是否注册
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

	// 注册etcd 到管理中心
	err := registry.RegisterPlugin(etcdRegistry)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Etcd 注册中心初始化完毕")
	}
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
func (e *EtcdRegistry) Register(ctx context.Context, service *registry.Node) (err error) {
	select {
	case e.serviceCh <- service:
	default:
		err = fmt.Errorf("register chan is full")
		return
	}
	return
}

// 服务反注册
func (e *EtcdRegistry) Unregister(ctx context.Context, service *registry.Node) (err error) {
	name, path := e.getServerName(service)
	e.registryServiceMap.Delete(name)

	// v0.0.1 废弃
	//defer e.registryServiceMap.Delete(name)
	//value, ok := e.registryServiceMap.Load(name)
	//if ok {
	//	node := value.(registry.Node)
	//	_, err = e.client.Delete(context.TODO(), node.PathName)
	//	if err != nil {
	//		clog.PrintWa(err)
	//	}
	//}

	_, err = e.client.Delete(context.TODO(), path)
	return err
}

// 缓存中查询
func (e *EtcdRegistry) getServiceByCache(ctx context.Context, name string) (service *registry.Service, err error) {
	//name = e.servicePath(name)
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

	response, err := e.client.Get(context.TODO(), pathName, clientv3.WithPrefix())
	if err != nil {
		// 如果不存在 则返回err
		clog.PrintWa(err)
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
	serverInfo.serviceMap[name] = service
	e.value.Store(serverInfo)

	return service, nil
}

// 服务发现
func (e *EtcdRegistry) GetService(ctx context.Context, name string) (service *registry.Service, err error) {
	// 服务发现 先想缓存中获取  如果缓存被穿透 就想etcd中获取
	//clog.PrintEr(name)
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
			// v0.0.1 弃用
			//// 先查询是否存在  如果存在则续租，反之进行注册
			//ser, ok := e.registryServiceMap.Load(item.Name)
			//if ok {
			//	// 存在 添加节点信息
			//	server := ser.(*RegisterService)
			//	for _, i := range item.Nodes {
			//		server.service.Nodes = append(server.service.Nodes, i)
			//	}
			//
			//	server.service.Nodes = append(server.service.Nodes, i)
			//
			//	server.registered = false
			//
			//	e.registryServiceMap.Store(item.Name, server)
			//	continue
			//} else {
			//	// 不存在 进行注册
			//	registryed := &RegisterService{
			//		service: item,
			//	}
			//	e.registryServiceMap.Store(item.Name, registryed)
			//}
			e.registerOrKeepAlive(item)
		case <-ticker.C:
			// 更新缓存
			e.syncUpdateCache()
			//default:   1.0.1 版本后就失效了
			//	// 续约
			//	e.registerOrKeepAlive()
			//	time.Sleep(time.Millisecond * 200)
		}
	}
}

// 定时更新缓存   这个是保护措施
func (e *EtcdRegistry) syncUpdateCache() {
	log.Println("更新缓存")
	serverInfo := e.value.Load().(*AllServiceInfo)
	for k, _ := range serverInfo.serviceMap {
		_, err := e.getServiceByEtcd(context.TODO(), k, 2)
		if err != nil {
			clog.PrintWa(err)
		}
	}
}

// 注册 或者 续租
func (e *EtcdRegistry) registerOrKeepAlive(ser *registry.Node) {
	// v 0.0.1 弃用
	//e.registryServiceMap.Range(func(key, value interface{}) bool {
	//	item := value.(*RegisterService)
	//
	//	// 如果注册就续约
	//	if item.registered {
	//		e.keepAlive(item)
	//		return true
	//	}
	//	// 反之就进行注册
	//	err := e.registerServer(item)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	return true // 一直往下找
	//})

	service := &RegisterService{
		name:       ser.Name,
		node:       ser,
		registered: false,
	}
	id, _ := e.getServerName(ser)
	value, ok := e.registryServiceMap.Load(id)
	if ok {
		service = value.(*RegisterService)
		if service.registered {
			// 如果存在 续约
			e.keepAlive(service)
			return
		} else {
			// 进行服务注册
			err := e.registerServer(service)
			if err != nil {
				clog.PrintWa("服务注册失败")
				clog.PrintWa(err)
			}
			return
		}
	}
	// 进行服务注册
	err := e.registerServer(service)
	if err != nil {
		clog.PrintWa("服务注册失败")
		clog.PrintWa(err)
	}
}

// 续约
func (e *EtcdRegistry) keepAlive(server *RegisterService) {
	log.Println(server.id)
	select {
	case <-server.keepAliveChan:
		//clog.PrintWa(server.id)
		//clog.PrintWa("续租的")

		// 查询服务是否存在 如果不存在 则创建
		// 服务宕机  再次收到注册新号 进行注册
		key, path := e.getServerName(server.node)
		response, err := e.client.Get(context.TODO(), path)
		if err != nil {
			clog.PrintWa(err)
		}
		if len(response.Kvs) == 0 {
			server.registered = false
			go e.registerServer(server)
		}
		e.registryServiceMap.Store(key, server)

		if e.options.Debug {
			clog.Println(fmt.Sprintf("续租 server: %v, node %v, port: %v", server.node.Name, server.node.Ip, server.node.Port))
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

	server.id = response.ID

	//clog.PrintEr(server.id)
	//clog.PrintEr("这是注册的id")

	key, path := e.getServerName(server.node)
	bytes, err := utils.Jsonp.Marshal(server.node)
	if err != nil {
		clog.PrintWa(err)
		return err
	}

	_, err = e.client.Put(context.TODO(), path, string(bytes), clientv3.WithLease(server.id))
	if err != nil {
		clog.PrintWa(err)
		return err
	}

	responses, err := e.client.KeepAlive(context.TODO(), server.id)
	if err != nil {
		return err
	}

	server.keepAliveChan = responses
	server.registered = true
	if e.options.Debug {
		log.Printf("注册完毕 server: %v  host: %v  port: %v", server.name, server.node.Ip, server.node.Port)
	}

	// 更新服务缓存
	e.registryServiceMap.Store(key, server)

	e.updateCache(server)

	return nil
}

//// 获取etcd 配置路径
//func (e *EtcdRegistry) serviceNodePath(service *registry.Node) string {
//	nodeIP := fmt.Sprintf("%s:%d", service.Ip, service.Port)
//	return path.Join(e.options.RegistryPath, "/", service.Name, "/", nodeIP)
//}

// 获取服务名称   唯一id,路径
func (e *EtcdRegistry) getServerName(service *registry.Node) (id, paths string) {
	nodeIP := fmt.Sprintf("%s:%d", service.Ip, service.Port)
	path := path.Join(e.options.RegistryPath, "/", service.Name, "/", nodeIP)

	return utils.Sha256Encode(path), path
}

// 获取 service 路径
func (e *EtcdRegistry) servicePath(name string) string {
	return path.Join(e.options.RegistryPath, "/", name, "/")
}

// 更新缓存
func (e *EtcdRegistry) updateCache(server *RegisterService) {
	_, err := e.getServiceByEtcd(context.TODO(), server.name, 2)
	if err != nil {
		clog.PrintWa(err)
	}
}
