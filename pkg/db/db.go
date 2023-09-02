package database

import (
	"log"
	"os"
)

// Make the database file
func MakeDB() {
	if _, err := os.Stat("database.db"); os.IsNotExist(err) {
		file, err := os.Create("database.db")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}

// Delete the database file
func DeleteDB() {
	if _, err := os.Stat("database.db"); os.IsNotExist(err) {
		os.Remove("database.db")
	}
}

// Check if the database file exists
func CheckDB() bool {
	if _, err := os.Stat("database.db"); os.IsNotExist(err) {
		return false
	}
	return true
}
