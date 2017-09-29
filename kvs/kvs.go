package kvs

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/youtangai/buslocation_api_server/config"
)

var (
	c *redis.Client
)

func init() {
	err := Connection()
	if err != nil {
		log.Fatal(err)
	}
}

//Connection is hogehoge
func Connection() error {
	host := config.GetRedisHost()
	port := config.GetRedisPort()

	c = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	return err
}

//SetBusStopID is hoges
func SetBusStopID(name, id string) error {
	err := c.Set(name, id, 0).Err()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//GetBusStopID is hoge
func GetBusStopID(name string) (string, error) {
	id, err := c.Get(name).Result()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return id, nil
}
