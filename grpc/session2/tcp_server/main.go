package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/distributed/grpc/pb"
	"github.com/thinkgos/distributed/grpc/services"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("../../ssl/no_password_server.crt", "../../ssl/no_password_server.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterArithServer(rpcServer, new(services.Arith))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
