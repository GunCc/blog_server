package system

import "blog_server/global"

type SysArticleType struct {
	global.BLOG_MODEL
	Title string `json:"title" gorm:"comment:标题名"`
	Icon  string `json:"icon" gorm:"comment:标题图标"`
}

func (SysArticleType) TableName() string {
	return "sys_article_types"
}
