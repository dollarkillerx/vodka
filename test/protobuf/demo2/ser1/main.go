/**
 * @Author: DollarKiller
 * @Description: server 1 test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 19:56 2019-10-03
 */
package main

import (
	"context"
	"github.com/dollarkillerx/vodka/test/protobuf/demo2/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *hello.HelloReq) (*hello.HelloResp, error) {

	// 逻辑处理
	log.Println(req)

	return &hello.HelloResp{
		Reply: "hello" + req.Name,
	}, nil
}

func main() {
	listener, e := net.Listen("tcp", ":50050")
	if e != nil {
		panic(e)
	}

	newServer := grpc.NewServer()

	hello.RegisterHelloServiceServer(newServer, &server{})

	newServer.Serve(listener)
}
