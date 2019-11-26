/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:49 2019-10-06
 */
package etcd

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func TestEt1(t *testing.T) {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 10 * time.Second,
	}

	client, e := clientv3.New(config)
	if e != nil {
		panic(e)
	}

	g1, e := client.Grant(context.TODO(), 8)
	if e != nil {
		panic(e)
	}

	p1, e := client.Put(context.TODO(), "/vv/v1", "vvv1", clientv3.WithLease(g1.ID))
	if e != nil {
		panic(e)
	}
	p1 = p1

	k1, e := client.KeepAlive(context.TODO(), g1.ID)
	if e != nil {
		panic(e)
	}

	go func() {
		g2, e := client.Grant(context.TODO(), 8)
		if e != nil {
			panic(e)
		}
		_, e = client.Put(context.TODO(), "/vv/v2", "vvv2", clientv3.WithLease(g2.ID))
		if e != nil {
			panic(e)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			resp, i := client.Get(context.TODO(), "/vv", clientv3.WithPrefix())
			if i != nil {
				panic(i)
			}
			for _, kv := range resp.Kvs {
				log.Println(kv.Value)
			}

		}
	}()

	for {
		time.Sleep(time.Second * 5)
		<-k1
	}

}
