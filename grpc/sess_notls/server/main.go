package main

import (
	"log"
	"net"

	"github.com/thinkgos/distributed/grpc/pb"
	"github.com/thinkgos/distributed/grpc/services"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	pb.RegisterArithServer(rpcServer, new(services.Arith))

	// tcp
	go func() {
		listen, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatalf("启动网络监听失败 %v\n", err)
		}
		rpcServer.Serve(listen)
	}()

	// http TODO: 不行
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	rpcServer.ServeHTTP(w, r)
	// })
	// err := http.ListenAndServe(":8082", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
