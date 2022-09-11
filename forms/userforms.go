package forms

// 注册表单
type UserRegisterForm struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role" form:"role"`
}

// 登录表单
type UserLoginForm struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 查询表单
type QueryUserForm struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
	PageSize int    `json:"page_size" form:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
}

// 修改表单
type UpdateUserForm struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
	Password string `json:"password" form:"password"`
}
