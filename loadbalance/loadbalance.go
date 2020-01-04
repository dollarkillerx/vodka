/**
 * @Author: DollarKillerX
 * @Description: loadbalance 负载均衡器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:23 2020/1/4
 */
package loadbalance

import (
	"context"
	"github.com/dollarkillerx/vodka/registered"
)

type LoadBalance interface {
	Name() string                                                                      // 返回负载器名词
	Select(ctx context.Context, service *registered.Service) (*registered.Node, error) // 传入获取到的服务列表 返回负载后的节点
}
