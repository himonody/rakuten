package xhttp

import (
	"github.com/gin-gonic/gin"
	"rakuten_backend/internal/context"
	"time"
)

type Handler func(c *context.Context)

func convert(relativePath string, h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now().UnixNano() / 1e6
		c1 := &context.Context{Context: c, Trace: c.GetHeader("X-Request-Id")}
		h(c1)
		costTime := time.Now().UnixNano()/1e6 - startTime
		if costTime > 500 {
			c1.Infof("%v, use time=%v", relativePath, costTime)
		}
	}
}

func POST(group *gin.RouterGroup, relativePath string, handler Handler) {
	group.POST(relativePath, convert(relativePath, handler))
}

func GET(group *gin.RouterGroup, relativePath string, handler Handler) {
	group.GET(relativePath, convert(relativePath, handler))
}

func PUT(group *gin.RouterGroup, relativePath string, handler Handler) {
	group.PUT(relativePath, convert(relativePath, handler))
}
func PATCH(group *gin.RouterGroup, relativePath string, handler Handler) {
	group.PATCH(relativePath, convert(relativePath, handler))
}

func DELETE(group *gin.RouterGroup, relativePath string, handler Handler) {
	group.DELETE(relativePath, convert(relativePath, handler))
}
