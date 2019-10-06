/**
 * @Author: DollarKiller
 * @Description: 服务注册插件接口
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:32 2019-09-23
 */
package registry

import "context"

type Registry interface {
	Name() string                                                              // 插件名称
	Init(ctx context.Context, opts ...Option) (err error)                      // 初始化
	Register(ctx context.Context, service *Node) (err error)                   // 服务注册
	Unregister(ctx context.Context, service *Node) (err error)                 // 服务反注册
	GetService(ctx context.Context, name string) (service *Service, err error) // 服务发现 通过服务的名称获取服务的信息 (ip,port)
}
