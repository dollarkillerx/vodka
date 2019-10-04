/**
 * @Author: DollarKiller
 * @Description: 目录生成器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:38 2019-10-04
 */
package main

import (
	"github.com/dollarkillerx/vodka/utils"
	"log"
	"os"
)

var dirList = []string{
	"controller",
	"idl",
	"scripts",
	"conf",
	"app/routers",
	"app/config",
	"datamodels",
	"generate",
}

type DirGenerator struct {
	dirList []string
}

func (d *DirGenerator) Run(opt *Option) error {
	for _,i := range dirList {
		opts := utils.PathSlash(opt.Output)
		path := opts + "/" + i
		err := os.MkdirAll(path,00755)
		if err != nil {
			log.Fatalf("Failed to create directory error: %v \n",err)
		}
	}
	return nil
}



func init() {
	dirGenerator := DirGenerator{
		dirList:dirList,
	}

	Register("dir_generator",&dirGenerator)
}