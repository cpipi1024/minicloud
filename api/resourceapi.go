package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"cpipi1024.com/minicloud/pkg/customerr"
	"cpipi1024.com/minicloud/service"
	"github.com/gin-gonic/gin"
)

const (
	defaultMemory = 32 << 20
)

// todo: 获取文件细节
func GetResourceDetail(c *gin.Context) {

	rawDir, _ := c.Get("baseDir")

	baseDir := rawDir.(string)

	relative := c.DefaultQuery("path", "")

	filename := c.DefaultQuery("name", "")

	data, err := service.ResourceService.ResourceDetail(baseDir, relative, filename)

	if err != nil {
		Fail(c, 0, "获取文件详情失败")
		c.Abort()
		return
	}
	Succes(c, data)
}

// todo: 删除文件
func DeleteResourceFile(c *gin.Context) {
	rawDir, _ := c.Get("baseDir")

	baseDir := rawDir.(string)

	relative := c.DefaultQuery("path", "")

	filename := c.DefaultQuery("name", "")

	// 缺少filename的检查

	err := service.ResourceService.DeleteResourceFile(baseDir, relative, filename)

	if err != nil {
		Fail(c, 0, "删除文件失败")
		c.Abort()
		return
	}

	Succes(c, "success")

}

// todo: 获取指定路径文件夹下的文件
func ListFiles(c *gin.Context) {
	rawDir, _ := c.Get("baseDir")

	baseDir := rawDir.(string)

	relativePath := c.DefaultQuery("path", "") // 获取的相对路径

	data, err := service.ResourceService.ListContents(baseDir, relativePath)

	if err != nil {
		Fail(c, 0, "获取文件夹信息失败")
		c.Abort()
		return
	}

	Succes(c, data)
}

// todo: 创建文件夹
func CreateResourceDir(c *gin.Context) {
	rawDir, _ := c.Get("baseDir")

	baseDir := rawDir.(string)

	relative := c.DefaultQuery("path", "")

	dirname := c.DefaultQuery("name", "")

	err := service.ResourceService.CreateResourceDir(baseDir, relative, dirname)

	if err != nil {
		Fail(c, 0, "创建文件夹失败")
		c.Abort()
		return
	}

	Succes(c, "success")
}

// todo: 删除文件夹
func DeleteResourceDir(c *gin.Context) {
	rawDir, _ := c.Get("baseDir")

	baseDir := rawDir.(string)

	relative := c.DefaultQuery("path", "")

	dirname := c.DefaultQuery("name", "")

	err := service.ResourceService.DeleteResourceDir(baseDir, relative, dirname)

	if err != nil {
		Fail(c, 0, "删除文件夹失败")
		c.Abort()
		return
	}

	Succes(c, "success")
}

// todo: 上传文件
func UploadFile(c *gin.Context) {

	rawDir, _ := c.Get("baseDir") // 获取用户目录

	baseDir := rawDir.(string)

	relativePath := c.DefaultQuery("path", "") // 当前相对路径

	err := c.Request.ParseMultipartForm(defaultMemory) // 配置最大文件大小

	if err != nil {
		Fail(c, customerr.CodeResourceUploadFailed, "文件上传失败")
		c.Abort()
		return
	}

	// 解析上传表单
	if c.Request.MultipartForm != nil {
		if fileheaders := c.Request.MultipartForm.File; fileheaders != nil {
			// 获取对应mimetype下的files
			for _, files := range fileheaders {
				// 获得上传文件句柄
				for _, file := range files {
					service.ResourceService.StreamUploadResource(baseDir, relativePath, file)
				}
			}
		}
	}

	Succes(c, "success")
}

// todo: 下载文件
func DownloaFile(c *gin.Context) {
	rawDir, _ := c.Get("baseDir")

	baseDir := rawDir.(string)

	relative := c.DefaultQuery("path", "")

	filename := c.DefaultQuery("name", "")

	// 获取下载文件信息
	info, err := service.ResourceService.ResourceDetail(baseDir, relative, filename)

	if err != nil {
		Fail(c, 0, "资源下载失败")
		c.Abort()
		return
	}

	fileSize := info.Size
	mimeType := info.Mime
	lastMidified := info.LastModified
	fileName := info.Name

	// 预设置响应头
	c.Header("Content-Length", strconv.Itoa(int(fileSize)))
	c.Header("Content-Type", mimeType)
	c.Header("Last-Modified", time.Unix(lastMidified, 0).UTC().Format(http.TimeFormat))
	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Disposition", "attachment;filename="+fileName)

	// 获取请求头
	ifRangeHeaderValue := c.GetHeader("If-Range") // 断点续传
	rangeHeaderValue := c.GetHeader("Range")      // 文件范围

	isHeaderRequest := c.Request.Method == "HEAD"

	// 文件起始位置和结束位置
	var start, end int64
	_, _ = fmt.Sscanf(rangeHeaderValue, "bytes=%d-%d", start, end)

	// 检验请求文件部分
	if start < 0 || start >= fileSize || end < 0 || end >= fileSize {
		Fail(c, 0, "文件下载失败")
		c.Abort()
		return
	}

	if end == 0 {
		end = fileSize - 1
	}
	if rangeHeaderValue != "" {
		if ifRangeHeaderValue != "" && ifRangeHeaderValue != time.Unix(lastMidified, 0).UTC().Format(http.TimeFormat) {
			// 如果if-range请求头存在，但是匹配不上文件修改时间，则直接返回完整文件
			c.Status(http.StatusOK)
			newpath := filepath.Join(baseDir, relative, fileName)
			c.File(newpath)
			return
		} else {
			// 匹配则返回指定部分数据
			// 响应状态码 206
			c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
			c.Status(http.StatusPartialContent)
		}
	}

	// 过滤head请求
	if !isHeaderRequest {
		path := filepath.Join(baseDir, relative, fileName)

		file, err := os.Open(path)

		if err != nil {
			Fail(c, 0, "下载文件失败")
			c.Abort()
			return
		}

		// 找到分块文件开始下标
		file.Seek(start, 0)

		// 传输数据
		_, err = io.CopyN(c.Writer, file, end-start+1)

		if err != nil {
			return
		}
	}

}
