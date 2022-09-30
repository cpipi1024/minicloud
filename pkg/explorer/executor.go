package explorer

import (
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

	// 设置模式
	SetMode(uint)

	// 设置buffersize
	SetBuffer(int)

	// 获取模式
	GetMode() uint

	// 获取buffersize
	GetBuffer() int

	// 创建文件
	CreateFile(string, io.Reader) error

	// 打开文件
	OpenFile(string) (io.Reader, error)

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

// todo: 设置mode
func (executor *NormalExecutor) SetMode(mode uint) {
	executor.mode = mode
}

// todo: 获取mode
func (executor *NormalExecutor) GetMode() uint {
	return executor.mode
}

// todo: 设置buffer
func (executor *NormalExecutor) SetBuffer(size int) {
	executor.buffersize = size
}

// todo: 获取buffer
func (executor *NormalExecutor) GetBuffer() int {
	return executor.buffersize
}

// todo: 创建文件
func (executor *NormalExecutor) CreateFile(path string, source io.Reader) error {
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

// todo: 打开文件
// 将文件句柄返回给gin进行下载
//
// 因此文件的句柄需要手动进行关闭
func (executor *NormalExecutor) OpenFile(path string) (io.Reader, error) {
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
