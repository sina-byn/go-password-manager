package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sina-byn/go-password-manager/auth"
	"github.com/sina-byn/go-password-manager/db"
	"github.com/sina-byn/go-password-manager/utils"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A command to encrypt and add a password to db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Password argument is required")
		}

		userId := auth.Login()

		password := strings.TrimSpace(args[0])

		encryptedPassword := utils.Encrypt(password)

		stmt, err := db.DB.Prepare(`INSERT INTO passwords (password, user_id) VALUES(?, ?)`)

		if err != nil {
			log.Fatalf("Failed to add new password: %v", err)
		}

		defer stmt.Close()
		_, err = stmt.Exec(encryptedPassword, userId)

		if err != nil {
			log.Fatalf("Failed to add new password: %v", err)
		}

		fmt.Println("Password added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
