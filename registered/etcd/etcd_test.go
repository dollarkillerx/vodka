/**
 * @Author: DollarKillerX
 * @Description: etcd_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:36 2020/1/3
 */
package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/dollarkillerx/vodka/registered"
	"log"
	"testing"
	"time"
)

// 测试成功!!!
func TestEtcd_GetService(t *testing.T) {
	etcd := Etcd{opt: &registered.Options{}}
	etcdConfig := clientv3.Config{
		Endpoints: []string{"0.0.0.0:2079"},
		Username:  "golang",
		Password:  "123456",
	}
	err := etcd.Init(context.TODO(), registered.WithConfig(etcdConfig))
	if err != nil {
		log.Println(err)
	}

	s, err := etcd.Registry(context.TODO(), &registered.Node{
		Name:   "PPC",
		Addr:   "0.0.0.0:2079",
		Weight: 1,
		Load:   0,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(s)

	for {
		service, err := etcd.GetService(context.TODO(), "PPC")
		if err != nil {
			log.Fatalln("over")
		}
		log.Println(service.Name)
		for _, v := range service.Node {
			log.Println(v)
		}
		time.Sleep(time.Millisecond * 200)
	}
}
