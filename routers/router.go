package routers

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Run() {

	router := gin.Default()
	router.Use(cors.New(GetCorsConfig()))
	store, _ := redis.NewStoreWithPool(middlewares.RedisClient, []byte("secret"))
	router.Use(sessions.Sessions("gosession", store))

	//登陆
	router.POST("/api/login", controllers.Login)
	router.POST("/login", controllers.Login)

	user := router.Group("/user")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.FetchAllUser)
		user.GET("/:id", controllers.FetchUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
	model := router.Group("/model")
	{
		model.POST("/", controllers.CreateModel)
		model.GET("/", controllers.FetchAllModel)
		model.GET("/:id", controllers.FetchModel)
		model.PUT("/:id", controllers.UpdateModel)
		model.DELETE("/:id", controllers.DeleteModel)
	}
	dataset := router.Group("/dataset")
	{
		dataset.POST("/", controllers.CreateDataset)
		dataset.GET("/", controllers.FetchAllDataset)
		dataset.GET("/:id", controllers.FetchDataset)
		dataset.PUT("/:id", controllers.UpdateDataset)
		dataset.DELETE("/:id", controllers.DeleteDataset)
	}
	label := router.Group("/label")
	{
		label.POST("/", controllers.CreateLabel)
		label.GET("/", controllers.FetchAllLabel)
		label.GET("/:id", controllers.FetchLabel)
		label.PUT("/:id", controllers.UpdateLabel)
		label.DELETE("/:id", controllers.DeleteLabel)
	}
	sample := router.Group("/sample")
	{
		sample.POST("/", controllers.CreateSample)
		sample.GET("/", controllers.FetchAllSample)
		sample.GET("/:id", controllers.FetchSample)
		sample.PUT("/:id", controllers.UpdateSample)
		sample.DELETE("/:id", controllers.DeleteSample)
	}
	labelset := router.Group("/labelset")
	{
		labelset.POST("/", controllers.CreateLabelset)
		labelset.GET("/", controllers.FetchAllLabelset)
		labelset.GET("/:id", controllers.FetchLabelset)
		labelset.PUT("/:id", controllers.UpdateLabelset)
		labelset.DELETE("/:id", controllers.DeleteLabelset)
	}
	annotation := router.Group("/annotation")
	{
		annotation.POST("/", controllers.CreateAnnotation)
		annotation.GET("/", controllers.FetchAllAnnotation)
		annotation.GET("/:id", controllers.FetchAnnotation)
		annotation.PUT("/:id", controllers.UpdateAnnotation)
		annotation.DELETE("/:id", controllers.DeleteAnnotation)
	}

	//router.Use(middlewares.Cors())

	router.Run(":8080")
}

func GetCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"x-requested-with", "Content-Type", "AccessToken", "X-CSRF-Token", "X-Token", "Authorization", "token"}
	return config
}
