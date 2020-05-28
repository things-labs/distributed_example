package main

import (
	"log"
	"net"

	"github.com/thinkgos/grpcexample/session1/services"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
