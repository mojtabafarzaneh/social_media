package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
	"github.com/mojtabafarzaneh/social_media/src/types"
)

type PostController struct {
	PostRepository repository.PostgresPostRepo
}

func NewPostController() *PostController {
	return &PostController{
		PostRepository: *repository.NewPostgresPostRepo(),
	}
}

func (pc *PostController) CreatePostHandler(c *gin.Context) {
	var posts types.Post
	if err := c.BindJSON(&posts); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user not found",
			"details": err.Error(),
		})
	}

	res, err := pc.PostRepository.InsertPost(c, posts)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user not found",
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, res)

}

func (pc *PostController) ListPostsHandler(c *gin.Context) {

	response, err := pc.PostRepository.GetAllPosts(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"details": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)

}

func (pc *PostController) UpdatePostsHandler(c *gin.Context) {
	var updatContent types.Post

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "couldn't find the id",
			"details": err.Error(),
		})
	}

	if err := c.BindJSON(&updatContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "insert the right data!",
			"details": err.Error(),
		})
	}

	response, err := pc.PostRepository.UpdatePost(c, updatContent.Content, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "couldn't find the post you wanted",
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusOK, &response)
}

func (pc *PostController) DeletePostHandler(c *gin.Context) {

	id := c.Params.ByName("id")

	_, err := pc.PostRepository.GetPost(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "user not found",
			"detail": err.Error(),
		})
	}

	if err := pc.PostRepository.DeletePost(c, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "user not found",
			"detail": err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, gin.H{})

}

func (pc *PostController) GetPostHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	post, err := pc.PostRepository.GetPost(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "user not found",
			"detail": err.Error(),
		})
	}
	c.JSON(http.StatusOK, post)
}
