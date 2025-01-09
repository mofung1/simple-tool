package routes

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/frontend/handler"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/internal/middleware"
)

// setupFrontendRoutes 前台路由
func setupFrontendRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.GET("/ping", func(c *gin.Context) {
		response.Ok(c)
	})

	frontend := apiGroup.Group("/v1")
	{
		parseHandler := new(handler.Parse)
		loginHandler := new(handler.Login)
		// 解析视频
		frontend.GET("/parse/url", middleware.JWTAuth(), parseHandler.Url)
		// 小程序登录
		frontend.POST("/user/login", loginHandler.MnpLogin)
	}

}
