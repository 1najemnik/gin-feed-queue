package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateAccessKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		requiredAccessKey := os.Getenv("ACCESS_KEY")
		queryAccessKey := c.Query("access_key")
		if requiredAccessKey != "" && queryAccessKey != requiredAccessKey {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: invalid or missing access_key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
