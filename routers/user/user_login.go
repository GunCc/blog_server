package user

import (
	v1 "blog_server/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	UserLoginApi := v1.ApiGroupApp.UserApiGroup.UserLoginApi
	userRouter := r.Group("/user")
	{
		userRouter.GET("/login", UserLoginApi.UserLogin)
	}

}
