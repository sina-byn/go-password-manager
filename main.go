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

	db.ConnectDB()

	enc := utils.Enctryption("password")

	fmt.Println(enc)

	dec := utils.Decryption(enc)

	fmt.Println(dec)

	enc_ := utils.Enctryption("password")

	fmt.Println(enc_)

	_dec := utils.Decryption(enc_)

	fmt.Println(_dec)
}
