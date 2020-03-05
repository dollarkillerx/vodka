package pb

import (
	"context"
	"github.com/dollarkillerx/vodka/cli/test/pb/out/generate"
)

type ServiceController struct {
}

func (s *ServiceController) Run1(ctx context.Context,req *pb.Req) (*pb.Resp,error) {
	return nil,nil
}

func (s *ServiceController) Run2(ser pb.Service_Run2Server) error {
	return nil
}

func (s *ServiceController) Run3(req *pb.Req,ser pb.Service_Run3Server) error {
	return nil
}

func (s *ServiceController) Run4(ser pb.Service_Run4Server) error {
	return nil
}
