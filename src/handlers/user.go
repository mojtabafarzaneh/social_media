package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary List all users
// @Description Retrieves a list of all users. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} types.ResponseUser "List of users"
// @Failure 400 {object} ErrorResponse "Bad request error"
// @Security BearerAuth
// @Router /users/list [get]
func (cl *UserControler) ListUserHandler(c *gin.Context) {

	user, err := cl.UserRepository.ListUser(c)
	if err != nil {
		ErrBadRequest(c, err.Error())
		return
	}
	response := types.UsersToUserResponses(user)
	c.JSON(http.StatusOK, response)
}

// @Summary Get a specific user
// @Description Retrieves details of a specific user by ID. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} types.ResponseUser "User details"
// @Failure 404 {object} ErrorResponse "User not found error"
// @Security BearerAuth
// @Router /users/{id} [get]
func (cl *UserControler) GetUserHandler(c *gin.Context) {

	var id = c.Params.ByName("id")

	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}
	user, err := cl.UserRepository.GetUserByID(c, userID)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}
	response := types.UsersToUserResponses(user)

	c.JSON(http.StatusOK, response)

}

// @Summary Create a new user
// @Description Creates a new user. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body types.CreateUserParams true "User data"
// @Success 200 {object} types.ResponseUser "Created user details"
// @Failure 400 {object} ErrorResponse "Bad request error"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Failure 404 {object} ErrorResponse "Record not found"
// @Security BearerAuth
// @Router /users [post]
func (cl *UserControler) InsertUserHandler(c *gin.Context) {
	var params types.CreateUserParams

	if err := c.BindJSON(&params); err != nil {
		ErrBadRequest(c, err.Error())
		return
	}

	if validationErrors := params.Validate(); len(validationErrors) > 0 {
		ErrValidationFailed(c, validationErrors)
		return
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		ErrBadRequest(c, err.Error())
		return
	}

	insertedUser, err := cl.UserRepository.CreateUser(c, *user)
	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	showcase := types.UsersToUserResponses(insertedUser)

	c.JSON(http.StatusOK, showcase)

}

// @Summary Delete a user
// @Description Deletes a user by ID. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 404 {object} ErrorResponse "User not found or failed to delete"
// @Security BearerAuth
// @Router /users/{id} [delete]
func (cl *UserControler) DeleteUserHandler(c *gin.Context) {

	var id = c.Params.ByName("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}
	_, err = cl.UserRepository.GetUserByID(c, userID)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	err = cl.UserRepository.DeleteUser(c, userID)
	if err != nil {
		ErrRecordNotFound(c, err.Error())
	}

	c.JSON(http.StatusNoContent, gin.H{"success": "deleted successfuly"})
}

// @Summary Update a user's username
// @Description Updates the username of a user by ID. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param username body types.UpdateUsernameParams true "New username"
// @Success 202 {object} map[string]interface{} "Username updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid request or bad user ID"
// @Failure 404 {object} ErrorResponse "Record not found"
// @Security BearerAuth
// @Router /users/{id}/username [put]
func (cl *UserControler) UpdateUsernameHandler(c *gin.Context) {
	var params types.UpdateUsernameParams

	id := c.Param("id")

	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}

	if err := c.BindJSON(&params); err != nil {
		ErrBadRequest(c, err.Error())
		return
	}

	err = cl.UserRepository.UpdateUsername(params.Username, userID)
	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"massage": "username updated"})
}
