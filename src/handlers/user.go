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

// @Summary List all users
// @Description Retrieves a list of all users. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} types.ResponseUser "List of users"
// @Failure 400 {object} map[string]string "Bad request error"
// @Security BearerAuth
// @Router /users [get]
func (cl *UserControler) ListUserHandler(c *gin.Context) {

	user, err := cl.UserRepository.ListUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request!"})
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
// @Failure 404 {object} map[string]string "User not found error"
// @Failure 401 {object} map[string]string "Unauthorized error"
// @Security BearerAuth
// @Router /users/{id} [get]
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

// @Summary Create a new user
// @Description Creates a new user. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body types.CreateUserParams true "User data"
// @Success 200 {object} types.ResponseUser "Created user details"
// @Failure 400 {object} map[string]string "Bad request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /users [post]
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

// @Summary Delete a user
// @Description Deletes a user by ID. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 404 {object} map[string]string "User not found or failed to delete"
// @Security BearerAuth
// @Router /users/{id} [delete]
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

// @Summary Update a user's username
// @Description Updates the username of a user by ID. Accessible only by admin users.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param username body types.UpdateUsernameParams true "New username"
// @Success 202 {object} map[string]string "Username updated successfully"
// @Failure 400 {object} map[string]string "Invalid request or bad user ID"
// @Failure 500 {object} map[string]string "Failed to update username"
// @Security BearerAuth
// @Router /users/{id}/username [put]
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
