package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Log the request body
	// log.Printf("Received request to create user. Request Body: %s\n", body)

	// Parse the JSON request body
	var userReq UserRequest
	err = json.Unmarshal(body, &userReq)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Extract username and password from the request body
	username := userReq.Username
	password := userReq.Password

	// Check if the username or password is empty
	if username == "" || password == "" {
		http.Error(w, "Username or password is missing", http.StatusBadRequest)
		return
	}

	// Create the user
	err = createUser(username, password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return a success message
	fmt.Fprint(w, "User created successfully!")
}

func loginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var loginReq UserRequest
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Extract username and password from the parsed request body
		username := loginReq.Username
		password := loginReq.Password

		// Retrieve the stored password hash from the database based on the username
		storedPasswordHash, err := getUserPasswordHash(db, username)
		if err != nil {
			fmt.Println(username)
			http.Error(w, "Failed to retrieve user credentials", http.StatusInternalServerError)
			return
		}

		// Compare the provided password with the stored password hash
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))
		if err != nil {
			// Passwords do not match
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Passwords match, login is successful
		// Implement your login logic here (e.g., generate authentication token, set session, etc.)
		// Create a token with the username
		tokenString, err := createToken(username)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			return
		}

		response := AuthResponse{
			Message: "Authentication successful",
			Token:   tokenString,
		}

		// Convert the response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
			return
		}

		// Create a new session cookie and set the token as its value
		cookie := http.Cookie{
			Name:     "session",
			Value:    tokenString,
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour), // Set an expiration time for the session cookie
		}

		// Set the session cookie in the HTTP response
		http.SetCookie(w, &cookie)

		// Create the authentication response
		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Write the response body
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func checkUserAuthHandler(w http.ResponseWriter, r *http.Request) {
	movies := getAllMovies()

	// Convert the slice of Movie structs to JSON
	jsonData, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}

	// Set the response headers and write the JSON data to the response body
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "session",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, &cookie)

	response := LogoutResponse{
		Message: "Logout Successful",
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
