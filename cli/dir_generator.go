/**
 * @Author: DollarKillerX
 * @Description: dir_generator 目录生成工具
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:41 2020/1/6
 */
package main

import (
	"log"
	"os"
	"path/filepath"
)

var dirList = []string{
	"controller",
	"idl",
	"main",
	"scripts",
	"conf",
	"app/router",
	"app/config",
	"datamodels",
	"generate",
}

type dirGenerator struct {
}

func (d *dirGenerator) Name() string {
	return "DirGenerator"
}

func (d *dirGenerator) Run(opt *Option) error {
	for _, v := range dirList {
		path := filepath.Join(opt.Output, v)
		err := os.MkdirAll(path, 00755)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func init() {
	generator := dirGenerator{}
	genMgr.RegisterMgr(generator.Name(), &generator)
}
