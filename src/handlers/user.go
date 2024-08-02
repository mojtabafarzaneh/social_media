package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
	"github.com/mojtabafarzaneh/social_media/src/types"
)

type Controler struct {
	UserRepository repository.UserRepository
}

func NewControler() *Controler {
	return &Controler{
		UserRepository: *repository.NewUserRepository(),
	}
}

func (cl *Controler) ListUserHandler(c *gin.Context) {

	user := cl.UserRepository.ListUser(c)
	showcase := types.UsersToUserResponses(user)
	c.JSON(http.StatusOK, showcase)
}

func (cl *Controler) GetUserHandler(c *gin.Context) {

	var id = c.Params.ByName("id")
	user, err := cl.UserRepository.GetUserByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})

	}

	c.JSON(http.StatusOK, user)
}

func (cl *Controler) InsertUserHandler(c *gin.Context) {
	var params types.CreateUserParams

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusFailedDependency, err)
	}

	if err := params.Validate(); len(err) > 0 {
		c.JSON(http.StatusBadRequest, err)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	insertedUser, err := cl.UserRepository.CreateUser(c, *user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	showcase := types.UsersToUserResponses(insertedUser)

	c.JSON(http.StatusOK, showcase)

}

func (cl *Controler) DeleteUserHandler(c *gin.Context) {

	var id = c.Params.ByName("id")
	c.JSON(http.StatusNoContent, cl.UserRepository.DeleteUser(c, id))
}

func (cl *Controler) UpdateUsernameHandler(c *gin.Context) {
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

	user, err := cl.UserRepository.UpdateUsername(params.Username, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update username"})
		return
	}

	c.JSON(http.StatusAccepted, user)
}
