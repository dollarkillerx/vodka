/**
 * @Author: DollarKiller
 * @Description: Generator 管理
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:43 2019-10-04
 */
package main

import (
	"fmt"
	"sync"
)

var genMgr = &GeneratorMgr{
	genMap: sync.Map{},
}

type GeneratorMgr struct {
	genMap sync.Map
}

func (g *GeneratorMgr) Run(opt *Option) error {
	var err error
	g.genMap.Range(func(key, value interface{}) bool {
		generator := value.(Generator)
		err = generator.Run(opt)
		if err != nil {
			return false
		}
		return true
	})
	return err
}

func Register(name string, generator Generator) error {
	_, ok := genMgr.genMap.Load(name)
	if ok {
		return fmt.Errorf("generator %s is exists", name)
	}

	genMgr.genMap.Store(name, generator)

	return nil
}
