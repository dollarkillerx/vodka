/**
 * @Author: DollarKillerX
 * @Description: weighted_polling 加权轮询
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:01 2020/1/4
 */
package loadbalance

import (
	"context"
	"errors"
	"github.com/dollarkillerx/vodka/registered"
)

type WeightedPolling struct {
	id int
}

func NewWeightedPolling() *WeightedPolling {
	return &WeightedPolling{
		id: 0,
	}
}

func (w *WeightedPolling) Name() string {
	return "WeightedPolling"
}

func (w *WeightedPolling) Select(ctx context.Context, server *registered.Service) (*registered.Node, error) {
	if len(server.Node) == 0 {
		return nil, errors.New("not work")
	}
	structure := w.structure(server)
	if w.id > len(structure)-1 {
		w.id = 0
	}
	defer func() {
		w.id++
	}()
	return structure[w.id], nil
}

func (w *WeightedPolling) structure(server *registered.Service) []*registered.Node {
	result := []*registered.Node{}

	for _, v := range server.Node {
		for i := 0; i < v.Weight; i++ {
			result = append(result, v)
		}
	}
	return result
}
