// +build ignore

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"github.com/relops/sqlc/test"
	"os"
)

func main() {
	dbFile := "x.db"
	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		os.Exit(1)
	}

	steps := sqlc.LoadBindata(test.AssetNames(), test.Asset)
	sqlc.Migrate(db, steps)

	err = sqlc.Generate(db)
	if err != nil {
		os.Exit(1)
	}
}
