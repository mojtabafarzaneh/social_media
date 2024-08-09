package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		user, err := utils.ValidateToken(tokenString, string(utils.SecretKey))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return

		}
		c.Set("user", user)

		c.Next()
	}
}
