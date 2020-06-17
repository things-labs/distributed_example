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
	go func() {
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal("listen error:", err)
		}

		arith := new(method.Arith)
		rpc.Register(arith)
		rpc.Accept(l)
	}()

	time.Sleep(time.Second)

	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &method.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)
}
