package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/grpcexample/grpc/session5/services"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../../cert/client.pem", "../../cert/client.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair失败 %v\n", err)
	}
	certPool := x509.NewCertPool()

	ca, _ := ioutil.ReadFile("../../cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
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
