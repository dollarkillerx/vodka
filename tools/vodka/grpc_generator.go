/**
 * @Author: DollarKiller
 * @Description: grpc生成
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 12:34 2019-10-04
 */
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/dollarkillerx/vodka/utils"
)

type GrpcGenerator struct {
}

func (g *GrpcGenerator) Run(opt *Option) error {
	// protoc --go_out=plugins=grpc:. protu.proto

	path := fmt.Sprintf("%s/generate", utils.PathSlash(opt.Output))
	os.MkdirAll(path, 00755)
	dir := fmt.Sprintf("plugins=grpc:%s", path)

	cmd := exec.Command("protoc", "--go_out", dir, opt.Proto3Filename)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("grpc generator failed, err:%v\n", err)
	}
	return nil
}

func init() {
	grpcGenerator := GrpcGenerator{}

	Register("grpc_generator", &grpcGenerator)
}
