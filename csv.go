package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func generateCsvData() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Open the CSV file
	file, err := os.Open("dataset.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Prepare the INSERT statement
	stmt, err := db.Prepare("INSERT INTO movies (id, title, genre, release_year, director, liveaction) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	// Insert each record into the database
	for _, record := range records {

		// Parse the fields from the CSV record
		releaseYear, _ := strconv.Atoi(record[3])
		liveAction, _ := strconv.Atoi(record[5])

		// Create a Movie object
		movie := Movie{
			ID:          record[0],
			Title:       record[1],
			Genre:       record[2],
			ReleaseYear: releaseYear,
			Director:    record[4],
			LiveAction:  liveAction,
		}

		// Insert the Movie object into the database
		_, err = stmt.Exec(movie.ID, movie.Title, movie.Genre, movie.ReleaseYear, movie.Director, movie.LiveAction)
		if err != nil {
			fmt.Println("Error inserting record:", err)
			return
		}
	}

	fmt.Println("CSV data inserted into the database successfully!")

}
