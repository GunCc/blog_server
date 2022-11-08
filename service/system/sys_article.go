package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/system"
	"errors"

	"gorm.io/gorm"
)

type ArticleService struct {
}

var ArticleServiceApp = new(ArticleService)

//@function: CreateArticle
//@description: 新增文章
//@param: article  system.ArticleBlog
//@return: err error

func (a *ArticleService) CreateArticle(article system.ArticleBlog) (err error) {
	if !errors.Is(global.BLOG_DB.Where("title =?", article.Title).Preload("Tags").First(&system.ArticleBlog{}).Error, gorm.ErrRecordNotFound) {
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
	var at system.ArticleBlog
	err = global.BLOG_DB.Delete("id = ?", id).Delete(&at).Error
	if err != nil {
		return err
	}
	err = global.BLOG_DB.Delete(&[]system.ArticleBlog{}, "id = ?", id).Error
	return err
}

//@function: EditArticle
//@description: 编辑文章
//@param: article  article.ID
//@return: err error

func (a *ArticleService) EditArticle(article system.ArticleBlog) (err error) {
	var oldAt system.ArticleBlog
	err = global.BLOG_DB.Where("id =?", article.ID).First(&oldAt).Error
	if oldAt.Title != article.Title {
		if !errors.Is(global.BLOG_DB.Where("title =?", article.Title).Preload("Tags").First(&system.ArticleBlog{}).Error, gorm.ErrRecordNotFound) {
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
	db := global.BLOG_DB.Model(&system.ArticleBlog{})
	var articleList []system.ArticleBlog
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleList).Error
	return articleList, total, err
}
