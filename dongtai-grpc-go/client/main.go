package main

import (
	"context"
	"fmt"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/grpc"
	DontaiGRPC "github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/pb/go-pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	client := DontaiGRPC.NewGRPCServiceClient(conn)
	req, _ := client.Send(context.Background(), &DontaiGRPC.Request{
		Text: "发送测试数据",
	})
	fmt.Println(req)
}
