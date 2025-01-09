package models

import (
	"simple-tool/server/internal/global"
)

// ParseRecord 视频解析记录表
type ParseRecord struct {
	global.Model
	UserId   int64  `gorm:"type:int unsigned;not null;index;comment:用户ID"`
	Author   string `json:"author" gorm:"type:varchar(255);comment:作者名称"`
	Avatar   string `json:"avatar" gorm:"type:varchar(250);comment:作者头像"`
	Title    string `json:"title" gorm:"type:varchar(500);comment:视频标题"`
	CoverUrl string `json:"cover_url" gorm:"type:varchar(800);comment:视频封面地址"`
	MusicUrl string `json:"music_url" gorm:"type:varchar(800);comment:音乐播放地址"`
	VideoUrl string `json:"video_url" gorm:"type:varchar(800);comment:视频播放地址"`
}

// TableName 表名
func (ParseRecord) TableName() string {
	return "parse_record"
}
