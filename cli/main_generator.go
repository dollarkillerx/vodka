/**
*@program: vodka
*@description: grpc main生成器
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 09:23
 */
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type MainGenerator struct {
}

func (m *MainGenerator) Name() string {
	return "MainGenerator"
}

func (m *MainGenerator) Run(opt *Option, data *RPCData) error {
	fileBody := fmt.Sprintf(MainTemplate, opt.GoMod, opt.GoMod, opt.GoMod)
	file := filepath.Join(opt.Output, "main", "main.go")
	return ioutil.WriteFile(file, []byte(fileBody), 00755)
}
