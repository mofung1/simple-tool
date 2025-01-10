package models

import (
	"simple-tool/server/internal/global"
)

// ParseRecord 视频解析记录表
type ParseRecord struct {
	global.Model
	UserId   int64  `gorm:"type:int unsigned;not null;index;comment:用户ID" json:"user_id"`
	Author   string `gorm:"type:varchar(255);comment:作者名称" json:"author"`
	Avatar   string `gorm:"type:varchar(250);comment:作者头像" json:"avatar"`
	Title    string `gorm:"type:varchar(500);comment:视频标题" json:"title"`
	CoverUrl string `gorm:"type:varchar(800);comment:视频封面地址" json:"cover_url" `
	MusicUrl string `gorm:"type:varchar(800);comment:音乐播放地址" json:"music_url"`
	VideoUrl string `gorm:"type:varchar(800);comment:视频播放地址" json:"video_url"`
}

// TableName 表名
func (ParseRecord) TableName() string {
	return "parse_record"
}
