package system

import (
	v1 "blog_server/api/v1"
	"blog_server/middlewares"

	"github.com/gin-gonic/gin"
)

type ArticleTypeRouter struct {
}

func (r *ArticleTypeRouter) InitArticleTypeRouter(Router *gin.RouterGroup) {
	articleTypeRouter := Router.Group("article/type").Use(middlewares.OperationRecord())
	articleTypeApi := v1.ApiGroupApp.SysApiGroup.ArticleTypeApi
	{
		articleTypeRouter.POST("/add", articleTypeApi.CreateArticleType)
		articleTypeRouter.DELETE("/delete", articleTypeApi.DeleteArticleType)
		articleTypeRouter.PUT("/edit", articleTypeApi.EditArticleType)
		articleTypeRouter.POST("/list", articleTypeApi.GetArticleType)
	}
}
