package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // use default Addr
		Password: "",           // no password set
		DB:       0,            // use default DB
	})

	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
}
