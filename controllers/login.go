package controllers

import (
	"backend/conf"
	"backend/models"
	"backend/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	//user_name := c.PostForm("user_name")
	//password := c.PostForm("password")

	var userloginreq = models.UserLoginReq{}
	err := c.BindJSON(&userloginreq)

	if err != nil {
		c.JSON(400,
			gin.H{
				"status":  400,
				"message": "Password or user_name is empty",
			})
		return
	}

	var user models.User
	db.First(&user, "user_name=?", userloginreq.Username)

	if user.Password == userloginreq.Password {
		session := sessions.Default(c)
		var data = make(map[string]interface{}, 0)
		v := session.Get("token")

		if v == nil {
			cur := time.Now()
			timestamps := cur.UnixNano()
			times := strconv.FormatInt(timestamps, 10)
			v = utils.Md5En(utils.GetRandomString(16) + times)
			session.Set(conf.Cfg.Token, v)
			session.Set(v, user.UserId)
			err = session.Save()
		}
		data["token"] = v

		c.JSON(http.StatusOK,
			gin.H{
				"status":  http.StatusOK,
				"message": "Login success",
				"token":   v,
			})
	} else {
		c.JSON(400,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "No user found or Combination error",
			})
	}
}
