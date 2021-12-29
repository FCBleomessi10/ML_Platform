package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	//"path/filepath"
	//"strconv"
	//"time"
)

func TestFiles(t *testing.T) {
	router := gin.Default()
	router.POST("/sample", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		//1.获取文件
		files := form.File["file"]
		//2.循环全部的文件
		for _, file := range files {
			// 3.根据时间鹾生成文件名
			//fileNameInt := time.Now().Unix()
			//fileNameStr := strconv.FormatInt(fileNameInt,10)
			//4.新的文件名(如果是同时上传多张图片的时候就会同名，因此这里使用时间鹾加文件名方式)
			//fileName := fileNameStr + file.Filename
			//5.保存上传文件
			filePath := fmt.Sprintf("/Users/sunzhiqiang/Desktop/%s", file.Filename)
			c.SaveUploadedFile(file, filePath)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "上传成功",
		})
	})
	router.Run(":8080")
}
