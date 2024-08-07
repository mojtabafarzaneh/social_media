package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
	"github.com/mojtabafarzaneh/social_media/src/types"
)

type UserControler struct {
	UserRepository repository.PostgresRep
}

func NewUserControler() *UserControler {
	return &UserControler{
		UserRepository: *repository.NewUserPostgresRep(),
	}
}

func (cl *UserControler) ListUserHandler(c *gin.Context) {

	user, err := cl.UserRepository.ListUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request!"})
		return
	}
	response := types.UsersToUserResponses(user)
	c.JSON(http.StatusOK, response)
}

func (cl *UserControler) GetUserHandler(c *gin.Context) {

	var id = c.Params.ByName("id")
	user, err := cl.UserRepository.GetUserByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
		return
	}
	response := types.UsersToUserResponses(user)

	c.JSON(http.StatusOK, response)

}

func (cl *UserControler) InsertUserHandler(c *gin.Context) {
	var params types.CreateUserParams

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if validationErrors := params.Validate(); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedUser, err := cl.UserRepository.CreateUser(c, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert user", "details": err.Error()})
		return
	}

	showcase := types.UsersToUserResponses(insertedUser)

	c.JSON(http.StatusOK, showcase)

}

func (cl *UserControler) DeleteUserHandler(c *gin.Context) {

	var id = c.Params.ByName("id")

	_, err := cl.UserRepository.GetUserByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user not found!",
			"details": err.Error(),
		})
		return
	}

	err = cl.UserRepository.DeleteUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "failed to find the user",
			"details": err.Error(),
		})
	}
}

func (cl *UserControler) UpdateUsernameHandler(c *gin.Context) {
	var params types.UpdateUsernameParams

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = cl.UserRepository.UpdateUsername(params.Username, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update username"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"massage": "username updated"})
}
