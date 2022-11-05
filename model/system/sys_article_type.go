package system

import "blog_server/global"

type SysArticleType struct {
	Title string `json:"title" gorm:"comment:标题名"`
	Icon  string `json:"icon" gorm:"comment:标题图标"`
	global.BLOG_MODEL
}

func (SysArticleType) TableName() string {
	return "sys_article_types"
}
