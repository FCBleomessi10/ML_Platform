package controllers

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func CreateSample(c *gin.Context) { //添加样本时创建文件夹

	//参数提取 dataset_id
	_datasetId := c.Query("dataset_id")
	datasetId, _ := strconv.Atoi(c.Query("dataset_id"))

	//使用 dataset_id 创建文件夹
	var path string = "/Users/sunzhiqiang/Desktop"
	path = path + "/" + c.Query("dataset_id")

	ok, _ := utils.PathExists(path)
	if ok == false { //创建对应的文件夹
		os.MkdirAll(path, os.ModePerm)
	}

	//参数提取 文件
	form, _ := c.MultipartForm()
	files := form.File["file"]

	//查询本文件夹已有数据量
	basepath := fmt.Sprintf("/Users/sunzhiqiang/Desktop/%s", _datasetId)
	num := utils.Filenums(basepath)
	fileprefix := ".jpg"

	//文件录入
	for _, file := range files {
		dst := fmt.Sprintf("/Users/sunzhiqiang/Desktop/%s/%s%s", _datasetId, strconv.Itoa(num), fileprefix)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		sample := models.Sample{
			DatasetId:  uint(datasetId),
			ImgPath:    dst,
			ImgName:    file.Filename,
			IsDeleted:  0,
			CreateTime: utils.TimeStringToGoTime(),
			UpdateTime: utils.TimeStringToGoTime(),
			Status:     0,
		}
		db.Save(&sample)
		num++
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
		"nums":    num,
	})
}

func FetchAllSample(c *gin.Context) {
	var sample []models.Sample
	var _sample []models.Sample
	db.Find(&sample)

	if len(sample) <= 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No sample found!",
			})
		return
	}

	for _, item := range sample {
		_sample = append(_sample,
			models.Sample{
				SampleId:   item.SampleId,
				DatasetId:  item.DatasetId,
				ImgName:    item.ImgName,
				ImgPath:    item.ImgPath,
				Status:     item.Status,
				CreateTime: item.CreateTime,
				UpdateTime: item.UpdateTime,
				IsDeleted:  item.IsDeleted,
			})
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _sample,
		})
}

//通过sample id 查到所有的   annotation

func FetchSample(c *gin.Context) {

	var sample models.Sample
	var annotations []models.Annotation
	var _annotationRes []models.AnnotationRes

	sampleId := c.Param("id")
	if sample.SampleId == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No sample found!",
			})
		return
	}

	db.First(&sample, "sample_id=?", sampleId)
	db.Find(&annotations, "sample_id=?", sampleId)

	res := models.Sample{
		SampleId:   sample.SampleId,
		DatasetId:  sample.DatasetId,
		ImgPath:    sample.ImgPath,
		ImgName:    sample.ImgName,
		Status:     sample.Status,
		UpdateTime: sample.UpdateTime,
		IsDeleted:  sample.IsDeleted,
	}

	for _, item := range annotations {
		if item.IsDeleted == 0 {
			_annotationRes = append(_annotationRes,
				models.AnnotationRes{
					LabelId:      item.LabelId,
					X1:           item.X1,
					X2:           item.X2,
					Y1:           item.Y1,
					Y2:           item.Y2,
					AnnotationId: item.AnnotationId,
				})
		}
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status":      http.StatusOK,
			"data":        res,
			"annotations": _annotationRes,
		})

	c.File(sample.ImgPath)
}

func UpdateSample(c *gin.Context) {
	var sample models.Sample
	sampleId := c.PostForm("id")
	db.First(&sample, sampleId)

	if sample.SampleId == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No sample found!",
			})
		return
	}

	db.Model(&sample).Update("dataset_id", c.PostForm("dataset_id"))
	db.Model(&sample).Update("status", c.PostForm("status"))
	db.Model(&sample).Update("update_time", utils.TimeStringToGoTime())

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Sample update successfully!",
		})
}

func DeleteSample(c *gin.Context) {
	var sample models.Sample
	sampleId := c.PostForm("id")
	db.First(&sample, sampleId)

	if sample.SampleId == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No sample found!",
			})
		return
	}

	db.Model(&sample).Update("is_deleted", 0)
	db.Model(&sample).Update("update_time", utils.TimeStringToGoTime())

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Sample delete successfully!",
		})
}
