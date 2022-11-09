package system

import (
	v1 "blog_server/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	// userRouter := Router.Group("user").Use(middlewares.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SysApiGroup.BaseApi
	{

	}
	{
		userRouterWithoutRecord.POST("register", baseApi.Register)

	}
	return userRouterWithoutRecord
}
