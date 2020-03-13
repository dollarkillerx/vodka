/**
*@program: vodka
*@description: prometheus generator
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-13 15:33
 */
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
)

type PrometheusGenerator struct {
}

func (p *PrometheusGenerator) Name() string {
	return "PrometheusGenerator"
}

func (p *PrometheusGenerator) Run(opt *Option, data *RPCData) error {
	bs := bytes.NewBufferString("")
	parse, err := template.New("main").Parse(PrometheusTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	err = parse.Execute(bs, map[string]interface{}{
		"Package": opt.GoMod,
	})
	if err != nil {
		log.Fatalln(err)
	}

	file := filepath.Join(opt.Output, "middleware", "prometheus.go")
	return ioutil.WriteFile(file, bs.Bytes(), 00755)
}
