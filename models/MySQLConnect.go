package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var DB, err = gorm.Open("mysql",
	"root:123456@(47.101.150.22:3306)/deeppumadl?charset=utf8&parseTime=True&loc=Local")

func init() {
	if err != nil {
		fmt.Println("连接MySQL数据库失败")
	} else {
		fmt.Println("连接MySQL数据库成功")
	}
}
