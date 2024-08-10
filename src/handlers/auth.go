package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"github.com/mojtabafarzaneh/social_media/src/utils"
)

type AuthControler struct {
	repository repository.AuthPostgresRepo
}

func NewAuthControler() *AuthControler {
	return &AuthControler{
		repository: *repository.NewAuthPostgresRepo(),
	}
}

// RegiserHandler handles the registration of a new user
// @Summary Register a new user
// @Description Register a new user and return an authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body types.CreateUserParams true "User registration details"
// @Success 201 {object} map[string]interface{} "successfuly registered"
// @Failure 400 {object} map[string]interface{} "Validation failed or bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/register [post]
func (ar *AuthControler) RegiserHandler(c *gin.Context) {
	var params types.CreateUserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if validationErrors := params.Validate(); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": validationErrors})
		return
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	if user, err = ar.repository.GetRegister(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(24*time.Hour, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "successfuly registerd",
		"token":    token,
		"username": user.Username,
	})
}

// LoginHandler handles user login
// @Summary User login
// @Description Authenticate a user and return an authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body types.LoginUser true "User login details"
// @Success 200 {object} map[string]interface{} "authenticated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid username or password"
// @Router /auth/login [post]
func (ar *AuthControler) LoginHandler(c *gin.Context) {

	var params types.LoginUser

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user := &types.User{
		Username: params.Username,
		Password: params.Password,
	}

	user, err := ar.repository.GetLogin(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "username or password is incorrect",
			"detail": err.Error(),
		})
		return
	}

	log.Print("the user is", user)
	genToken, err := utils.GenerateToken(24*time.Hour, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "couldn't generate the token",
			"detail": err.Error(),
		})
		return
	}

	log.Print(genToken)

	c.JSON(http.StatusOK, gin.H{
		"message": "authenticated successfully",
		"token":   genToken,
	})

}

// GetAdminRegisterHandler handles the registration of a new admin user
// @Summary Register a new admin user
// @Description Register a new admin user and return an authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body types.AdminRegisterParams true "Admin registration details"
// @Success 200 {string} string "Token generated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input or couldn't generate token"
// @Router /auth/admin/register [post]
func (ac *AuthControler) GetAdminRegisterHandler(c *gin.Context) {
	var params *types.AdminRegisterParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user := &types.User{
		Username: params.Username,
		Password: params.Password,
	}

	user, err := ac.repository.GetAdminRegister(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "username or password is incorrect",
			"detail": err.Error(),
		})
		return
	}

	gentoken, err := utils.GenerateToken(24*time.Hour, *user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "couldn't generate the token",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gentoken)

}
