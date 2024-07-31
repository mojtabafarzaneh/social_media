package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mojtabafarzaneh/social_media/src/config"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	rootCmd.AddCommand(ServeCmd)
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "this command will serve the app",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	e := echo.New()
	config.Set()
	configs := config.Get()
	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)))
}
