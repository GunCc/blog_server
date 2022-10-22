package utils

import (
	"blog_server/global"
	sysRequest "blog_server/model/system/request"

	"github.com/gin-gonic/gin"
)

// 解析Header 是否能成功解析token
func GetClaims(c *gin.Context) (*sysRequest.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	cc, err := j.ParseToken(token)
	if err != nil {
		global.BLOG_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")

	}
	return cc, err
}
