package middlewares

import (
	"net/http"

	"cpipi1024.com/minicloud/api"
	"cpipi1024.com/minicloud/pkg/customerr"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// todo: 中间件处理错误
func ErrHandle(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			if custErr, ok := err.(*customerr.CustomError); ok {
				// 处理自定义错误
				value, _ := c.Get("zapinfo") //从context上下文获取info
				info := value.(string)
				logger.Error(
					info,
					zap.String("Method", c.Request.Method),
					zap.String("Path", c.Request.URL.Path),
					zap.Int("ErrCode", custErr.Code),
					zap.String("ErrMsg", custErr.Error()),
				)
				c.JSON(http.StatusOK, api.Response{
					Code: custErr.Code,
					Msg:  custErr.Msg,
					Data: nil,
				})
			} else {
				// 处理内置错误
				c.JSON(http.StatusOK, api.Response{
					Code: 500,
					Msg:  "服务器异常 ",
					Data: err.Error(),
				})
			}
			break
		}
	}
}
