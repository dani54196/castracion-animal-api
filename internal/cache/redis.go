package cache

import (
	"context"
	"log"
	_ "time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis(addr, password string, db int) error {
	//
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	// Test the connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	log.Println("Redis connected successfully!")
	return nil
}
