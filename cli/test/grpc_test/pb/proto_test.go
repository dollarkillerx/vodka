/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 18:11
 */
package pb

import (
	"bytes"
	"fmt"
	"github.com/dollarkillerx/proto"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"text/template"
)

func TestCmd(t *testing.T) {
	dir := fmt.Sprintf("plugins=grpc:%s", filepath.Join("out", "generate"))
	command := exec.Command("protoc", "--go_out", dir, "test.proto")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestProtoEncoding(t *testing.T) {
	open, err := os.Open("test.proto")
	if err != nil {
		log.Fatalln(err)
	}

	parser := proto.NewParser(open)
	parse, e := parser.Parse()
	if e != nil {
		log.Fatalln(e)
	}
	c := grpcInit{}
	proto.Walk(
		parse,
		proto.WithService(c.withService),
		proto.WithMessage(c.withMessage),
		proto.WithRPC(c.withRPC),
		proto.WithPackage(c.withPackage),
	)

	for _, v := range c.rpc {
		fmt.Printf("rpc name: %s  RequestType:%s ReturnsType:%s StreamsRequest: %v StreamsRequest: %v \n", v.Name, v.RequestType, v.ReturnsType,
			v.StreamsRequest, v.StreamsReturns)
	}

	for _, v := range c.message {
		fmt.Printf("message: %s \n", v.Name)
	}

	fmt.Printf("Server: %s  Pkg: %s\n", c.service.Name, c.pkg.Name)
}

type grpcInit struct {
	rpc     []*proto.RPC
	service *proto.Service
	message []*proto.Message
	pkg     *proto.Package
}

func (c *grpcInit) Name() string {
	return "ControllerGenerator"
}

func (c *grpcInit) withPackage(pkg *proto.Package) {
	c.pkg = pkg
}

func (c *grpcInit) withRPC(rpc *proto.RPC) {
	c.rpc = append(c.rpc, rpc)
}

func (c *grpcInit) withMessage(message *proto.Message) {
	c.message = append(c.message, message)
}

func (c *grpcInit) withService(service *proto.Service) {
	c.service = service
}

func TestEncoding2(t *testing.T) {
	open, err := os.Open("test.proto")
	if err != nil {
		log.Fatalln(err)
	}

	parser := proto.NewParser(open)
	parse, e := parser.Parse()
	if e != nil {
		log.Fatalln(e)
	}
	c := grpcInit{}
	proto.Walk(
		parse,
		proto.WithService(c.withService),
		proto.WithMessage(c.withMessage),
		proto.WithRPC(c.withRPC),
		proto.WithPackage(c.withPackage),
	)
	result := RPCData{}
	result.Message = c.message
	result.Service = c.service
	result.Rpc = c.rpc
	result.Pkg = c.pkg

	generator := ControllerGenerator{}
	err = generator.Run(&Option{GoMod: "test"}, &result)
	if err != nil {
		log.Fatalln(err)
	}

}

type RPCData struct {
	Rpc     []*proto.RPC
	Service *proto.Service
	Message []*proto.Message
	Pkg     *proto.Package
}

type ControllerGenerator struct {
}

func (c *ControllerGenerator) Name() string {
	return "ControllerGenerator"
}

func (c *ControllerGenerator) Run(opt *Option, data *RPCData) error {
	fileBody := ""
	header := c.getHeader(opt, data)
	fileBody += header

	for _, v := range data.Rpc {
		fileBody += c.getBody(data, v)
	}
	ioutil.WriteFile("test.go", []byte(fileBody), 000666)
	return nil
}

func (c *ControllerGenerator) getHeader(opt *Option, data *RPCData) string {
	return fmt.Sprintf(ControllerTemplateHeader, opt.GoMod, data.Pkg.Name, data.Service.Name)
}

func (c *ControllerGenerator) getBody(data *RPCData, rpc *proto.RPC) string {
	fileBody := ""
	switch {
	// 双向流
	case rpc.StreamsReturns && rpc.StreamsRequest:
		fileBody = fmt.Sprintf(ControllerFunctionType3, data.Service.Name, rpc.Name, data.Pkg.Name, data.Service.Name, rpc.Name)
	// 普通
	case rpc.StreamsReturns == false && rpc.StreamsRequest == false:
		fileBody = fmt.Sprintf(ControllerFunctionType1, data.Service.Name, rpc.Name, data.Pkg.Name, rpc.RequestType, data.Pkg.Name, rpc.ReturnsType)
	// 客户端流
	case rpc.StreamsReturns == false && rpc.StreamsRequest == true:
		fileBody = fmt.Sprintf(ControllerFunctionType2, data.Service.Name, rpc.Name, data.Pkg.Name, rpc.RequestType, data.Pkg.Name, data.Service.Name, rpc.Name)
	// 服务端流
	case rpc.StreamsReturns == true && rpc.StreamsRequest == false:
		fileBody = fmt.Sprintf(ControllerFunctionType3, data.Service.Name, rpc.Name, data.Pkg.Name, data.Service.Name, rpc.Name)
	}

	return fileBody
}

type Option struct {
	ProtoFileName string // protoFile目录
	Output        string // 输出目录
	GenClientCode bool   // 生成client
	GenServerCode bool   // 生成server
	GoMod         string // go mod
	Prefix        string
}

var ControllerTemplateHeader = `
package controller

import (
	"context"
	"%s/generate/%s"
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

func TestEnd3(t *testing.T) {
	open, err := os.Open("test.proto")
	if err != nil {
		log.Fatalln(err)
	}

	parser := proto.NewParser(open)
	parse, e := parser.Parse()
	if e != nil {
		log.Fatalln(e)
	}
	c := grpcInit{}
	proto.Walk(
		parse,
		proto.WithService(c.withService),
		proto.WithMessage(c.withMessage),
		proto.WithRPC(c.withRPC),
		proto.WithPackage(c.withPackage),
	)

	bufferString := bytes.NewBufferString("")

	t2, err := template.New("main").Parse(CoreRouterTemplateHeader)
	if err != nil {
		log.Fatalln(err)
	}

	err = t2.Execute(bufferString, map[string]interface{}{
		"GoMod":       "pb",
		"RPC":         c.rpc,
		"Pkg":         "pb",
		"ServiceName": c.service.Name,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ioutil.WriteFile("aa.go", bufferString.Bytes(), 00755)
}

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
