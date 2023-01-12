package initialize

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化Gorm
func Gorm() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mangoblognew?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
		return nil
	} else {
		fmt.Println("数据库打开成功，dsn：", dsn)
		return db
	}
}
