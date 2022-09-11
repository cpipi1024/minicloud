package bootstrap

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type injecotor func(*gorm.DB)

// todo: 加载mysql
func LoadMysql() {
	msConf := SrvConf.Mysql

	dsn := genDsn(msConf.Host, msConf.Port, msConf.User, msConf.Password, msConf.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	// 注入依赖
	calinject(db)

	MSDB = db
}

// todo: mysql依赖注入
func calinject(db *gorm.DB) {
	for _, f := range injectors {
		f(db)
	}
}

// todo: 注册mysql依赖
func RegisterInjector(f injecotor) {
	injectors = append(injectors, f)
}

// todo: 拼接dsn
func genDsn(host string, port int, username string, password string, database string) string {
	fmtstr := "%s:%v@tcp(%v:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	return fmt.Sprintf(fmtstr, host, port, username, password, database)

}
