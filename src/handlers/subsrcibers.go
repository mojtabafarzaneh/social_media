package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/repository"
	"github.com/mojtabafarzaneh/social_media/src/types"
)

type SubsController struct {
	SubsRepository repository.PostgresSubsRepo
}

func NewSubsController() *SubsController {
	return &SubsController{
		SubsRepository: *repository.NewPostgresSubsRepo(),
	}
}

func (sc *SubsController) GetAllSubscriptions(c *gin.Context) {
	id := c.Params.ByName("id")

	subs, err := sc.SubsRepository.GetAllSubscriptions(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user not found",
			"details": err.Error(),
		})
	}

	if len(subs) == 0 {
		c.JSON(http.StatusOK, gin.H{"details": "this user has not subscribe to any user"})

	} else {
		res := types.UserToSubscriberResponse(subs)

		c.JSON(http.StatusOK, gin.H{fmt.Sprintf("all the subscriptions of the user %v", id): res})
	}

}

func (sc *SubsController) GetAllSubscribed(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	subscriber, err := sc.SubsRepository.GetAllSubscribed(c, uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "user not found",
			"details": err.Error(),
		})
	}

	if len(subscriber) == 0 {
		c.JSON(http.StatusOK, gin.H{"details": "there are no records of any subscriber for this user"})
	} else {
		response := types.UserToSubscriberResponse(subscriber)

		c.JSON(http.StatusOK, gin.H{fmt.Sprintf("all the subscribers of the user %v", id): response})
	}
}

func (sc *SubsController) CreateSubs(c *gin.Context) {
	var requestedUser types.SubscriptionResponse
	if err := c.BindJSON(&requestedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("subscriber")

	if len(requestedUser.Username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request is empty"})
		return
	}

	err := sc.SubsRepository.CreateSubscription(c, requestedUser.Username, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"details": "subscribed!"})
}
