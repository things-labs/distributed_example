// Package main
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	reg := consul.NewRegistry(consul.Config(&api.Config{Address: "127.0.0.1:8500"}))
	server := web.NewService(
		web.Name("hello"),
		web.Handler(routers()), // 将gin的router放入
		web.Registry(reg),
		web.Address(":8080"),
	)
	server.Init()
	server.Run()
}

func routers() http.Handler {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"info": "index"})
	})
	return router
}
