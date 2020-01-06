/**
 * @Author: DollarKillerX
 * @Description: generator_mgr 生产插件管理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:07 2020/1/6
 */
package main

import (
	"sync"
)

var genMgr = &generatorMgr{}

// 插件管理
type generatorMgr struct {
	mgr sync.Map
}

func (g *generatorMgr) Run(opt *Option) error {
	var err error
	genMgr.mgr.Range(func(key, value interface{}) bool {
		generator := value.(Generator)
		err := generator.Run(opt)
		if err != nil {
			return false
		}
		return true
	})
	return err
}

// 注册插件
func (g *generatorMgr) RegisterMgr(name string, mgr Generator) {
	genMgr.mgr.Store(name, mgr)
}
