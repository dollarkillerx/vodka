/**
 * @Author: DollarKillerX
 * @Description: controller_generator grpc 调用逻辑生成
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:45 2020/1/6
 */
package main

import (
	"github.com/dollarkillerx/proto"
	"os"
)

type controllerGenerator struct {
	rpc     []*proto.RPC
	service *proto.Service
	message []*proto.Message
}

func (c *controllerGenerator) Name() string {
	return "ControllerGenerator"
}

func (c *controllerGenerator) Run(opt *Option) error {
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
	)

	return nil
}

func (c *controllerGenerator) withRPC(rpc *proto.RPC) {
	c.rpc = append(c.rpc, rpc)
}

func (c *controllerGenerator) withMessage(message *proto.Message) {
	c.message = append(c.message, message)
}

func (c *controllerGenerator) withService(service *proto.Service) {
	c.service = service
}

// 获取rpc基本信息
func getRPCData(opt *Option) (*RPCData, error) {
	generator := controllerGenerator{
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
	return &result, nil
}
