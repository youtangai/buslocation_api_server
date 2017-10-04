package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youtangai/buslocation_api_server/controller"
	"github.com/youtangai/buslocation_api_server/kvs"
)

func main() {
	router := gin.Default()

	kvs.ImportRedis()
	defer kvs.ExportRedis()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/list", controller.GetBusStopList)
	router.POST("/info", controller.GetInfo)
	router.GET("/kvs/all", controller.GetAllRedis)

	router.Run(":8080")
}
