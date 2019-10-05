/**
 * @Author: DollarKiller
 * @Description: controller 生成器
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 17:20 2019-10-04
 */
package main

import (
	"fmt"
	"github.com/dollarkillerx/proto"
	"github.com/dollarkillerx/vodka/utils"
	"log"
	"os"
)

type ControllerGenerator struct {
	rpc     []*proto.RPC
	service *proto.Service
	message []*proto.Message
}

func (c *ControllerGenerator) Run(opt *Option) error {
	file, e := os.Open(opt.Proto3Filename)
	if e != nil {
		log.Fatal(e)
	}

	parser := proto.NewParser(file)
	parse, e := parser.Parse()
	if e != nil {
		log.Fatal(e)
	}
	proto.Walk(
		parse,
		proto.WithRPC(c.withRPC),
		proto.WithService(c.withService),
		proto.WithMessage(c.withMessage),
	)

	return c.generator(opt)
}

func (c *ControllerGenerator) withRPC(rpc *proto.RPC) {
	c.rpc = append(c.rpc, rpc)
}

func (c *ControllerGenerator) withService(service *proto.Service) {
	c.service = service
}

func (c *ControllerGenerator) withMessage(message *proto.Message) {
	c.message = append(c.message, message)
}

func (c *ControllerGenerator) generator(opt *Option) error {
	filename := utils.PathSlash(opt.Output) + "/controller/" + c.service.Name + ".go"

	file, e := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 00755)
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	fmt.Fprintf(file, "package controller \n")
	fmt.Fprintf(file, "import ( \n")
	fmt.Fprintf(file, "\"context\" \n")
	fmt.Fprintf(file, "\"github.com/dollarkillerx/vodka/test/protobuf/demo2/hello\" \n")
	fmt.Fprintf(file, ") \n")

	for _, k := range c.rpc {
		fmt.Fprintf(file, "func %s(ctx context.Context,req *hello.%s) (*hello.%s,error) {\n }\n", k.Name, k.RequestType, k.ReturnsType)
	}

	return nil
}

func init() {
	conGenerator := ControllerGenerator{}

	Register("controller_generator", &conGenerator)
}
