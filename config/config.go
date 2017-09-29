package config

import "github.com/kelseyhightower/envconfig"

//Configuration is hoge
type Configuration struct {
	RedisHost  string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort  string `envconfig:"REDIS_PORT" default:"6379"`
	ServerPort string `default:"8080"`
}

var (
	c Configuration
)

const (
	prefix = "API"
)

func init() {
	envconfig.MustProcess(prefix, &c)
}

func reload() {
	envconfig.Process(prefix, &c)
}

//GetRedisHost is hogehoge
func GetRedisHost() string {
	return c.RedisHost
}

//GetRedisPort is hogehoge
func GetRedisPort() string {
	return c.RedisPort
}

//GetServerPort is hogehoge
func GetServerPort() string {
	return c.ServerPort
}
