package repository

import (
	"github.com/mojtabafarzaneh/social_media/src/db"
	"gorm.io/gorm"
)

type Profile interface {
}

type PostgresProfileRepo struct {
	DB *gorm.DB
}

func NewPostgresProfileRepo() *PostgresPostRepo {
	return &PostgresPostRepo{
		DB: db.ConnectToDB(),
	}
}
