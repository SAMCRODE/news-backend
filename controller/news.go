package controller

import (
	"net/http"

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
