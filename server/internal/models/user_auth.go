package models

import (
	"simple-tool/server/internal/global"
)

// UserAuth 用户授权信息
type UserAuth struct {
	global.Model
	UserId   int64  `gorm:"type:int unsigned;not null;index;comment:用户ID" json:"user_id"`
	AuthType string `gorm:"type:varchar(20);not null;comment:授权类型 mnp-小程序" json:"auth_type"`
	UnionId  string `gorm:"type:varchar(100);comment:微信unionId" json:"union_id"`
	Openid   string `gorm:"type:varchar(100);comment:微信openid" json:"openid"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}
