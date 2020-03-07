/**
*@program: vodka
*@description: router 核心部件生成
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 19:56
 */
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

type CoreRouterGenerator struct {
}

func (c *CoreRouterGenerator) Name() string {
	return "CoreRouterGenerator"
}

func (c *CoreRouterGenerator) Run(opt *Option, data *RPCData) error {
	fileBody := c.getBody(opt, data)
	file := filepath.Join(opt.Output, "core", "router", strings.ToLower(data.Service.Name)+"_router.go")
	return ioutil.WriteFile(file, fileBody, 00755)
}

func (c *CoreRouterGenerator) getBody(opt *Option, data *RPCData) []byte {
	bufferString := bytes.NewBufferString("")
	t2, err := template.New("main").Parse(CoreRouterTemplateHeader)
	if err != nil {
		log.Fatalln(err)
	}
	err = t2.Execute(bufferString, map[string]interface{}{
		"GoMod":       opt.GoMod,
		"RPC":         data.Rpc,
		"Pkg":         data.Pkg.Name,
		"ServiceName": data.Service.Name,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return bufferString.Bytes()
}
