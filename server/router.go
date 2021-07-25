package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/news-backend/controller"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	newsGroup := router.Group("news")
	{
		news := new(controller.NewController)
		newsGroup.POST("/create", news.Save)
		newsGroup.GET("/collection/:page", news.SearchNewsByPage)
		newsGroup.GET("/lastest", news.SearchLastestNews)
		newsGroup.GET("/detail/:id", news.SearchNewsDetail)
		newsGroup.GET("/home", news.SearchHomePageNews)
	}

	userGroup := router.Group("users")
	{
		users := new(controller.UserController)
		userGroup.POST("/create-account", users.CreateAccount)
	}

	authGroup := router.Group("auth")
	{
		auth := new(controller.AuthController)
		authGroup.POST("/", auth.Auth)
	}

	return router
}
