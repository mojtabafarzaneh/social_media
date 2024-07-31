package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
)

type Controler struct {
	UserRepository repository.UserRepository
}

func NewControler() *Controler {
	return &Controler{
		UserRepository: *repository.NewUserRepository(),
	}
}

func (cl *Controler) ListUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cl.UserRepository.ListUser(8))

}
