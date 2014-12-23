package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/relops/sqlc/sqlc"
	. "github.com/relops/sqlc/test/generated/postgres"
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

	deletePostgres(t, db)
	generic.RunCallRecordGroupTests(t, sqlc.Postgres, db)

	// POSTGRES specific integration tests

	deletePostgres(t, db)
	testReturning(t, db)
}

func deletePostgres(t *testing.T, db *sql.DB) {
	_, err := db.Exec("TRUNCATE call_records;")
	_, err = db.Exec("TRUNCATE key_value;")
	assert.NoError(t, err)
}

// POSTGRES specific RETURNING syntax
func testReturning(t *testing.T, db *sql.DB) {

	d := sqlc.Postgres

	var returnedId, selectedId int

	row, err := sqlc.InsertInto(KEY_VALUE).SetString(KEY_VALUE.VALUE, "foo").Returning(KEY_VALUE.ID).Fetch(d, db)
	assert.NoError(t, err)

	err = row.Scan(&returnedId)
	assert.NoError(t, err)

	// TODO ideally we need to implement LIMIT for SELECT
	row, err = sqlc.Select(KEY_VALUE.ID).From(KEY_VALUE).Where(KEY_VALUE.VALUE.Eq("foo")).QueryRow(d, db)
	assert.NoError(t, err)

	err = row.Scan(&selectedId)
	assert.NoError(t, err)

	assert.Equal(t, returnedId, selectedId, "Returned id was not the same as the subsequent selected id")

}
