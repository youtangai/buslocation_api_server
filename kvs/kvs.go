package kvs

import (
	"log"

	"../config"
	"github.com/go-redis/redis"
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
