package repository

import (
	"errors"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type PostgresMiddlewareRepo struct {
	DB *gorm.DB
}

func NewPostgresMiddlewareRepo() *PostgresMiddlewareRepo {
	return &PostgresMiddlewareRepo{
		DB: db.ConnectToDB(),
	}
}

func (pm *PostgresMiddlewareRepo) GetUserId(id string) (*types.User, error) {
	var user types.User

	if err := pm.DB.First(&user, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}
