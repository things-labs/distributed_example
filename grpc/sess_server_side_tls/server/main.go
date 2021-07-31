package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/distributed/grpc/pb"
	"github.com/thinkgos/distributed/grpc/services"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("../../cert/server.crt", "../../cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterArithServer(rpcServer, new(services.Arith))

	// tcp server
	go func() {
		listen, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatalf("启动网络监听失败 %v\n", err)
		}
		rpcServer.Serve(listen)
	}()

	// http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rpcServer.ServeHTTP(w, r)
	})
	err = http.ListenAndServeTLS(":8082", "../../cert/server.crt", "../../cert/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
