// 请求日志中间件

package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 记录 HTTP 请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		// 记录日志
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method

		log.Printf("[%d] %s %s?%s | %v | %s",
			statusCode,
			method,
			path,
			query,
			latency,
			c.ClientIP(),
		)
	}
}
