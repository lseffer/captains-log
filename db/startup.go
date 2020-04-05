//+build !testsuite

package db

// Startup ...
// Run the database startup logic
func Startup() {
	connection := getDatabaseConnection()
	revisions := parseRevisions("sql")
	currentRevision := getCurrentRevision(connection)
	applyRevisions(connection, currentRevision, revisions)
}
