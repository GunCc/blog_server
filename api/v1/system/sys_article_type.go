package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/commen/response"
	"blog_server/model/system"
	"blog_server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleTypeApi struct {
}

// @Tags ArticleType
// @Summary 新增博客一级标题
// @Produce application/json
// @Param data body system.SysArticleType true "标题，图标"
// @Success 200 {object} response.Response{msg=string} "新建标题"
// @Router /article/type/add  [post]
func (A *ArticleTypeApi) CreateArticleType(c *gin.Context) {
	var at system.SysArticleType
	_ = c.ShouldBindJSON(&at)
	fmt.Println("--------------------", at.Title)
	if err := utils.Verify(at, utils.ArticleTypeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ArticleTypeService.CreateArticleType(at); err != nil {
		global.BLOG_LOG.Error("标题新增失败", zap.Error(err))
		response.FailWithDetailed(err, "标题新增失败", c)
		return
	} else {
		response.OkWithMessage("新增成功", c)
	}
}

// @Tags ArticleType
// @Summary 删除博客一级标题
// @Produce application/json
// @Param data body system.SysArticleType true "标题，图标"
// @Success 200 {object} response.Response{msg=string} "删除标题"
// @Router /article/type/delete  [delete]
func (A *ArticleTypeApi) DeleteArticleType(c *gin.Context) {
	var at system.SysArticleType
	_ = c.ShouldBindJSON(&at)
	if err := utils.Verify(at.ID, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ArticleTypeService.DeleteArticleType(at.ID); err != nil {
		global.BLOG_LOG.Error("标题删除失败", zap.Error(err))
		response.FailWithDetailed(err, "标题删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ArticleType
// @Summary 修改博客一级标题
// @Produce application/json
// @Param data body system.SysArticleType true "标题，图标"
// @Success 200 {object} response.Response{msg=string} "修改标题"
// @Router /article/type/edit [put]
func (A *ArticleTypeApi) EditArticleType(c *gin.Context) {
	var at system.SysArticleType
	_ = c.ShouldBindJSON(&at)
	if err := utils.Verify(at, utils.ArticleTypeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ArticleTypeService.EditArticleType(at); err != nil {
		global.BLOG_LOG.Error("标题修改失败", zap.Error(err))
		response.FailWithDetailed(err, "标题修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags ArticleType
// @Summary 获取博客一级标题
// @Produce application/json
// @Param data body system.SysArticleType true "标题，图标"
// @Success 200 {object} response.Response{msg=string} "修改标题"
// @Router /article/type/list   [get]
func (A *ArticleTypeApi) GetArticleType(c *gin.Context) {
	var pageInfo request.PageInfo

	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := ArticleTypeService.GetArticleType(pageInfo); err != nil {
		global.BLOG_LOG.Error("标题获取失败", zap.Error(err))
		response.FailWithDetailed(err, "标题获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
