package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// @Param body body types.CreatePostParams true "Post details"
// @Success 201 {object} types.Post "Post created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input or error creating post"
// @Failure 404 {object} ErrorResponse "Record not found"
// @Router /posts/{user} [post]
func (pc *PostController) CreatePostHandler(c *gin.Context) {
	var params types.CreatePostParams
	if err := c.ShouldBindJSON(&params); err != nil {
		ErrBadRequest(c, err.Error())
		return
	}

	post := types.CreatePostFromParams(params)

	res, err := pc.PostRepository.InsertPost(c, *post)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)

}

// ListPostsHandler handles the retrieval of posts, optionally filtering by content
// @Summary Retrieve all posts or filter by content
// @Description Retrieves a list of posts. Optionally, you can filter posts by providing a 'content' query parameter.
// @security BearerAuth
// @Tags posts
// @Accept json
// @Produce json
// @Param content query string false "Filter posts by content"
// @Success 200 {array} types.Post "List of posts"
// @Failure 400 {object} ErrorResponse "Invalid request or error retrieving posts"
// @Failure 404 {object} ErrorResponse "Record not fount"
// @Router /posts [get]
func (pc *PostController) ListPostsHandler(c *gin.Context) {
	content, ok := c.GetQuery("content")
	if !ok {
		response, err := pc.PostRepository.GetAllPosts(c)
		if err != nil {
			ErrBadRequest(c, err.Error())
			return
		}

		c.JSON(http.StatusOK, response)
	}

	query, err := pc.PostRepository.FindPost(c, content)
	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, query)

}

// UpdatePostsHandler godoc
// @Summary      Update a Post
// @Description  Update the content of an existing post by its UUID.
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id    path     string          true  "Post ID"
// @Param        user  path     string          true  "User ID"
// @Param        post  body     types.Post      true  "Post Content"
// @Success      200  {object}  types.Post      "Updated Post"
// @Failure      400  {object}  ErrorResponse   "Bad Request"
// @Failure      404  {object}  ErrorResponse   "Post Not Found"
// @security BearerAuth
// @Router       /posts/{user}/{id} [put]
func (pc *PostController) UpdatePostsHandler(c *gin.Context) {
	var updatContent types.Post

	id := c.Params.ByName("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}

	if err := c.BindJSON(&updatContent); err != nil {
		ErrBadRequest(c, err.Error())
		return
	}

	response, err := pc.PostRepository.UpdatePost(c, updatContent.Content, userID)
	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
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
// @Success 204 {object} ErrorResponse "Post deleted successfully"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Router /posts/{user}/{id} [delete]
// @Security BearerAuth
func (pc *PostController) DeletePostHandler(c *gin.Context) {

	id := c.Params.ByName("id")

	postID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}

	_, err = pc.PostRepository.GetPost(c, postID)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	if err := pc.PostRepository.DeletePost(c, postID); err != nil {
		ErrRecordNotFound(c, err.Error())
		return
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
// @Failure 404 {object} ErrorResponse "Not Found"
// @Router /posts/{id} [get]
// @Security BearerAuth
func (pc *PostController) GetPostHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}
	post, err := pc.PostRepository.GetPost(c, userID)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}
