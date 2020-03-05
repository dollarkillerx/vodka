/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 19:57
 */
package main

var ControllerTemplateHeader = `
package controller

import (
	"context"
	"%s/generate"
)

type %sController struct {
}
`

// 普通类型
var ControllerFunctionType1 = `
func (s *%sController) %s(ctx context.Context,req *%s.%s) (*%s.%s,error) {
	return nil,nil
}
`

// 客户端流
var ControllerFunctionType2 = `
func (s *%sController) %s(req *%s.%s,ser %s.%s_%sServer) error {
	return nil
}
`

// 服务端流 & 双向流
var ControllerFunctionType3 = `
func (s *%sController) %s(ser %s.%s_%sServer) error {
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
