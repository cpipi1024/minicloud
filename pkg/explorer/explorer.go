package explorer

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

// 文件抽象
type File struct {
	Name         string `json:"name"`         // 文件名
	IsDirectory  bool   `json:"isDirectory"`  // 是否是文件夹
	Size         int64  `json:"size"`         // 大小
	Extension    string `json:"extension"`    // 扩展名
	Mime         string `json:"mime"`         // 媒体类型
	Path         string `json:"path"`         // 相对路径(相对于storagePath)
	LastModified int64  `json:"lastModified"` // 最后修改时间 Unix Timestamp seconds
}

// 文件管理器
type Explorer struct {
	rootDir    string   // 管理根目录
	currentDir string   // 当前目录
	executor   Executor // 文件操作执行器
}

// todo: 返回管理器实例
func NewDefaultExplorer(root string) *Explorer {

	return &Explorer{
		rootDir:    root,
		currentDir: root,
		executor:   NewNormalExecutor(),
	}
}

// todo: 进入下一级目录
func (exp *Explorer) EnterDir(dirname string) error {
	newpath := filepath.Join(exp.currentDir, dirname)

	if !dirExist(newpath) {
		return errors.New("文件夹不存在")
	}

	exp.SetCurrentDir(newpath)

	return nil
}

//todo: 创建目录
func (exp *Explorer) CreateDir(dirname string) error {
	newpath := filepath.Join(exp.currentDir, dirname)

	if dirExist(newpath) {
		return errors.New("目录已存在")
	}

	err := os.Mkdir(newpath, 0775)

	if err != nil {
		return err
	}

	return nil
}

// todo: 删除目录
func (exp *Explorer) DeleteDir(dirname string) error {
	newpath := filepath.Join(exp.currentDir, dirname)

	if !dirExist(newpath) {
		return errors.New("目录不存在")
	}

	err := os.RemoveAll(newpath)

	if err != nil {
		return err
	}

	return nil
}

// todo: 删除文件
func (exp *Explorer) DeleteFile(filename string) error {
	newpath := filepath.Join(exp.currentDir, filename)

	err := os.Remove(newpath)

	if err != nil {
		return err
	}

	return nil
}

// todo: 返回当前目录
func (exp *Explorer) GetCurrentDir() string {
	return exp.currentDir
}

// todo: 设置当前目录
func (exp *Explorer) SetCurrentDir(path string) {
	exp.currentDir = path
}

// todo: 返回根目录
func (exp *Explorer) GetRootDir() string {
	return exp.rootDir
}

// todo: 列出当前目录内容
func (exp *Explorer) ListContents() ([]*File, error) {
	// 返回 []direntity
	files, err := os.ReadDir(exp.currentDir)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return []*File{}, nil
	}

	ret := []*File{}

	for i := 0; i < len(files); i++ {
		// 获取fileinfo
		fileinfo, _ := files[i].Info()

		file := exp.normalizeInfo(fileinfo)

		if file != nil {
			ret = append(ret, file)
		}
	}

	return ret, nil
}

// todo: 规范化数据
func (exp *Explorer) normalizeInfo(info fs.FileInfo) *File {

	name := info.Name()

	// 过滤
	if name == "." || name == ".." {
		return nil
	}

	afterSplit := strings.Split(name, ".")

	// 获取文件扩展名
	extension := afterSplit[len(afterSplit)-1]

	filepath := filepath.Join(exp.currentDir, name)

	// 获取mimetype
	mime, _ := mimetype.DetectFile(filepath)

	file := File{
		Name:         name,
		IsDirectory:  info.IsDir(),
		Size:         info.Size(),
		Path:         exp.currentDir,
		Extension:    extension,
		Mime:         mime.String(),
		LastModified: info.ModTime().Unix(),
	}

	return &file
}

// todo: 返回执行器
func (exp *Explorer) GetExecutor() Executor {
	return exp.executor
}

// todo: 配置执行器
func (exp *Explorer) SetExecutor(executor Executor) {
	exp.executor = executor
}

// todo: 流上传
func (exp *Explorer) StreamUpload(path string, source io.Reader) error {
	return exp.GetExecutor().StreamUpload(path, source)
}

// todo: 流下载
func (exp *Explorer) StreamDownload(path string) (io.Reader, error) {
	return exp.GetExecutor().StreamDownload(path)
}

// todo: 获取文件信息
func (exp *Explorer) GetFileInfo(path string) (*File, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	file := exp.normalizeInfo(info)

	return file, nil
}

// todo: 判断目录是否存在
func dirExist(dirPath string) bool {
	_, err := os.Stat(dirPath)

	return err == nil || os.IsExist(err)
}
