/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 17:06
 */
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
)

type RouterGenerator struct {
}

func (r *RouterGenerator) Name() string {
	return "RouterGenerator"
}

func (r *RouterGenerator) Run(opt *Option, data *RPCData) error {
	fileBody := r.getBody(opt, data)
	file := filepath.Join(opt.Output, "router", "app.go")
	return ioutil.WriteFile(file, fileBody, 00755)
}

func (r *RouterGenerator) getBody(opt *Option, data *RPCData) []byte {
	bufferString := bytes.NewBufferString("")
	parse, err := template.New("main").Parse(RouterGeneratorTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	err = parse.Execute(bufferString, map[string]interface{}{
		"GoMod": opt.GoMod,
		"RPC":   data.Rpc,
	})
	if err != nil {
		log.Fatalln("Router Generator: ", err)
	}
	return bufferString.Bytes()
}
