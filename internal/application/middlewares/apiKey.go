package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func ApiKeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("APP_ENV") == "dev" {
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
