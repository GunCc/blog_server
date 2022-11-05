package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/commen/response"
	"blog_server/model/system"
	"blog_server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleTagApi struct {
}

// @Tags SysArticleTag
// @Summary 新增博客标签
// @Produce application/json
// @Param data body system.SysArticleTag{msg=string} "新建标签"
// @Router /article/tag/add [post]

func (A *ArticleTagApi) CreateArticleTag(c *gin.Context) {
	var at system.SysArticleTag
	_ = c.ShouldBindJSON(at)
	if err := utils.Verify(at, utils.ArticleTagVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ArticleTagService.CreateArticleTag(at); err != nil {
		global.BLOG_LOG.Error("创建失败", zap.Error(err))
		response.FailWithDetailed(err, "创建失败", c)
		return
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysArticleTag
// @Summary  修改博客标签
// @Produce application/json
// @Param data body system.SysArticleTag{msg=string} "修改标签"
// @Router /article/tag/edit [PUT]

func (A *ArticleTagApi) EditArticleTag(c *gin.Context) {
	var at system.SysArticleTag
	_ = c.ShouldBindJSON(at)
	if err := utils.Verify(at, utils.ArticleTagVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ArticleTagService.EditArticleTag(at); err != nil {
		global.BLOG_LOG.Error("修改失败", zap.Error(err))
		response.FailWithDetailed(err, "修改失败", c)
		return
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysArticleTag
// @Summary  删除博客标签
// @Produce application/json
// @Param data body system.SysArticleTag{msg=string} "删除标签"
// @Router /article/tag/delete [delete]

func (A *ArticleTagApi) DeleteArticleTag(c *gin.Context) {
	var at system.SysArticleTag
	_ = c.ShouldBindJSON(at)
	if err := utils.Verify(at.ID, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ArticleTagService.DeleteArticleTag(at.ID); err != nil {
		global.BLOG_LOG.Error("删除失败", zap.Error(err))
		response.FailWithDetailed(err, "删除失败", c)
		return
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (A *ArticleTagApi) GetArticleTag(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := ArticleTagService.GetArticleTag(pageInfo); err != nil {
		response.FailWithDetailed(err, "查询失败", c)
		global.BLOG_LOG.Error(err.Error(), zap.Error(err))
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "查询成功", c)
	}
}
