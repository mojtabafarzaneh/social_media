package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary Get all subscriptions of a user
// @Description Retrieves all subscriptions for a user by ID or username. Accessible to authenticated users.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string optional "User ID"
// @Param username query string optional "Username"
// @Success 200 {object} map[string]interface{} "List of subscriptions"
// @Failure 404 {object} ErrorResponse "User not found error"
// @Security BearerAuth
// @Router /subs/subscriptions/{id} [get]
func (sc *SubsController) GetAllSubscriptions(c *gin.Context) {

	username, ok := c.GetQuery("username")

	if !ok {
		id := c.Params.ByName("id")

		userID, err := uuid.Parse(id)
		if err != nil {
			ErrBadRequest(c, err.Error())
		}

		subs, err := sc.SubsRepository.GetAllSubscriptions(c, userID)
		if err != nil {
			ErrRecordNotFound(c, err.Error())
			return
		}

		if len(subs) == 0 {
			c.JSON(http.StatusOK, gin.H{"details": "this user has not subscribed to any user"})
			return
		}

		res := types.UserToSubscriberResponse(subs)
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("all the subscriptions of the user %v", id): res,
		})
		return
	}
	query, err := sc.SubsRepository.FindUsernames(c, username)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	response := types.UserToSubscriberResponse(query)
	c.JSON(http.StatusOK, response)
}

// @Summary Get all users subscribed to a user
// @Description Retrieves all users subscribed to a specific user by ID. Accessible to authenticated users.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{} "List of subscribers"
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 404 {object} ErrorResponse "User not found error"
// @Security BearerAuth
// @Router /subs/subscribers/{id} [get]
func (sc *SubsController) GetAllSubscribed(c *gin.Context) {
	id := c.Params.ByName("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}

	subscriber, err := sc.SubsRepository.GetAllSubscribed(c, userID)

	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	if len(subscriber) == 0 {
		c.JSON(http.StatusOK, gin.H{"details": "there are no records of any subscriber for this user"})
		return
	} else {
		response := types.UserToSubscriberResponse(subscriber)
		c.JSON(http.StatusOK, gin.H{fmt.Sprintf("all the subscribers of the user %v", id): response})
		return
	}
}

// @Summary Create a new subscription
// @Description Creates a new subscription for a user. Accessible to authenticated users.
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param subscriber path string true "Subscriber ID"
// @Param subscription body types.SubscriptionResponse true "Subscription data"
// @Success 201 {object} ErrorResponse "Subscription created successfully"
// @Failure 400 {object} ErrorResponse "Bad request error"
// @Failure 404 {object} ErrorResponse "Record not found"
// @Security BearerAuth
// @Router /subs/{subscriber} [post]
func (sc *SubsController) CreateSubs(c *gin.Context) {
	var requestedUser types.SubscriptionResponse
	if err := c.BindJSON(&requestedUser); err != nil {
		ErrBadRequest(c, err.Error())
		return
	}

	id := c.Params.ByName("subscriber")

	if len(requestedUser.Username) == 0 {
		ErrBadRequest(c, "should provid the username")
		return
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		ErrBadRequest(c, err.Error())
	}

	err = sc.SubsRepository.CreateSubscription(c, requestedUser.Username, userID)
	if err != nil {
		ErrRecordNotFound(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"details": "subscribed!"})
}
