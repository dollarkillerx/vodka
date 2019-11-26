/**
 * @Author: DollarKiller
 * @Description: random随机
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:53 2019-10-03
 */
package loadbalance

import (
	"context"
	"math/rand"
	"time"

	"github.com/dollarkillerx/vodka/registry"
)

type Random struct {
}

func (r *Random) Name() string {
	return "random"
}

func (r *Random) Select(ctx context.Context, nodes *registry.Service) (node *registry.Node, err error) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(nodes.Nodes))

	return nodes.Nodes[num], nil
}
