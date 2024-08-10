package types

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primarykey" json:"Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
	Author    uint      `json:"author"`
}

type PostResponse struct {
	ID        uint      `json:"Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}
