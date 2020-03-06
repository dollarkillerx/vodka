/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 19:56
 */
package main

import (
	"fmt"
	"github.com/dollarkillerx/proto"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type RouterGenerator struct {
}

func (c *RouterGenerator) Name() string {
	return "RouterGenerator"
}

func (c *RouterGenerator) Run(opt *Option, data *RPCData) error {
	fileBody := ""
	header := c.getHeader(opt, data)
	fileBody += header

	for _, v := range data.Rpc {
		fileBody += c.getBody(data, v)
	}

	file := filepath.Join(opt.Output, "router", strings.ToLower(data.Service.Name)+"_router.go")
	return ioutil.WriteFile(file, []byte(fileBody), 00755)
}

func (c *RouterGenerator) getHeader(opt *Option, data *RPCData) string {
	return fmt.Sprintf(RouterTemplateHeader, opt.GoMod, data.Service.Name)
}

func (c *RouterGenerator) getBody(data *RPCData, rpc *proto.RPC) string {
	fileBody := ""
	switch {
	// 双向流
	case rpc.StreamsReturns && rpc.StreamsRequest:
		fileBody = fmt.Sprintf(RouterFunctionType3, data.Service.Name, rpc.Name, data.Pkg.Name, data.Service.Name, rpc.Name)
	// 普通
	case rpc.StreamsReturns == false && rpc.StreamsRequest == false:
		fileBody = fmt.Sprintf(RouterFunctionType1, data.Service.Name, rpc.Name, data.Pkg.Name, rpc.RequestType, data.Pkg.Name, rpc.ReturnsType)
	// 客户端流
	case rpc.StreamsReturns == false && rpc.StreamsRequest == true:
		fileBody = fmt.Sprintf(RouterFunctionType3, data.Service.Name, rpc.Name, data.Pkg.Name, data.Service.Name, rpc.Name)
	// 服务端流
	case rpc.StreamsReturns == true && rpc.StreamsRequest == false:
		fileBody = fmt.Sprintf(RouterFunctionType2, data.Service.Name, rpc.Name, data.Pkg.Name, rpc.RequestType, data.Pkg.Name, data.Service.Name, rpc.Name)
	}

	return fileBody
}
