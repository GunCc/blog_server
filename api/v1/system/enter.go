package system

import "blog_server/service"

type ApiGroup struct {
	SystemApiApi
	BaseApi
	ArticleTypeApi
	ArticleTagApi
}

var (
	apiService         = service.ServiceGroupApp.SystemServiceGroup.ApiService
	userService        = service.ServiceGroupApp.SystemServiceGroup.UserService
	JwtService         = service.ServiceGroupApp.SystemServiceGroup.JwtService
	ArticleTypeService = service.ServiceGroupApp.SystemServiceGroup.ArticleTypeService
	ArticleTagService  = service.ServiceGroupApp.SystemServiceGroup.ArticleTagService
)
