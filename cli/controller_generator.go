/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-07 10:57
 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type ControllerGenerator struct {
}

func (c *ControllerGenerator) Name() string {
	return "ControllerGenerator"
}

func (c *ControllerGenerator) Run(opt *Option, data *RPCData) error {
	for _, v := range data.Rpc {
		filename := filepath.Join(opt.Output, "controller", strings.ToLower(v.Name)+"_controller.go")
		fileBody := fmt.Sprintf(ControllerTemplate, opt.GoMod, v.Name, v.Name)
		err := ioutil.WriteFile(filename, []byte(fileBody), 00755)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}
