package configs

import (
	"github.com/go-redis/redis/v8"
)

var cacheConn *redis.Client

func InitCache() {
	cacheConn = redis.NewClient(
		&redis.Options{
			Addr:     "redis-15523.c60.us-west-1-2.ec2.cloud.redislabs.com:15523",
			Password: "v4aVmDaPc05SSwoBiBchRJfsku65t0gu",
			DB:       0,
		},
	)
}
