package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Define a function for creating a new user in the database
func createUser(username, password string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

// Define a function for hashing a password
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), costFactor)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Helper function to retrieve the stored password hash from the database
func getUserPasswordHash(db *sql.DB, username string) (string, error) {
	// Implement your database query to retrieve the stored password hash based on the username
	// You can use the database/sql package to execute the query and retrieve the result

	// Example query using SQLite
	query := "SELECT password FROM users WHERE username = ?"
	var passwordHash string
	err := db.QueryRow(query, username).Scan(&passwordHash)
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}
