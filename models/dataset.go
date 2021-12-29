package models

import "time"

type (
	Dataset struct {
		DatasetId   uint      `json:"dataset_id"`
		LabelsetId  uint      `json:"labelset_id"`
		DatasetName string    `json:"dataset_name"`
		Status      int       `json:"status"`
		CreateTime  time.Time `json:"create_time"`
		UpdateTime  time.Time `json:"update_time"`
		IsDeleted   int       `json:"is_deleted"`
	}
	DatasetRes struct {
		DatasetId   uint   `json:"dataset_id"`
		LabelsetId  uint   `json:"labelset_id"`
		DatasetName string `json:"dataset_name"`
		Status      int    `json:"status"`
	}
)
