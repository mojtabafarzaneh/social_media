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

// @security BearerAuth
// @Summary Create a new post
// @Description Create a new post. Only accessible by authorized users with a valid JWT token.
// @Tags posts
// @Accept json
// @Produce json
// @Param user path string true "User ID"
// @Param body body types.Post true "Post details"
// @Success 201 {object} types.Post "Post created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input or error creating post"
// @Failure 401 {object} map[string]interface{} "Unauthorized access"
// @Router /posts/{user} [post]
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
	content, ok := c.GetQuery("content")
	if !ok {
		response, err := pc.PostRepository.GetAllPosts(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"details": err.Error(),
			})
		}

		c.JSON(http.StatusOK, response)
	}

	query, err := pc.PostRepository.FindPost(c, content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"details": err.Error(),
		})
	}
	c.JSON(http.StatusOK, query)

}

// ListPostsHandler handles the retrieval of posts, optionally filtering by content
// @Summary Retrieve all posts or filter by content
// @Description Retrieves a list of posts. Optionally, you can filter posts by providing a 'content' query parameter.
// @Tags posts
// @Accept json
// @Produce json
// @Param content query string false "Filter posts by content"
// @Success 200 {array} types.Post "List of posts"
// @Failure 400 {object} map[string]interface{} "Invalid request or error retrieving posts"
// @Router /posts [get]
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

// DeletePostHandler handles the deletion of a specific post by its ID
// @Summary Delete a specific post
// @Description Delete a post by its ID. Only accessible by authorized users with a valid JWT token.
// @Tags posts
// @Accept json
// @Produce json
// @Param user path string true "User ID"
// @Param id path string true "Post ID"
// @Success 204 {object} map[string]interface{} "Post deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /posts/{user}/{id} [delete]
// @Security BearerAuth
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

	c.JSON(http.StatusNoContent, gin.H{"message": "success"})

}

// GetPostHandler retrieves a specific post by its ID
// @Summary Get a specific post by ID
// @Description Retrieve a post by its ID. Only accessible by authorized users with a valid JWT token.
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} types.Post "Post retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /posts/{id} [get]
// @Security BearerAuth
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
