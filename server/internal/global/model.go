package global

import (
	"gorm.io/gorm"
	"time"
)

// Model 基础模型
type Model struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;comment:ID"`
	CreatedAt time.Time      `gorm:"type:datetime(3);comment:创建时间;index"`
	UpdatedAt time.Time      `gorm:"type:datetime(3);comment:更新时间;index"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(3);comment:删除时间;index"`
}
