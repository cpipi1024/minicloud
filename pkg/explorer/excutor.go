package explorer

import (
	"bytes"
	"io"
	"io/fs"
	"os"
)

var (
	defaultsize int  = 10 << 10 // 默认buffersize 10MB
	defaultmode uint = 0766     // 默认文件模式
)

// 文件操作执行器
type Executor interface {

	// 文件上传
	Upload(string, string) error

	// 流上传
	StreamUpload(string, io.Reader) error

	// 文件下载
	Download(string) (string, error)

	// 流下载
	StreamDownload(string) (io.Reader, error)

	// 复制
	Copy(string, string) error
}

// 默认执行器
type NormalExecutor struct {
	mode       uint // 文件模式
	buffersize int  // buffer大小
}

var _ Executor = (*NormalExecutor)(nil)

// todo: 新建执行器
func NewNormalExecutor() *NormalExecutor {
	exec := NormalExecutor{
		mode:       defaultmode,
		buffersize: defaultsize,
	}

	return &exec
}

// todo: 上传文件
func (executor *NormalExecutor) Upload(path, content string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, fs.FileMode(executor.mode))

	if err != nil {
		return err
	}

	defer file.Close()

	// 使用字符串初始化buffer
	buf := bytes.NewBufferString(content)

	_, err = io.Copy(file, buf)

	if err != nil {
		return err
	}

	return nil
}

// todo: 流上传
func (executor *NormalExecutor) StreamUpload(path string, source io.Reader) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.FileMode(executor.mode))

	if err != nil {
		return err
	}

	defer file.Close()

	// 手动创建buffer
	buf := make([]byte, executor.buffersize)

	_, err = io.CopyBuffer(file, source, buf)

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// todo: 下载文件
func (executor *NormalExecutor) Download(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, fs.FileMode(executor.mode))

	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil && err != io.EOF {
		return "", err
	}

	return string(data), nil
}

// todo: 流下载
// 将文件句柄返回给gin进行下载
//
// 因此文件的句柄需要手动进行关闭
func (executor *NormalExecutor) StreamDownload(path string) (io.Reader, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, fs.FileMode(executor.mode))

	if err != nil {
		return nil, err
	}

	return file, nil
}

// todo: 复制文件
func (executor *NormalExecutor) Copy(source, dest string) error {
	sourceFile, err := os.OpenFile(source, os.O_RDONLY, fs.FileMode(executor.mode))

	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, fs.FileMode(executor.mode))

	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}
