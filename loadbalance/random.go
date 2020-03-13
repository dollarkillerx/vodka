/**
 * @Author: DollarKillerX
 * @Description: random 随机负载器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:26 2020/1/4
 */
package loadbalance

import (
	"context"
	"errors"
	"github.com/dollarkillerx/vodka/registered"
	"math/rand"
	"time"
)

type RandomLoadBalance struct {
}

func NewRandomLoadBalance() *RandomLoadBalance {
	return &RandomLoadBalance{}
}

func (r *RandomLoadBalance) Name() string {
	return "RandomLoadBalance"
}

func (r *RandomLoadBalance) Select(ctx context.Context, server *registered.Service) (*registered.Node, error) {
	if len(server.Node) == 0 {
		return nil, errors.New("not work")
	}
	rand.Seed(time.Now().UnixNano())
	return server.Node[rand.Intn(len(server.Node))], nil
}
