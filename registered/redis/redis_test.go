/**
 * @Author: DollarKillerX
 * @Description: redis_test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:40 2020/1/4
 */
package redis

import (
	"context"
	"github.com/dollarkillerx/easylog"
	"github.com/dollarkillerx/vodka/registered"
	"log"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	easylog.SetLevel("ERROR")
	redis := NewRedis()

	err := redis.Init(
		context.TODO(),
		registered.WithAddr([]string{"0.0.0.0:6379"}),
		registered.WithConfig("C9C8B3D369A83E57932EAF52C904C1C6"), // 配置注册中心密码
	)
	if err != nil {
		log.Fatalln(err)
	}

	s, err := redis.Registry(context.TODO(), &registered.Node{
		Name:   "Rpx",
		Addr:   "0.0.0.0:2079",
		Weight: 1,
		Load:   0,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(s)
	s, err = redis.Registry(context.TODO(), &registered.Node{
		Name:   "Rpx",
		Addr:   "0.0.0.0:2079",
		Weight: 1,
		Load:   0,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(s)

	for {
		service, err := redis.GetService(context.TODO(), "Rpx")
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
