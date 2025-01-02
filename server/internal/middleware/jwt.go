package middleware

import (
	"simple-tool/server/internal/global/cache"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/pkg/jwt"
	"simple-tool/server/pkg/redis"

	"github.com/gin-gonic/gin"
)

// JWTAuth 基于JWT的认证中间件
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMsg(c, "请登录")
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(token)
		if err != nil {
			response.FailWithMsg(c, "授权已到期，请重新登录"+err.Error())
			c.Abort()
			return
		}

		// 校验当前token是否在黑名单中
		key := cache.KeyLoginBlack + token
		if redis.GetString(key) != "" {
			response.FailWithMsg(c, "已被禁用")
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("userId", mc.UserID)
		c.Next()
	}
}
