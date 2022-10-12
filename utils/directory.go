package utils

import (
	"errors"
	"os"
)

// 判断文件目录是否存在
func PathExists(path string) (bool, error) {
	// os 判断文件是否存在的方法
	fi, err := os.Stat(path)
	// 如果文件存在
	if err == nil {
		// 文件是否为文件夹
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	// 根据错误判断是不是文件不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
