package system

import (
	v1 "blog_server/api/v1"
	"blog_server/middlewares"

	"github.com/gin-gonic/gin"
)

type BlogArticleRouter struct {
}

func (B *BlogArticleRouter) InitBlogArticleRouter(c *gin.RouterGroup) {
	RouterGroup := c.Group("/article").Use(middlewares.OperationRecord())
	ArticleApi := v1.ApiGroupApp.SysApiGroup.ArticleApi
	{
		RouterGroup.POST("/add", ArticleApi.CreateArticle)
		RouterGroup.PUT("/edit", ArticleApi.EditArticle)
		RouterGroup.DELETE("/delete", ArticleApi.DeleteArticle)
		RouterGroup.GET("/list", ArticleApi.GetArticleList)
	}
}
