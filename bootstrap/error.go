package bootstrap

import (
	"errors"
)

// 错误code
const (
	// jwt auth
	CodeJWTAuthFailed  = 10001 // token签发失败
	CodeJWTAuthInvalid = 10002 // token无效
	CodeJWTExpired     = 10003 // token失效

	// userserivce
	CodeRegisterFailed = 20001 // 用户注册失败
	CodeLoginFailed    = 20002 // 用户登录失败

	// mysql
	CodeMySQLOptFailed = 30001 // 数据库操作错误

	// 请求验证失败
	CodeRequestParmaFailed = 40001 // 请求参数验证失败
)

// 错误集
var (
	// jwt auth
	ErrJWTAuthFailed = errors.New("token签发失败")

	// userservice
	ErrRegisterFailed = errors.New("用户注册失败")
	ErrLoginFailed    = errors.New("用户登录失败")

	// mysql
	ErrMySQLOptFailed = errors.New("Mysql数据库操作失败")
)

var _ error = (*CustomError)(nil)

// 自定义错误
type CustomError struct {
	Inner error  // 内置错误
	Code  int    // 错误码
	Msg   string // 错误信息
}

// todo: 创建自定义函数
func NewCustomError(code int, msg string) *CustomError {
	return &CustomError{
		Code: code,
		Msg:  msg,
	}
}

// todo: 实现error接口
func (err *CustomError) Error() string {
	if err.Inner != nil {
		return err.Inner.Error()
	} else if err.Msg != "" {
		return err.Msg
	} else {
		return "unkown err"
	}
}

// todo: 判断错误类别
func (err *CustomError) Is(other error) bool {
	if errors.Is(errors.Unwrap(err), other) {
		return true
	}

	switch other {
	case ErrRegisterFailed:
		return err.Code == CodeRegisterFailed
	case ErrLoginFailed:
		return err.Code == CodeLoginFailed
	case ErrMySQLOptFailed:
		return err.Code == CodeMySQLOptFailed
	case ErrJWTAuthFailed:
		return err.Code == CodeJWTAuthFailed
	}

	return false
}

// todo: 返回嵌套的err
func (err *CustomError) UnWrap() error {
	return err.Inner
}
