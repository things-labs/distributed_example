package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/distributed/grpc/pb"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../../cert/ca.crt", "www.thinkgos.cn")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	arithClient := pb.NewArithClient(conn)
	response, err := arithClient.Mul(context.Background(), &pb.ArithRequest{A: 12, B: 11})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(response.Result)
}
