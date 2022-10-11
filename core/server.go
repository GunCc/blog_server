// 初始化 服务器

package core

import (
	"blog_server/initialize"
	"fmt"
)

type server interface {
	ListenAndServe() error
}

// 运行服务
func RunWindowsServer() {

	// 初始化路由
	Router := initialize.Routers()
	// fmt.Sprintf(":%d", global)
	s := InitServer(":8080", Router)
	fmt.Errorf(s.ListenAndServe().Error())
}
