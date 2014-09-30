package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgres(t *testing.T) {

	db, err := sql.Open("postgres", "postgres://sqlc:sqlc@localhost/sqlc?sslmode=disable")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)
}
