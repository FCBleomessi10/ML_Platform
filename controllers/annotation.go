package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAnnotation(c *gin.Context) {
	//AnnotationId, _ := strconv.Atoi(c.Query("annotation_id"))
	LabelId, _ := strconv.Atoi(c.Query("label_id"))
	SampleId, _ := strconv.Atoi(c.Query("sample_id"))
	x1, _ := strconv.Atoi(c.Query("x1"))
	x2, _ := strconv.Atoi(c.Query("x2"))
	y1, _ := strconv.Atoi(c.Query("y1"))
	y2, _ := strconv.Atoi(c.Query("y2"))

	if LabelId == 0 || SampleId == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Annotation created failed!,Param missing or wrong",
			})
		return
	}

	annotation := models.Annotation{
		//AnnotationId: uint(AnnotationId),
		LabelId:    uint(LabelId),
		SampleId:   uint(SampleId),
		X1:         x1,
		X2:         x2,
		Y1:         y1,
		Y2:         y2,
		CreateTime: utils.TimeStringToGoTime(),
		UpdateTime: utils.TimeStringToGoTime(),
		IsDeleted:  0,
	}

	db.Save(&annotation)

	c.JSON(http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Annotation created successfully!",
			"ID":      annotation.AnnotationId,
		})
}

func FetchAllAnnotation(c *gin.Context) {
	var annotation []models.Annotation
	var _annotation []models.Annotation
	db.Find(&annotation)

	//if len(annotation) <= 0 {
	//	c.JSON(http.StatusNotFound,
	//		gin.H{
	//			"status":  http.StatusNotFound,
	//			"message": "No annotation found!",
	//		})
	//	return
	//}

	for _, item := range annotation {
		if item.IsDeleted == 0 {
			_annotation = append(_annotation,
				models.Annotation{
					AnnotationId: item.AnnotationId,
					LabelId:      item.LabelId,
					SampleId:     item.SampleId,
					X1:           item.X1,
					X2:           item.X2,
					Y1:           item.Y1,
					Y2:           item.Y2,
					CreateTime:   item.CreateTime,
					UpdateTime:   item.UpdateTime,
					IsDeleted:    item.IsDeleted,
				})
		}
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _annotation,
		})
}

func FetchAnnotation(c *gin.Context) {
	var annotation models.Annotation
	AnnotationId := c.Param("id")

	db.First(&annotation, "annotation_id=?", AnnotationId)

	if annotation.AnnotationId == 0 || annotation.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No annotation found!",
			})
		return
	}

	res := models.Annotation{
		AnnotationId: annotation.AnnotationId,
		LabelId:      annotation.LabelId,
		SampleId:     annotation.SampleId,
		X1:           annotation.X1,
		X2:           annotation.X2,
		Y1:           annotation.Y1,
		Y2:           annotation.Y2,
		CreateTime:   annotation.CreateTime,
		UpdateTime:   annotation.UpdateTime,
		IsDeleted:    annotation.IsDeleted,
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   res,
		})

}

func UpdateAnnotation(c *gin.Context) {
	var annotation models.Annotation

	AnnotationId := c.Param("id")

	db.First(&annotation, "annotation_id=?", AnnotationId)

	if annotation.AnnotationId == 0 || annotation.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No model found!",
			})
		return
	}

	X1, _ := strconv.Atoi(c.DefaultQuery("x1", strconv.Itoa(int(annotation.X1))))
	X2, _ := strconv.Atoi(c.DefaultQuery("x2", strconv.Itoa(int(annotation.X2))))
	Y1, _ := strconv.Atoi(c.DefaultQuery("y1", strconv.Itoa(int(annotation.Y1))))
	Y2, _ := strconv.Atoi(c.DefaultQuery("y2", strconv.Itoa(int(annotation.Y2))))

	db.Model(&annotation).Where("annotation_id=?", annotation.AnnotationId).Update(models.Annotation{
		UpdateTime: utils.TimeStringToGoTime(),
		IsDeleted:  0,
		X1:         int(X1),
		X2:         int(X2),
		Y1:         int(Y1),
		Y2:         int(Y2),
		SampleId:   annotation.SampleId,
	})

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Update model successfully!",
		})

}

func DeleteAnnotation(c *gin.Context) {
	var annotation models.Annotation //也是结构体类型

	db.Model(&annotation).Where("annotation_id=?", c.Param("id")).Update(models.Annotation{
		IsDeleted:  1,
		UpdateTime: utils.TimeStringToGoTime(),
	})

	//OUTDATED
	//AnnotationId := c.Param("id")
	//db.First(&annotation, "annotation_id=?", AnnotationId)
	//db.Model(&annotation).Where("annotation_id=?", annotation.AnnotationId).Update(models.Annotation{
	//	IsDeleted:  1,
	//	UpdateTime: common.TimeStringToGoTime(

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Model delete successfully!",
		})
}
