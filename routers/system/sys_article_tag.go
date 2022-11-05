package system

import (
	v1 "blog_server/api/v1"
	"blog_server/middlewares"

	"github.com/gin-gonic/gin"
)

type ArticleTagRouter struct {
}

func (a *ArticleTypeRouter) InitArticleTagRouter(Router *gin.RouterGroup) {
	ATRouter := Router.Group("article/tag").Use(middlewares.OperationRecord())
	ArticleTagApi := v1.ApiGroupApp.SysApiGroup.ArticleTagApi
	{
		ATRouter.POST("/add", ArticleTagApi.CreateArticleTag)
		ATRouter.PUT("/edit", ArticleTagApi.EditArticleTag)
		ATRouter.DELETE("/delete", ArticleTagApi.DeleteArticleTag)
		ATRouter.GET("/list", ArticleTagApi.GetArticleTag)
	}
}
