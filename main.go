package main

import (
	"blog_server_new/core"
	"blog_server_new/global"
	"blog_server_new/initialize"
)

func main() {

	// 初始化数据库
	global.BLOG_DB = initialize.Gorm()
	// 启动服务
	core.RunWindowServer()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
