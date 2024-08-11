package types

import "github.com/google/uuid"

type Subscription struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"Id"`
	SubscriberID uuid.UUID `json:"subscriberId"`
	TargetID     uuid.UUID `json:"targetId"`
}

type SubscriberResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SubscriptionResponse struct {
	Username string `json:"username"`
}

func (u *User) SubcriberResponse() *SubscriberResponse {
	return &SubscriberResponse{
		Username: u.Username,
		Email:    u.Email,
	}
}

func UserToSubscriberResponse(user []*User) []SubscriberResponse {
	subResponse := make([]SubscriberResponse, len(user))

	for i, users := range user {
		subResponse[i] = *users.SubcriberResponse()
	}
	return subResponse
}
