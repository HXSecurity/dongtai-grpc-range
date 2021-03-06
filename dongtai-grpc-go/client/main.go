package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"

	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/grpc"
	_ "github.com/HXSecurity/DongTai-agent-go/run/http"
	DongtaiGRPC "github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/pb/go-pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	http.HandleFunc("/grpcCmd1", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmdStr := values.Get("cmd")
		str := done(cmdStr)
		cmdiOne(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd2", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := io.ReadAll(request.Body)
		fmt.Println(body)
		req := make(map[string]string)
		json.Unmarshal(body, &req)
		cmdStr := req["cmd"]
		str := done(cmdStr)
		cmdiTwo(str)
		writer.Write([]byte("hello" + string(body)))
	})

	http.HandleFunc("/grpcCmd3", func(writer http.ResponseWriter, request *http.Request) {
		cmdStr := request.Header.Get("Cmd")
		str := done(cmdStr)
		cmdiThree(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd4", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmdStr := values.Get("cmd")
		str := done(cmdStr)
		cmdiFour(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd5", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmdStr := values.Get("cmd")
		str := done(cmdStr)
		cmdiFive(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd6", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmdStr := values.Get("cmd")
		str := doneJava(cmdStr)
		cmdiOne(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd7", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := io.ReadAll(request.Body)
		fmt.Println(body)
		req := make(map[string]string)
		json.Unmarshal(body, &req)
		cmdStr := req["cmd"]
		str := doneJava(cmdStr)
		cmdiTwo(str)
		writer.Write([]byte("hello" + string(body)))
	})

	http.HandleFunc("/grpcCmd8", func(writer http.ResponseWriter, request *http.Request) {
		cmdStr := request.Header.Get("Cmd")
		str := doneJava(cmdStr)
		cmdiThree(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd9", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmdStr := values.Get("cmd")
		str := doneJava(cmdStr)
		cmdiFour(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpcCmd10", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		cmdStr := values.Get("cmd")
		str := doneJava(cmdStr)
		cmdiFive(str)
		writer.Write([]byte("hello" + cmdStr))
	})

	http.HandleFunc("/grpc", func(writer http.ResponseWriter, request *http.Request) {
		f, err := ioutil.ReadFile("./index.html")
		if err != nil {
			writer.Write([]byte(err.Error()))
		} else {
			writer.Write(f)
		}
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func done(cmd string) string {
	conn, err := grpc.Dial("177.7.0.11:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	client := DongtaiGRPC.NewGRPCServiceClient(conn)
	req, _ := client.Send(context.Background(), &DongtaiGRPC.Request{
		Text: cmd,
	})
	fmt.Println(req.Message)
	return req.Message
}

func doneJava(cmd string) string {
	conn, err := grpc.Dial("177.7.0.13:6565", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	client := DongtaiGRPC.NewGRPCServiceClient(conn)
	req, _ := client.Send(context.Background(), &DongtaiGRPC.Request{
		Text: cmd,
	})
	fmt.Println(req.Message)
	return req.Message
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
