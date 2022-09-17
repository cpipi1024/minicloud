package explorer

import "os"

// 文件抽象
type File struct {
	Name         string `json:"name"`         // 文件名
	IsDirectory  bool   `json:"isDirectory"`  // 是否是文件夹
	Size         int    `json:"size"`         // 大小
	Extension    string `json:"extension"`    // 扩展名
	Mime         string `json:"mime"`         // 媒体类型
	Path         string `json:"path"`         // 相对路径(相对于storagePath)
	LastModified int64  `json:"lastModified"` // 最后修改时间 Unix Timestamp seconds
}

// 文件管理器
type Explorer struct {
	rootDir    string   // 管理根目录
	CurrentDir string   // 当前目录
	executor   Executor // 文件操作执行器
}

// todo: 返回管理器实例
func NewDefaultExplorer(root string) *Explorer {

}

// todo: 进入下一级目录
func (exp *Explorer) EnterDir(dirname string) error {

}

// todo: 列出当前目录内容
func (exp *Explorer) ListContents() []*File {

}

// todo: 返回执行器
func (exp *Explorer) GetExecutor() Executor {
	return exp.executor
}

// todo: 配置执行器
func (exp *Explorer) SetExecutor(executor Executor) {
	exp.executor = executor
}

// todo: 判断目录是否存在
func dirExist(dirPath string) bool {
	_, err := os.Stat(dirPath)

	return err == nil || os.IsExist(err)
}
