package must

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

func GinListener(l *Limit) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println(l.ReqCount)
		if !l.IsAble() {

			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": "访问已到上限"})
			return
		}
		l.Inc()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word\n123")
		fmt.Println("url", c.Request.Host)
	})
	r.GET("/123", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word\n123")
		fmt.Println("url", c.Request.Host)
	})

	//监听端口默认为8080
	r.Run(":8000")
}

type Limit struct {
	Interval time.Duration
	MaxCount int
	Lock     sync.Mutex
	ReqCount int
}

func NewLimitTicker(interval time.Duration, maxCnt int) *Limit {
	reqLimit := &Limit{
		Interval: interval,
		MaxCount: maxCnt,
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			reqLimit.Lock.Lock()
			reqLimit.ReqCount = 0
			reqLimit.Lock.Unlock()
		}
	}()

	return reqLimit
}

func (r *Limit) Inc() {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	r.ReqCount += 1
}

func (r *Limit) IsAble() bool {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	return r.ReqCount < r.MaxCount
}
