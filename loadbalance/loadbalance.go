/**
 * @Author: DollarKiller
 * @Description: 负载均衡器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:03 2019-10-03
 */
package loadbalance

import (
	"context"

	"github.com/dollarkillerx/vodka/registry"
)

type LoadBalance interface {
	Name() string                                                                         // 算法名称
	Select(ctx context.Context, nodes *registry.Service) (node *registry.Node, err error) // 返回选中
}
