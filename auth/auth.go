package auth

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/sina-byn/go-password-manager/db"
)

func Login() int {
	var id int
	var hashedPassword string
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

	row := db.DB.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err := row.Scan(&id, &username, &hashedPassword)

	if err != nil {
		log.Fatalf("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		log.Fatalf("Invalid credentials")
	}

	return id
}
