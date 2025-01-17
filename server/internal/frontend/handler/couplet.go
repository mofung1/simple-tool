package handler

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/frontend/service"
	"simple-tool/server/internal/global/response"
)

type Couplet struct{}

// GenerateCoupletRequest 生成对联的请求参数
type GenerateCoupletRequest struct {
	// Name 用户姓名
	Name string `json:"name" binding:"required"`
	// Style 图片风格
	Style string `json:"style" binding:"required"`
}

func (cp *Couplet) GenerateCoupletImage(c *gin.Context) {
	var req GenerateCoupletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, "参数验证失败")
		return
	}

	// 调用AI服务生成对联内容
	aiService := new(service.AIService)
	coupletText, err := aiService.GenerateCouplet(c.Request.Context(), req.Name, req.Style)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	result := map[string]interface{}{
		"name":      req.Name,
		"style":     req.Style,
		"couplet":   coupletText,
		"image_url": "", // 这里后续需要填充真实的图片URL
	}

	response.OkWithData(c, result)
}
