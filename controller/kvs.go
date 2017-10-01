package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/buslocation_api_server/kvs"
)

//GetAllRedis is hoge
func GetAllRedis(c *gin.Context) {
	m, err := kvs.GetAllKeys()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, m)
}
