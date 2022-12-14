package system

import (
	"blog_server/global"
	"time"
)

// 操作日志模型
type SysOperationRecord struct {
	global.BLOG_MODEL
	// 操作ip地址
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	Path         string        `json:"path" form:"path"  gorm:"column:path;comment:请求路径"`
	Status       int           `json:"status" form:"status"  gorm:"column:status;comment:请求状态"`
	Latency      time.Duration `json:"latency" form:"latency"  gorm:"column:latency;comment:延迟" swaggertype:"string"`
	Agent        string        `json:"agent"  form:"agent"  gorm:"column:agent;comment:代理"`
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`
	Body         string        `json:"body" form:"boby" gorm:"type:text;column:body;comment:请求Body"`
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应body" `
	UserId       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`
	User         SysUser       `json:"user"`
}
