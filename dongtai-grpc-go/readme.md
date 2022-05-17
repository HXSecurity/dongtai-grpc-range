
## 探针安装

到洞态网站 add Agent 下 

选择go 下载 config.yaml 

放到 dongtai-grpc-go 目录下即可

进入dongtai-grpc-go 项目

## 项目配置

### 建议配置代理

#### Bash (Linux or macOS)
```cmd
# 配置 GOPROXY 环境变量
export GOPROXY=https://proxy.golang.com.cn,direct
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
export GOPRIVATE=git.mycompany.com,github.com/my/private
```
#### PowerShell (Windows)
```cmd
# 配置 GOPROXY 环境变量
$env:GOPROXY = "https://proxy.golang.com.cn,direct"
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
$env:GOPRIVATE = "git.mycompany.com,github.com/my/private"
```
执行命令 go mod tidy

进入 server 执行 go run main.go 启动服务端 默认端口为 8888

服务端端口修改可以进入 server下 main.go文件 第33行修改

进入client端口 执行 go run main.go  访问浏览器 http://127.0.0.1:2234/grpc

端口修改可以进入 server下 main.go文件 第28行修改

客户端（网页）端口可以进入client下main.go 修改第70行

客户端请求的服务端端口可以到 client下main.go  请求go服务
116行  conn, err := grpc.Dial("192.168.0.45:6565", grpc.WithTransportCredentials(insecure.NewCredentials()))

客户端请求的服务端端口可以到 client下main.go  请求java服务
130行  conn, err := grpc.Dial("192.168.0.45:6565", grpc.WithTransportCredentials(insecure.NewCredentials()))

## 项目运行

在 dongtai-grpc-go 目录下执行

go run -gcflags "all=-N -l" ./server/main.go  运行服务端

go run -gcflags "all=-N -l" ./client/main.go  运行客户端