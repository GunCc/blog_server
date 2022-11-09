package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/system"
	"errors"

	"gorm.io/gorm"
)

type ArticleTagService struct {
}

var ArticleTagServiceApp = new(ArticleTagService)

func (a ArticleTagService) CreateArticleTag(at system.SysArticleTag) (err error) {
	if !errors.Is(global.BLOG_DB.Where("color = ? or content = ?", at.Color, at.Content).First(&system.SysArticleTag{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同内容或者颜色的博客标签")
	}
	return global.BLOG_DB.Create(&at).Error
}

func (a ArticleTagService) EditArticleTag(at system.SysArticleTag) (err error) {
	var oldAt system.SysArticleTag
	err = global.BLOG_DB.Where("id = ?", at.TagId).First(&oldAt).Error
	if oldAt.Color != at.Color || oldAt.Content != at.Content {
		if !errors.Is(global.BLOG_DB.Where("(color = ? or content = ?) and id != ?", at.Color, at.Content, at.TagId).First(&system.SysArticleTag{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同内容或者颜色的博客标签")
		}
	}

	if err != nil {
		return err
	} else {
		return global.BLOG_DB.Updates(&at).Error
	}
}

func (a ArticleTagService) DeleteArticleTag(id uint) (err error) {
	var at system.SysArticleTag
	err = global.BLOG_DB.Where("id = ?", id).Delete(&at).Error
	if err != nil {
		return err
	}
	err = global.BLOG_DB.Delete(&[]system.SysArticleTag{}, "id = ?", id).Error
	return err
}

func (A *ArticleTagService) GetArticleTag(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.BLOG_DB.Model(&system.SysArticleTag{})
	var articleTagList []system.SysArticleTag
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleTagList).Error
	return articleTagList, total, err
}
