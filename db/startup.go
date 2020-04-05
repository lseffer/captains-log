package db

import (
	"database/sql"
	"fmt"
	"log"
)

var updateRevisionQuery = []string{"DELETE FROM REVISIONS WHERE id < ?", "INSERT INTO REVISIONS(id) VALUES(?)"}

// ApplyRevisions ...
// Set up database and run revisions
func ApplyRevisions() {
	connection := getDatabaseConnection()

	revisions := parseRevisions()

	currentRevision := getCurrentRevision(connection)
	log.Println(fmt.Sprintf("Current revision id: %d", currentRevision.id))
	for _, revision := range revisions {
		if revision.id > currentRevision.id {
			statement := prepareStatement(connection, revision.sql)
			_, err := statement.Exec()
			if err != nil {
				log.Panic(fmt.Sprintf("Something went wrong applying revision %d, statement was %s", revision.id, revision.sql))
			}
			updateRevision(connection, revision)
			log.Println(fmt.Sprintf("Applied revision: %d", revision.id))
		} else {
			log.Println(fmt.Sprintf("Revision %d is same or newer than revision %d, nothing to do...", currentRevision.id, revision.id))
		}
	}
}

func getCurrentRevision(connection *sql.DB) revision {
	var rev revision
	rows := connection.QueryRow("SELECT id, '' AS sql FROM revisions LIMIT 1")
	err := rows.Scan(&rev.id, &rev.sql)
	if err != nil {
		return revision{0, ""}
	}
	return rev
}

func updateRevision(connection *sql.DB, rev revision) {
	for _, query := range updateRevisionQuery {
		statement := prepareStatement(connection, query)
		_, err := statement.Exec(rev.id)
		if err != nil {
			log.Panic(fmt.Sprintf("Error updating revision %d. Statement was %s", rev.id, query))
		}
	}
}
