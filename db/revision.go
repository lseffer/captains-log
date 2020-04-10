package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

var updateRevisionQuery = []string{"DELETE FROM REVISIONS WHERE id < ?", "INSERT INTO REVISIONS(id) VALUES(?)"}

type revision struct {
	id  uint16
	sql string
}

func parseRevisions(directory string) ([]revision, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	var result []revision
	for _, file := range files {
		fileName := file.Name()
		id, err := strconv.ParseUint(strings.Replace(fileName, ".sql", "", -1), 10, 16)
		if err != nil {
			return nil, err
		}
		fileContents, err := ioutil.ReadFile(filepath.Join(directory, fileName))
		if err != nil {
			return nil, err
		}
		result = append(result, revision{uint16(id), string(fileContents)})
	}
	return result, nil
}

func applyRevisions(connection *sql.DB, currentRevision revision, revisions []revision) {
	log.Println(fmt.Sprintf("Current revision id: %d", currentRevision.id))
	for _, revision := range revisions {
		if revision.id > currentRevision.id {
			_, err := connection.Exec(revision.sql)
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
		_, err := connection.Exec(query, rev.id)
		if err != nil {
			log.Panic(fmt.Sprintf("Error updating revision %d. Statement was %s", rev.id, query))
		}
	}
}
