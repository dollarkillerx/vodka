/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 19:57
 */
package main

var RouterTemplateHeader = `
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package router

import (
	"%s/generate"
	"context"
)

type %sRouter struct {
}
`

// 普通类型
var RouterFunctionType1 = `
func (s *%sRouter) %s(ctx context.Context,req *%s.%s) (*%s.%s,error) {
	return nil,nil
}
`

// 客户端流
var RouterFunctionType2 = `
func (s *%sRouter) %s(req *%s.%s,ser %s.%s_%sServer) error {
	return nil
}
`

// 服务端流 & 双向流
var RouterFunctionType3 = `
func (s *%sRouter) %s(ser %s.%s_%sServer) error {
	return nil
}
`

//var ControllerTemplateHeader = `
//package controller
//
//import (
//	"context"
//	"{{.GoMod}}/generate/{{.PkgName}}"
//)
//
//type {{.ServerName}}Controller struct {
//}
//`
//
//// 普通类型
//var ControllerFunctionType1 = `
//func (s *{{.ServerName}}Controller) {{.Name}}(ctx context.Context,req *{{.PkgName}}.RequestType) (*{{.PkgName}}.ReturnsType,error) {
//	return nil,nil
//}
//`
//
//// 客户端流
//var ControllerFunctionType2 = `
//func (s *{{.ServerName}}Controller) {{.Name}}(req *{{.PkgName}}.RequestType,ser {{.PkgName}}.{{.ServerName}}_{{.Name}}Server) error {
//	return nil
//}
//`
//
//// 服务端流 & 双向流
//var ControllerFunctionType3 = `
//func (s *{{.ServerName}}Controller) {{.Name}}(ser {{.PkgName}}.{{.ServerName}}_{{.Name}}Server) error {
//	return nil
//}
//`
