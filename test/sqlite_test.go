package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"github.com/relops/sqlc/test/generic"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSqlite(t *testing.T) {

	dbFile := "sqlc.db"

	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)
	assert.NoError(t, err)

	filtered := sqlc.FilterBindata("test/db/sqlite", AssetDir)
	steps := sqlc.LoadBindata(filtered, Asset)
	err = sqlc.Migrate(db, sqlc.Sqlite, steps)
	assert.NoError(t, err)

	deleteSqlite(t, db)
	generic.RunCallRecordTests(t, db)

	deleteSqlite(t, db)
	generic.RunCallRecordGroupTests(t, db)
}

func deleteSqlite(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM call_records;")
	assert.NoError(t, err)
}
