package cmd

import (
	"fmt"
	"log"

	"github.com/sina-byn/go-password-manager/auth"
	"github.com/sina-byn/go-password-manager/db"
	"github.com/sina-byn/go-password-manager/utils"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log stored passwords",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		userId := auth.Login()

		rows, err := db.DB.Query(`SELECT * FROM passwords WHERE user_id = ?`, userId)
		hasPassowrds := false

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			var id int
			var password string

			hasPassowrds = true

			if err := rows.Scan(&id, &password, &userId); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("ID: %d - password: %s\n", id, utils.Decrypt(password))
		}

		if !hasPassowrds {
			fmt.Println("No passwords were found")
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
