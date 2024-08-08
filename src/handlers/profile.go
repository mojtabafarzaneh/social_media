package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
)

type ProfileControler struct {
	ProfileRepository repository.PostgresProfileRepo
}

func NewProfileControler() *ProfileControler {
	return &ProfileControler{
		ProfileRepository: *repository.NewPostgresProfileRepo(),
	}
}

func (pc *ProfileControler) GetUserProfileHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
	}

	profile, err := pc.ProfileRepository.GetUserProfile(c, uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user profile": profile,
	})
}
