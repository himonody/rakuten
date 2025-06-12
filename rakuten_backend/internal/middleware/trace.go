package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c.Set("trace_id", traceID)
		c.Writer.Header().Set("X-Trace-ID", traceID) // 加入响应头
		c.Next()
	}
}
