package service

import (
	"fmt"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/models"
	"simple-tool/server/pkg/parser"
)

type ParseService struct {
}

// ParseUrl 视频地址解析
func (p *ParseService) ParseUrl(paramUrl string, userId int64) (*parser.VideoParseInfo, error) {
	// 解析视频
	parseRes, err := parser.ParseVideoShareUrlByRegexp(paramUrl)
	if err != nil {
		if err.Error() == "str not have url" {
			err = fmt.Errorf("无效链接")
		}
		return nil, err
	}

	// 记录解析日志
	parseRecord := models.ParseRecord{
		UserId:   userId,
		Author:   parseRes.Author.Name,
		Avatar:   parseRes.Author.Avatar,
		Title:    parseRes.Title,
		CoverUrl: parseRes.CoverUrl,
		VideoUrl: parseRes.VideoUrl,
		MusicUrl: parseRes.MusicUrl,
	}

	// 保存到数据库
	if err := global.DB.Create(&parseRecord).Error; err != nil {
		global.ZapLog.Error("解析记录写入错误: " + err.Error())
		return nil, err
	}

	return parseRes, nil
}
