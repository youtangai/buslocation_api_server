package kvs

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "函館駅前"
	id   = "3"
)

func TestSetBusStopID(t *testing.T) {
	log.Println("name is", name)
	log.Println("id is", id)
	err := SetBusStopID(name, id)
	assert.Nil(t, err)
}

func TestGetBusStopID(t *testing.T) {
	val, err := GetBusStopID(name)
	log.Println("name is", name)
	log.Println("id is", val)
	assert.Nil(t, err)
	assert.Equal(t, val, id)
}

func TestExportRedis(t *testing.T) {
	err := ExportRedis("redis.json")
	assert.Nil(t, err)
}

func TestImportRedis(t *testing.T) {
	err := ImportRedis("redis.json")
	assert.Nil(t, err)
}
