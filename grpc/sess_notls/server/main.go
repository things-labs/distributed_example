package main

import (
	"log"
	"net"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/thinkgos/distributed/grpc/pb"
	"github.com/thinkgos/distributed/grpc/services"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterArithServer(grpcServer, new(services.Arith))

	// tcp server
	go func() {
		log.Println("监听TCP: 8081")
		listen, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatalf("启动网络监听失败 %v\n", err)
		}
		grpcServer.Serve(listen)
	}()

	// http
	log.Println("监听HTTP: 8082")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		grpcServer.ServeHTTP(w, r)
	})
	err := http.ListenAndServe(":8082", h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	}), &http2.Server{}))
	if err != nil {
		log.Fatal(err)
	}
}
