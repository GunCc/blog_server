package initialize

import (
	"blog_server/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	userRouter := routers.RouterGroupApp.User

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
	}

	// 权限路由
	PrivateGroup := Router.Group("")
	{
		userRouter.InitUserRouter(PrivateGroup) // 注册用户相关路由
	}
	return Router
}
