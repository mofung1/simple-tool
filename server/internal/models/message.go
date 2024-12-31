package models

import (
	"simple-tool/server/internal/global"
)

// Message 定义消息结构体
type Message struct {
	global.Model
	FormID   int64  `json:"form_id" gorm:"not null;default:0;comment:发送人"`
	TargetID int64  `json:"target_id" gorm:"not null;default:0;comment:接收人"`
	Type     int8   `json:"type" gorm:"not null;default:1;comment:聊天类型：1-群聊 2-私聊  3-广播"`
	Media    int8   `json:"media" gorm:"not null;default:1;comment:信息类型：1-文字 2-图片 3-音频"`
	Content  string `json:"content" gorm:"not null;type:text;comment:消息内容"`
	Image    string `json:"image" gorm:"not null;size:200;default:'';comment:图片"`
	File     string `json:"file" gorm:"not null;size:200;default:'';comment:文件"`
	Desc     string `json:"desc" gorm:"not null;size:255;default:'';comment:描述"`
}

// TableName 指定数据库表名
func (Message) TableName() string {
	return "message"
}
