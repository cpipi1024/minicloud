package routers

import (
	"cpipi1024.com/minicloud/api"
	"cpipi1024.com/minicloud/middlewares"
	"github.com/gin-gonic/gin"
)

// todo: 注册user模块录音
func UserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")

	// 注册服务
	user.POST("/register", api.UserRegister)
	// 登录服务
	user.POST("/login", api.UserLogin)

	auth := user.Group("").Use(middlewares.JWTMiddlewares("minicloud"))

	{
		// 通过UUID得到user
		auth.GET("/auth/:uuid", api.GetUserByUUID)
	}

}
