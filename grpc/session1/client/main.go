package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/thinkgos/distributed/grpc/session1/services"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProductStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes.ProdStock)
}
