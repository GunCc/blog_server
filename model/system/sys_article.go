package system

import "blog_server/global"

type ArticleBlog struct {
	global.BLOG_MODEL
	Title          string `json:"title" gorm:"comment:文章标题"`
	SysArticleType `json:"parent_id" gorm:"comment:一级类别"`
	Tags           []SysArticleTag `json:"tags" gorm:"comment:文章标签"`
	Content        string          `json:"content" gorm:"comment:文章内容;type:text;"`
	Enable         bool            `json:"enable" gorm:"comment:文章是否显示;default:true"`
	Sort           int             `json:"sort" gorm:"comment:排序标记;default:50"`
}
