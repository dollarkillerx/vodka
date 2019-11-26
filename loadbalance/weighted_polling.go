/**
 * @Author: DollarKiller
 * @Description: 加权轮询
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:07 2019-10-03
 */
package loadbalance

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/dollarkillerx/vodka/registry"
)

type WeightedPolling struct {
	num atomic.Value
}

func NewWeightedPolling() *WeightedPolling {
	item := &WeightedPolling{}
	item.num.Store(0)
	return item
}

func (w *WeightedPolling) Name() string {
	return "WeightedPolling"
}

func (w *WeightedPolling) Select(ctx context.Context, nodes *registry.Service) (node *registry.Node, err error) {
	if len(nodes.Nodes) <= 0 {
		return nil, fmt.Errorf("not data")
	}
	var nodec []*registry.Node
	for _, c := range nodes.Nodes {
		if c.Weight == 0 {
			c.Weight = 1
		}
		for i := 0; i < c.Weight; i++ {
			nodec = append(nodec, c)
		}
	}

	i := w.num.Load().(int)
	if i >= len(nodec) {
		i = 0
		node = nodec[i]
		i += 1
		w.num.Store(i)
		return node, nil
	} else {
		node = nodec[i]
		i += 1
		w.num.Store(i)
		return node, nil
	}
}
