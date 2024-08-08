package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type Profile interface {
}

type PostgresProfileRepo struct {
	DB *gorm.DB
}

func NewPostgresProfileRepo() *PostgresProfileRepo {
	return &PostgresProfileRepo{
		DB: db.ConnectToDB(),
	}
}

func (pr *PostgresProfileRepo) GetUserProfile(ctx context.Context, id uint) (*types.Profile, error) {

	var profile types.Profile

	subscriberCount, targetCount, err := pr.GetSubsriptionsCount(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := pr.DB.WithContext(ctx).Where("user_id = ?", id).First(&profile).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		profile = types.Profile{
			UserId:            id,
			SubscriberCount:   uint(targetCount),
			SubscriptionCount: uint(subscriberCount),
		}
		err = pr.DB.WithContext(ctx).Create(&profile).Error
		if err != nil {
			return nil, err
		}
	}

	switch {
	case profile.SubscriberCount != uint(targetCount) || profile.SubscriptionCount != uint(subscriberCount):
		profile.SubscriberCount = uint(targetCount)
		profile.SubscriptionCount = uint(subscriberCount)

		if err := pr.DB.WithContext(ctx).Where("user_id = ?", id).Save(&profile).Error; err != nil {
			return nil, fmt.Errorf("couldn't update profile of the user: %w", err)
		}
	default:
		return &profile, nil
	}

	return &profile, nil
}

func (pr *PostgresProfileRepo) GetSubsriptionsCount(ctx context.Context, id uint) (int64, int64, error) {
	var subscriberCount int64
	var targetCount int64

	if err := pr.DB.WithContext(ctx).Find(&types.User{}, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, 0, fmt.Errorf("user doesn't exists %w", err)
	}

	if err := pr.DB.WithContext(ctx).Model(&types.Subscription{}).Where("subscriber_id = ?", id).Count(&subscriberCount).Error; err != nil {
		return 0, 0, err
	}

	if err := pr.DB.WithContext(ctx).Model(&types.Subscription{}).Where("target_id = ?", id).Count(&targetCount).Error; err != nil {
		return 0, 0, err
	}

	return subscriberCount, targetCount, nil

}
