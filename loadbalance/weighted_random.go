/**
 * @Author: DollarKillerX
 * @Description: weighted_random 加权随机
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:48 2020/1/4
 */
package loadbalance

import (
	"context"
	"errors"
	"github.com/dollarkillerx/vodka/registered"
	"math/rand"
	"time"
)

type WeightedRandom struct {
}

func NewWeightedRandom() *WeightedRandom {
	return &WeightedRandom{}
}

func (w *WeightedRandom) Name() string {
	return "WeightedRandom"
}

func (w *WeightedRandom) Select(ctx context.Context, server *registered.Service) (*registered.Node, error) {
	if len(server.Node) == 0 {
		return nil, errors.New("not work")
	}
	rand.Seed(time.Now().UnixNano())
	structure := w.structure(server)
	return structure[rand.Intn(len(structure))], nil
}

func (w *WeightedRandom) structure(server *registered.Service) []*registered.Node {
	result := []*registered.Node{}

	for _, v := range server.Node {
		for i := 0; i < v.Weight; i++ {
			result = append(result, v)
		}
	}
	return result
}
