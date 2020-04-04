package db

import (
	"fmt"
)

// InitializeDatabase ...
// Set up database and run revisions
func InitializeDatabase() {
	revisions := parseRevisions()
	fmt.Println(revisions)
	// db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
}
