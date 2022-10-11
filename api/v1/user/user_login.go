package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginApi struct{}

func (u *UserLoginApi) UserLogin(c *gin.Context) {
	c.String(http.StatusOK, "测试数据")
}
