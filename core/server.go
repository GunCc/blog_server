package core

import (
	"blog_server_new/initialize"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWindowServer() {

	var blog_gin *gin.Engine
	blog_gin = gin.Default()
	// 初始化路由
	initialize.Routers(blog_gin)
	blog_gin.Run(":8080")
	fmt.Println(`
		运行成功
	`)
}
