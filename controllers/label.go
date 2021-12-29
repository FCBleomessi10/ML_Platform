package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateLabel(c *gin.Context) {
	labelName := c.Query("label_name")
	labelsetId, _ := strconv.Atoi(c.Query("labelset_id"))
	label := models.Label{
		LabelName:  labelName,
		LabelsetId: uint(labelsetId),
		CreateTime: utils.TimeStringToGoTime(),
		UpdateTime: utils.TimeStringToGoTime(),
		IsDeleted:  0,
	}
	if labelName == "" || labelsetId == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Label created failed!",
			})
		return
	}
	db.Save(&label)

	c.JSON(http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Label created successfully!",
		})
}

func FetchAllLabel(c *gin.Context) {
	var label []models.Label
	var _labelRes []models.LabelRes
	db.Find(&label)

	for _, item := range label {
		if item.IsDeleted == 0 {
			_labelRes = append(_labelRes,
				models.LabelRes{
					LabelId:    item.LabelId,
					LabelName:  item.LabelName,
					LabelsetId: item.LabelsetId,
				})
		}
	}
	if _labelRes == nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No label found!",
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _labelRes,
		})
}

func FetchLabel(c *gin.Context) {
	var label models.Label
	labelId := c.Param("id")
	db.First(&label, "label_id=?", labelId)

	if label.LabelId == 0 || label.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No label found!",
			})
		return
	}

	res := models.LabelRes{
		LabelId:    label.LabelId,
		LabelName:  label.LabelName,
		LabelsetId: label.LabelsetId,
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   res,
		})
}

func UpdateLabel(c *gin.Context) {
	var label models.Label
	labelId := c.Param("id")
	db.First(&label, "label_id=?", labelId)

	if label.LabelId == 0 || label.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No label found!",
			})
		return
	}

	labelsetId, _ := strconv.Atoi(c.DefaultQuery("labelset_id", string(label.LabelsetId)))
	db.Model(&label).Where("label_id=?", label.LabelId).Updates(models.Label{
		LabelName:  c.DefaultQuery("label_name", label.LabelName),
		LabelsetId: uint(labelsetId),
		UpdateTime: utils.TimeStringToGoTime(),
	})
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Update label successfully!",
		})
}

func DeleteLabel(c *gin.Context) {
	var label models.Label
	labelId := c.Param("id")
	db.First(&label, "label_id=?", labelId)

	if label.LabelId == 0 || label.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No label found!",
			})
		return
	}

	db.Model(&label).Where("label_id=?", label.LabelId).Updates(models.Label{
		IsDeleted:  1,
		UpdateTime: utils.TimeStringToGoTime(),
	})
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Delete label successfully!",
		})
}
