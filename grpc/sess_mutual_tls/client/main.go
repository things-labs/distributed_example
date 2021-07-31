package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

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

	// TCP
	conn1, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn1.Close()

	arithClient1 := pb.NewArithClient(conn1)
	req := &pb.ArithRequest{A: 12, B: rand.Int31n(20)}
	resp, err := arithClient1.Mul(context.Background(), req)
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Printf("TCP端口请求: A: %d, B: %d , 结果: %d\r\n", req.A, req.B, resp.Result)

	// HTTP
	conn2, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn2.Close()

	arithClient2 := pb.NewArithClient(conn2)
	req = &pb.ArithRequest{A: 8, B: rand.Int31n(20)}
	resp, err = arithClient2.Mul(context.Background(), req)
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Printf("HTTP端口请求: A: %d, B: %d , 结果: %d\r\n", req.A, req.B, resp.Result)

}
