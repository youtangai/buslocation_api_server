package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/buslocation_api_server/kvs"
)

//GetAllKeys is hoge
func GetAllKeys(c *gin.Context) {
	keys, err := kvs.GetAllKeys()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, keys)
}

//GetAllValues is hoge
func GetAllValues(c *gin.Context) {
	keys, err := kvs.GetAllKeys()
	log.Println("keys len is", len(keys))
	values := make([]string, len(keys))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	for i, key := range keys {
		value, err := kvs.GetBusStopID(key)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, err)
		}
		values[i] = value
	}
	c.JSON(http.StatusOK, values)
}
