/**
*@program: vodka
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-05 16:39
 */
package main

import (
	"context"
	"github.com/dollarkillerx/vodka/cli/test/grpc_test/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func main() {
	server := grpc.NewServer()
	pb.RegisterServiceServer(server, &Server{})
	dial, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(dial)
}

func (s *Server) Run1(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return &pb.Resp{}, nil
}

func (s *Server) Run2(req *pb.Req, ser pb.Service_Run2Server) error {
	return nil
}

func (s *Server) Run3(ser pb.Service_Run3Server) error {
	return nil
}

func (s *Server) Run4(ser pb.Service_Run4Server) error {
	return nil
}
