package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "file:db.sqlite")

	if err != nil {
		log.Fatalf("Failed to establish connection to database: %v", err)
	}

	DB = db

	query := `
	CREATE TABLE IF NOT EXISTS passwords(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		password VARCHAR(255) NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatalf("Failed to create table \"passwords\": %v", err)
	}

	defer stmt.Close()
	_, err = stmt.Exec()

	if err != nil {
		log.Fatalf("Failed to create table \"passwords\": %v", err)
	}

	query = `
	CREATE TABLE IF NOT EXISTS users(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) UNIQUE,
		password VARCHAR(255)
	);
	`

	stmt, err = db.Prepare(query)

	if err != nil {
		log.Fatalf("Failed to create table \"users\": %v", err)
	}

	defer stmt.Close()
	_, err = stmt.Exec()

	if err != nil {
		log.Fatalf("Failed to create table \"users\": %v", err)
	}

	return db
}
