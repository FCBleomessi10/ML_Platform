package models

import "time"

type (
	Sample struct {
		SampleId   uint      `json:"sample_id"`
		DatasetId  uint      `json:"dataset_id"`
		ImgPath    string    `json:"img_path"`
		ImgName    string    `json:"img_name"`
		Status     int       `json:"status"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
		IsDeleted  int       `json:"is_deleted"`
	}

	SampleRes struct {
		SampleId uint `json:"sample_id"`
		//DatasetId  uint      `json:"dataset_id"`
		ImgPath string `json:"img_path"`
	}
)
