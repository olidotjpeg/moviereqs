package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	ID          string
	Title       string
	Genre       string
	ReleaseYear int
	Director    string
	LiveAction  int
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

const (
	costFactor = 12 // Controls the computational cost of hashing
)

func main() {
	router := mux.NewRouter()

	// Open the SQLite database
	db, err := connectDB()

	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Register the create user endpoint
	router.HandleFunc("/users", createUserHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler(db)).Methods("POST")

	// Create a table in the database
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS movies (
			id TEXT PRIMARY KEY,
			title TEXT,
			genre TEXT,
			release_year INTEGER,
			director TEXT,
			liveaction INTEGER
		)
	`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
