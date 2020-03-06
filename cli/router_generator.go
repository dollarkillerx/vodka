/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 17:06
 */
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type RouterGenerator struct {
}

func (r *RouterGenerator) Name() string {
	return "RouterGenerator"
}

func (r *RouterGenerator) Run(opt *Option, data *RPCData) error {
	fileBody := fmt.Sprintf(RouterGeneratorTemplate, opt.GoMod)
	file := filepath.Join(opt.Output, "router", "app.go")
	return ioutil.WriteFile(file, []byte(fileBody), 00755)
}
