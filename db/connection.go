//+build !testsuite

package db

import (
	"database/sql"
	"log"

	"github.com/mattn/go-sqlite3"
)

func init() {
	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"sqlite3_mod_regexp",
			},
		})
}

// GetDatabaseConnection retrieve the main db connection object
func GetDatabaseConnection() (*sql.DB, error) {
	connection, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared&mode=rwc")
	return connection, err
}

// Startup runs the database revisions
func Startup() {
	connection, err := GetDatabaseConnection()
	if err != nil {
		log.Panic(err)
	}
	revisions, err := parseRevisions("sql")
	if err != nil {
		log.Panic(err)
	}
	currentRevision := getCurrentRevision(connection)
	applyRevisions(connection, currentRevision, revisions)
}
