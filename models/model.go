package models

import "time"

type (
	Model struct {
		ModelId     uint      `json:"model_id"`
		ModelType   string    `json:"model_type"`
		ModelPath   string    `json:"model_path"`
		ModelName   string    `json:"model_name"`
		DatasetId   uint      `json:"dataset_id"`
		Version     string    `json:"version"`
		TrainStatus int       `json:"train_status"`
		CreateTime  time.Time `json:"create_time"`
		UpdateTime  time.Time `json:"update_time"`
		IsDeleted   int       `json:"is_deleted"`
	}
	ModelRes struct {
		ModelId     uint   `json:"model_id"`
		ModelType   string `json:"model_type"`
		ModelPath   string `json:"model_path"`
		ModelName   string `json:"model_name"`
		DatasetId   uint   `json:"dataset_id"`
		Version     string `json:"version"`
		TrainStatus int    `json:"train_status"`
		LabelId     uint   `json:"label_id"`
	}
)
