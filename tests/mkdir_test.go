package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMkdirandPathExists(t *testing.T) {
	// filename := "device/sdk/CMakeLists.txt"
	//    filenameall := path.Base(filename)
	//    filesuffix := path.Ext(filename)
	//    fileprefix := filenameall[0:len(filenameall) - len(filesuffix)]
	//    //fileprefix, err := strings.TrimSuffix(filenameall, filesuffix)
	//
	//    fmt.Println("file name:", filenameall)
	//    fmt.Println("file prefix:", fileprefix)
	//    fmt.Println("file suffix:", filesuffix) 关于文件头尾

	//err:=os.Mkdir("/Users/sunzhiqiang/Desktop/1",os.ModePerm)
	//Mkdirnew("/Users/sunzhiqiang/Desktop/1")
	//
	//if err==nil{
	//	fmt.Println("success created")
	//}else{
	//	fmt.Println(err)
	//}
	Filenum("/Users/sunzhiqiang/Desktop/1")
	//ok1,_:=PathExists("/Users/sunzhiqiang/Desktop")//正确写法
	//ok2:=PathExists_("/Users/sunzhiqiang/Desktop") //错误写法
	//fmt.Println(ok1)
	//fmt.Println(ok2)
	//path,err:=os.Open("file_test.go")
	//fmt.Println(path,err)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

// 判断所给路径文件/文件夹是否存在
func PathExists_(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Mkdirnew(basePath string) string { //在指定的文件夹内创建文件
	//获取当前时间,并且格式化时间
	//folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, "4")

	//使用mkdir会创建多层级目录
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}

func Filenum(basePath string) {
	files, _ := ioutil.ReadDir(basePath)
	var count int
	for _, f := range files {
		fmt.Println(f.Name())
		count++
	}
	fmt.Println(len(files))
}
