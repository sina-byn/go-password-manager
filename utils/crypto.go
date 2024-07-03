package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
)

func Enctryption(plainText string) string {
	secret := os.Getenv("SECRET")
	aes, err := aes.NewCipher([]byte(secret))

	if err != nil {
		log.Fatalf(err.Error())
	}

	gcm, err := cipher.NewGCM(aes)

	if err != nil {
		log.Fatalf(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)

	if err != nil {
		log.Fatalf(err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func Decryption(cipherText string) string {
	secret := os.Getenv("SECRET")
	aes, err := aes.NewCipher([]byte(secret))

	if err != nil {
		log.Fatalf(err.Error())
	}

	gcm, err := cipher.NewGCM(aes)

	if err != nil {
		log.Fatalf(err.Error())
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)

	if err != nil {
		log.Fatalf(err.Error())
	}

	nonceSize := gcm.NonceSize()
	cipherText = string(cipherBytes)
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)

	if err != nil {
		panic(err)
	}

	return string(plainText)
}
