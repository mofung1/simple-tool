package models

import (
	"simple-tool/server/internal/global"
)

// Group 群
type Group struct {
	global.Model
	Sn      int64  `gorm:"type:int(10);not null;default:0;comment:sn"`
	OwnerID int64  `gorm:"not null;default:0;comment:所属人id" json:"owner_id"`
	Status  int    `gorm:"not null;default:1;comment:状态 1-正常 2-已解散 3-冻结"`
	Name    string `gorm:"type:varchar(100);not null;default:'';comment:群名称" json:"name"`
	Desc    string `gorm:"type:varchar(300);not null;default:'';comment:群简介" json:"desc"`
}

// GroupRelation 群关系
type GroupRelation struct {
	global.Model
	UserID  int64  `gorm:"not null;default:0;comment:用户id"`
	GroupID int64  `gorm:"not null;default:0;comment:群id"`
	Status  int    `gorm:"not null;default:1;comment:状态 1-群成员 2-已退群"`
	Desc    string `gorm:"type:varchar(300);not null;default:'';comment:描述"`
}

func (Group) TableName() string {
	return "group"
}

func (GroupRelation) TableName() string {
	return "group_relation"
}
