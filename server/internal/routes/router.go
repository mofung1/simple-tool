package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/internal/middleware"
)

// SetupRouter 注册路由
func SetupRouter() *gin.Engine {
	if global.Conf.App.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// 设置全局中间件
	registerGlobalMiddleWare(r)
	// 注册前台分组路由
	setupFrontendRoutes(r)
	// 处理404请求
	setupNotFoundHandler(r)
	return r
}

// registerGlobalMiddleWare 注册全局中间件
func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Cors(),
		middleware.Translations("zh"),
	)
}

// setupNotFoundHandler 处理未定义请求
func setupNotFoundHandler(r *gin.Engine) {
	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		method := c.Request.Method
		switch method {
		case http.MethodGet:
			response.FailWithMsg(c, "GET请求的路径不存在")
		case http.MethodPost:
			response.FailWithMsg(c, "POST请求的路径不存在")
		default:
			response.FailWithMsg(c, "请求的路径不存在")
		}
	})
}
