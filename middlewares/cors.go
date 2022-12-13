package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
		context.Header("Access-Control-Allow-Origin", "*")
		// context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
		fmt.Println("===================", context.GetHeader("Origin"))
		// 必须，设置服务器支持的所有跨域请求的方法
		context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 服务器支持的所有头信息字段，不限于浏览器在“预检”中请求的字段
		context.Header("Access-Control-Allow-Headers", "Content-Type,Content-Length,Token")
		// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
		context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers,Token")
		// 可选，是否允许后续请求带Cookie，该值只能是true，不需要就不要设置了
		context.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有的 options 方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
