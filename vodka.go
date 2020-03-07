/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-07 14:58
 */
package vodka

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	log.Println("Vodka is initialized")
}

type vodka struct {
	gRPC *grpc.Server
}

func New() *vodka {
	return &vodka{
		gRPC: grpc.NewServer()}
}

func (v *vodka) RegisterServer() *grpc.Server {
	return v.gRPC
}

func (v *vodka) Run(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return v.gRPC.Serve(listen)
}
