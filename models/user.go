package models

import "time"

type (
	User struct {
		UserId     uint      `json:"user_id"`
		UserName   string    `json:"user_name"`
		Phone      string    `json:"phone"`
		Password   string    `json:"password"`
		Salt       string    `json:"salt"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
		IsDeleted  int       `json:"is_deleted"`
	}

	UserRes struct {
		UserId   uint   `json:"user_id"`
		UserName string `json:"user_name"`
		Phone    string `json:"phone"`
	}

	UserLoginReq struct {
		Username string `json:"user_name"`
		Password string `json:"password"`
	}
)
