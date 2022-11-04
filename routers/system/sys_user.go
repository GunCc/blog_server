package system

import (
	v1 "blog_server/api/v1"
	"blog_server/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middlewares.OperationRecord())
	baseApi := v1.ApiGroupApp.SysApiGroup.BaseApi
	{
		userRouter.POST("register", baseApi.Register)
	}
}
