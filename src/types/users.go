package types

import (
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"unique" json:"username"`
	Password  string `json:"password"`
	Email     string `gorm:"unique" json:"email"`
	Post      []Post `gorm:"foreignKey:Author"`
	IsAdmin   bool
}

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Post      []Post    `json:"posts"`
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

type UpdateUsernameParams struct {
	Username string `json:"username"`
}
