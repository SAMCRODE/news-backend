package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/news-backend/config"
	"github.com/news-backend/models"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		decr := jwt.MapClaims{}

		jwt.ParseWithClaims(token, decr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JWTKEY), nil
		})

		// var user models.User

		jsonbody, err := json.Marshal(decr["user"])

		if err != nil {
			fmt.Println(err)
		}

		user := models.User{}
		if err := json.Unmarshal(jsonbody, &user); err != nil {
			// do error check
			fmt.Println(err)
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
