/**
 * @Author: DollarKiller
 * @Description: etcd 注册中心  (其他注册中心可以参考它)
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 12:51 2019-09-24
 */
package etcd

import (
	"context"
	"dkg/registry"
	"fmt"
	"go.etcd.io/etcd/clientv3"
)

type EtcdRegistry struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service
}

var (
	etcdRegistry = &EtcdRegistry{
		serviceCh: make(chan *registry.Service, 8),
	}
)

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

}
