package main

import (
	"context"
	"fmt"
	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/grpc"
	_ "github.com/HXSecurity/DongTai-agent-go/run/http"
	DongtaiGRPC "github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/pb/go-pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/grpc", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmd := values.Get("cmd")
		done(cmd)
		writer.Write([]byte("hello"))
	})
	http.ListenAndServe("127.0.0.1:1234", nil)
}

func done(cmd string) {
	//192.168.0.45:6565
	//conn, err := grpc.Dial("192.168.0.45:6565", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	client := DongtaiGRPC.NewGRPCServiceClient(conn)
	req, _ := client.Send(context.Background(), &DongtaiGRPC.Request{
		Text: cmd,
	})

	comd := req.Message
	cmds := exec.Command(comd)
	cmds.Run()
	fmt.Println(req)
}
