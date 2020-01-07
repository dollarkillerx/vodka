/**
 * @Author: DollarKillerX
 * @Description: test_server.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:55 2020/1/7
 */
package main

import (
	"context"
	test "github.com/dollarkillerx/vodka/cli/output/generate"
	"google.golang.org/grpc"
	"net"
)

type server struct {
}

func (s *server) Run(ctx context.Context, req *test.Req) (*test.Resp, error) {
	return &test.Resp{}, nil
}

func main() {
	listener, e := net.Listen("tcp", "0.0.0.0:8083")
	if e != nil {
		panic(e)
	}
	serve := grpc.NewServer()
	test.RegisterServiceServer(serve, &server{})
	e = serve.Serve(listener)
	if e != nil {
		panic(e)
	}
}
