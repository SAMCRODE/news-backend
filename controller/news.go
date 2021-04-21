package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/news-backend/models"
)

type NewController struct{}

func (n NewController) Save(c *gin.Context) {
	new := models.New{}

	if errA := c.ShouldBind(&new); errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid new object"})
		return
	}

	err := new.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gift Save"})
}

func (n NewController) SearchNewsByPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again later1"})
		return
	}

	news, err := models.SearchNewsPaginated(page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later2"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}

func (n NewController) SearchLastestNews(c *gin.Context) {
	news, err := models.SearchLastestNews(10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later2"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}

func (n NewController) SearchHomePageNews(c *gin.Context) {
	news, err := models.SearchLastestNews(10)
	hotnews, err1 := models.SearchHotNews()

	if err != nil || err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later2"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lastest": news, "hot": hotnews})
}

func (n NewController) SearchNewsDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again later1"})
		return
	}

	new, err := models.SearchNewWithRows(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later2"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": new})
}
