package middleware

import (
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/global/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		global.ZapLog.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.ZapLog.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("validator", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					global.ZapLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("validator", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.ZapLog.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("validator", string(httpRequest)),
					)
				}

				response.FailWithInfo(c, "系统错误", err)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
