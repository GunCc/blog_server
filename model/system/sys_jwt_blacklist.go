package system

import "blog_server/global"

type JwtBlackList struct {
	global.BLOG_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
