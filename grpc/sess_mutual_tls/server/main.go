package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/distributed/grpc/pb"
	"github.com/thinkgos/distributed/grpc/services"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../../cert/server.crt", "../../cert/server.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair失败 %v\n", err)
	}
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../../cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequestClientCert,
		ClientCAs:    certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterArithServer(rpcServer, new(services.Arith))

	// TCP
	go func() {
		listen, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatalf("启动网络监听失败 %v\n", err)
		}
		rpcServer.Serve(listen)
	}()

	// http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rpcServer.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServeTLS(":8082", "../../cert/server.crt", "../../cert/server.key", nil))
}
