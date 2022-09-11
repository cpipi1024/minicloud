package middlewares

import "github.com/gin-gonic/gin"

// todo: 自定义日志中间件
func GinLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// todo: 自定义panic recover
func GinRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
