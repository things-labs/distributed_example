package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/thinkgos/distributed/grpc/pb"
)

func main() {
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
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
