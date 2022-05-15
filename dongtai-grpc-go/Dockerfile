FROM golang:alpine as builder

WORKDIR /go/src/github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && cd server/ && go build -gcflags "all=-N -l" \
    && cd ../client/ && go build -gcflags "all=-N -l" 

FROM alpine:latest
WORKDIR /dongtai
COPY --from=0 /go/src/github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/server/server ./
COPY --from=0 /go/src/github.com/HXSecurity/dongtai-grpc-range/dongtai-grpc-go/client/client ./
EXPOSE 8080
EXPOSE 8888
ENTRYPOINT ./server
