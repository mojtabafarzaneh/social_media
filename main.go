package main

import (
	"github.com/mojtabafarzaneh/social_media/src/cmd"
	"github.com/mojtabafarzaneh/social_media/src/db"
)

func init() {
	db.ConnectToDB()
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
func main() {

	cmd.Execute()
}
