/**
 * @Author: DollarKiller
 * @Description: 基于cache的注册中心  (这个是单机注册中心)
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:01 2019-10-07
 */
package cache

import (
	"context"
	"fmt"
	"github.com/bluele/gcache"
	"github.com/dollarkillerx/vodka/registry"
	"github.com/dollarkillerx/vodka/utils"
	"github.com/dollarkillerx/vodka/utils/clog"
	"log"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxServiceNum       = 8
	SyncServerCacheTime = 20 * time.Second // 为防止意外的定时缓存更新  分钟
)

type CacheRegistry struct {
	options   *registry.Options
	gcache    gcache.Cache
	serviceCh chan *registry.Node

	lock               sync.Mutex   // nutex 防止并发超量
	value              atomic.Value // 原子缓存  高效
	registryServiceMap sync.Map
}

var (
	// 单利 饿汉式
	cacheRegistry = &CacheRegistry{
		serviceCh: make(chan *registry.Node, MaxServiceNum),
		gcache:    gcache.New(20).LRU().Build(),
	}
)

type RegisterService struct {
	name       string         // 服务名称
	path       string         // 服务path
	node       *registry.Node // 服务体
	registered bool           // 是否注册
}

// 定义缓存结构体
type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

// 外部初始化
func CacheInit() {}

func init() {
	serverCache := &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}
	// 初始化缓存
	cacheRegistry.value.Store(serverCache)

	// 注册etcd 到管理中心
	err := registry.RegisterPlugin(cacheRegistry)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Etcd 注册中心初始化完毕")
	}
	go cacheRegistry.run()
}

func (c *CacheRegistry) Name() string {
	return "cache"
}

func (c *CacheRegistry) Init(ctx context.Context, opts ...registry.Option) (err error) {
	//基础设置
	options := registry.Options{}

	for _, k := range opts {
		k(&options)
	}

	c.options = &options

	return
}

// 服务注册
func (c *CacheRegistry) Register(ctx context.Context, service *registry.Node) (err error) {
	select {
	case c.serviceCh <- service:
	default:
		err = fmt.Errorf("register chan is full")
		return
	}
	return
}

// 服务反注册
func (c *CacheRegistry) Unregister(ctx context.Context, service *registry.Node) (err error) {
	name, _ := c.getServerName(service)
	c.registryServiceMap.Delete(name)

	//_, err = c.client.Delete(context.TODO(), path)
	return err
}

func (c *CacheRegistry) run() {
	//ticker := time.NewTicker(SyncServerCacheTime)
	for {
		select {
		case item := <-c.serviceCh:

			c.registerOrKeepAlive(item)
			//case <-ticker.C:
			//	// 更新缓存
			//	c.syncUpdateCache()
		}
	}
}

// 注册 or 续租
func (c *CacheRegistry) registerOrKeepAlive(node *registry.Node) {
	id, _ := c.getServerName(node)
	err := c.gcache.SetWithExpire(id, node, time.Second*time.Duration(c.options.HeartBeat))

	if err != nil {
		clog.PrintWa(err)
	}

	if c.options.Debug {
		clog.Println(fmt.Sprintf("注册OR续租 server: %v, node %v, port: %v", node.Name, node.Ip, node.Port))
	}
}

//// 更新缓存    考虑到这个是一个小系统  就不用缓存了
//func (c *CacheRegistry) syncUpdateCache() {
//	log.Println("更新缓存")
//	// 获取中心数据
//	data := &AllServiceInfo{}
//	data.serviceMap = map[string]*registry.Service{}
//	all := c.gcache.GetALL(true)
//	for _, i := range all {
//		node, bo := i.(*registry.Node)
//		if bo {
//			_,ok := data.serviceMap[node.Name]
//			if ok {
//				// 如果转码正确
//				data.serviceMap[node.Name].Nodes = append(data.serviceMap[node.Name].Nodes, node)
//			}else {
//				data.serviceMap[node.Name] = &registry.Service{
//					Name:node.Name,
//				}
//				data.serviceMap[node.Name].Nodes = append(data.serviceMap[node.Name].Nodes,node)
//			}
//		}
//	}
//
//	c.value.Store(data)
//}

// 服务发现
func (c *CacheRegistry) GetService(ctx context.Context, name string) (service *registry.Service, err error) {
	// 考虑到小系统就不用cache了
	//info := c.value.Load().(*AllServiceInfo)
	//
	//i,ok := info.serviceMap[name]
	//if ok {
	//	return i,nil
	//}
	//return nil, fmt.Errorf("not data")

	data := &AllServiceInfo{}
	data.serviceMap = map[string]*registry.Service{}
	all := c.gcache.GetALL(true)
	for _, i := range all {
		node, bo := i.(*registry.Node)
		if bo {
			_, ok := data.serviceMap[node.Name]
			if ok {
				// 如果转码正确
				data.serviceMap[node.Name].Nodes = append(data.serviceMap[node.Name].Nodes, node)
			} else {
				data.serviceMap[node.Name] = &registry.Service{
					Name: node.Name,
				}
				data.serviceMap[node.Name].Nodes = append(data.serviceMap[node.Name].Nodes, node)
			}
		}
	}

	i, ok := data.serviceMap[name]
	if ok {
		return i, nil
	}
	return nil, fmt.Errorf("not data")
}

// 获取服务名称   唯一id,路径
func (c *CacheRegistry) getServerName(service *registry.Node) (id, paths string) {
	nodeIP := fmt.Sprintf("%s:%d", service.Ip, service.Port)
	path := path.Join(c.options.RegistryPath, "/", service.Name, "/", nodeIP)

	return utils.Sha256Encode(path), path
}
