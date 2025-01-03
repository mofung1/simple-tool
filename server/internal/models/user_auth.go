package models

import (
	"simple-tool/server/internal/global"
)

// UserAuth 用户授权信息
type UserAuth struct {
	global.Model
	UserId   int64  `gorm:"type:int unsigned;not null;index;comment:用户ID"`
	AuthType string `gorm:"type:varchar(20);not null;comment:授权类型 mnp-小程序"`
	UnionId  string `gorm:"type:varchar(100);comment:微信unionId" json:"union_id"`
	Openid   string `gorm:"type:varchar(100);comment:微信openid"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}
