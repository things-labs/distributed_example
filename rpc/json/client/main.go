package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"

	"github.com/thinkgos/distributed/rpc/method"
)

func main() {
	// 方法1
	// conn, err := net.Dial("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal("dial error:", err)
	// }
	//
	// // 这里，这里😁
	// client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	// 方法2,其实是方法1的封装而已
	client, err := jsonrpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer client.Close()

	// Arith.Multiply
	args := &method.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)

	// Arith.Divide
	args = &method.Args{15, 6}
	var quo method.Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("Divide error:", err)
	}
	fmt.Printf("Divide: %d/%d=%d...%d\n", args.A, args.B, quo.Quo, quo.Rem)

}
