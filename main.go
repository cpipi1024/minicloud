package main

import (
	"cpipi1024.com/minicloud/bootstrap"
	"cpipi1024.com/minicloud/db"
	"cpipi1024.com/minicloud/middlewares"
	"cpipi1024.com/minicloud/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	// 加载conf
	bootstrap.LoadConf()

	// 加载logger
	bootstrap.LoadLogger()

	// 加载 mysql
	bootstrap.RegisterInjector(db.MigrateTables())
	bootstrap.RegisterInjector(db.CreateAdminUser())
	bootstrap.LoadMysql()

	e := gin.New()

	r := e.Group("/v1")

	r.Use(middlewares.GinRecover(bootstrap.Logger, true), middlewares.GinLogger(bootstrap.Logger))

	routers.UserRouter(r)

	routers.ResourceRouter(r)

	e.Run(":9000")

}
