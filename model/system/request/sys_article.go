package request

type ReqArticleBlog struct {
	ID                uint   `json:"id"`
	Title             string `json:"title" gorm:"comment:文章标题"`
	SysArticleTypeIds uint   `json:"article_type"  gorm:"comment:一级标题"`
	TagsIds           []uint `json:"tags" gorm:"many2many:sys_article_to_tags"`
	Content           string `json:"content" gorm:"comment:文章内容;type:text"`
	Enable            bool   `json:"enable" gorm:"comment:文章是否显示;default:true"`
	Sort              int    `json:"sort" gorm:"comment:排序标记;default:50"`
}
