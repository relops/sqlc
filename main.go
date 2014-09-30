package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

	if err := opts.Validate(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db, dialect, err := dataSource()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	opts.Dialect = dialect

	err = sqlc.Generate(db, &opts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func dataSource() (*sql.DB, sqlc.Dialect, error) {
	if opts.File != "" {
		db, err := sql.Open("sqlite3", opts.File)
		return db, sqlc.Sqlite, err
	} else {
		db, err := sql.Open("mysql", opts.Url)
		return db, sqlc.MySQL, err
	}
}
