package global

import (
	"gorm.io/gorm"
	"time"
)

// Model 基础模型
type Model struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;comment:ID" json:"id"`
	CreatedAt time.Time      `gorm:"type:datetime(3);comment:创建时间;index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime(3);comment:更新时间;index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(3);comment:删除时间;index" json:"deleted_at"`
}
