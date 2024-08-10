package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type Subscription interface {
	GetAllSubs(ctx context.Context, id string) ([]*types.User, error)
	GetAllSubscribed(ctx context.Context, id uint) ([]*types.User, error)
	CreateSubscription(ctx context.Context, targetUsername string, subscriberId string) error
}

type PostgresSubsRepo struct {
	DB *gorm.DB
}

func NewPostgresSubsRepo() *PostgresSubsRepo {
	return &PostgresSubsRepo{
		DB: db.ConnectToDB(),
	}
}

func (ps *PostgresSubsRepo) GetAllSubscriptions(ctx context.Context, id string) ([]*types.User, error) {
	var subs []*types.Subscription
	var targetId []uint
	var subcripted []*types.User

	//searching for all the users that this user subsribed to.
	if err := ps.DB.WithContext(ctx).Model(&subs).
		Where("subscriber_id = ?", id).
		Pluck("target_id", &targetId).Error; err != nil {
		return nil, err
	}

	if len(targetId) == 0 {
		return nil, nil
	}

	//retrieving useraname of the founded users.
	if err := ps.DB.WithContext(ctx).Find(&subcripted, targetId).Error; err != nil {
		return nil, err
	}

	return subcripted, nil
}

func (ps *PostgresSubsRepo) GetAllSubscribed(ctx context.Context, id uint) ([]*types.User, error) {
	var subscriberId []uint
	var subscribers []*types.User
	var subs []*types.Subscription

	if err := ps.DB.WithContext(ctx).Model(subs).
		Where("target_id = ?", id).
		Pluck("subscriber_id", &subscriberId).Error; err != nil {
		return nil, err

	}

	if len(subscriberId) == 0 {
		return nil, nil
	}

	if err := ps.DB.WithContext(ctx).Find(&subscribers, subscriberId).Error; err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (ps *PostgresSubsRepo) CreateSubscription(ctx context.Context, targetUsername string, subscriberId string) error {
	var target, subscriber types.User
	var existingSubscription types.Subscription

	//getting the targetusername
	if err := ps.DB.WithContext(ctx).
		First(&target, "username = ?", targetUsername).
		Error; err != nil {
		return fmt.Errorf("this username does not exists %w", err)
	}
	//getting subscribers id
	if err := ps.DB.WithContext(ctx).First(&subscriber, subscriberId).Error; err != nil {
		return fmt.Errorf("the subscriber does not exists %w", err)
	}

	//making sure target and subscriber does not exists
	if err := ps.DB.WithContext(ctx).
		Where("subscriber_id = ? AND target_id = ?", subscriber.ID, target.ID).
		First(&existingSubscription).
		Error; err == nil {
		return fmt.Errorf("subscription already exists")
	}

	subscribe := types.Subscription{
		TargetID:     target.ID,
		SubscriberID: subscriber.ID,
	}

	//creating subscriber
	if err := ps.DB.WithContext(ctx).Create(&subscribe).Error; err != nil {
		return fmt.Errorf("couldn't create this object %w", err)
	}

	return nil
}

func (ps *PostgresSubsRepo) FindUsernames(ctx context.Context, useraname string) ([]*types.User, error) {
	var user []*types.User

	if err := ps.DB.Model(&user).
		Where("username LIKE ?", "%"+useraname+"%").
		Find(&user).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return user, nil
}

//TODO: Delete functionality!
