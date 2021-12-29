package models

import "time"

type (
	Annotation struct {
		AnnotationId uint      `json:"annotation_id"`
		LabelId      uint      `json:"label_id"`
		SampleId     uint      `json:"sample_id"`
		X1           int       `json:"x1"`
		Y1           int       `json:"y1"`
		X2           int       `json:"x2"`
		Y2           int       `json:"y2"`
		CreateTime   time.Time `json:"create_time"`
		UpdateTime   time.Time `json:"update_time"`
		IsDeleted    int       `json:"is_deleted"`
	}

	AnnotationRes struct {
		AnnotationId uint `json:"annotation_id"`
		X1           int  `json:"x1"`
		Y1           int  `json:"y1"`
		X2           int  `json:"x2"`
		Y2           int  `json:"y2"`
		LabelId      uint `json:"label_id"`
	}
)
