/**
 * @Author: DollarKillerX
 * @Description: polling 轮询负载均衡
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:31 2020/1/4
 */
package loadbalance

import (
	"context"
	"errors"
	"github.com/dollarkillerx/vodka/registered"
)

type PollingLoadBalance struct {
	id int // 上次所在空间的位置
}

func NewPollingLoadBalance() *PollingLoadBalance {
	return &PollingLoadBalance{
		id: 0,
	}
}

func (p *PollingLoadBalance) Name() string {
	return "PollingLoadBalanc"
}

func (p *PollingLoadBalance) Select(ctx context.Context, server *registered.Service) (*registered.Node, error) {
	if len(server.Node) == 0 {
		return nil, errors.New("not work")
	}

	if p.id > len(server.Node)-1 {
		p.id = 0
	}
	defer func() {
		p.id++
	}()
	return server.Node[p.id], nil
}
