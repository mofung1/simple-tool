package routes

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/frontend/handler"
	"simple-tool/server/internal/global/response"
)

// setupFrontendRoutes 前台api
func setupFrontendRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.GET("/ping", func(c *gin.Context) {
		response.Ok(c)
	})

	parseGroup := apiGroup.Group("/parse")
	{
		parseHandler := new(handler.Parse)
		parseGroup.GET("/url", parseHandler.Url)
	}

}
