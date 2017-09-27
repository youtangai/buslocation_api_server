package controller

import (
	"log"
	"net/http"

	"../japanese"
	"../model"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const (
	//WEBHOST is hogehoge
	WEBHOST = "http://www.hakobus.jp"
	//INFOPATH is hogehoge
	INFOPATH = "/result.php"
)

//GetInfo is hogehoge
func GetInfo(c *gin.Context) {
	startID := c.Query("start_id")
	endID := c.Query("end_id")
	result, err := ScrapeInfos(startID, endID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusFound, nil)
	}
	c.JSON(http.StatusOK, result)
}

// ScrapeInfos is hogehoge
func ScrapeInfos(startID, endID string) ([]model.Info, error) {
	var infos []model.Info
	info := model.Info{}
	doc, err := goquery.NewDocument(WEBHOST + INFOPATH + "?in=" + startID + "&out=" + endID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Find the review items
	doc.Find("table tr td").Each(func(i int, s *goquery.Selection) {
		text, err := japanese.SjisToUTF8(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		index := (i - 11) % 9
		switch index {
		case 0:
			info.Time = text
			break
		case 1:
			info.Via = text
			break
		case 2:
			info.Landing = text
			break
		case 3:
			info.Dest = text
			break
		case 6:
			info.Status = text
			key := (i - 11) / 9
			infos[key] = info
			info = model.Info{}
			break
		default:
			break
		}
	})
	return infos, nil
}