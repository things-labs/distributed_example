package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/thinkgos/distributed/rpc/method"
)

func main() {
	rpc.Register(new(method.Arith))

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		// 注意这一行
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
