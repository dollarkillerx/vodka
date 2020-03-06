/**
 * @Author: DollarKillerX
 * @Description: grpc_generator.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:04 2020/1/6
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// protoc --go_out=plugins=grpc:. *.proto

type grpcGenerator struct {
}

func (g *grpcGenerator) Name() string {
	return "grpcGenerator"
}

func (g *grpcGenerator) Run(opt *Option, data *RPCData) error {
	dir := fmt.Sprintf("plugins=grpc:%s", filepath.Join(opt.Output, "generate"))
	command := exec.Command("protoc", "--go_out", dir, opt.ProtoFileName)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func init() {
	generator := grpcGenerator{}
	genMgr.RegisterMgr(generator.Name(), &generator)

	routerGenerator := RouterGenerator{}
	genMgr.RegisterMgr(routerGenerator.Name(), &routerGenerator)

	mainGenerator := MainGenerator{}
	genMgr.RegisterMgr(mainGenerator.Name(), &mainGenerator)

}
