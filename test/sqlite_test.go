package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/relops/sqlc"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var foo = Table("foo")
var quux = Table("quux")
var bar = Varchar(foo, "bar")
var baz = Varchar(foo, "baz")
var id = Varchar(quux, "id")

func TestIntegration(t *testing.T) {

	dbFile := "sqlc.db"

	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)
	assert.NoError(t, err)

	steps := LoadBindata(AssetNames(), Asset)
	err = Migrate(db, steps)
	assert.NoError(t, err)

	_, err = InsertInto(foo).Set(baz, "quux").Set(bar, "gorp").Exec(db)
	assert.NoError(t, err)

	row, err := Select(bar).From(foo).Where(baz.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	var barScan string
	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "gorp", barScan)

	_, err = Update(foo).Set(bar, "porg").Where(baz.Eq("quux")).Exec(db)
	assert.NoError(t, err)

	row, err = Select(bar).From(foo).Where(baz.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "porg", barScan)

}
