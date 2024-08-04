package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "host=localhost user=user password=password dbname=social_media port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("couldn't connect to the database!")
	}

	return db
}
