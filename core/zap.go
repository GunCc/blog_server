package core

import (
	"blog_server/core/internal"
	"blog_server/global"
	"blog_server/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger

func Zap() *zap.Logger {
	if ok, _ := utils.PathExists(global.BLOG_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.BLOG_CONFIG.Zap.Director)
		// 创建文件夹 0777 最高权限
		os.Mkdir(global.BLOG_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))
	if global.BLOG_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
