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

	if err := ar.repository.GetRegister(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(24*time.Hour, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	//log.Print(token)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "successfuly registerd",
		"token":    token,
		"username": params.Username,
	})
}

func (ar *AuthControler) LoginHandler(c *gin.Context) {
	var user *types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
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

func (ac *AuthControler) GetAdminRegisterHandler(c *gin.Context) {
	var user *types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user, err := ac.repository.GetAdminRegister(user.IsAdmin, user.Username, user.Password)
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
