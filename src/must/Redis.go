package must

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"zhanyia/src/common"
)

type Redis struct {
	Pool *redis.Pool
}

// 创建redis组件实例
func init() {
	common.AllGlobal["Redis"] = &Redis{}
	common.AllGlobal["Redis"].(*Redis).Pool = &redis.Pool{
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
}
