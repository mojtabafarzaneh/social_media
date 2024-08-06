package cmd

import (
	"fmt"
	"log"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(MigrateCmd)
}

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "it will migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		Migration()
	},
}

func Migration() {
	db := db.ConnectToDB()

	if err := db.AutoMigrate(&types.User{}, &types.Post{}, &types.Subscription{}, &types.Profile{}); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("the migration has been done successfully!")
}
