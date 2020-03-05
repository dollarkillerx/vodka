/**
 * @Author: DollarKillerX
 * @Description: 初始化grpc基本信息
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:45 2020/1/6
 */
package main

import (
	"github.com/dollarkillerx/proto"
	"os"
)

type grpcInit struct {
	rpc     []*proto.RPC
	message []*proto.Message
	service *proto.Service
	pkg     *proto.Package
}

func (c *grpcInit) Name() string {
	return "GrpcInit"
}

func (c *grpcInit) Run(opt *Option) error {
	file, e := os.Open(opt.ProtoFileName)
	if e != nil {
		return e
	}
	parser := proto.NewParser(file)
	parse, e := parser.Parse()
	if e != nil {
		return e
	}
	proto.Walk(
		parse,
		proto.WithService(c.withService),
		proto.WithMessage(c.withMessage),
		proto.WithRPC(c.withRPC),
		proto.WithPackage(c.withPackage),
	)

	return nil
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

func (c *grpcInit) withPackage(pkg *proto.Package) {
	c.pkg = pkg
}

// 获取rpc基本信息
func getRPCData(opt *Option) (*RPCData, error) {
	generator := grpcInit{
		rpc:     []*proto.RPC{},
		message: []*proto.Message{},
	}
	err := generator.Run(opt)
	if err != nil {
		return nil, err
	}
	result := RPCData{}
	result.Message = generator.message
	result.Service = generator.service
	result.Rpc = generator.rpc
	result.Pkg = generator.pkg
	return &result, nil
}
