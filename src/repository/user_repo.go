package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type User interface {
	ListUser(ctx context.Context) []*types.User
	GetUserByID(ctx context.Context, id string) *types.User
	CreateUser(ctx context.Context, user types.User) ([]*types.User, error)
	DeleteUser(ctx context.Context, id string) []*types.User
	UpdateUsername(ctx context.Context, username string) []*types.User
}

type PostgresRep struct {
	DB *gorm.DB
}

func NewUserPostgresRep() *PostgresRep {
	return &PostgresRep{
		DB: db.ConnectToDB(),
	}
}

func (r *PostgresRep) ListUser(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	err := r.DB.WithContext(ctx).Preload("Post").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresRep) GetUserByID(ctx context.Context, id uuid.UUID) ([]*types.User, error) {
	var user []*types.User

	err := r.DB.Preload("Post").First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *PostgresRep) CreateUser(ctx context.Context, user types.User) ([]*types.User, error) {
	var users = []*types.User{{Username: user.Username, Email: user.Email, Password: user.Password}}

	err := r.DB.Create(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresRep) DeleteUser(ctx context.Context, id uuid.UUID) error {
	var user []*types.User
	if err := r.DB.Delete(user, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostgresRep) UpdateUsername(username string, id uuid.UUID) error {
	var user []*types.User
	err := r.DB.Model(&user).Where("id = ?", id).Update("username", username).Error

	if err != nil {
		return err
	}
	return nil
}
