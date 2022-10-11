package routers

import "blog_server/routers/user"

type RouterGroup struct {
	User user.UserRouter
}

var RouterGroupApp = new(RouterGroup)
