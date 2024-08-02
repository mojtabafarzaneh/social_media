package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/config"
	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/handlers"
	"github.com/spf13/cobra"
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
	app := gin.Default()
	config.Set()
	configs := config.Get()
	db.ConnectToDB()
	hc := handlers.NewControler()
	//app.GET("/docs/*", echoSwagger.WrapHandler)
	app.GET("/", hc.ListUserHandler)
	app.GET("/:id", hc.GetUserHandler)
	app.POST("/", hc.InsertUserHandler)
	app.DELETE("/:id", hc.DeleteUserHandler)
	app.PUT("/:id/username", hc.UpdateUsernameHandler)

	app.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
