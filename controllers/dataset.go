package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateDataset(c *gin.Context) {

	labelsetId, _ := strconv.Atoi(c.Query("labelset_id"))

	datasetName := c.Query("dataset_name")

	dataset := models.Dataset{
		LabelsetId:  uint(labelsetId),
		DatasetName: datasetName,
		Status:      0,
		CreateTime:  utils.TimeStringToGoTime(),
		UpdateTime:  utils.TimeStringToGoTime(),
		IsDeleted:   0,
	}
	db.Save(&dataset)

	c.JSON(http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Dataset created successfully!",
		})
}

func FetchAllDataset(c *gin.Context) {
	var dataset []models.Dataset
	var _datasetRes []models.DatasetRes
	db.Find(&dataset)

	for _, item := range dataset {
		if item.IsDeleted == 0 {
			_datasetRes = append(_datasetRes,
				models.DatasetRes{
					DatasetId:   item.DatasetId,
					LabelsetId:  item.LabelsetId,
					DatasetName: item.DatasetName,
					Status:      item.Status,
				})
		}
	}

	if _datasetRes == nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No dataset found!",
			})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _datasetRes,
		})
}

//通过数据集id 查到本数据集所有图片（sample）
func FetchDataset(c *gin.Context) {
	var dataset models.Dataset
	var samples []models.Sample        //用于存放找到的东西
	var _samplesRes []models.SampleRes //用于存放回复的结构体数组

	datasetId := c.Param("id")

	db.First(&dataset, "dataset_id=?", datasetId)
	db.Find(&samples, "dataset_id=?", datasetId)

	if dataset.DatasetId == 0 || dataset.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No dataset found!",
			})
		return
	} //判断条件

	for _, item := range samples {
		if item.IsDeleted != 1 {
			_samplesRes = append(_samplesRes,
				models.SampleRes{
					SampleId: item.SampleId,
					ImgPath:  item.ImgPath,
				})
		}
	}

	res := models.DatasetRes{
		DatasetId:   dataset.DatasetId,
		LabelsetId:  dataset.LabelsetId,
		DatasetName: dataset.DatasetName,
		Status:      dataset.Status,
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   res,
			"labels": _samplesRes,
		})
}

func UpdateDataset(c *gin.Context) {
	var dataset models.Dataset
	datasetId := c.Param("id")
	db.First(&dataset, "dataset_id=?", datasetId)

	if dataset.DatasetId == 0 || dataset.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No dataset found!",
			})
		return
	}

	labelsetId, _ := strconv.Atoi(c.DefaultQuery("labelset_id", strconv.Itoa(int(dataset.LabelsetId))))
	status, _ := strconv.Atoi(c.DefaultQuery("status", strconv.Itoa(dataset.Status)))

	db.Model(&dataset).Where("dataset_id=?", dataset.DatasetId).Updates(models.Dataset{
		LabelsetId:  uint(labelsetId),
		DatasetName: c.DefaultQuery("dataset_name", dataset.DatasetName),
		Status:      status,
		UpdateTime:  utils.TimeStringToGoTime(),
	})

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Update dataset successfully!",
		})
}

func DeleteDataset(c *gin.Context) {
	var dataset models.Dataset
	datasetId := c.Param("id")
	db.First(&dataset, "dataset_id=?", datasetId)

	if dataset.DatasetId == 0 || dataset.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No dataset found!",
			})
		return
	}

	db.Model(&dataset).Where("dataset_id=?", dataset.DatasetId).Updates(models.Dataset{
		IsDeleted:  1,
		UpdateTime: utils.TimeStringToGoTime(),
	})
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Delete dataset successfully!",
		})
}
