package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/thinkgos/distributed/rpc/method"
)

// rpc http async client
func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Arith.Multiply
	args1 := &method.Args{7, 8}
	var reply int
	multiplyReply := client.Go("Arith.Multiply", args1, &reply, nil)

	// Arith.Divide
	args2 := &method.Args{15, 6}
	var quo method.Quotient
	divideReply := client.Go("Arith.Divide", args2, &quo, nil)

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	var multiplyReplied, divideReplied bool
	for !multiplyReplied || !divideReplied {
		select {
		case replyCall := <-multiplyReply.Done:
			if err := replyCall.Error; err != nil {
				fmt.Println("Multiply error:", err)
			} else {
				fmt.Printf("Multiply: %d*%d=%d\n", args1.A, args1.B, reply)
			}
			multiplyReplied = true
		case replyCall := <-divideReply.Done:
			if err := replyCall.Error; err != nil {
				fmt.Println("Divide error:", err)
			} else {
				fmt.Printf("Divide: %d/%d=%d...%d\n", args2.A, args2.B, quo.Quo, quo.Rem)
			}
			divideReplied = true
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
