package main

import (
	"context"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/grpc"
	DontaiGRPC "github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/pb/go-pb"
	"google.golang.org/grpc"
	"net"
	"os/exec"
)

type GRPCServer struct {
	DontaiGRPC.UnimplementedGRPCServiceServer
}

func (c *GRPCServer) Send(context context.Context, request *DontaiGRPC.Request) (*DontaiGRPC.Response, error) {
	comd := request.Text
	cmd := exec.Command(comd)
	cmd.Run()
	return &DontaiGRPC.Response{
		State:   200,
		Message: "someThing",
	}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	DontaiGRPC.RegisterGRPCServiceServer(s, &GRPCServer{})
	err := s.Serve(l)
	if err != nil {
		return
	}
}
