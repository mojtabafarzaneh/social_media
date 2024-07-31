package cmd

import (
	"fmt"
	"log"

	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rootCmd.AddCommand(SeedCmd)
}

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "it will seed the database",
	Run: func(cmd *cobra.Command, args []string) {
		Seed()
	},
}

func Seed() {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("mojtaba7878"), 12)
	if err != nil {
		log.Fatal("couldn't hash the password", err)
	}
	db := db.ConnectToDB()

	for i := 0; i <= 10; i++ {
		user := types.User{Username: fmt.Sprintf("the random mojtaba %d", i),
			Password: string(hashedPassword),
			Email:    fmt.Sprintf("mojtaba%d@gmail.com", i)}

		db.Create(&user)
	}

	fmt.Println("the objects has been created!")
}
