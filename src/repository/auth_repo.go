package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthPostgresRepo struct {
	DB *gorm.DB
}

func NewAuthPostgresRepo() *AuthPostgresRepo {
	return &AuthPostgresRepo{
		DB: db.ConnectToDB(),
	}
}

func (ar *AuthPostgresRepo) GetRegister(user *types.User) (*types.User, error) {
	var existingUser types.User

	err := ar.DB.Where("username = ?", user.Username).
		Or("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return nil, fmt.Errorf("user already exists %w", err)
	}

	if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("error checking for existing user: %w", err)
	}

	if err := ar.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ar *AuthPostgresRepo) GetLogin(username, password string) (*types.User, error) {
	var user types.User

	if err := ar.DB.First(&user, "username = ?", username).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}

func (ar *AuthPostgresRepo) GetAdminRegister(username, password string) (*types.User, error) {
	var user types.User
	if err := ar.DB.Model(&user).Where("username = ?", username).Update("is_admin", true).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found %w", err)
	}

	if err := ar.DB.First(&user, "username = ?", username).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found %w", err)
	}
	log.Print("the database user is", user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	return &user, nil

}
