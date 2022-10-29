package system

import (
	"blog_server/global"
	"blog_server/model/commen/response"
	"blog_server/model/system"
	"blog_server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApiApi struct{}

func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	// 解析json
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiSerice.CreateApi(api); err != nil {
		global.BLOG_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}
