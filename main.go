package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/sina-byn/go-password-manager/db"
	"github.com/sina-byn/go-password-manager/utils"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(err.Error())
	}

	db := db.ConnectDB()
	defer db.Close()

	enc := utils.Encrypt("password")

	fmt.Println(enc)

	dec := utils.Decrypt(enc)

	fmt.Println(dec)

	enc_ := utils.Encrypt("password")

	fmt.Println(enc_)

	_dec := utils.Decrypt(enc_)

	fmt.Println(_dec)
}
