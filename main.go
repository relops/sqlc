package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
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

	switch opts.Type {
	case "sqlite":
		db, err := sql.Open("sqlite3", opts.File)
		return db, sqlc.Sqlite, err
	case "mysql":
		db, err := sql.Open("mysql", opts.Url)
		return db, sqlc.MySQL, err
	case "postgres":
		db, err := sql.Open("postgres", opts.Url)
		return db, sqlc.Postgres, err
	default:
		return nil, sqlc.Sqlite, errors.New("Invalid Db type")
	}
}
