package system

import (
	"blog_server/global"
	"blog_server/model/commen/response"
	"blog_server/model/system"
	"blog_server/model/system/request"
	"fmt"

	systemRes "blog_server/model/system/response"
	"blog_server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {object} systemRes.LoginResponse{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&l)
	fmt.Println(l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// if store.Verify(l.CaptchaId, l.Captcha, true) {
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	if user, err := userService.Login(u); err != nil {
		global.BLOG_LOG.Error("登陆失败！用户名不存在或者密码错误！", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		if user.Enable != 1 {
			global.BLOG_LOG.Error("登陆失败！用户被禁止登录！")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(c, *user)
	}
	// } else {
	// 	response.FailWithMessage("验证码错误", c)
	// }
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce application/json
// @Param data body request.Register true "用户米，昵称，密码，角色ID"
// @Success 200 {object} systemRes.SysUserResponse{data=systemRes.SysUserResponse,msg=string} "用户注册账号，返回包括用户信息"
// @Router /user/register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		fmt.Println("报错：", err)
		return
	}
	fmt.Println("数据：=========", r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{
		Username: r.Username,
		NickName: r.NickName,
		Password: r.Password,
		// HeaderImg:   r.HeaderImg,
		AuthorityId: r.AuthorityId,
		Authorities: authorities,
		Enable:      r.Enable,
	}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.BLOG_LOG.Error("注册失败！", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{
			User: userReturn,
		}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{
			User: userReturn,
		}, "注册成功", c)
	}
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	jwt := &utils.JWT{SigningKey: []byte(global.BLOG_CONFIG.JWT.SigningKey)}
	clamis := jwt.CreateClamis(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NikckName:   user.NickName,
		AuthorityId: user.AuthorityId,
	})
	token, err := jwt.CreateToken(clamis)
	if err != nil {
		global.BLOG_LOG.Error("获取token失败！", zap.Error(err))
		response.FailWithMessage("获取token失败！", c)
		return
	}
	// 多点登录
	if !global.BLOG_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clamis.StandardClaims.ExpiresAt * 1000,
		}, "登陆成功", c)
		return
	}

	if redisJWT, err := JwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := JwtService.SetRedisJWT(token, user.Username); err != nil {
			global.BLOG_LOG.Error("设置登陆状态失败！", zap.Error(err))
			response.FailWithMessage("设置登陆状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clamis.StandardClaims.ExpiresAt * 1000,
		}, "登陆成功", c)

	} else if err != nil {
		global.BLOG_LOG.Error("设置登陆状态失败！", zap.Error(err))
		response.FailWithMessage("设置登陆状态失败", c)
	} else {
		var blackJWT system.JwtBlackList
		blackJWT.Jwt = redisJWT
		// JwtService.JSON
		if err := JwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登陆状态失败", c)
			// global.BLOG_REDIS.Options().TLSConfig.Time().UTC().Round()
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clamis.StandardClaims.ExpiresAt * 1000,
		}, "登陆成功", c)
	}

}
