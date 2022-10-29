package system

import (
	"blog_server/global"
	"blog_server/model/system"
	"errors"

	"gorm.io/gorm"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

func (ApiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.BLOG_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.BLOG_DB.Create(&api).Error
}
