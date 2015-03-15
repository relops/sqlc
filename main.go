package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shutej/sqlc/sqlc"
	"log"
	"os"
)

var VERSION string = "0.1.5"

var opts sqlc.Options
var parser = flags.NewParser(&opts, flags.Default)

func init() {
	opts.Version = printVersionAndExit
}

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

	err = sqlc.Generate(db, VERSION, &opts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func dataSource() (*sql.DB, sqlc.Dialect, error) {

	d, err := opts.DbType()
	if err != nil {
		return nil, sqlc.Sqlite, err
	}

	switch d {
	case sqlc.Sqlite:
		db, err := sql.Open("sqlite3", opts.File)
		return db, d, err
	case sqlc.MySQL:
		db, err := sql.Open("mysql", opts.Url)
		return db, d, err
	case sqlc.Postgres:
		db, err := sql.Open("postgres", opts.Url)
		return db, d, err
	default:
		return nil, sqlc.Sqlite, errors.New("Invalid Db type")
	}
}

func printVersionAndExit() {
	fmt.Fprintf(os.Stderr, "%s %s\n", "sqlc", VERSION)
	os.Exit(0)
}
