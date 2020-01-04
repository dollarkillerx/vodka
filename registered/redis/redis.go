/**
 * @Author: DollarKillerX
 * @Description: redis 注册中心
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:01 2020/1/4
 */
package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/dollarkillerx/easylog"
	"github.com/dollarkillerx/vodka/registered"
	"github.com/dollarkillerx/vodka/utils"
	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	option *registered.Options
	pool   *redis.Pool
}

func RedisInit() {}

func NewRedis() *Redis {
	return &Redis{
		option: &registered.Options{HeartBeat: 10},
	}
}

func init() {
	registered.RegistryMar(NewRedis())
}

func (r *Redis) Name() string {
	return "REDIS"
}

// 初始化Etcd注册中心
func (r *Redis) Init(ctx context.Context, set ...registered.SetOption) error {
	for _, v := range set {
		v(r.option)
	}
	r.pool = &redis.Pool{
		MaxIdle:     50,                // 池中的最大空闲连接数
		MaxActive:   30,                // 最大连接数
		IdleTimeout: 300 * time.Second, // 超时回收
		Dial: func() (conn redis.Conn, e error) {
			// 1. 打开连接
			dial, e := redis.Dial("tcp", r.option.Addr[0])
			if e != nil {
				easylog.PrintWarning(e.Error())
				return nil, e
			}
			// 2. 访问认证
			s, b := r.option.Config.(string)
			if b == true {
				if s != "" {
					dial.Do("AUTH", s)
				}
			}
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 定时检查连接是否可用
			// time.Since(t) 获取离现在过了多少时间
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			easylog.PrintWarning(err)
			return err
		},
	}
	return nil
}

// 服务注册
func (r *Redis) Registry(ctx context.Context, node *registered.Node) (string, error) {
	// 如果node.ID == "" 说明该node未注册
	if node.ID == "" {
		s, err := utils.SonyFlakeGetId()
		if err != nil {
			return "", err
		}
		node.ID = s
		go r.registry(node) // 进行服务注册 或这 心跳
		return s, nil
	}
	go r.registry(node) // 进行服务注册 或在 心跳
	return node.ID, nil
}

// 服务反注册
func (r *Redis) UnRegistry(ctx context.Context, node *registered.Node) error {
	path := r.getRedisPath(node)
	redis := r.pool.Get()
	defer redis.Close()
	_, err := redis.Do("del", path)
	return err
}

// 服务发现
func (r *Redis) GetService(ctx context.Context, name string) (ser *registered.Service, err error) {
	path := fmt.Sprintf("/registry/%s/*", name)
	red := r.pool.Get()
	defer red.Close()
	result := registered.Service{Name: name}
	resultMap := map[string]*registered.Node{}
	for {
		values, err := redis.Values(red.Do("SCAN", "0", "MATCH", path))
		if err != nil {
			return nil, err
		}
		nodes, err := redis.Strings(values[1], nil)
		if err != nil {
			return nil, err
		}
		for _, id := range nodes {
			s, err := redis.Bytes(red.Do("get", id))
			if err != nil {
				easylog.PrintError(err)
				continue
			}
			node := &registered.Node{}
			err = utils.Json.Unmarshal(s, node)
			if err != nil {
				easylog.PrintError(err)
			}
			resultMap[node.ID] = node
		}
		i, err := redis.Int(values[0], nil)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			break
		}
	}

	for _, v := range resultMap {
		result.Node = append(result.Node, v)
	}
	return &result, nil
}

// 服务注册&&心跳  (逻辑解耦)
func (r *Redis) registry(node *registered.Node) {
	path := r.getRedisPath(node)
	bytes, err := utils.Json.Marshal(node)
	if err != nil {
		easylog.PrintError(err)
	}
	redis := r.pool.Get()
	defer redis.Close()
	_, err = redis.Do("setex", path, r.option.HeartBeat, string(bytes))
	if err != nil {
		easylog.PrintError(err)
	}
}

// 获取redis 存储 路径 [redis]中存储的格式 /registry/服务名称/服务id
func (r *Redis) getRedisPath(node *registered.Node) string {
	return fmt.Sprintf("/registry/%s/%s", node.Name, node.ID)
}
