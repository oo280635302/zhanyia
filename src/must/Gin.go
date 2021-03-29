package must

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func GinListener(l *Limit) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		if v, ok := l.CountCup.Load("1"); ok {
			v1 := v.(*int64)
			fmt.Println(*v1)
		}
		if !l.IsAble("1") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": "访问已到上限"})
			return
		}
		l.Inc("1")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word\n123")
	})
	r.GET("/123", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word\n123")
	})

	//监听端口默认为8080
	r.Run(":8080")
}

type Limit struct {
	Interval time.Duration
	MaxCount int64
	Lock     sync.Mutex
	CountCup *sync.Map
}

func NewLimitTicker(interval time.Duration, maxCnt int64) *Limit {
	reqLimit := &Limit{
		Interval: interval,
		MaxCount: maxCnt,
		CountCup: new(sync.Map),
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			reqLimit.Lock.Lock()
			reqLimit.CountCup = new(sync.Map)
			reqLimit.Lock.Unlock()
		}
	}()

	return reqLimit
}

func (r *Limit) Inc(key string) {
	p := new(int64)
	*p = 1
	v, ok := r.CountCup.LoadOrStore(key, p)
	if !ok {
		return
	}
	p = v.(*int64)
	v1 := atomic.LoadInt64(p)
	atomic.CompareAndSwapInt64(p, v1, v1+1)
}

func (r *Limit) IsAble(key string) bool {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	v1 := int64(0)
	v, ok := r.CountCup.Load(key)
	if ok {
		v1 = atomic.LoadInt64(v.(*int64))
	}
	return v1 < r.MaxCount
}
