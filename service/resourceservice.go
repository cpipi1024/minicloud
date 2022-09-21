package service

import (
	"io"
	"mime/multipart"
	"path/filepath"

	"cpipi1024.com/minicloud/pkg/explorer"
)

type resourceService struct{}

var ResourceService = new(resourceService)

// todo: 创建文件夹
func (service *resourceService) CreateResourceDir(path, relative, dirname string) error {
	exp := explorer.NewDefaultExplorer(path)

	err := exp.EnterDir(relative)

	if err != nil {
		return err
	}

	err = exp.CreateDir(dirname)

	if err != nil {
		return err
	}

	return nil
}

//todo: 删除文件夹
func (service *resourceService) DeleteResourceDir(path, relative, dirname string) error {
	exp := explorer.NewDefaultExplorer(path)

	err := exp.EnterDir(relative)

	if err != nil {
		return err
	}

	return exp.DeleteDir(dirname)

}

// todo: 删除文件
func (service *resourceService) DeleteResourceFile(path, relative, filename string) error {
	exp := explorer.NewDefaultExplorer(path)

	err := exp.EnterDir(relative)

	if err != nil {
		return err
	}

	return exp.DeleteFile(filename)
}

// todo: 列出当前目录下的内容
func (service *resourceService) ListContents(path, relative string) ([]*explorer.File, error) {
	exp := explorer.NewDefaultExplorer(path)

	// 进入对应目录
	err := exp.EnterDir(relative)

	if err != nil {
		return nil, err
	}

	return exp.ListContents()
}

// todo: 获取文件info
func (service *resourceService) ResourceDetail(path, relative, fileaname string) (*explorer.File, error) {
	exp := explorer.NewDefaultExplorer(path)

	exp.EnterDir(relative)

	newpath := filepath.Join(exp.GetCurrentDir(), fileaname)

	return exp.GetFileInfo(newpath)

}

// todo: 下载文件
func (service *resourceService) StreamDownloadResource(path, relative, resourceName string) (io.Reader, error) {

	exp := explorer.NewDefaultExplorer(path)

	err := exp.EnterDir(relative)

	if err != nil {
		return nil, err
	}

	newpath := filepath.Join(exp.GetCurrentDir(), resourceName)

	return exp.StreamDownload(newpath)
}

// todo: 上传文件
func (service *resourceService) StreamUploadResource(path string, relative string, fh *multipart.FileHeader) error {

	exp := explorer.NewDefaultExplorer(path)

	err := exp.EnterDir(relative)

	if err != nil {
		return err
	}

	newpath := filepath.Join(exp.GetCurrentDir(), fh.Filename)

	src, err := fh.Open()

	if err != nil {
		return err
	}

	return exp.StreamUpload(newpath, src)
}
