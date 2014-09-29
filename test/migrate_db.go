// +build ignore

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"github.com/relops/sqlc/test"
	"os"
)

// TODO paramterize this path
var dbFile = "test/test.db"

func main() {

	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		os.Exit(1)
	}

	steps := sqlc.LoadBindata(test.AssetNames(), test.Asset)
	err = sqlc.Migrate(db, steps)

	if err != nil {
		os.Exit(1)
	}
}
