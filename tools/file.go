// 文件相关操作

package tools

import "os"

// todo: 检查文件是否存在
func FileExists(path string) bool {

	_, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	}

	return true
}
