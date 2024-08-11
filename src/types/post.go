package types

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
	Author    uuid.UUID `json:"author"`
}

type PostResponse struct {
	ID        uuid.UUID `json:"Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}

type CreatePostParams struct {
	Author  uuid.UUID `json:"author"`
	Content string    `json:"content"`
}

func CreatePostFromParams(params CreatePostParams) *Post {
	return &Post{
		Author:  params.Author,
		Content: params.Content,
	}
}
