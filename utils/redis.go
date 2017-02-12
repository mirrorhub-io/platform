package utils

import (
	redis "gopkg.in/redis.v5"
	"os"
)

var redisconn *redis.Client

func Redis() *redis.Client {
	addr := "redis:6379"
	if os.Getenv("REDIS_ADDR") != "" {
		addr = os.Getenv("REDIS_ADDR")
	}
	if redisconn == nil {
		return redis.NewClient(
			&redis.Options{
				Addr:     addr,
				Password: "",
				DB:       0,
				PoolSize: 100,
			},
		)
	}
	return redisconn
}
