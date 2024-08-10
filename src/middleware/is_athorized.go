package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"github.com/mojtabafarzaneh/social_media/src/utils"
)

type Controler struct {
	repo repository.PostgresMiddlewareRepo
}

func NewControler() *Controler {
	return &Controler{
		repo: *repository.NewPostgresMiddlewareRepo(),
	}
}

func (uc *Controler) IsUserAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *types.User

		id := c.Params.ByName("user")

		user, err := uc.repo.GetUserId(id)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error()})
			return
		}

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you have to be authenticated to enter this page"})
			return
		}

		claims, err := utils.ValidateToken(tokenString, string(utils.SecretKey))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "please enter the right authorization credintionals" + err.Error()})
			c.Abort()
			return
		}

		userIDToken, ok := claims["sub"].(float64)
		if !ok || uint(userIDToken) != user.ID && claims["isAdmin"] != true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you can not access this page!",
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
