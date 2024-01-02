package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var baseToken string

func init() {
	baseToken = os.Getenv("AUTHORIZATION_TOKEN")
	if baseToken == "" {
		panic("AUTHORIZATION_TOKEN is not set")
	}
}

func CheckTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), " ")[1]

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		} else {
			if token != baseToken {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
