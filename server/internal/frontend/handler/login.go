package handler

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/frontend/request"
	"simple-tool/server/internal/frontend/service"
	"simple-tool/server/internal/global/response"
)

type Login struct {
}

// MnpLogin 小程序登录
func (l *Login) MnpLogin(c *gin.Context) {
	var req request.MnpLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, "参数错误")
		return
	}

	loginService := new(service.LoginService)
	result, err := loginService.MnpLogin(req.Code, c.ClientIP())
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, result)
}
