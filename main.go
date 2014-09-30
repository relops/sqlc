package main

import (
	"database/sql"
	"github.com/jessevdk/go-flags"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"log"
	"os"
)

var opts sqlc.Options
var parser = flags.NewParser(&opts, flags.Default)

func main() {

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", opts.File)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	opts.Dialect = sqlc.Sqlite

	err = sqlc.Generate(db, &opts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
