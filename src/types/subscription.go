package types

type Subscription struct {
	ID           uint `gorm:"PrimaryKey" json:"Id"`
	SubscriberID uint `json:"subscriberId"`
	TargetID     uint `json:"targetId"`
}

type SubscriberResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
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
