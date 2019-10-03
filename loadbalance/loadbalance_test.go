/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:26 2019-10-03
 */
package loadbalance

import (
	"context"
	"fmt"
	"github.com/dollarkillerx/vodka/registry"
	"log"
	"testing"
)

func TestWeightedPolling(t *testing.T) {
	polling := NewWeightedPolling()

	var nodes []*registry.Node
	for i := 0; i < 8; i++ {
		node := &registry.Node{
			Ip:   fmt.Sprintf("127.0.0.%d", i),
			Port: 8080,
		}
		nodes = append(nodes, node)
	}

	data := registry.Service{Name: "ppc", Nodes: nodes}

	for i := 0; i < 100; i++ {
		node, err := polling.Select(context.TODO(), &data)
		if err == nil {
			log.Println(node)
		} else {
			log.Println("===")
		}
	}
}
