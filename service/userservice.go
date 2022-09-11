package service

import (
	"cpipi1024.com/minicloud/bootstrap"
	"cpipi1024.com/minicloud/db"
	"cpipi1024.com/minicloud/forms"
)

type userService struct{}

var UserService = new(userService)

// todo: 用户注册
func (service *userService) RegisterUser(form forms.UserRegisterForm) error {

	exist := db.UserMobileExist(form.Mobile)

	if exist {
		return bootstrap.NewCustomError(bootstrap.CodeRegisterFailed, "注册手机号已存在")
	}

	data := map[string]interface{}{
		"name":     form.Name,
		"email":    form.Email,
		"mobile":   form.Email,
		"password": form.Password,
		"role":     db.RoleUser,
	}

	err := db.CreateUser(data)

	if err != nil {
		return &bootstrap.CustomError{Inner: err, Code: bootstrap.CodeMySQLOptFailed, Msg: "注册用户失败"}
	}

	return nil
}

// todo: 用户登录
func (service *userService) UserLogin(form forms.UserLoginForm) (TokenData, error) {

	u, err := db.QueryUserByMobile(form.Mobile)

	if err != nil {
		return TokenData{}, &bootstrap.CustomError{Inner: err, Code: bootstrap.CodeMySQLOptFailed, Msg: "用户登录失败"}
	}

	if form.Password != u.Password {
		return TokenData{}, bootstrap.NewCustomError(bootstrap.CodeLoginFailed, "用户密码错误")
	}

	_, tokendata, err := JwtService.CreateToken(u)

	if err != nil {
		return TokenData{}, &bootstrap.CustomError{Inner: err, Code: bootstrap.CodeJWTAuthFailed, Msg: "token授权失败"}
	}

	return tokendata, nil
}

// todo: 根据ID获取用户信息
func (service *userService) GetUserByID(id int) (*db.User, error) {
	u, err := db.QueryUserByID(id)

	if err != nil {
		return nil, &bootstrap.CustomError{Inner: err, Code: bootstrap.CodeMySQLOptFailed, Msg: "根据ID查询用户失败"}
	}

	return u, nil

}

// todo: 根据UUID获取用户信息
func (service *userService) GetUserByUUID(uuid string) (*db.User, error) {
	u, err := db.QueryUserByUUID(uuid)

	if err != nil {
		return nil, &bootstrap.CustomError{Inner: err, Code: bootstrap.CodeMySQLOptFailed, Msg: "根据UUID查询用户失败"}
	}

	return u, nil
}

// todo: 查询用户信息
func (service *userService) QueyUsers(form forms.QueryUserForm) ([]*db.User, error) {
	query := map[string]interface{}{}

	if form.Name != "" {
		query["name"] = form.Name
	}

	if form.Mobile != "" {
		query["mobile"] = form.Mobile
	}

	if form.Email != "" {
		query["email"] = form.Email
	}

	if form.PageNum == 0 {
		form.PageNum = 1
	}

	if form.PageSize == 0 {
		form.PageSize = 10
	}

	users, err := db.QueryUsers(query, form.PageNum, form.PageSize)

	if err != nil {
		return nil, &bootstrap.CustomError{Inner: err, Code: bootstrap.CodeMySQLOptFailed, Msg: "查询用户信息失败"}
	}

	return users, nil
}
