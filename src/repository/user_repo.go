package repository

import (
	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: db.ConnectToDB(),
	}
}

type User interface {
	ListUser(limit int) *[]types.User
}

func (r *UserRepository) ListUser(limit int) *[]types.User {
	var users []types.User
	r.DB.Limit(limit).Find(&users)
	return &users
}
