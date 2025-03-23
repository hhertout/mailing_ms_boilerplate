package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("GO_ENV") == "development" {
			c.Next()
			return
		}

		apiKey := os.Getenv("API_KEY")
		bearer := c.GetHeader("Authorization")
		if bearer == "" || bearer != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
		}
		c.Next()
	}
}
