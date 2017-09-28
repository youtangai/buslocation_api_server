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
	router.GET("/list", controller.GetBusStopList)
	router.GET("/info", controller.GetInfo)

	router.Run(":8080")
}
