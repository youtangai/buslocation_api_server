package kvs

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"encoding/json"

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

//GetKeys is hoge
func GetKeys(keyword string) ([]string, error) {
	result := []string{}
	keys, err := GetAllKeys()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, key := range keys {
		if strings.Contains(key, keyword) {
			result = append(result, key)
		}
	}
	return result, nil
}

//GetAllKeys is hoge
func GetAllKeys() ([]string, error) {
	keys, err := c.Keys("*").Result()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return keys, nil
}

//GetAllKeyValues is hoge
func GetAllKeyValues() (map[string]string, error) {
	m := map[string]string{}
	keys, err := GetAllKeys()
	log.Println("keys len is", len(keys))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, key := range keys {
		value, err := GetBusStopID(key)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		m[key] = value
	}
	return m, nil
}

//ExportRedis is hoge
func ExportRedis() error {
	path := config.GetRedisPath()
	m, err := GetAllKeyValues()
	if err != nil {
		log.Fatal(err)
		return err
	}
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	ioutil.WriteFile(path, data, os.ModePerm)
	return nil
}

//ImportRedis is hoge
func ImportRedis() error {
	path := config.GetRedisPath()
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	m := map[string]string{}
	json.Unmarshal(file, &m)

	for key, val := range m {
		err := SetBusStopID(key, val)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
