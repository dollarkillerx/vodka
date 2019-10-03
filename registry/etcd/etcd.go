/**
 * @Author: DollarKiller
 * @Description: etcd 注册中心  (其他注册中心可以参考它)
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 12:51 2019-09-24
 */
package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dollarkillerx/vodka/registry"
	"go.etcd.io/etcd/clientv3"
	"log"
	"path"
	"sync"
	"time"
)

const (
	MaxServiceNum = 8
)

type EtcdRegistry struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service

	registryServiceMap sync.Map
}

var (
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

// 外部初始化
func EtcdInit() {}

func init() {
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
	return
}

// 获取服务
func (e *EtcdRegistry) GetService(ctx context.Context, name string) (service *registry.Service, err error) {
	return nil, nil
}

func (e *EtcdRegistry) run() {
	for {
		select {
		case item := <-e.serviceCh:
			// 先查询是否存在  如果存在则续租，反之进行注册
			_, ok := e.registryServiceMap.Load(item.Name)
			if ok {
				// 存在
				continue
			} else {
				// 不存在 进行注册
				registryed := &RegisterService{
					service: item,
				}
				e.registryServiceMap.Store(item.Name, registryed)
			}
		default:
			// 续约
			e.registerOrKeepAlive()
			time.Sleep(time.Millisecond * 200)
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

		return true  // 一直往下找
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
			log.Printf("续租 server: %v, node %v, port: %v",server.service.Name,server.service.Nodes[0].Ip,server.service.Nodes[0].Port)
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


	for _,node := range server.service.Nodes {
		// 为每个节点够着一个server 注入  /dkg/server1/127.0.0.1:8081  /dkg/server1/127.0.0.2:8081
		ser := &registry.Service{
			Name:server.service.Name,
			Nodes:[]*registry.Node{
				node,
			},
		}
		path := e.serviceNodePath(ser)
		bytes, err := json.Marshal(ser)
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
		log.Printf("注册完毕 server: %v  host: %v  port: %v",server.service.Name,server.service.Nodes[0].Ip,server.service.Nodes[0].Port)
	}
	return nil
}


// 获取etcd 配置路径
func (e *EtcdRegistry) serviceNodePath(service *registry.Service) string {
	nodeIP := fmt.Sprintf("%s:%d", service.Nodes[0].Ip, service.Nodes[0].Port)
	return path.Join(e.options.RegistryPath,"/", service.Name,"/", nodeIP)
}