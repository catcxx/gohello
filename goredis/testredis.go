package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Hello, World!")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.15.1.48:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}