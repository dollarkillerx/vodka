/**
 * @Author: DollarKillerX
 * @Description: etcd 注册中心
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:34 2020/1/3
 */
package etcd

import (
	"context"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/dollarkillerx/easylog"
	"github.com/dollarkillerx/vodka/registered"
	"github.com/dollarkillerx/vodka/utils"
)

type Etcd struct {
	opt    *registered.Options // 用户传入参数
	config clientv3.Config     // 配置模块
	client *clientv3.Client    // 连接维护
}

/**
 * 服务注册
 * 服务发现
 * 服务反注册
 */

func EtcdInit() {}

func (e *Etcd) Name() string {
	return "etcd"
}

func init() {
	etcd := Etcd{
		opt: &registered.Options{HeartBeat: 10},
	}
	registered.RegistryMar(&etcd)
}

// 初始化注册中心
func (e *Etcd) Init(ctx context.Context, set ...registered.SetOption) error {
	for _, v := range set {
		v(e.opt)
	}
	config, ok := e.opt.Config.(clientv3.Config)
	if !ok {
		return errors.New("config error")
	}
	client, err := clientv3.New(config)
	if err != nil {
		return err
	}
	e.client = client
	e.config = config
	return nil
}

// 服务注册
func (e *Etcd) Registry(ctx context.Context, node *registered.Node) (string, error) {
	// 如果node.ID == "" 说明该node未注册
	if node.ID == "" {
		s, err := utils.SonyFlakeGetId()
		if err != nil {
			return "", err
		}
		node.ID = s
		go e.registry(node) // 进行服务注册 或这 心跳
		return s, nil
	}
	go e.registry(node) // 进行服务注册 或在 心跳
	return node.ID, nil
}

// 服务反注册
func (e *Etcd) UnRegistry(ctx context.Context, node *registered.Node) error {
	path := e.getEtcdPath(node)
	_, err := clientv3.NewKV(e.client).Delete(context.TODO(), path)
	return err
}

// 服务发现
func (e *Etcd) GetService(ctx context.Context, node string) (*registered.Service, error) {
	path := fmt.Sprintf("/registry/%s/", node)
	kv := clientv3.NewKV(e.client)
	response, err := kv.Get(context.TODO(), path, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	result := registered.Service{Name: node}
	for _, v := range response.Kvs {
		node := &registered.Node{}

		err := utils.Json.Unmarshal(v.Value, node)
		if err != nil {
			easylog.PrintError(err)
		}
		result.Node = append(result.Node, node)
	}

	return &result, nil
}

// 服务注册&&心跳  (逻辑解耦)
func (e *Etcd) registry(node *registered.Node) {
	path := e.getEtcdPath(node)
	lease, err := clientv3.NewLease(e.client).Grant(context.TODO(), int64(e.opt.HeartBeat))
	if err != nil {
		easylog.PrintError(err)
		return
	}
	bytes, err := utils.Json.Marshal(node)
	if err != nil {
		easylog.PrintError(err)
	}
	_, err = clientv3.NewKV(e.client).Put(context.TODO(), path, string(bytes), clientv3.WithLease(lease.ID))
	if err != nil {
		easylog.PrintError(err)
	}
}

// 获取etcd 存储 路径 [etcd]中存储的格式 /registry/服务名称/服务id
func (e *Etcd) getEtcdPath(node *registered.Node) string {
	return fmt.Sprintf("/registry/%s/%s", node.Name, node.ID)
}
