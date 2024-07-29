package main

import (
	"log"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
)

func main() {

	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&types.User{}); err != nil {
		log.Fatal(err)
	}
}
