package service

import (
	"fmt"
	"simple-tool/server/internal/frontend/response"
	"simple-tool/server/internal/global"
	"simple-tool/server/internal/models"
	"simple-tool/server/pkg/parser"
)

type ParseService struct {
}

// ParseUrl 视频地址解析
func (p *ParseService) ParseUrl(userId int64, paramUrl string) (*parser.VideoParseInfo, error) {
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

// GetParseRecordLists 获取解析记录列表
func (p *ParseService) GetParseRecordLists(userId int64, pagination *global.Pagination) (*response.ParseRecordListResponse, error) {
	var total int64
	var list []models.ParseRecord

	// 查询总数
	if err := global.DB.Model(&models.ParseRecord{}).Where("user_id = ?", userId).Count(&total).Error; err != nil {
		global.ZapLog.Error("获取解析记录总数失败: " + err.Error())
		return nil, err
	}

	// 查询列表数据
	if err := global.DB.Model(&models.ParseRecord{}).
		Where("user_id = ?", userId).
		Order("created_at desc").
		Scopes(global.Paginate(pagination)).
		Find(&list).Error; err != nil {
		global.ZapLog.Error("获取解析记录列表失败: " + err.Error())
		return nil, err
	}

	return &response.ParseRecordListResponse{
		Total:    total,
		PageNo:   pagination.PageNo,
		PageSize: pagination.PageSize,
		List:     list,
	}, nil
}
