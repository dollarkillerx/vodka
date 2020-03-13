/**
*@program: vodka
*@description: grpc main生成器
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 09:23
 */
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
)

type MainGenerator struct {
}

func (m *MainGenerator) Name() string {
	return "MainGenerator"
}

func (m *MainGenerator) Run(opt *Option, data *RPCData) error {
	bs := bytes.NewBufferString("")
	parse, err := template.New("main").Parse(MainTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	err = parse.Execute(bs, map[string]interface{}{
		"Package": opt.GoMod,
		"Pkg":     data.Pkg.Name,
	})
	if err != nil {
		log.Fatalln(err)
	}

	file := filepath.Join(opt.Output, "main", "main.go")
	return ioutil.WriteFile(file, bs.Bytes(), 00755)
}
