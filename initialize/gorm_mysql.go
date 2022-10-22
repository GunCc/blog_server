package initialize

import (
	"blog_server/config"
	"blog_server/global"
	"blog_server/initialize/internal"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	// 拿到关于数据库的所有配置
	m := global.BLOG_CONFIG.Mysql
	// 如果配置中没有设置数据库名称直接放回错误

	fmt.Println("查看mysqldsn--------", m.Dsn())
	if m.Dbname == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if d, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := d.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return d
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置

func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   // 字符串默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if d, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := d.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return d
	}
}
