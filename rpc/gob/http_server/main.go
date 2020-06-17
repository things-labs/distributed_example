package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/thinkgos/distributed/rpc/method"
)

func main() {
	rpc.Register(new(method.Arith))
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}
