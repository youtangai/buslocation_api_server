package controller

import (
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

const (
	//SEARCHPATH is hogehoge
	SEARCHPATH = "/search.php"
)

//GetBusStopList is hogehoge
func GetBusStopList(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")
	c.JSON(http.StatusOK, start+end)
}

//ScrapeBusStopList is hogehoge
func ScrapeBusStopList(start, end string) (model.List, error) {
	var startMap map[string]string
	startMap["156"] = "ほげ"
	var endMap map[string]string
	endMap["124"] = "foo"
	var list model.List
	list.StartMap = startMap
	list.EndMap = endMap
	return list, nil
}
