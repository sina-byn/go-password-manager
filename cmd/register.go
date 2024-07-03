/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/sina-byn/go-password-manager/db"
	"github.com/sina-byn/go-password-manager/utils"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		username := ""
		password := ""

		for len(username) < 1 {
			fmt.Print("Username: ")
			fmt.Scanln(&username)
		}

		for len(password) < 1 {
			fmt.Print("Password: ")
			fmt.Scanln(&password)
		}

		stmt, err := db.DB.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")

		if err != nil {
			log.Fatalf("Failed to register user: %v", err)
		}

		defer stmt.Close()
		_, err = stmt.Exec(username, utils.HashPassword(password))

		if err != nil {
			log.Fatalf("Failed to register user: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
