package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/news-backend/config"
	"github.com/news-backend/models"
)

type AuthController struct{}

func (u AuthController) Auth(c *gin.Context) {
	user := models.User{}

	if errA := c.ShouldBind(&user); errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid user object"})
		return
	}

	suser, err := models.SearchUserByEmail(user.Email)

	if err != nil || user.Password != suser.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not identified"})
		return
	}

	suser.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": suser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(config.GetConfig().JWTKEY))

	c.Header("Authorization", tokenString)
	c.JSON(http.StatusAccepted, gin.H{"Name": suser.Name,
		"Email": suser.Email, "ImageUrl": suser.ImageUrl, "Permissions": suser.Permissions,
		"Jwt": tokenString})
}
