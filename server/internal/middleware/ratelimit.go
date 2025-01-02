package middleware

import (
	"simple-tool/server/internal/global/response"
	"simple-tool/server/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimit 登录限流中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		// 获取当前计数
		count := redis.GetInt(key)
		if count >= 60 {
			response.FailWithMsg(c, "请求过于频繁，请稍后再试")
			c.Abort()
			return
		}

		redis.Set(key, count+1, int64(60*time.Second))

		c.Next()
	}
}
