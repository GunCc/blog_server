package initialize

import (
	"blog_server/global"
	"blog_server/model/system"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	// switch globa

	// return GormMysql();
	switch global.BLOG_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	// 注册表
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.SysOperationRecord{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
	)
	if err != nil {
		global.BLOG_LOG.Error("注册表失败", zap.Error(err))
		os.Exit(0)
	}
	global.BLOG_LOG.Info("注册表成功")
}
