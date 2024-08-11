package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mojtabafarzaneh/social_media/src/handlers"
	"github.com/mojtabafarzaneh/social_media/src/utils"
)

func IsUserAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			handlers.ErrUnauthorizedUser(c, "authorization header required!")
			c.Abort()
			return
		}

		tok, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
			}
			return utils.SecretKey, nil
		})
		if err != nil {
			handlers.ErrUnauthorizedUser(c, err.Error())
			c.Abort()
			return
		}

		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok || !tok.Valid {
			handlers.ErrUnauthorizedUser(c, "invalid token!")
			c.Abort()
			return
		}
		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok || !isAdmin {
			handlers.ErrUnauthorizedUser(c, "access denied! only admins can access this page.")
			c.Abort()
			return
		}

		c.Next()

	}
}
