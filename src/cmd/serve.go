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
	hc := handlers.NewUserControler()
	//app.GET("/docs/*", echoSwagger.WrapHandler)

	//users router
	user := app.Group("/users")
	user.GET("/", hc.ListUserHandler)
	user.GET("/:id", hc.GetUserHandler)
	user.POST("/", hc.InsertUserHandler)
	user.DELETE("/:id", hc.DeleteUserHandler)
	user.PUT("/:id/username", hc.UpdateUsernameHandler)

	//posts router
	post := app.Group("/posts")
	pc := handlers.NewPostController()
	post.GET("/", pc.ListPostsHandler)
	post.GET("/:id", pc.GetPostHandler)
	post.POST("/", pc.CreatePostHandler)
	post.PUT("/:id", pc.UpdatePostsHandler)
	post.DELETE("/:id", pc.DeletePostHandler)

	app.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
