package repository

import (
	"github.com/mojtabafarzaneh/social_media/src/db"
	"gorm.io/gorm"
)

type Post interface {
}

type PostgresPostRepo struct {
	DB *gorm.DB
}

func NewPostgresPostRepo() *PostgresPostRepo {
	return &PostgresPostRepo{
		DB: db.ConnectToDB(),
	}
}
