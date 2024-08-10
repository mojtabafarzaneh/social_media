package main

import (
	"github.com/mojtabafarzaneh/social_media/src/cmd"
)

// @title Social Media API
// @version 1.0
// @description This is a sample server for a social media application.
// @host 172.24.78.105:3000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	cmd.Execute()
}
