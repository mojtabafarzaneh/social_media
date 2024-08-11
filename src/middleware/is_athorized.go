package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/handlers"
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
			handlers.ErrUnauthorizedUser(c, err.Error())
			c.Abort()
			return
		}

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			handlers.ErrUnauthorizedUser(c, "provide the right token!")
			return
		}

		claims, err := utils.ValidateToken(tokenString, string(utils.SecretKey))

		if err != nil {
			handlers.ErrUnauthorizedUser(c, err.Error())
			c.Abort()
			return
		}

		userIDToken, ok := claims["sub"].(float64)
		if !ok || uint(userIDToken) != user.ID && claims["isAdmin"] != true {
			handlers.ErrUnauthorizedUser(c, "you can not access this page!")
			c.Abort()
			return
		}

		c.Next()

	}
}
