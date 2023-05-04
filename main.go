package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// file, err := os.Open("dataset.csv")
	// if err != nil {
	// 	fmt.Println("Error opening file: ", err)
	// 	return
	// }

	// defer file.Close()

	// reader := csv.NewReader(file)

	// records, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println("Error reading CSV: ", err)
	// 	return
	// }

	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	databaseSetup()
}

func databaseSetup() {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create a table in the database
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS movies (
			id INTEGER PRIMARY KEY,
			title TEXT,
			genre TEXT,
			release_year INTEGER,
			director TEXT,
			animated BOOLEAN,
			liveaction BOOLEAN
		)
	`)

	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("SQLite table created successfully!")
}
