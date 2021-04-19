package server

import (
	"github.com/gin-gonic/gin"
	"github.com/news-backend/controller"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
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
		newsGroup.GET("/detail/:id", news.SearchNewsDetail)
	}

	return router
}
