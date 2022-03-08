package must

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func getConn() *redis.Client {
	// 连接redis
	redisConn := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		//Password: "G9_I3pT_g2nGb87_v59sd", // no password set
		DB: 0, // use default DB
	})
	return redisConn
}

func PTTL(key string) {
	redisConn := getConn()
	t, _ := redisConn.PTTL(context.TODO(), key).Result()
	fmt.Println(t, int64(t)/1e9)
}
