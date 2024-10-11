package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-gin-template/common/response"
	"go-gin-template/common/tools"
	"go-gin-template/config/log"
	"time"
)

// CustomLogger 自定义日志中间件
func CustomLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()

		// 自己实现一个 gin.ResponseWriter
		writer := response.ResponseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer

		c.Next()

		endTime := time.Now()

		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		duration := endTime.Sub(startTime)

		// 使用结构化日志记录请求信息
		log.Logger.Infof("\n ==================请求信息================== \n {\"method\": \"%s\", \"URL\": \"%s\", \"状态码\": %d, \"持续时间(ms)\": %d} \n ==================响应结果================== \n %s",
			method,
			path,
			statusCode,
			duration.Milliseconds(),
			tools.JsonFmt(writer.B.String()),
		)
	}
}
