package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/distributed/grpc/pb"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../../cert/ca.crt", "www.thinkgos.cn")
	if err != nil {
		log.Fatal(err)
	}

	// TCP
	conn1, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn1.Close()

	arithClient1 := pb.NewArithClient(conn1)
	req := &pb.ArithRequest{A: 3, B: rand.Int31n(50)}
	resp, err := arithClient1.Mul(context.Background(), req)
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Printf("TCP端口请求: A: %d, B: %d , 结果: %d\r\n", req.A, req.B, resp.Result)

	// HTTP
	conn2, err := grpc.Dial(":8082", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn2.Close()

	arithClient2 := pb.NewArithClient(conn2)
	req = &pb.ArithRequest{A: 6, B: rand.Int31n(50)}
	resp, err = arithClient2.Mul(context.Background(), req)
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Printf("HTTP端口请求: A: %d, B: %d , 结果: %d\r\n", req.A, req.B, resp.Result)
}
