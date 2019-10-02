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
	"go.etcd.io/etcd/clientv3"
	"sync"
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
		case item := <- e.serviceCh:
			// 先查询是否存在  如果存在则续租，反之进行注册
			_, ok := e.registryServiceMap.Load(item.Name)
			if ok {
				// 存在
			}else {
				// 不存在 进行注册
				
			}
		}
	}
}
