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
		chatHandler := new(handler.Chat)
		coupletHandler := new(handler.Couplet)
		hotHandler := new(handler.Hot)
		// 小程序登录
		frontend.POST("/user/login", loginHandler.MnpLogin)
		// 解析视频
		frontend.GET("/parse/url", middleware.JWTAuth(), parseHandler.Handle)
		// 解析记录
		frontend.GET("/parse/lists", middleware.JWTAuth(), parseHandler.Lists)
		// 对联测试
		frontend.POST("/ai/couplet", coupletHandler.GenerateCoupletImage)
		// AI聊天
		frontend.POST("/ai/chat", chatHandler.StreamChat)

		// 热门
		frontend.GET("/hot/lists", hotHandler.Lists)
	}

}
