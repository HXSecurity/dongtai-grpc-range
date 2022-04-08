package main

import (
	"context"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/grpc"
	DongtaiGRPC "github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/pb/go-pb"
	"google.golang.org/grpc"
	"net"
	"os/exec"
)

type GRPCServer struct {
	DongtaiGRPC.UnimplementedGRPCServiceServer
}

func (c *GRPCServer) Send(context context.Context, request *DongtaiGRPC.Request) (*DongtaiGRPC.Response, error) {
	comd := request.Text
	cmd := exec.Command(comd)
	cmd.Run()
	return &DongtaiGRPC.Response{
		State:   200,
		Message: "someThing",
	}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	DongtaiGRPC.RegisterGRPCServiceServer(s, &GRPCServer{})
	err := s.Serve(l)
	if err != nil {
		return
	}
}
