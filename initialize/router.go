package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	// 公共路由（不需要jwt检测）
	PublicGroup := router.Group("")
	{
		// 健康检测
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, "成功")
		})
	}
	{
		// 基础路由
	}

	// 私有路由（需要正常登录才可以执行的操作）
	// PrivateGroup := router.Group("")
	// {

	// }
	fmt.Println("路由注册成功")
}
