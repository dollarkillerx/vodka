/**
 * @Author: DollarKiller
 * @Description: 插件管理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:47 2019-09-23
 */
package registry

import (
	"context"
	"fmt"
	"sync"
)

var (
	pluginMar = &PluginMgr{}  // 饿汉式 单利 保证多线程的安全性
)

// 插件管理
type PluginMgr struct {
	plugins sync.Map
}

// 插件注册
func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	_, ok := p.plugins.Load(plugin.Name())
	// 如果插件存在 返回以存在
	if ok {
		err = fmt.Errorf("duplicate registry plugin")
		return
	}
	p.plugins.Store(plugin.Name(), plugin)
	return
}

// 插件初始化
func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	plugin, ok := p.plugins.Load(name)
	if !ok {
		// 如果不存在 这跑错
		err = fmt.Errorf("plugin %s not exists", name)
		return
	}
	registry = plugin.(Registry)
	err = registry.Init(ctx, opts...)
	return
}

// 外部使用注册插件
func RegisterPlugin(registry Registry) (err error) {
	return pluginMar.registerPlugin(registry)
}

// 外部使用 初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMar.initRegistry(ctx, name, opts...)
}
