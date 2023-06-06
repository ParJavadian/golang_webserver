package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client
var initialized = false

func startRedisConnection() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
	initialized = true
}

func GetValue(ctx context.Context, key string) string {
	if !initialized {
		startRedisConnection()
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func cacheData(ctx context.Context, key string, value string, expirationDuration time.Duration) {
	if !initialized {
		startRedisConnection()
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	if expirationDuration == -1 {
		expirationDuration = redis.KeepTTL
	}
	err := client.Set(ctx, key, value, expirationDuration).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func deleteValue(ctx context.Context, key string) {
	if !initialized {
		startRedisConnection()
	}
	if ctx == nil {
		ctx = context.TODO()
	}
	err := client.Del(ctx, key).Err()
	if err != nil {
		fmt.Println(err)
	}
}
