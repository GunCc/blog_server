package system

import "blog_server/service"

type ApiGroup struct {
	SystemApiApi
	BaseApi
}

var (
	apiService  = service.ServiceGroupApp.SystemServiceGroup.ApiService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	JwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
