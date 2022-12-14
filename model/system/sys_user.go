package system

import (
	"blog_server/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.BLOG_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"username" gorm:"comment:用户登陆名"`
	Password    string       `json:"-" gorm:"comment:用户登陆密码"`
	NickName    string       `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	Phone       string       `json:"phone" gorm:"comment:用户手机号"`
	AuthorityId uint         `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`

	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Email       string         `json:"email" gorm:"comment:用户邮箱"`
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
