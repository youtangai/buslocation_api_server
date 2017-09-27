package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/list", controller.GetBusStopList)
	router.POST("/info", controller.GetInfo)

	router.Run(":8080")
}
