package routers

import (
	"blog_server/routers/system"
	"blog_server/routers/user"
)

type RouterGroup struct {
	User   user.UserRouter
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
