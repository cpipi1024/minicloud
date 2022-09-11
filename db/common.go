package db

import (
	"cpipi1024.com/minicloud/bootstrap"
	"gorm.io/gorm"
)

type RoleType uint16

// 角色级别
const (
	RoleAdmin RoleType = iota //管理员
	RoleUser                  //普通用户
	RoleGuest                 //游客
)

var (
	tables []interface{} = []interface{}{
		&User{},
		&Resource{},
	}
)

// todo: 自动迁移表
func MigrateTables() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if bootstrap.SrvConf.App.Mode == "dev" {
			db.AutoMigrate(tables...)
		}
	}
}

// todo:创建admin用户
func CreateAdminUser() func(db *gorm.DB) {

	return func(db *gorm.DB) {
		if adminExist(db) {
			return
		}
		admin := User{
			Name:     "admin",
			Email:    "1461481767@qq.com",
			Mobile:   "15575382165",
			Password: "minicloud",
			Role:     RoleAdmin,
		}

		db.Create(&admin)
	}
}

// todo: admin用户是否存在
func adminExist(db *gorm.DB) bool {

	var exist int64

	db.Model(User{}).Where("role = ? AND email = ?", RoleAdmin, "1461481767@qq.com").Count(&exist)

	return exist > 0
}
