package handler

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/frontend/service"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/global/response"
)

type Parse struct{}

// Handle 处理视频解析请求
func (p *Parse) Handle(c *gin.Context) {
	// 获取URL参数
	paramUrl := c.Query("url")
	if paramUrl == "" {
		response.FailWithMsg(c, "URL不能为空")
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		response.FailWithMsg(c, "登录信息异常")
		return
	}

	parseService := new(service.ParseService)
	result, err := parseService.ParseUrl(userId.(int64), paramUrl)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	// 返回解析结果
	response.OkWithData(c, result)
	return
}

// Lists 获取解析记录列表
func (p *Parse) Lists(c *gin.Context) {
	var pagination global.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.FailWithMsg(c, "参数错误")
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		response.FailWithMsg(c, "登录信息异常")
		return
	}

	parseService := new(service.ParseService)
	result, err := parseService.GetParseRecordLists(userId.(int64), &pagination)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	// 返回列表数据
	response.OkWithData(c, result)
	return
}
