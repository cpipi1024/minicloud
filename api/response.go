package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应
type Response struct {
	Code int         `json:"code"`           // 自定义code
	Msg  interface{} `json:"msg"`            // 自定义msg
	Data interface{} `json:"data,omitempty"` // data
}

// todo: 成功响应
func Succes(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: data,
		},
	)
}

// todo: 失败响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(
		http.StatusOK,
		Response{
			Code: code,
			Msg:  msg,
		},
	)
}
