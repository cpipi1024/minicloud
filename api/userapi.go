package api

import (
	"errors"
	"strconv"

	"cpipi1024.com/minicloud/bootstrap"
	"cpipi1024.com/minicloud/db"
	"cpipi1024.com/minicloud/forms"
	"cpipi1024.com/minicloud/pkg/customerr"
	"cpipi1024.com/minicloud/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// todo: 通过ID获取user
func GetUserByID(c *gin.Context) {
	param_id, _ := c.Params.Get("id")

	id, err := strconv.Atoi(param_id)

	if err != nil {
		Fail(c, customerr.CodeRequestParmaFailed, "id类型错误")
		c.Abort()
		return
	}

	u, err := service.UserService.GetUserByID(id)

	if err != nil {
		e, _ := err.(*customerr.CustomError)
		bootstrap.Logger.Error(
			"[query user by ID failed]",
			zap.Int("code", e.Code),
			zap.String("err", e.Error()),
		)
		Fail(c, e.Code, e.Msg)
		c.Abort()
		return
	}

	Succes(c, u)
}

// todo: 通过UUID获取user
func GetUserByUUID(c *gin.Context) {

	uuid, _ := c.Params.Get("uuid")

	u, err := db.QueryUserByUUID(uuid)

	if err != nil {
		e, _ := err.(*customerr.CustomError)
		bootstrap.Logger.Error(
			"query user by uuid failed",
			zap.Int("code", e.Code),
			zap.String("err", e.Error()),
		)

		Fail(c, e.Code, e.Msg)
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
		Fail(c, customerr.CodeRequestParmaFailed, "查询表单错误")
		c.Abort()
		return
	}

	users, err := service.UserService.QueyUsers(queryForm)

	if err != nil {
		e, _ := err.(*customerr.CustomError)

		bootstrap.Logger.Error(
			"[query users failed]",
			zap.Int("code", e.Code),
			zap.String("err", e.Error()),
		)
		Fail(c, e.Code, e.Msg)
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
		Fail(c, customerr.CodeRequestParmaFailed, "登录表单错误")
		c.Abort()
		return
	}

	token, err := service.UserService.UserLogin(loginForm)

	if err != nil {
		e, _ := err.(*customerr.CustomError)

		if errors.Is(e, customerr.ErrMySQLOptFailed) {
			bootstrap.ApiError("[user login failed]", c.Request.Method, c.Request.URL.Path, e.Code, e.Error())
		}

		if errors.Is(e, customerr.ErrLoginFailed) {
			bootstrap.ApiError("[user login failed]", c.Request.Method, c.Request.URL.Path, e.Code, e.Error())
		}

		if errors.Is(e, customerr.ErrJWTAuthFailed) {
			bootstrap.ApiError("[user login failed]", c.Request.Method, c.Request.URL.Path, e.Code, e.Error())
		}

		Fail(c, e.Code, e.Msg)
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
		Fail(c, customerr.CodeRequestParmaFailed, "注册表单错误")
		c.Abort()
		return
	}

	err = service.UserService.RegisterUser(registerForm)

	if err != nil {
		e, _ := err.(*customerr.CustomError)

		if errors.Is(e, customerr.ErrRegisterFailed) {
			bootstrap.ApiError("[register user failed]", c.Request.Method, c.Request.URL.Path, e.Code, e.Error())
		}

		if errors.Is(e, customerr.ErrMySQLOptFailed) {
			bootstrap.ApiError("[register user failed]", c.Request.Method, c.Request.URL.Path, e.Code, e.Error())
		}

		if errors.Is(e, customerr.ErrCreateResourceDirFailed) {
			bootstrap.ApiError("[create user root directory failed]", c.Request.Method, c.Request.URL.Path, e.Code, e.Error())
		}
		Fail(c, e.Code, e.Msg)
		c.Abort()
		return
	}

	Succes(c, "sucess")

}
