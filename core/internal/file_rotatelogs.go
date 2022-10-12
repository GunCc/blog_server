package internal

import (
	"blog_server/global"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// 这个文件主要是使用 rotatelogs 做 日志分割
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	// 创建一个新 分割日志文件 第一个参数是 文件路径 使用path.Join创建文件
	fileWritter, err := rotatelogs.New(
		path.Join(global.BLOG_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		// 设置时区 Local 获取当前地区时间
		rotatelogs.WithClock(rotatelogs.Local),
		// 设置最大年龄 也就是多久删除这个文件
		rotatelogs.WithMaxAge(time.Duration(global.BLOG_CONFIG.Zap.MaxAge)*24*time.Hour),
		// 多长时间创建一个新的 这是一天
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	// 日志是否打印出来
	if global.BLOG_CONFIG.Zap.LogInConsole {
		// 同步写入
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWritter)), err

	}
	// 同步写入
	return zapcore.AddSync(fileWritter), err
}
