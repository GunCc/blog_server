package system

import (
	"blog_server/global"
	"blog_server/model/system"
)

type OperationRecordService struct{}

// 创建记录
func (o *OperationRecordService) CreateSysOperationRecord(SysOperationRecord system.SysOperationRecord) (err error) {
	err = global.BLOG_DB.Create(&SysOperationRecord).Error
	return err
}
