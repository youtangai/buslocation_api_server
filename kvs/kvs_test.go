package kvs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "函館駅前"
	id   = "3"
)

func TestSetBusStopID(t *testing.T) {
	err := SetBusStopID(name, id)
	assert.Nil(t, err)
}

func TestGetBusStopID(t *testing.T) {
	val, err := GetBusStopID(name)
	assert.Nil(t, err)
	assert.Equal(t, val, id)
}
