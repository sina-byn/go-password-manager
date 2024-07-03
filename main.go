package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/sina-byn/go-password-manager/cmd"
	"github.com/sina-byn/go-password-manager/db"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(err.Error())
	}

	db := db.ConnectDB()
	defer db.Close()

	cmd.Execute()
}
