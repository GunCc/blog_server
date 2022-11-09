package initialize

import (
	"blog_server/global"
	"blog_server/routers"
	"net/http"

	_ "blog_server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	userRouter := routers.RouterGroupApp.User
	systemRouter := routers.RouterGroupApp.System

	// swagger配置
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.BLOG_LOG.Info("register swagger handler")
	// 做一个健康检测
	// 公共路由
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
	{
		// 注册基础路由,这些是不用权限控制的

		// useruserRouter.userRouter
		systemRouter.InitUserRouter(PublicGroup)
		systemRouter.InitBaseRouter(PublicGroup)

	}

	// 权限路由
	PrivateGroup := Router.Group("")
	{
		userRouter.InitUserRouter(PrivateGroup) // 用户相关路由
		systemRouter.InitApiRouter(PrivateGroup)
		systemRouter.InitArticleTypeRouter(PrivateGroup) // 文章标题相关路由
		systemRouter.InitArticleTagRouter(PrivateGroup)
		systemRouter.InitBlogArticleRouter(PrivateGroup)
	}
	return Router
}
