package main

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB:   0,
    })
}

func SaveURL(code, url string) error {
    return rdb.Set(ctx, "short:"+code, url, 0).Err()
}

func GetOriginalURL(code string) (string, error) {
    return rdb.Get(ctx, "short:"+code).Result()
}
