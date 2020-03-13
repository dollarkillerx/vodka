/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 18:02
 */
package main

var RouterGeneratorTemplate = ` 
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package router

import (
	"{{.GoMod}}/controller"
	"{{.GoMod}}/core/router"
)

func Registry(app *router.Router) {
{{range $k,$v := .RPC}}
	app.{{$v.Name}}(controller.{{$v.Name}})
{{end}}
}
`
