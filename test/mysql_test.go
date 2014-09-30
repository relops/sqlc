package test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/relops/sqlc/sqlc"
	"github.com/relops/sqlc/test/generic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysql(t *testing.T) {
	db, err := sql.Open("mysql", "sqlc:sqlc@/sqlc")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)

	filtered := sqlc.FilterBindata("test/db/mysql", AssetDir)
	steps := sqlc.LoadBindata(filtered, Asset)
	err = sqlc.Migrate(db, sqlc.Sqlite, steps)
	assert.NoError(t, err)

	deleteMySQL(t, db)
	generic.RunCallRecordTests(t, db)

	deleteMySQL(t, db)
	generic.RunCallRecordGroupTests(t, db)
}

func deleteMySQL(t *testing.T, db *sql.DB) {
	_, err := db.Exec("TRUNCATE call_records;")
	assert.NoError(t, err)
}
