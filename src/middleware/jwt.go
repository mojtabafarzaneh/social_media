package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/handlers"
	"github.com/mojtabafarzaneh/social_media/src/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			handlers.ErrUnauthorizedUser(c, "you have to provide the Authorization header")
			c.Abort()
			return
		}

		user, err := utils.ValidateToken(tokenString, string(utils.SecretKey))
		if err != nil {
			handlers.ErrUnauthorizedUser(c, err.Error())
			c.Abort()
			return

		}
		c.Set("user", user)

		c.Next()
	}
}
