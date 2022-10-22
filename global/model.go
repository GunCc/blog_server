package global

import (
	"time"

	"gorm.io/gorm"
)

// 全局模型
type BLOG_MODEL struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
