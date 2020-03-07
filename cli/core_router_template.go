/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 19:57
 */
package main

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
type {{$v.Name}}FuncContext struct {
	{{if and (eq $v.StreamsRequest true) (eq $v.StreamsReturns true)}}
	Ser {{$.Pkg}}.{{$.ServiceName}}_{{$v.Name}}Server
	{{else if and (eq $v.StreamsRequest false) (eq $v.StreamsReturns false)}}
	Ctx  context.Context
	Req  *{{$.Pkg}}.{{$v.RequestType}}
	Resp *{{$.Pkg}}.{{$v.ReturnsType}}
	{{else if and (eq $v.StreamsRequest true) (eq $v.StreamsReturns false)}}
	Ser {{$.Pkg}}.{{$.ServiceName}}_{{$v.Name}}Server
	{{else if and (eq $v.StreamsRequest false) (eq $v.StreamsReturns true)}}
	Req  *{{$.Pkg}}.{{$v.RequestType}}
	Ser {{$.Pkg}}.{{$.ServiceName}}_{{$v.Name}}Server	
	{{end}}
	Err  error
}
{{end}}

{{range $k,$v := .RPC}}
func (r *{{$v.Name}}FuncContext) _routerContext() {}
{{end}}
type RunFunc func(ctx *RouterContext)

// 下面是主题方法
{{range $k,$v := .RPC}}
{{if and (eq $v.StreamsRequest true) (eq $v.StreamsReturns true)}}
func (s *serviceRouter) {{$v.Name}}(ser {{$.Pkg}}.{{$.ServiceName}}_{{$v.Name}}Server) error {
	routerContext := RouterContext{
		Ctx: &{{$v.Name}}FuncContext{
			Err: nil,
			Ser: ser,
		},
		funcList: s.{{$v.Name}}FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*{{$v.Name}}FuncContext)
	return funcContext.Err
}
{{else if and (eq $v.StreamsRequest false) (eq $v.StreamsReturns false)}}
func (s *serviceRouter) {{$v.Name}}(ctx context.Context, req *{{$.Pkg}}.{{$v.RequestType}}) (*{{$.Pkg}}.{{$v.ReturnsType}}, error) {
	routerContext := RouterContext{
		Ctx: &{{$v.Name}}FuncContext{
			Ctx:  ctx,
			Req:  req,
			Resp: nil,
			Err:  nil,
		},
		funcList: s.{{$v.Name}}FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*{{$v.Name}}FuncContext)
	return funcContext.Resp, funcContext.Err
}
{{else if and (eq $v.StreamsRequest true) (eq $v.StreamsReturns false)}}
func (s *serviceRouter) {{$v.Name}}(ser {{$.Pkg}}.{{$.ServiceName}}_{{$v.Name}}Server) error {
	routerContext := RouterContext{
		Ctx: &{{$v.Name}}FuncContext{
			Err: nil,
			Ser: ser,
		},
		funcList: s.{{$v.Name}}FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*{{$v.Name}}FuncContext)
	return funcContext.Err
}
{{else if and (eq $v.StreamsRequest false) (eq $v.StreamsReturns true)}}
func (s *serviceRouter) {{$v.Name}}(req *{{$.Pkg}}.{{$v.RequestType}}, ser {{$.Pkg}}.{{$.ServiceName}}_{{$v.Name}}Server) error {
	routerContext := RouterContext{
		Ctx: &{{$v.Name}}FuncContext{
			Req: req,
			Ser: ser,
			Err: nil,
		},
		funcList: s.{{$v.Name}}FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*{{$v.Name}}FuncContext)
	return funcContext.Err
}
{{end}}
{{end}}
`
