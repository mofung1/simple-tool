package models

import (
	"simple-tool/server/internal/global"
	"time"
)

// User 用户
type User struct {
	global.Model
	Sn            int64     `gorm:"type:int(10);not null;default:0;comment:sn"`
	Nickname      string    `gorm:"type:varchar(200);not null;default:'';comment:昵称"`
	Avatar        string    `gorm:"type:varchar(200);not null;default:'';comment:头像"`
	Username      string    `gorm:"type:varchar(100);not null;default:'';comment:账号"`
	Password      string    `gorm:"type:varchar(100);not null;default:'';comment:密码" json:"-"`
	Salt          string    `gorm:"type:varchar(100);not null;default:'';comment:密码盐"`
	Phone         string    `gorm:"type:varchar(20);not null;default:'';comment:手机"`
	ClientIp      string    `gorm:"type:varchar(30);not null;default:'';comment:客户端ip" json:"client_ip"`
	ClientPort    string    `gorm:"type:varchar(10);not null;default:'';comment:客户端端口" json:"client_port"`
	DeviceInfo    string    `gorm:"type:varchar(200);not null;default:'';comment:设备信息" json:"device_info"`
	Gender        int       `gorm:"type:tinyint(1);not null;default:0;comment:0-未知 1-男 2-女"`
	IsLogin       int       `gorm:"type:tinyint(1);not null;default:0;comment:0-未登录 1-已登录" json:"is_login"`
	IsDisable     int       `gorm:"type:tinyint(1);not null;default:0;comment:0-正常 1-禁用" json:"is_disable"`
	Openid        string    `gorm:"type:varchar(200);comment:openid"`
	LoginTime     time.Time `gorm:"type:datetime(3);comment:登录时间"`
	HeartBeatTime time.Time `gorm:"type:datetime(3);comment:心跳时间"`
	LoginOutTime  time.Time `gorm:"type:datetime(3);comment:退出时间"`
}

func (User) TableName() string {
	return "user"
}
