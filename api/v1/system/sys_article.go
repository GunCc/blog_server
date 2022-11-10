package system

import (
	"blog_server/global"
	"blog_server/model/commen/request"
	"blog_server/model/commen/response"
	"blog_server/model/system"
	sysReq "blog_server/model/system/request"
	"blog_server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleApi struct {
}

// @Tags Article
// @Summary 新增文章
// @Produce application/json
// @Param data body system.SysArticleBlog true "标题，类别，标签，内容，标记"
// @Success 200 {object} response.Response{msg=string} "新建标题"
// @Router /article/add  [post]
func (A *ArticleApi) CreateArticle(c *gin.Context) {
	var reqData sysReq.ReqArticleBlog
	_ = c.ShouldBindJSON(&reqData)

	if err := utils.Verify(reqData, utils.ArticleVerify); err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}
	var a = system.SysArticleBlog{
		Title:   reqData.Title,
		Content: reqData.Content,
	}
	if typeInfo, err := ArticleService.GetArticleType(reqData); err != nil {
		global.BLOG_LOG.Error("操作失败", zap.Error(err))
		response.FailWithDetailed(err, "操作失败", c)
		return
	} else {
		a.SysArticleTypeIds = typeInfo
	}
	if tags, err := ArticleService.GetArticleTags(reqData); err != nil {
		global.BLOG_LOG.Error("操作失败", zap.Error(err))
		response.FailWithDetailed(err, "操作失败", c)
		return
	} else {
		a.TagsIds = tags
	}
	if err := ArticleService.CreateArticle(a); err != nil {
		global.BLOG_LOG.Error("操作失败", zap.Error(err))
		response.FailWithDetailed(err, "操作失败", c)
		return
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// @Tags Article
// @Summary 删除文章
// @Produce application/json
// @Param data body system.SysArticleBlog.ID true "文章ID"
// @Success 200 {object} response.Response{msg=string} "删除文章"
// @Router /article/delete [delete]
func (A *ArticleApi) DeleteArticle(c *gin.Context) {
	var a system.SysArticleBlog
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.IdVerify); err != nil {
		response.FailWithDetailed(err, "操作失败", c)
		return
	}
	fmt.Println(a)
	if err := ArticleService.DeleteArticle(a.ID); err != nil {
		response.FailWithDetailed(err, "操作失败", c)
		return
	}
}

// @Tags Article
// @Summary 修改文章
// @Produce application/json
// @Param data body system.SysArticleBlog true "标题，类别，标签，内容，标记"
// @Success 200 {object} response.Response{msg=string} "修改文章"
// @Router /article/edit [put]
func (A *ArticleApi) EditArticle(c *gin.Context) {
	var a sysReq.ReqArticleBlog
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.IdVerify); err != nil {
		response.FailWithDetailed(err, "无ID标识符", c)
		return
	}
	if err := utils.Verify(a, utils.ArticleVerify); err != nil {
		response.FailWithMessage("操作失败", c)
		return
	}
	if err := ArticleService.EditArticle(a); err != nil {
		global.BLOG_LOG.Error("操作失败", zap.Error(err))
		response.FailWithDetailed(err, "操作失败", c)
		return
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// @Tags Article
// @Summary 获取博客列表
// @Produce application/json
// @Param data body request.PageInfo true "筛选条件"
// @Success 200 {object} response.Response{msg=string} "获取表"
// @Router /article/list   [get]
func (A *ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)

	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := ArticleService.GetArticleList(pageInfo); err != nil {
		global.BLOG_LOG.Error("操作失败", zap.Error(err))
		response.FailWithDetailed(err, "操作失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "操作成功", c)
	}
}
