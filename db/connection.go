//+build !testsuite

package db

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

func getDatabaseConnection() (*sql.DB, error) {
	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"sqlite3_mod_regexp",
			},
		})
	connection, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared&mode=rwc")
	return connection, err
}
