package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/social_media/src/config"
	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/handlers"
	"github.com/mojtabafarzaneh/social_media/src/middleware"
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
	mc := middleware.NewControler()
	db.ConnectToDB()
	hc := handlers.NewUserControler()
	//app.GET("/docs/*",)

	//users router
	user := app.Group("/users")
	user.Use(middleware.JWTAuthMiddleware())
	user.GET("/", middleware.IsUserAdminMiddleware(), hc.ListUserHandler)
	user.GET("/:id", hc.GetUserHandler)
	user.POST("/", hc.InsertUserHandler)
	user.DELETE("/:id", hc.DeleteUserHandler)
	user.PUT("/:id/username", hc.UpdateUsernameHandler)

	//posts router
	post := app.Group("/posts")
	pc := handlers.NewPostController()
	post.Use(middleware.JWTAuthMiddleware())
	post.GET("/", pc.ListPostsHandler)
	post.GET("/:id", pc.GetPostHandler)
	post.POST("/:user", mc.IsUserAuthorized(), pc.CreatePostHandler)
	post.PUT("/:user/:id", mc.IsUserAuthorized(), pc.UpdatePostsHandler)
	post.DELETE("/:user/:id", mc.IsUserAuthorized(), pc.DeletePostHandler)

	//subs router
	subs := app.Group("/subs")
	sc := handlers.NewSubsController()
	subs.Use(middleware.JWTAuthMiddleware())
	subs.Use(middleware.JWTAuthMiddleware())
	subs.GET("/subscribers/:id", sc.GetAllSubscribed)
	subs.GET("/subscriptions/:id", sc.GetAllSubscriptions)
	subs.POST("/:subscriber", sc.CreateSubs)

	//profile router
	profile := app.Group("/profile")
	profile.Use(middleware.JWTAuthMiddleware())
	proc := handlers.NewProfileControler()
	profile.GET("/:id", proc.GetUserProfileHandler)

	//auth router
	auth := app.Group("/auth")
	ac := handlers.NewAuthControler()
	auth.POST("/register", ac.RegiserHandler)
	auth.POST("/login", ac.LoginHandler)
	auth.POST("/admin/register", ac.GetAdminRegisterHandler)

	app.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
