package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateModel(c *gin.Context) {
	modelType := c.Query("model_type")
	modelPath := c.Query("model_path")
	modelName := c.Query("model_name")
	datasetId, _ := strconv.Atoi(c.Query("dataset_id"))
	version := c.Query("version")
	if modelType == "" || modelName == "" || version == "" {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Model created failed!",
			})
		return
	}

	model := models.Model{
		ModelType:   modelType,
		ModelPath:   modelPath,
		ModelName:   modelName,
		DatasetId:   uint(datasetId),
		Version:     version,
		TrainStatus: 0,
		CreateTime:  utils.TimeStringToGoTime(),
		UpdateTime:  utils.TimeStringToGoTime(),
		IsDeleted:   0,
	}
	db.Save(&model)

	c.JSON(http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Model created successfully!",
		})
}

func FetchAllModel(c *gin.Context) {
	var model []models.Model
	var _modelRes []models.ModelRes
	db.Find(&model)

	for _, item := range model {
		if item.IsDeleted == 0 {
			_modelRes = append(_modelRes,
				models.ModelRes{
					ModelId:     item.ModelId,
					ModelType:   item.ModelType,
					ModelPath:   item.ModelPath,
					ModelName:   item.ModelName,
					DatasetId:   item.DatasetId,
					Version:     item.Version,
					TrainStatus: item.TrainStatus,
				})
		}
	}
	if _modelRes == nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No model found!",
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _modelRes,
		})
}

func FetchModel(c *gin.Context) {
	var model models.Model
	modelId := c.Param("id")
	db.First(&model, "model_id=?", modelId)

	if model.ModelId == 0 || model.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No model found!",
			})
		return
	}

	res := models.ModelRes{
		ModelId:     model.ModelId,
		ModelType:   model.ModelType,
		ModelPath:   model.ModelPath,
		ModelName:   model.ModelName,
		DatasetId:   model.DatasetId,
		Version:     model.Version,
		TrainStatus: model.TrainStatus,
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   res,
		})
}

func UpdateModel(c *gin.Context) {
	var model models.Model
	modelId := c.Param("id")
	db.First(&model, "model_id=?", modelId)

	if model.ModelId == 0 || model.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No model found!",
			})
		return
	}

	datasetId, _ := strconv.Atoi(c.DefaultQuery("dataset_id", strconv.Itoa(int(model.DatasetId))))
	trainStatus, _ := strconv.Atoi(c.DefaultQuery("train_status", strconv.Itoa(model.TrainStatus)))

	db.Model(&model).Where("model_id=?", model.ModelId).Updates(models.Model{
		ModelType:   c.DefaultQuery("model_type", model.ModelType),
		ModelPath:   c.DefaultQuery("model_path", model.ModelPath),
		ModelName:   c.DefaultQuery("model_name", model.ModelName),
		DatasetId:   uint(datasetId),
		Version:     c.DefaultQuery("version", model.Version),
		TrainStatus: trainStatus,
		UpdateTime:  utils.TimeStringToGoTime(),
	})
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Update model successfully!",
		})
}

func DeleteModel(c *gin.Context) {
	var model models.Model
	modelId := c.Param("id")
	db.First(&model, "model_id=?", modelId)

	if model.ModelId == 0 || model.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No model found!",
			})
		return
	}

	db.Model(&model).Where("model_id=?", model.ModelId).Updates(models.Model{
		IsDeleted:  1,
		UpdateTime: utils.TimeStringToGoTime(),
	})
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Model delete successfully!",
		})
}
