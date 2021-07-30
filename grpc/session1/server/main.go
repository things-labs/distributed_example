package main

import (
	"log"
	"net"

	"github.com/thinkgos/distributed/grpc/pb"
	"github.com/thinkgos/distributed/grpc/services"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	pb.RegisterArithServer(rpcServer, new(services.Arith))

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
