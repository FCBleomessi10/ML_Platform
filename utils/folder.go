package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Mkdir 定义一个创建文件目录的方法
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	//使用mkdir会创建多层级目录
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

//_, err := os.Stat(path)

func Filenums(basePath string) int {
	files, _ := ioutil.ReadDir(basePath)
	return len(files)
}
