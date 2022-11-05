package system

import "blog_server/global"

type SysArticleTag struct {
	global.BLOG_MODEL
	Color   string `json:"color" gorm:"comment:标签颜色"`
	Content string `json:"content" gorm:"comment:标签内容"`
}

func (SysArticleTag) TableName() string {
	return "article_tags"
}
