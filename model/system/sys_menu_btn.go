package system

import "blog_server/global"

type SysBaseMenuBtn struct {
	global.BLOG_MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuId" gorm:"comment:菜单ID"`
}
