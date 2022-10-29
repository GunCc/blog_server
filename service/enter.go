package service

import (
	"blog_server/service/system"
	"blog_server/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup system.SystemService
	UserServiceGroup   user.UserGroup
}

var ServiceGroupApp = new(ServiceGroup)
