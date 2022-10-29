package system

import "blog_server/service"

type ApiGroup struct {
	SystemApiApi
}

var (
	apiSerice = service.ServiceGroupApp.SystemServiceGroup.ApiService
)
