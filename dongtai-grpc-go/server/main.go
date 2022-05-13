package main

import (
	"context"
	"fmt"
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
	cmdiOne(comd)
	return &DongtaiGRPC.Response{
		State:   200,
		Message: "someThing",
	}, nil
}

func cmdiOne(cmdStr string) {
	cmd := exec.Command(cmdStr)
	cmd.Run()
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	fmt.Println(err)
	s := grpc.NewServer()
	DongtaiGRPC.RegisterGRPCServiceServer(s, &GRPCServer{})
	fmt.Println(s.Serve(l))

}
