/**
 * @Author: DollarKillerX
 * @Description: dir_generator 目录生成工具
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:41 2020/1/6
 */
package main

import (
	"io/ioutil"
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
	"router",
	"app/config",
	"datamodels",
	"core/router",
	"generate",
	"middleware",
}

func dirGenerator(opt *Option) error {
	// 创建基础目录
	for _, v := range dirList {
		path := filepath.Join(opt.Output, v)
		err := os.MkdirAll(path, 00755)
		if err != nil {
			return err
		}
	}

	// 迁移 proto
	proto, err := ioutil.ReadFile(opt.ProtoFileName)
	if err != nil {
		log.Fatalln("Proto file read error Please check permissions or other !!! Err: ", err)
	}

	path := filepath.Join(opt.Output, "idl", opt.ProtoFileName)
	create, err := os.Create(path)
	if err != nil {
		log.Println("create proto err: ", err)
		return err
	}
	defer create.Close()
	_, err = create.Write(proto)
	if err != nil {
		log.Println("write proto err: ", err)
		return err
	}
	return nil
}
