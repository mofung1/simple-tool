package handler

import (
	"github.com/gin-gonic/gin"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/global/response"
	"simple-tool/server/internal/models"
	"simple-tool/server/pkg/parser"
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

	// 解析视频
	parseRes, err := parser.ParseVideoShareUrlByRegexp(paramUrl)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		response.FailWithMsg(c, "登录信息异常")
		return
	}

	// 记录解析日志
	parseRecord := &models.ParseRecord{
		UserId:   userId.(int64),
		Author:   parseRes.Author.Name,
		Avatar:   parseRes.Author.Avatar,
		Title:    parseRes.Title,
		CoverUrl: parseRes.CoverUrl,
		VideoURL: parseRes.VideoUrl,
		MusicUrl: parseRes.MusicUrl,
	}

	// 保存到数据库
	if err := global.DB.Create(parseRecord).Error; err != nil {
		response.FailWithMsg(c, "记录异常"+err.Error())
		return
	}

	// 返回解析结果
	response.OkWithData(c, parseRes)
	return
}
