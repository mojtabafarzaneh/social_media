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

// GetUserProfileHandler handles the retrieval of a user profile by ID.
// @Summary Get user profile
// @Description Retrieve the profile information for a specific user by ID.
// @Tags profile
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User profile retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 406 {object} map[string]interface{} "Not Acceptable"
// @Security BearerAuth
// @Router /profile/{id} [get]
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
