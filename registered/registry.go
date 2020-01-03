/**
 * @Author: DollarKillerX
 * @Description: registry 服务注册插件定义 用于规范注册中心
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:26 2020/1/3
 */
package registered

import "context"

type Registry interface {
	Name() string                                                          // 返回注册中心名词
	Init(ctx context.Context, set ...SetOption) error                      // 初始化注册中心
	Registry(ctx context.Context, node *Node) (string, error)              // 注册服务
	UnRegistry(ctx context.Context, node *Node) error                      // 服务反注册
	GetService(ctx context.Context, name string) (ser *Service, err error) // 服务发现
}
