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
	"%s/core/router"
)

func Registry(app *router.Router) {
	
}
`
