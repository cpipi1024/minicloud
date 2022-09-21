package middlewares

import (
	"strings"

	"cpipi1024.com/minicloud/api"
	"cpipi1024.com/minicloud/pkg/customerr"
	"cpipi1024.com/minicloud/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	secretKey = "cpipi1024"
)

func JWTMiddlewares(issuer string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.Request.Header.Get("Authorization") // 获取token串

		if len(data) == 0 {
			api.Fail(c, customerr.CodeJWTAuthFailed, "token鉴权失败")
			c.Abort()
			return
		}

		tokenType := strings.ToLower(data[:6]) // token类型

		if tokenType != "bearer" {
			api.Fail(c, customerr.CodeJWTAuthInvalid, "token类型错误")
			c.Abort()
			return
		}
		tokenstr := data[7:] // 获取加密后的token串

		// decode token
		token, err := jwt.ParseWithClaims(tokenstr, &service.CustomClaim{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			api.Fail(c, customerr.CodeJWTAuthInvalid, "token捡钱失败")
			c.Abort()
			return
		}

		// 获得token负载内容
		claims := token.Claims.(*service.CustomClaim)
		c.Set("uuid", claims.UUID)
		c.Set("userName", claims.UserName)
		c.Set("role", claims.Role)
		c.Set("baseDir", claims.BaseDir) // 用户默认路径
	}
}
