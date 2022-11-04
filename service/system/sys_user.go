package system

import (
	"blog_server/global"
	"blog_server/model/system"
	"blog_server/utils"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

func (UserService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.BLOG_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	err = global.BLOG_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

// 用户注册
func (UserService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.BLOG_DB.Where("username =?", user.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已经注册")
	}

	u.Password = utils.BcrypHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.BLOG_DB.Create(&u).Error
	return u, err
}
