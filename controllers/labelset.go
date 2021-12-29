package controllers

import (
	"ML_Platform/backend/models"
	"ML_Platform/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateLabelset(c *gin.Context) {
	labelsetName := c.Query("labelset_name")
	labelset := models.Labelset{
		LabelsetName: labelsetName,
		CreateTime:   utils.TimeStringToGoTime(),
		UpdateTime:   utils.TimeStringToGoTime(),
		IsDeleted:    0,
	}
	if labelsetName == "" {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Labelset created failed!",
			})
		return
	}

	db.Save(&labelset)

	c.JSON(http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Labelset created successfully!",
		})
}

func FetchAllLabelset(c *gin.Context) {
	var labelset []models.Labelset
	var _labelsetRes []models.LabelsetRes
	db.Find(&labelset)

	for _, item := range labelset {
		if item.IsDeleted == 0 {
			_labelsetRes = append(_labelsetRes,
				models.LabelsetRes{
					LabelsetId:   item.LabelsetId,
					LabelsetName: item.LabelsetName,
				})
		}
	}
	if _labelsetRes == nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No labelset found!",
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _labelsetRes,
		})
}

func FetchLabelset(c *gin.Context) {
	var labelset models.Labelset
	var labels []models.Label
	var _labelRes []string

	labelsetId := c.Param("id")
	db.First(&labelset, "labelset_id=?", labelsetId)
	db.Find(&labels, "labelset_id=?", labelsetId)

	if labelset.LabelsetId == 0 || labelset.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No labelset found!",
			})
		return
	}

	res := models.LabelsetRes{
		LabelsetId:   labelset.LabelsetId,
		LabelsetName: labelset.LabelsetName,
	}
	for _, item := range labels {
		if item.IsDeleted == 0 {
			_labelRes = append(_labelRes, item.LabelName)
		}
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   res,
			"labels": _labelRes,
		})
}

func UpdateLabelset(c *gin.Context) {
	var labelset models.Labelset
	labelsetId := c.Param("id")
	db.First(&labelset, "labelset_id=?", labelsetId)

	if labelset.LabelsetId == 0 || labelset.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No labelset found!",
			})
		return
	}

	db.Model(&labelset).Where("labelset_id=?", labelset.LabelsetId).Updates(models.Labelset{
		LabelsetName: c.DefaultQuery("labelset_name", labelset.LabelsetName),
		UpdateTime:   utils.TimeStringToGoTime(),
	})

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Labelset update successfully!",
		})
}

func DeleteLabelset(c *gin.Context) {
	var labelset models.Labelset
	labelsetId := c.Param("id")
	db.First(&labelset, "labelset_id=?", labelsetId)

	if labelset.LabelsetId == 0 || labelset.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No labelset found!",
			})
		return
	}
	db.Model(&labelset).Where("labelset_id=?", labelset.LabelsetId).Updates(models.Labelset{
		IsDeleted:  1,
		UpdateTime: utils.TimeStringToGoTime(),
	})
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Labelset delete successfully!",
		})
}
