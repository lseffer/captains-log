//+build !testsuite

package db

import "log"

// Startup ...
// Run the database startup logic
func Startup() {
	connection, err := getDatabaseConnection()
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
