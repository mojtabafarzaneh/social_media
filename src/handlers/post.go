package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
)

type PostController struct {
	PostRepository repository.PostgresPostRepo
}

func NewUserController() *PostController {
	return &PostController{
		PostRepository: *repository.NewPostgresPostRepo(),
	}
}

func (pc *PostController) CreatePostHandler(c *gin.Context) {

}

func (pc *PostController) ListPostsHandler(c *gin.Context) {

}

func (pc *PostController) UpdatePostsHandler(c *gin.Context) {

}

func (pc *PostController) DeletePostHandler(c *gin.Context) {

}

func (pc *PostController) GetPostHandler(c *gin.Context) {

}
