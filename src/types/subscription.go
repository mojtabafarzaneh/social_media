package types

type Subscription struct {
	ID           uint `gorm:"PrimaryKey" json:"Id"`
	SubscriberID uint `json:"subscriberId"`
	TargetID     uint `json:"targetId"`
}
