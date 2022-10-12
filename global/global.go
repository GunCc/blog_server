package global

import (
	"blog_server/service/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 全局变量函数
var (
	BLOG_DB     *gorm.DB
	BLOG_VP     *viper.Viper
	BLOG_CONFIG config.Server
)
