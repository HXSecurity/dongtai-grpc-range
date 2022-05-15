## 探针安装

到洞态网站 add Agent 页面下载以下两个agent放到根目录下即可：

1.选择go 下载 config.yaml 

2.选择java 下载 agent.jar



```cmd
启动 docker-compose up -d

访问：go-clinet：http://127.0.0.1:8082/grpc

访问：java-client：http://127.0.0.1:8083/grpc/send?text=1     
```


