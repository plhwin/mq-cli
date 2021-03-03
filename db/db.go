package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisCli *redis.Client
var Ctx = context.Background()

func InitRedis(addr, password string) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	if pong, err := RedisCli.Ping(Ctx).Result(); err != nil {
		log.Fatalln("RedisCli ping error:", err, pong)
	}
}
