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

	"github.com/thinkgos/distributed/grpc/pb"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../../cert/client.crt", "../../cert/client.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair失败 %v\n", err)
	}
	certPool := x509.NewCertPool()

	ca, _ := ioutil.ReadFile("../../cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "www.thinkgos.cn",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	arithClient := pb.NewArithClient(conn)
	response, err := arithClient.Mul(context.Background(), &pb.ArithRequest{A: 12, B: 20})
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(response.Result)
}
