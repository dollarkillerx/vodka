/**
 * @Author: DollarKillerX
 * @Description: plugin_mar.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:26 2020/1/3
 */
package registered

import (
	"context"
	"errors"
	"sync"
)

var registryMarInstantiate = &registryMar{} // 单例

// 服务管理
type registryMar struct {
	persistence sync.Map
}

// 插件注册
func (r *registryMar) registry(reg Registry) {
	r.persistence.LoadOrStore(reg.Name(), reg)
}

// 初始化插件
func (r *registryMar) initRegistry(name string, option ...SetOption) error {
	value, ok := r.persistence.Load(name)
	if !ok {
		return errors.New("not registry")
	}
	registry, ok := value.(Registry)
	if !ok {
		return errors.New("registry error")
	}
	return registry.Init(context.TODO(), option...)
}

// 以下是对外暴露的

func RegistryMar(reg Registry) {
	registryMarInstantiate.registry(reg)
}

func InitRegistryMar(name string, option ...SetOption) error {
	return registryMarInstantiate.initRegistry(name, option...)
}
