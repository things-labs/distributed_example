package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/thinkgos/distributed/rpc/method"
)

func main() {
	// rpc tcp server
	go func() {
		rpc.Register(new(method.Arith))

		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal("listen error:", err)
		}
		rpc.Accept(l)
	}()

	time.Sleep(time.Second * 2)

	// rpc tcp client
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()

	// Arith.Multiply
	args := &method.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d = %d\n", args.A, args.B, reply)

	// Arith.Divide
	args = &method.Args{15, 6}
	var quo method.Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("Divide error:", err)
	}
	fmt.Printf("Divide: %d/%d=%d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
