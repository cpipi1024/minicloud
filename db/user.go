package db

import (
	"cpipi1024.com/minicloud/bootstrap"
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	UUID     string   `gorm:"column:uuid;type:varchar(100)" json:"uuid"`
	Name     string   `gorm:"column:name;type:varchar(100)" json:"name"`
	Email    string   `gorm:"column:email;type:varchar(100)" json:"email"`
	Mobile   string   `gorm:"column:mobile;type:varchar(100)" json:"mobile"`
	Password string   `gorm:"column:password;type:varchar(100)" json:"password,omitempty"`
	Role     RoleType `gorm:"column:role;type:integer unsigned" json:"role"`
	BaseDir  string   `gorm:"column:base_dir;type:varchar(100)" json:"base_dir"`
}

// todo: user生成uuid的回调
/* func (u *User) BeforeCreate(db *gorm.DB) error {

	u.UUID = uuid.NewString()
	return nil
} */

// todo: 返回uuid
//
// 实现jwtuser接口
func (u User) GetUUID() string {
	return u.UUID
}

// todo: 返回用户角色
func (u User) GetRole() RoleType {
	return u.Role
}

// todo: 返回用户名
func (u User) GetName() string {
	return u.Name
}

// todo: 返回用户默认目录
func (u User) GetBaseDir() string {
	return u.BaseDir
}

// todo: 创建用户
func CreateUser(data map[string]interface{}) error {
	u := User{
		UUID:     data["uuid"].(string),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Mobile:   data["mobile"].(string),
		Password: data["password"].(string),
		BaseDir:  data["base_dir"].(string),
		Role:     RoleUser,
	}

	res := bootstrap.MSDB.Create(&u)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// todo: 根据id查询用户
func QueryUserByID(id int) (*User, error) {

	user := User{}

	res := bootstrap.MSDB.Where("id = ?", id).Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

// todo: 根据uuid查询用户
func QueryUserByUUID(uuid string) (*User, error) {
	user := User{}

	res := bootstrap.MSDB.Where("uuid = ?", uuid).Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

// todo: 根据mobile查询用户
func QueryUserByMobile(mobile string) (*User, error) {
	user := User{}

	res := bootstrap.MSDB.Where("mobile = ?", mobile).Find(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

// todo: 分页查询用户
func QueryUsers(query map[string]interface{}, pageNum, pageSize int) ([]*User, error) {

	users := []*User{}

	res := bootstrap.MSDB.Limit(pageSize).Offset(pageNum).Where(query).Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// todo: 根据ID修改用户
func UpdateUserByID(id int, data map[string]interface{}) error {
	res := bootstrap.MSDB.Model(User{}).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// todo: 根据UUID修改用户
func UpdateUserByUUID(uuid string, data map[string]interface{}) error {
	res := bootstrap.MSDB.Model(User{}).Where("uuid = ?", uuid).Updates(data)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// todo: 注册手机号是否存在
func UserMobileExist(mobile string) bool {
	var exist int64

	bootstrap.MSDB.Model(User{}).Where("mobile = ?", mobile).Count(&exist)

	return exist > 0
}
