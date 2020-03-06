/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 19:57
 */
package main

// 普通类型
var CoreRouterFunctionType1 = `
func (s *%sRouter) %s(ctx context.Context,req *%s.%s) (*%s.%s,error) {
	return nil,nil
}
`

// 客户端流
var CoreRouterFunctionType2 = `
func (s *%sRouter) %s(req *%s.%s,ser %s.%s_%sServer) error {
	return nil
}
`

// 服务端流 & 双向流
var CoreRouterFunctionType3 = `
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

var CoreRouterTemplateHeader = ` 
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package router

import (
	pb "{{.GoMod}}/generate"
	"context"
	"log"
)

type Router struct {
	router *serviceRouter
}

func New() *Router {
	return &Router{
		router: &serviceRouter{
			{{range $k,$v := .RPC}}
			{{$v.Name}}FuncSlice: make([]RunFunc, 0),
			{{end}}
		},
	}
}

func (r *Router) RegistryGRPC() *serviceRouter {
	return r.router
}
{{range $k,$v := .RPC}}
func (r *Router) {{$v.Name}}({{$v.Name}}func ...RunFunc) {
	r.router.{{$v.Name}}FuncSlice = append(r.router.{{$v.Name}}FuncSlice, {{$v.Name}}func...)
}
{{end}}


type serviceRouter struct {
{{range $k,$v := .RPC}}
	{{$v.Name}}FuncSlice []RunFunc
{{end}}
}

type RouterContextItem interface {
	_routerContext()
}

type RouterContext struct {
	Ctx      RouterContextItem
	funcList []RunFunc
	index    int
}

func (r *RouterContext) Next() {
	r.index++
	if r.index <= len(r.funcList) {
		r.funcList[r.index-1](r)
	} else {
		log.Println("RouterContext Next  what ???")
	}
}
{{range $k,$v := .RPC}}
type {{%v.Name}}FuncContext struct {
	{{if $v.StreamsRequest and $v.StreamsReturns}}
	Ctx  context.Context
	Req  *pb.Req
	Resp *pb.Resp
	Err  error
	{{else if eq $v.StreamsRequest true}}
	
	{{end}}
}
{{end}}
type Run2FuncContext struct {
	Req *pb.Req
	Ser pb.Service_Run2Server
	Err error
}
type Run3FuncContext struct {
	Ser pb.Service_Run3Server
	Err error
}
type Run4FuncContext struct {
	Ser pb.Service_Run4Server
	Err error
}

{{range $k,$v := .RPC}}
func (r *{{$v.Name}}FuncContext) _routerContext() {}
{{end}}
type RunFunc func(ctx *RouterContext)
`