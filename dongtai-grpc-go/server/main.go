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
	cmdiTwo(comd)
	cmdiThree(comd)
	cmdiFour(comd)
	cmdiFive(comd)
	return &DongtaiGRPC.Response{
		State:   200,
		Message: "someThing",
	}, nil
}

func cmdiOne(cmdStr string) {
	cmd := exec.Command(cmdStr)
	cmd.Run()
}

func cmdiTwo(cmdStr string) {
	cmd := exec.Command("curl", cmdStr)
	cmd.Run()
}

func cmdiThree(cmdStr string) {
	endCmd := "-a" + cmdStr
	cmd := exec.Command("curl", endCmd)
	cmd.Run()
}

func cmdiFour(cmdStr string) {
	endCmd := []string{"-a", cmdStr}
	cmd := exec.Command("curl", endCmd...)
	cmd.Run()
}

func cmdiFive(cmdStr string) {
	endCmd := make(map[string]string)
	endCmd["one"] = "-a"
	endCmd["two"] = cmdStr
	cmd := exec.Command(endCmd["two"], endCmd["one"])
	cmd.Run()
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	fmt.Println(err)
	s := grpc.NewServer()
	DongtaiGRPC.RegisterGRPCServiceServer(s, &GRPCServer{})
	fmt.Println(s.Serve(l))

}
