/**
 * @Author: DollarKillerX
 * @Description: defs.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:32 2019/12/31
 */
package _interface

import (
	"context"
	"github.com/dollarkillerx/vodka/register_test/service"
)

// 注册中心职责
// 注册
// 反注册
// 负载
// 健康
// 负载的概念不一定明确 , 给予用户传入特定数据的权力

type Register interface {
	Init(opts ...Option)                                     // 初始化注册中心
	Register(ctx context.Context, server *service.Service)   // 注册服务
	Unregister(ctx context.Context, server *service.Service) // 反注册服务
}
