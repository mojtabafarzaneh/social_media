package types

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"Id"`
	CreatedAt  time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"not null" json:"updated_at"`
	Username   string         `gorm:"unique;size:100" json:"username"`
	Password   string         `gorm:"size:100" json:"password"`
	Email      string         `gorm:"unique;size:100" json:"email"`
	Post       []Post         `gorm:"foreignKey:Author;constraint:OnDelete:CASCADE"`
	IsAdmin    bool           `gorm:"default:false"`
	Subscriber []Subscription `gorm:"foreignKey:SubscriberID;constraint:OnDelete:CASCADE"`
	Target     []Subscription `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
	Profile    Profile
}
type Profile struct {
	UserId            uuid.UUID
	SubscriberCount   uint `json:"subscriberCount"`
	SubscriptionCount uint `json:"subscriptionCount"`
}

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Post      []Post    `json:"posts"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type AdminRegisterParams struct {
	IsAdmin  bool   `json:"is_admin"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateUsernameParams struct {
	Username string `json:"username"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)

	if err != nil {
		return nil, err
	}

	return &User{
		Username: params.Username,
		Email:    params.Email,
		Password: string(encpw),
	}, nil
}

func IsValidPassword(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil

}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func (p *CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	minPasswordLen := 8

	if len(p.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password atleast should have atleast %d charecter", minPasswordLen)
	}
	if !isEmailValid(p.Email) {
		errors["email"] = fmt.Sprintln("email is invalid")
	}

	return errors

}

func (u *User) ResponseUser() *ResponseUser {

	return &ResponseUser{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Post:      u.Post,
	}
}

func UsersToUserResponses(users []*User) []ResponseUser {
	userResponses := make([]ResponseUser, len(users))
	for i, user := range users {
		userResponses[i] = *user.ResponseUser()
	}
	return userResponses
}
