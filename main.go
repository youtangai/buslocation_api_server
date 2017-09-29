package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youtangai/buslocation_api_server/controller"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/list", controller.GetBusStopList)
	router.GET("/info", controller.GetInfo)

	router.Run(":8080")
}
