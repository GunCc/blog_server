package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/initialize"
	"fmt"
)

// @title 博客系统 Swagger
// @version 0.0.1
// @description Go 语言编程之旅：一起用 Go 做项目
// @in header
// @name x-token
// @BasePath /
func main() {
	// 初始化Viper Viper就是一个可以解析.yaml 文件的一个golang 包
	global.BLOG_VP = core.Viper()
	// 增加zap日志库
	global.BLOG_LOG = core.Zap()
	// 初始化数据库配置
	global.BLOG_DB = initialize.Gorm()

	fmt.Println("数据------------", global.BLOG_DB)
	// 初始化数据库
	initialize.DBList()

	if global.BLOG_DB != nil {
		initialize.RegisterTables(global.BLOG_DB)
		d, _ := global.BLOG_DB.DB()
		defer d.Close()
	}
	// 启动服务
	core.RunWindowsServer()
}
