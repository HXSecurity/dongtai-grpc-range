## 探针安装

到洞态网站 add Agent 下 
选择go 下载 config.yaml
选择java 下载 agent.jar
放到根目录下即可

启动 docker-compose up -d

访问：go-clinet：http://127.0.0.1:8082/grpc
   java-client：curl http://127.0.0.1:8080/grpc/send?text=1
       
