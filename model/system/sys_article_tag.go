package system

import "time"

type SysArticleTag struct {
	// global.
	CreatedAt time.Time        // 创建时间
	UpdatedAt time.Time        // 更新时间
	DeletedAt *time.Time       `sql:"index"`
	TagId     uint             `json:"tagId" gorm:"not null;unique;primary_key;comment:角色ID;"`
	Articles  []SysArticleBlog `json:"-" gorm:"many2many:sys_article_to_tags;"`
	Color     string           `json:"color" gorm:"comment:标签颜色"`
	Content   string           `json:"content" gorm:"comment:标签内容"`
}

func (SysArticleTag) TableName() string {
	return "sys_article_tags"
}
