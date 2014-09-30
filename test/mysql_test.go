package test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysql(t *testing.T) {
	db, err := sql.Open("mysql", "sqlc:sqlc@/sqlc")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)
}
