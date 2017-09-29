package controller

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/youtangai/buslocation_api_server/model"
	"github.com/youtangai/buslocation_api_server/sjis"
)

const (
	//SEARCHPATH is hogehoge
	SEARCHPATH = "/search02.php"
)

//GetBusStopList is hogehoge
func GetBusStopList(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")
	log.Println("start is", start)
	log.Println("end is", end)
	result, err := ScrapeBusStopList(start, end)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
}

//ScrapeBusStopList is hogehoge
func ScrapeBusStopList(start, end string) (model.List, error) {
	params, err := GetURLEncofing(start, end)
	if err != nil {
		log.Fatal(err)
	}

	var list model.List
	startMap := map[string]string{}
	endMap := map[string]string{}

	doc, err := goquery.NewDocument(WEBHOST + SEARCHPATH + "?" + params)
	if err != nil {
		log.Fatal(err)
		return list, err
	}

	// Find the review items
	doc.Find(".select").Each(func(i int, s *goquery.Selection) {
		name, ok := s.Attr("name")
		if !ok {
			return
		}
		if name == "in" {
			s.Find("option").Each(func(j int, s *goquery.Selection) {
				value, ok := s.Attr("value")
				if !ok {
					return
				}
				key, err := sjis.SjisToUTF8(s.Text())
				if err != nil {
					log.Fatal(err)
				}
				startMap[key] = value
			})
		} else if name == "out" {
			s.Find("option").Each(func(j int, s *goquery.Selection) {
				value, ok := s.Attr("value")
				if !ok {
					return
				}
				key, err := sjis.SjisToUTF8(s.Text())
				if err != nil {
					log.Fatal(err)
				}
				endMap[key] = value
			})
		}
	})

	list.StartMap = startMap
	list.EndMap = endMap
	return list, nil
}

//GetURLEncofing is hoge
func GetURLEncofing(start, end string) (string, error) {
	startSjis, err := sjis.UTF8ToSjis(start)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	endSjis, err := sjis.UTF8ToSjis(end)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	params := url.Values{}
	params.Set("stopname_f", startSjis)
	params.Set("stopname_t", endSjis)
	return params.Encode(), nil
}
