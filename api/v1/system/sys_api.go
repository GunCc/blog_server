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

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {object} response.Response{msg=string} "创建基础api"
// @Router /api/createApi [post]
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	// 解析json
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiService.CreateApi(api); err != nil {
		global.BLOG_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}
