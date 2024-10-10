package request_timeout

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// RequestTimeout 创建请求超时时间中间件
func RequestTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
