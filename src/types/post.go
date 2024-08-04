package types

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Content   string
	Author    uint
}
