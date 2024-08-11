package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//type httpStatus int

type ErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func ErrBadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   http.StatusBadRequest,
		Message: "bad request " + err,
	})
}

func ErrRecordNotFound(c *gin.Context, err string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error:   http.StatusNotFound,
		Message: "record not found " + err,
	})
}

func ErrValidationFailed(c *gin.Context, err map[string]string) {
	c.JSON(http.StatusUnprocessableEntity, ErrorResponse{
		Error:   http.StatusUnauthorized,
		Message: err["error"],
	})
}

func ErrDatabaseFailed(c *gin.Context, err string) {
	c.JSON(http.StatusConflict, ErrorResponse{
		Error:   http.StatusConflict,
		Message: "database error: " + err,
	})
}

func ErrFailedGeneratingToken(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error:   http.StatusInternalServerError,
		Message: "internal server error " + err,
	})
}

func ErrNotAuthenticated(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error:   http.StatusUnauthorized,
		Message: "please provide the right password and username " + err,
	})
}

func ErrUnauthorizedUser(c *gin.Context, err string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error:   http.StatusUnauthorized,
		Message: "Authorization failed " + err,
	})
}
