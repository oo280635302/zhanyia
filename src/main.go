package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	fmt.Println("Hello")
	c := redis.Pool{
		MaxIdle:     16,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	conn := c.Get()
	defer conn.Close()
	hashSlice, err := redis.Strings(conn.Do("keys", "*"))
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range hashSlice {
		fmt.Println(k, v)
	}
}
