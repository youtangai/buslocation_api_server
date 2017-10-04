package controller

import (
	"log"
	"net/http"
	"net/url"

	"github.com/youtangai/buslocation_api_server/kvs"

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
	requestList := new(model.RequestList)
	c.BindJSON(requestList)
	start := requestList.Start
	end := requestList.End
	log.Println("start is", start)
	log.Println("end is", end)
	result, err := GetBusstopRedis(start, end)
	if err != nil {
		log.Fatal(err)
	}
	startMap := result.StartMap
	endMap := result.EndMap
	err = SetMapRedis(startMap)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	err = SetMapRedis(endMap)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}

//GetBusstopRedis is hoge
func GetBusstopRedis(start, end string) (model.List, error) {
	list := model.List{}
	s, err := GetBusStopMap(start)
	if err != nil {
		log.Fatal(err)
		return model.List{}, err
	}
	e, err := GetBusStopMap(end)
	if err != nil {
		log.Fatal(err)
		return model.List{}, err
	}
	list.StartMap = s
	list.EndMap = e
	return list, nil
}

//GetBusStopMap is hoge
func GetBusStopMap(str string) ([]model.BusStop, error) {
	stopSlice := []model.BusStop{}
	keys, err := kvs.GetKeys(str)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, key := range keys {
		val, err := kvs.GetBusStopID(key)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		busStop := model.BusStop{}
		busStop.ID = key
		busStop.Name = val
		stopSlice = append(stopSlice, busStop)
	}
	return stopSlice, nil
}

//ScrapeBusStopList is hogehoge
func ScrapeBusStopList(start, end string) (model.List, error) {
	params, err := GetURLEncofing(start, end)
	if err != nil {
		log.Fatal(err)
	}

	var list model.List
	startMap := []model.BusStop{}
	endMap := []model.BusStop{}

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
				key, err := sjis.ToUTF8(s.Text())
				if err != nil {
					log.Fatal(err)
				}
				busStop := model.BusStop{}
				busStop.ID = key
				busStop.Name = value
				startMap = append(startMap, busStop)
			})
		} else if name == "out" {
			s.Find("option").Each(func(j int, s *goquery.Selection) {
				value, ok := s.Attr("value")
				if !ok {
					return
				}
				key, err := sjis.ToUTF8(s.Text())
				if err != nil {
					log.Fatal(err)
				}
				busStop := model.BusStop{}
				busStop.ID = key
				busStop.Name = value
				endMap = append(endMap, busStop)
			})
		}
	})

	list.StartMap = startMap
	list.EndMap = endMap
	return list, nil
}

//GetURLEncofing is hoge
func GetURLEncofing(start, end string) (string, error) {
	startSjis, err := sjis.ToSjis(start)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	endSjis, err := sjis.ToSjis(end)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	params := url.Values{}
	params.Set("stopname_f", startSjis)
	params.Set("stopname_t", endSjis)
	return params.Encode(), nil
}

//SetMapRedis is hoge
func SetMapRedis(m []model.BusStop) error {
	for _, busStop := range m {
		// log.Println("key is", key)
		// log.Println("value is", value)
		err := kvs.SetBusStopID(busStop.ID, busStop.Name)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
