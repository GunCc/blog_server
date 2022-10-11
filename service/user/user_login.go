package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
}

func (u *UserLoginService) login(c *gin.Context) {
	c.String(http.StatusOK, "测试数据")
}
