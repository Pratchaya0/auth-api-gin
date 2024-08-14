package middlewares

import (
	"net/http"
	"os"

	"github.com/Pratchaya0/auth-api-gin/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_SECRET")), nil
	})

	if err != nil {
		helpers.WebResponseWithJSON(c, http.StatusUnauthorized, "unauthenticated", nil)
		return
	}

	c.Next()
}
