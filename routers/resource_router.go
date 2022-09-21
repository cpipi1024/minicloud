package routers

import (
	"cpipi1024.com/minicloud/api"
	"cpipi1024.com/minicloud/middlewares"
	"github.com/gin-gonic/gin"
)

func ResourceRouter(r *gin.RouterGroup) {
	resource := r.Group("/resource")

	// 获取文件信息

	auth := resource.Group("").Use(middlewares.JWTMiddlewares("mincloud"))
	{
		// 获取文件信息
		auth.GET("/file", api.GetResourceDetail)
		// 获取文件列表
		auth.GET("/files", api.ListFiles)
		// 上传文件
		auth.POST("/upload", api.UploadFile)
		// 下载文件
		auth.GET("/download", api.DownloaFile)
		// 删除文件
		auth.GET("/deleteFile", api.DeleteResourceFile)
		// 删除文件夹
		auth.GET("/deleteDir", api.DeleteResourceDir)
	}
}
