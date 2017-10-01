package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/buslocation_api_server/kvs"
)

//GetAllKeyValues is hoge
func GetAllKeyValues(c *gin.Context) {
	m := map[string]string{}
	keys, err := kvs.GetAllKeys()
	log.Println("keys len is", len(keys))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	for _, key := range keys {
		value, err := kvs.GetBusStopID(key)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, err)
		}
		m[key] = value
	}
	c.JSON(http.StatusOK, m)
}
