package models

import (
	"simple-tool/server/internal/global"
	"time"
)

// UserAuth 用户授权信息
type UserAuth struct {
	global.Model
	UserId       int64     `gorm:"type:int unsigned;not null;index;comment:用户ID"`
	AuthType     string    `gorm:"type:varchar(20);not null;comment:授权类型 miniapp-小程序"`
	AuthId       string    `gorm:"type:varchar(100);not null;unique;comment:第三方ID"`
	UnionId      string    `gorm:"type:varchar(100);comment:微信开放平台ID"`
	AccessToken  string    `gorm:"type:varchar(200);comment:访问令牌"`
	RefreshToken string    `gorm:"type:varchar(200);comment:刷新令牌"`
	ExpiresIn    int       `gorm:"type:int;comment:过期时间"`
	Scope        string    `gorm:"type:varchar(100);comment:授权范围"`
	ExpireTime   time.Time `gorm:"type:datetime;comment:令牌过期时间"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}
