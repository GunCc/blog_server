package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/system"
	sysReq "blog_server/model/system/request"
	"errors"

	"gorm.io/gorm"
)

type ArticleService struct {
}

var ArticleServiceApp = new(ArticleService)

//@function: CreateArticle
//@description: 新增文章
//@param: article  system.SysArticleBlog
//@return: err error

func (a *ArticleService) CreateArticle(article system.SysArticleBlog) (err error) {
	if !errors.Is(global.BLOG_DB.Where("title =?", article.Title).First(&system.SysArticleBlog{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同标题")
	}

	// global.BLOG_DB.Model(&system.SysArticleTag{})

	return global.BLOG_DB.Create(&article).Error
}

//@function: CreateDeleteArticleArticle
//@description: 删除文章
//@param: id uint
//@return: err error

func (a *ArticleService) DeleteArticle(id uint) (err error) {
	var at system.SysArticleBlog
	err = global.BLOG_DB.Delete("id = ?", id).Delete(&at).Error
	if err != nil {
		return err
	}
	err = global.BLOG_DB.Delete(&[]system.SysArticleBlog{}, "id = ?", id).Error
	return err
}

//@function: EditArticle
//@description: 编辑文章
//@param: article  article.ID
//@return: err error

func (a *ArticleService) EditArticle(article system.SysArticleBlog) (err error) {
	var oldAt system.SysArticleBlog
	err = global.BLOG_DB.Where("id =?", article.ID).First(&oldAt).Error
	if oldAt.Title != article.Title {
		if !errors.Is(global.BLOG_DB.Where("title =?", article.Title).Preload("Tags").First(&system.SysArticleBlog{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同标题")
		}
	}
	if err != nil {
		return err
	}
	return global.BLOG_DB.Updates(&article).Error
}

func (a *ArticleService) GetArticleList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.BLOG_DB.Model(&system.SysArticleBlog{})
	var articleList []system.SysArticleBlog
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleList).Error
	return articleList, total, err
}

//@function: GetArticleType
//@description: 新建文章 - 获取一级分类
//@param: info  sysReq.ReqArticleBlog
//@return: typeInfo system.SysArticleType, err error

func (a *ArticleService) GetArticleType(info sysReq.ReqArticleBlog) (typeInfo system.SysArticleType, err error) {
	if errors.Is(global.BLOG_DB.Where("id = ?", info.SysArticleTypeIds).First(&system.SysArticleType{}).Error, gorm.ErrRecordNotFound) {
		return typeInfo, errors.New("未找到指定一级类别")
	}
	err = global.BLOG_DB.Where("id = ?", info.SysArticleTypeIds).First(&typeInfo).Error
	return typeInfo, err
}

//@function: GetArticleTags
//@description: 新建文章 - 获取标签数据
//@param: info  sysReq.ReqArticleBlog
//@return:  tags []system.SysArticleTag, err error

func (a *ArticleService) GetArticleTags(info sysReq.ReqArticleBlog) (tags []system.SysArticleTag, err error) {

	for _, v := range info.TagsIds {
		tags = append(tags, system.SysArticleTag{
			TagId: v,
		})
	}

	return tags, err
}
