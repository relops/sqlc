package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/relops/sqlc/sqlc"
	"github.com/relops/sqlc/test/generic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgres(t *testing.T) {

	db, err := sql.Open("postgres", "postgres://sqlc:sqlc@localhost/sqlc?sslmode=disable")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)

	deletePostgres(t, db)
	generic.RunCallRecordTests(t, sqlc.Postgres, db)
}

func deletePostgres(t *testing.T, db *sql.DB) {
	_, err := db.Exec("TRUNCATE call_records;")
	assert.NoError(t, err)
}
