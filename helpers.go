package main

import (
	"log"
	"math/rand"
)

func getAllMovies() []Movie {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database to retrieve all movies
	rows, err := db.Query("SELECT * FROM movies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Create a slice to hold the Movie structs
	movies := []Movie{}

	// Iterate over the result set and scan the values into a Movie struct
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.ReleaseYear, &movie.Director, &movie.LiveAction)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

func getAllRatingsForUserById(userId int) []Rating {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database for all ratings from user 1
	rows, err := db.Query("SELECT id, user_id, movie_id, rating FROM ratings WHERE user_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows and create Rating objects
	var ratings []Rating
	for rows.Next() {
		var r Rating
		err := rows.Scan(&r.Id, &r.UserId, &r.MovieId, &r.Rating)
		if err != nil {
			log.Fatal(err)
		}
		ratings = append(ratings, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return ratings
}

func generateFakeRatings() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get all movies
	movies := getAllMovies()

	// Generate fake ratings for each movie and insert into database
	for _, movie := range movies {
		// Generate a random rating between 1 and 10
		rating := rand.Intn(10) + 1

		// Create a new Rating object
		newRating := Rating{
			UserId:  1,        // Generate a random user ID between 1 and 100
			MovieId: movie.ID, // Use the ID of the current movie
			Rating:  rating,   // Use the random rating generated above
		}

		// Insert the new rating into the database
		_, err := db.Exec("INSERT INTO ratings (userId, movieId, rating) VALUES (?, ?, ?)", newRating.UserId, newRating.MovieId, newRating.Rating)
		if err != nil {
			log.Fatal(err)
		}
	}
}
