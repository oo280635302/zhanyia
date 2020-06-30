package must

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinListener() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word\n123")
	})

	//监听端口默认为8080
	r.Run(":8000")
}
