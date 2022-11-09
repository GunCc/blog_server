package system

import "blog_server/global"

type SysArticleBlog struct {
	global.BLOG_MODEL
	Title             string          `json:"title" gorm:"comment:文章标题"`
	SysArticleTypeIds SysArticleType  `json:"article_type" gorm:"comment:一级标题;foreignKey:ID;"`
	TagsIds           []SysArticleTag `json:"tags" gorm:"many2many:sys_article_to_tags;"`
	Content           string          `json:"content" gorm:"comment:文章内容;type:text"`
	Enable            bool            `json:"enable" gorm:"comment:文章是否显示;default:true"`
	Sort              int             `json:"sort" gorm:"comment:排序标记;default:50"`
}

func (SysArticleBlog) TableName() string {
	return "sys_articles"
}
