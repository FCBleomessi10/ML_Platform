package models

import "time"

type (
	Labelset struct {
		LabelsetId   uint      `json:"labelset_id"`
		LabelsetName string    `json:"labelset_name"`
		CreateTime   time.Time `json:"create_time"`
		UpdateTime   time.Time `json:"update_time"`
		IsDeleted    int       `json:"is_deleted"`
	}

	LabelsetRes struct {
		LabelsetId   uint   `json:"labelset_id"`
		LabelsetName string `json:"labelset_name"`
	}
)
