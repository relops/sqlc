package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {

	dbFile := "sqlc.db"

	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)
	assert.NoError(t, err)

	steps := sqlc.LoadBindata(AssetNames(), Asset)
	err = sqlc.Migrate(db, steps)
	assert.NoError(t, err)

	_, err = sqlc.InsertInto(FOO).Set(FOO.BAZ, "quux").Set(FOO.BAR, "gorp").Exec(db)
	assert.NoError(t, err)

	row, err := sqlc.Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	var barScan string
	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "gorp", barScan)

	_, err = sqlc.Update(FOO).Set(FOO.BAR, "porg").Where(FOO.BAZ.Eq("quux")).Exec(db)
	assert.NoError(t, err)

	row, err = sqlc.Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "porg", barScan)

	_, err = sqlc.Delete(FOO).Where(FOO.BAZ.Eq("quux")).Exec(db)
	assert.NoError(t, err)

	row, err = sqlc.Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	err = row.Scan(&barScan)
	assert.Equal(t, err, sql.ErrNoRows)
}
