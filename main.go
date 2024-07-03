package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(err.Error())
	}

	ConnectDB()

	enc := Enctryption("password")

	fmt.Println(enc)

	dec := Decryption(enc)

	fmt.Println(dec)

	enc_ := Enctryption("password")

	fmt.Println(enc_)

	_dec := Decryption(enc_)

	fmt.Println(_dec)
}
