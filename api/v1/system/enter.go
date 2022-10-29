package system

import "blog_server/service"

type ApiGroup struct {
	SystemApiApi
	BaseApi
}

var (
	apiSerice = service.ServiceGroupApp.SystemServiceGroup.ApiService
)
