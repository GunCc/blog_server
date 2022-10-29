package v1

import (
	"blog_server/api/v1/system"
	"blog_server/api/v1/user"
)

type ApiGroup struct {
	UserApiGroup user.ApiGroup
	SysApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
