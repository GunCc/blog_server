package system

import (
	"blog_server/global"
	"blog_server/model/commen/response"
	systemRes "blog_server/model/system/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

// 单个服务器使用这个服务器, 如果使用多个服务器 使用captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// CaptchaConfig := global.BLOG_CONFIG.Captcha
	// 字符，公式，验证码配置
	// 生成默认数字的driver
	// driver := base64Captcha.NewDriverDigit(global.BLOG_CONFIG.Captcha.ImgHeight, global.BLOG_CONFIG.Captcha.ImgWidth, global.BLOG_CONFIG.Captcha.KeyLong, 0.7, 80)
	driver := base64Captcha.DefaultDriverDigit
	cp := base64Captcha.NewCaptcha(driver, store)
	// 生成

	if id, b64s, err := cp.Generate(); err != nil {
		global.BLOG_LOG.Error("验证码获取失败！", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(&systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.BLOG_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", c)
	}
}
