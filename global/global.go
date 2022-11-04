package global

import (
	"blog_server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量函数
var (
	BLOG_DB     *gorm.DB
	BLOG_VP     *viper.Viper
	BLOG_CONFIG config.Server
	BLOG_LOG    *zap.Logger
	BLOG_DBList map[string]*gorm.DB
	BLOG_REDIS  *redis.Client
)
