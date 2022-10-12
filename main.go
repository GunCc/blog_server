package main

import (
	"blog_server/core"
	"blog_server/global"
)

func main() {
	// 初始化Viper Viper就是一个可以解析.yaml 文件的一个golang 包
	global.BLOG_VP = core.Viper()
	// 增加zap日志库
	global.BLOG_LOG = core.Zap()
	// 链接数据库
	// 启动服务
	core.RunWindowsServer()
}
