package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/system"
	"errors"

	"gorm.io/gorm"
)

type ArticleTypeService struct {
}

var ArticleTypeServiceApp = new(ArticleTypeService)

//@function: CreateArticleType
//@description: 新增文章标题api
//@param: api  system.SysArticleType
//@return: err error

func (articleTypeService *ArticleTypeService) CreateArticleType(at system.SysArticleType) (err error) {
	if !errors.Is(global.BLOG_DB.Where("title = ?", at.Title).First(&system.SysArticleType{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同标题")
	}
	return global.BLOG_DB.Create(&at).Error
}

//@function: DeleteArticleType
//@description: 删除文章标题api
//@param: api  system.SysArticleType
//@return: err error

func (articleTypeService *ArticleTypeService) DeleteArticleType(id uint) (err error) {
	var at system.SysArticleType
	err = global.BLOG_DB.Where("id = ?", id).Delete(&at).Error
	if err != nil {
		return err
	}
	err = global.BLOG_DB.Delete(&[]system.SysArticleType{}, "id = ?", id).Error
	return err
}

//@function: EditArticleType
//@description: 修改文章标题api
//@param: api  system.SysArticleType
//@return: err error

func (ArticleTypeService *ArticleTypeService) EditArticleType(at system.SysArticleType) (err error) {
	var oldAt system.SysArticleType
	err = global.BLOG_DB.Where("id = ?", at.ID).First(&oldAt).Error
	if oldAt.Title != at.Title {
		if !errors.Is(global.BLOG_DB.Where("title = ?", at.Title).First(&system.SysArticleType{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的标题名")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.BLOG_DB.Updates(&at).Error
	}
	return err
}

//@function: getArticleType
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int6
func (ArticleTypeService *ArticleTypeService) GetArticleType(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.BLOG_DB.Model(&system.SysArticleType{})
	var articleList []system.SysArticleType
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleList).Error
	return articleList, total, err
}
