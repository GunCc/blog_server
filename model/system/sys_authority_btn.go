package system

type SysAuthorityBtn struct {
	Authority        uint           `gorm:"comment:角色Id"`
	SysMenuID        uint           `gorm:"comment:菜单Id"`
	SysBaseMenuBtnId uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn `gorm:"comment:按钮详情"`
}
