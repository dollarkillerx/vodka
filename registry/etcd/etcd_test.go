/**
 * @Author: DollarKiller
 * @Description: etcd 插件注册test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 12:51 2019-09-24
 */
package etcd

import (
	"context"
	"github.com/dollarkillerx/vodka/registry"
	"log"
	"testing"
	"time"
)

func TestRegistry(t *testing.T) {
	// 初始化etcd
	EtcdInit()

	regis, err := registry.InitRegistry(
		context.TODO(),
		"etcd",
		registry.WithDebug(true),
		registry.WithHeartBeat(8),
		registry.WithAddrs([]string{"127.0.0.1:2379"}),
		registry.WithTimeout(time.Second),
		registry.WithRegistryPath("/dkg/demo1/"),
	)

	if err != nil {
		t.Errorf("init registry failed, err:%v", err)
		return
	}

	node := registry.Node{
		Name:   "api",
		Ip:     "127.0.0.1",
		Port:   8081,
		Weight: 1,
	}

	regis.Register(context.TODO(), &node)

	go func() {
		for {
			time.Sleep(time.Second * 9)
			regis.Register(context.TODO(), &node)
		}
	}()

	go func() {
		time.Sleep(time.Second * 8)
		node := registry.Node{
			Name:   "api",
			Ip:     "127.0.0.1",
			Port:   8080,
			Weight: 1,
		}

		regis.Register(context.TODO(), &node)
		return
	}()

	for {
		time.Sleep(time.Second * 9)
		service, err := regis.GetService(context.TODO(), "api")
		if err != nil {
			log.Println(err)
		} else {
			service = service
			//log.Println(service.Name)
			//for _, item := range service.Nodes {
			//	log.Println(item)
			//}
		}
	}

}
