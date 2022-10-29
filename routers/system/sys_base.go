package system

import (
	v1 "blog_server/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SysApiGroup.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
