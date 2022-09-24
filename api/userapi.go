package api

import (
	"strconv"

	"cpipi1024.com/minicloud/forms"
	"cpipi1024.com/minicloud/pkg/customerr"
	"cpipi1024.com/minicloud/service"
	"github.com/gin-gonic/gin"
)

// todo: 通过ID获取user
func GetUserByID(c *gin.Context) {
	param_id, _ := c.Params.Get("id")

	id, err := strconv.Atoi(param_id)

	if err != nil {
		custErr := &customerr.CustomError{
			Inner: err,
			Code:  customerr.CodeRequestParmaFailed,
			Msg:   "URL参数验证失败",
		}
		c.Set("zapinfo", "[userservice failed]")
		c.Error(custErr)
		c.Abort()
		return
	}

	u, err := service.UserService.GetUserByID(id)

	if err != nil {
		c.Set("zapinfo", "[userservice failed]")
		c.Error(err)
		c.Abort()
		return
	}

	Succes(c, u)
}

// todo: 通过UUID获取user
func GetUserByUUID(c *gin.Context) {

	uuid, _ := c.Params.Get("uuid")

	u, err := service.UserService.GetUserByUUID(uuid)

	if err != nil {
		c.Set("zapinfo", "[userservice failed]")
		c.Error(err)
		c.Abort()
		return
	}

	Succes(c, u)

}

// todo: 条件查询users
func QueryUsers(c *gin.Context) {
	queryForm := forms.QueryUserForm{}

	err := c.ShouldBind(&queryForm)

	if err != nil {
		custErr := &customerr.CustomError{
			Inner: err,
			Code:  customerr.CodeRequestParmaFailed,
			Msg:   "查询表单错误",
		}
		c.Set("zapinfo", "[userservice failed]")
		c.Error(custErr)
		c.Abort()
		return
	}

	users, err := service.UserService.QueyUsers(queryForm)

	if err != nil {
		c.Set("zapinfo", "[userservice failed]")
		c.Error(err)
		c.Abort()
		return
	}
	Succes(c, users)
}

// todo: 用户登录
func UserLogin(c *gin.Context) {
	loginForm := forms.UserLoginForm{}

	err := c.ShouldBind(&loginForm)

	if err != nil {
		custErr := &customerr.CustomError{
			Inner: err,
			Code:  customerr.CodeRequestParmaFailed,
			Msg:   "登录表单错误",
		}
		c.Error(custErr)
		c.Abort()
		return
	}

	token, err := service.UserService.UserLogin(loginForm)

	if err != nil {
		c.Set("zapinfo", "[userservice failed]")
		c.Error(err)
		c.Abort()
		return
	}

	c.Header("Authorization", token.TokenStr)

	Succes(c, token)

}

// todo: 用户注册
func UserRegister(c *gin.Context) {
	registerForm := forms.UserRegisterForm{}

	err := c.ShouldBind(&registerForm)

	if err != nil {
		custErr := &customerr.CustomError{
			Inner: err,
			Code:  customerr.CodeRequestParmaFailed,
			Msg:   "注册表单错误",
		}
		c.Error(custErr)
		c.Abort()
		return
	}

	err = service.UserService.RegisterUser(registerForm)

	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	Succes(c, nil)
}
