package repository

import (
	"context"

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

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: db.ConnectToDB(),
	}
}

func (r *UserRepository) ListUser(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) ([]*types.User, error) {
	var user []*types.User

	err := r.DB.WithContext(ctx).First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *UserRepository) CreateUser(ctx context.Context, user types.User) ([]*types.User, error) {
	var users = []*types.User{{Username: user.Username, Email: user.Email, Password: user.Password}}

	err := r.DB.Create(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	var user []*types.User
	if err := r.DB.Delete(user, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUsername(username string, id uint) error {
	var user []*types.User
	err := r.DB.Model(&user).Where("id = ?", id).Update("username", username).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAll() ([]types.Post, error) {
	var posts []types.Post
	err := r.DB.Model(&types.Post{}).Preload("User").Find(&posts).Error

	return posts, err
}
