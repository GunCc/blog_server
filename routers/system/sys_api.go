package system

import (
	v1 "blog_server/api/v1"
	"blog_server/middlewares"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middlewares.OperationRecord())
	// apiRouerWithoutRecord := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SysApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)
	}
}
