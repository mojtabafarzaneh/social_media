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

	//users router
	user := app.Group("/users")
	user.GET("/", hc.ListUserHandler)
	user.GET("/:id", hc.GetUserHandler)
	user.POST("/", hc.InsertUserHandler)
	user.DELETE("/:id", hc.DeleteUserHandler)
	user.PUT("/:id/username", hc.UpdateUsernameHandler)

	post := app.Group("/posts")
	post.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"details": "working just fine",
		})
	})

	app.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
