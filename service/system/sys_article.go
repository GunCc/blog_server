package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/system"
	sysReq "blog_server/model/system/request"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err = global.BLOG_DB.Where("id = ?", id).First(&at).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = global.BLOG_DB.Select(clause.Associations).Delete(&at).Error
	return err
}

//@function: EditArticle
//@description: 编辑文章
//@param: article  article.ID
//@return: err error

func (a *ArticleService) EditArticle(info sysReq.ReqArticleBlog) (err error) {
	var oldAt system.SysArticleBlog
	var newAt system.SysArticleBlog
	newAt.Title = info.Title
	newAt.Content = info.Content
	newAt.ID = info.ID
	err = global.BLOG_DB.Where("id =?", info.ID).Preload("TagsIds").First(&oldAt).Error
	if oldAt.Title != info.Title {
		if !errors.Is(global.BLOG_DB.Where("title =? and id != ?", newAt.Title, newAt.ID).First(&system.SysArticleBlog{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同标题")
		}
	}
	if oldAt.SysArticleTypeIds.ID != info.SysArticleTypeIds {
		if typeInfo, err := a.GetArticleType(info); err != nil {
			return err
		} else {
			newAt.SysArticleTypeIds = typeInfo
		}
	}

	if tags, err := a.GetArticleTags(info); err != nil {
		fmt.Println("报错啦！")
		return err
	} else {
		newAt.TagsIds = tags
	}
	if err != nil {
		return err
	}
	fmt.Println("最后一步：", newAt)
	return global.BLOG_DB.Updates(&newAt).Error
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
		var tag system.SysArticleTag
		fmt.Println("v:", v)
		err = global.BLOG_DB.Where("tag_id = ?", v).First(&tag).Error
		fmt.Println("tag:", tag)
		if err != nil {
			break
		} else if tag.Color == "" || tag.Content == "" {
			err = errors.New("标签不存在")
			break
		} else {
			tags = append(tags, system.SysArticleTag{
				TagId: v,
			})
		}
	}
	// err = errors.New("报错")
	return tags, err
}
