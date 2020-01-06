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
	"time"
)

// protoc --go_out=plugins=grpc:. *.proto

type grpcGenerator struct {
}

func (g *grpcGenerator) Name() string {
	return "grpcGenerator"
}

func (g *grpcGenerator) Run(opt *Option) error {
	dir := fmt.Sprintf("plugins=grpc:%s/generate/", opt.Output)
	command := exec.Command("protoc", "--go_out", dir, opt.ProtoFileName)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func init() {
	time.Sleep(time.Millisecond * 50) // 延迟等待 目录创建完毕
	generator := grpcGenerator{}
	genMgr.RegisterMgr(generator.Name(), &generator)
}
