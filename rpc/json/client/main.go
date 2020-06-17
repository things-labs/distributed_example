package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"

	"github.com/thinkgos/distributed/rpc/method"
)

func main() {
	// conn, err := net.Dial("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal("dial error:", err)
	// }
	//
	// // 这里，这里😁
	// client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	client, err := jsonrpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	args := &method.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)
}
