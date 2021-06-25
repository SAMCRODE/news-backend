package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/news-backend/models"
)

type UserController struct{}

func (u UserController) CreateAccount(c *gin.Context) {
	user := models.User{}

	if errA := c.ShouldBind(&user); errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid user object"})
		return
	}

	err := user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gift Save"})
}
