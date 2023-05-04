package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Define a secret key for JWT signing
var jwtSecret = []byte("your-secret-key")

// Create a new token with JWT signing method and custom claims
func createToken(username string) (string, error) {
	// Create a new token object with JWT signing method and custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
	})

	// Sign the token with the secret key and get the complete encoded token as a string
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware function to check the JWT token from the session cookie
// func authenticate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Retrieve the token from the session cookie
// 		cookie, err := r.Cookie("session")
// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := cookie.Value

// 		// Verify and parse the token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Verify the signing method and return the secret key
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return jwtSecret, nil
// 		})

// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Check if the token is valid
// 		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Token is valid, proceed with the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }
