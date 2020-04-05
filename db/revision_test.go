package db

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestParseRevisions(t *testing.T) {
	// setup test directory
	testDir := "testDir"
	defer os.RemoveAll(testDir)
	fmt.Println(testDir)
	os.Mkdir(testDir, 0770)
	os.Create(filepath.Join(testDir, "1.sql"))
	os.Create(filepath.Join(testDir, "2.sql"))
	os.Create(filepath.Join(testDir, "3.sql"))
	os.Create(filepath.Join(testDir, "5.sql"))
	// Check that the parser returns correct objects
	results := parseRevisions(testDir)
	var expected = []revision{{1, ""}, {2, ""}, {3, ""}, {5, ""}}
	assert.Equal(t, expected, results)
}
func TestParseRevisionsInvalidFiles(t *testing.T) {
	// setup test directory
	testDir := "testDir"
	defer os.RemoveAll(testDir)
	fmt.Println(testDir)
	os.Mkdir(testDir, 0770)
	os.Create(filepath.Join(testDir, "afafs.sql"))
	os.Create(filepath.Join(testDir, "ufiheiu.sql"))
	// Since we don't raise error we expect the parseUint to return the nil int 0
	results := parseRevisions(testDir)
	var expected = []revision{{0, ""}, {0, ""}}
	assert.Equal(t, expected, results)
}

func TestApplyRevisionsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var mockRevisions = []revision{{1, "CREATE BLABLA"}}
	var mockCurrentRevision = revision{0, ""}
	mock.ExpectExec("CREATE BLABLA").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM REVISIONS WHERE id < ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("INSERT INTO REVISIONS").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	applyRevisions(db, mockCurrentRevision, mockRevisions)

	// current revision is newer than what we're trying to run, therefore nothing should happen
	mockRevisions = []revision{{1, "CREATE BLABLA"}}
	mockCurrentRevision = revision{2, ""}

	applyRevisions(db, mockCurrentRevision, mockRevisions)
}

func TestCurrentRevision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// Test correctly returned revision
	rows := sqlmock.NewRows([]string{"id", "sql"}).AddRow(1, "blabla")
	mock.ExpectQuery("SELECT id, '' AS sql FROM revisions LIMIT 1").WillReturnRows(rows)
	rev := getCurrentRevision(db)
	assert.Equal(t, rev.id, uint16(1))
	assert.Equal(t, rev.sql, "blabla")

	// Test we get the 0 revision if the query fails
	mock.ExpectQuery("SELECT id, '' AS sql FROM revisions LIMIT 1").WillReturnError(fmt.Errorf("some err"))
	revErr := getCurrentRevision(db)
	assert.Equal(t, revErr.id, uint16(0))
	assert.Equal(t, revErr.sql, "")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateRevision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var mockRevision = revision{1, "CREATE BLABLA"}
	mock.ExpectExec("DELETE FROM").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO REVISIONS").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	updateRevision(db, mockRevision)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
