package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	reg := consul.NewRegistry(consul.Config(&api.Config{Address: "127.0.0.1:8500"}))

	for {
		srvs, err := reg.GetService("hello")
		if err != nil {
			log.Fatal(err)
		}
		next := selector.Random(srvs)
		rsp, err := next()
		log.Println(rsp)
		time.Sleep(time.Second)
	}
}

func routers() http.Handler {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"info": "index"})
	})
	return router
}
