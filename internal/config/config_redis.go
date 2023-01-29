package config

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func (c *Config) InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.RedisCofig.Addr,
		Password: c.RedisCofig.Password,
		DB:       c.RedisCofig.Db,
	})

	if redisClient == nil {
		panic("failed to nit redis client")
	}
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func Redis() *redis.Client {
	return redisClient
}
