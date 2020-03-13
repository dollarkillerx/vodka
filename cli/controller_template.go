/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-07 10:58
 */
package main

var ControllerTemplate = `
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package controller

import (
	"%s/core/router"
)

func %s(ctx *router.RouterContext) {
	context := ctx.Ctx.(*router.%sFuncContext)
	context = context
}
`
