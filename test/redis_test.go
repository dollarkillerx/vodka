/**
 * @Author: DollarKillerX
 * @Description: redis_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:16 2020/1/4
 */
package test

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
	"time"
)

var (
	RedisConn *redis.Pool
)

func init() {
	RedisConn = newRedisPool()
}

// newRedisPool:创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,                // 池中的最大空闲连接数
		MaxActive:   30,                // 最大连接数
		IdleTimeout: 300 * time.Second, // 超时回收
		Dial: func() (conn redis.Conn, e error) {
			// 1. 打开连接
			dial, e := redis.Dial("tcp", "0.0.0.0:6379")
			if e != nil {
				fmt.Println(e.Error())
				return nil, e
			}
			// 2. 访问认证
			dial.Do("AUTH", "C9C8B3D369A83E57932EAF52C904C1C6")
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 定时检查连接是否可用
			// time.Since(t) 获取离现在过了多少时间
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func TestRedis(t *testing.T) {
	pool := newRedisPool()
	get := pool.Get()
	defer get.Close()
	reply, err := get.Do("set", "name", "ac")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(reply)
	reply2, err := redis.String(get.Do("get", "name"))
	if err != nil {
		log.Println(err)
	}
	log.Println(reply2)

	// set password
	reply2, err = redis.String(get.Do("config", "set", "requirepass", "C9C8B3D369A83E57932EAF52C904C1C6"))
	if err != nil {
		log.Println(err)
	}
	log.Println(reply2)
}

func TestSearch(t *testing.T) {
	pool := newRedisPool()
	get := pool.Get()
	defer get.Close()

	_, err := get.Do("set", "/col/ppr", "ok1")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = get.Do("set", "/col/pff", "ok2")
	if err != nil {
		log.Fatalln(err)
	}

	reply, err := redis.Values(get.Do("SCAN", "0", "MATCH", "/col/*"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(reply)

	log.Println(redis.Int(reply[0], nil))
	strings, err := redis.Strings(reply[1], nil)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range strings {
		log.Println(v)
	}
	//iter := 0
	//keys := []string{}
	//for {
	//	arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
	//	if err != nil {
	//		return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
	//	}
	//
	//	iter, _ = redis.Int(arr[0], nil)
	//	k, _ := redis.Strings(arr[1], nil)
	//	keys = append(keys, k...)
	//
	//	if iter == 0 {
	//		break
	//	}
	//}
}
