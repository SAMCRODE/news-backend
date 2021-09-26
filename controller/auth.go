package controller

import (
	"crypto/md5"
	"encoding/hex"
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

	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	if err != nil || user.Password != suser.Password || suser.Id == 0 {
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
