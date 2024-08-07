package repository

import (
	"context"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type Subscription interface {
	GetAllSubs(ctx context.Context, id string) ([]*types.User, error)
	GetAllSubscribed(ctx context.Context, id uint) ([]*types.User, error)
}

type PostgresSubsRepo struct {
	DB *gorm.DB
}

func NewPostgresSubsRepo() *PostgresSubsRepo {
	return &PostgresSubsRepo{
		DB: db.ConnectToDB(),
	}
}

func (ps *PostgresSubsRepo) GetAllSubs(ctx context.Context, id string) ([]*types.User, error) {
	var subs []*types.Subscription
	var targetId []uint
	var subcripted []*types.User

	if err := ps.DB.WithContext(ctx).Model(&subs).
		Where("subscriber_id = ?", id).
		Pluck("target_id", &targetId).Error; err != nil {
		return nil, err
	}

	if len(targetId) == 0 {
		return nil, nil
	}

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
