package controllers

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = models.DB

func CreateUser(c *gin.Context) {
	userName := c.Query("user_name")
	phone := c.Query("phone")
	password := c.Query("password")
	salt := c.Query("salt")

	if userName == "" || phone == "" || password == "" {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "User created failed!",
			})
		return
	}

	user := models.User{
		UserName:   userName,
		Phone:      phone,
		Password:   password,
		Salt:       salt,
		CreateTime: utils.TimeStringToGoTime(),
		UpdateTime: utils.TimeStringToGoTime(),
		IsDeleted:  0,
	}
	db.Save(&user)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "User created successfully!",
		})
}

func FetchAllUser(c *gin.Context) {
	var user []models.User
	var _userRes []models.UserRes
	db.Find(&user)

	for _, item := range user {
		if item.IsDeleted == 0 {
			_userRes = append(_userRes,
				models.UserRes{
					UserId:   item.UserId,
					UserName: item.UserName,
					Phone:    item.Phone,
				})
		}
	}

	if _userRes == nil {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!",
			})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _userRes,
		})
}

func FetchUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")

	db.First(&user, "user_id=?", userId)

	if user.UserId == 0 || user.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!",
			})
		return
	}

	res := models.UserRes{
		UserId:   user.UserId,
		UserName: user.UserName,
		Phone:    user.Phone,
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   res,
		})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")

	db.First(&user, "user_id=?", userId)

	if user.UserId == 0 || user.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!",
			})
		return
	}

	db.Model(&user).Where("user_id=?", user.UserId).Updates(models.User{
		UserName:   c.DefaultQuery("user_name", user.UserName),
		Phone:      c.DefaultQuery("phone", user.Phone),
		UpdateTime: utils.TimeStringToGoTime(),
	})

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Update User successfully!",
		})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	db.First(&user, "user_id=?", userId)

	if user.UserId == 0 || user.IsDeleted == 1 {
		c.JSON(http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No user found!",
			})
		return
	}

	db.Model(&user).Where("user_id=?", user.UserId).Updates(models.User{
		IsDeleted:  1,
		UpdateTime: utils.TimeStringToGoTime(),
	})

	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "User deleted successfully!",
		})
}
