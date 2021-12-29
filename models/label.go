package models

import "time"

type (
	Label struct {
		LabelId    uint      `json:"label_id"`
		LabelName  string    `json:"label_name"`
		LabelsetId uint      `json:"labelset_id"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
		IsDeleted  int       `json:"is_deleted"`
	}
	LabelRes struct {
		LabelId    uint   `json:"label_id"`
		LabelName  string `json:"label_name"`
		LabelsetId uint   `json:"labelset_id"`
	}
)
