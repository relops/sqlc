// +build ignore

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relops/sqlc/sqlc"
	"github.com/relops/sqlc/test"
	"log"
	"os"
)

type opts struct {
	driver  string
	url     string
	dialect sqlc.Dialect
	filter  string
}

// TODO paramterize this path
var dbFile = "test/test.db"

var sqlite = opts{
	driver:  "sqlite3",
	url:     dbFile,
	dialect: sqlc.Sqlite,
	filter:  "sqlite",
}

var mysql = opts{
	driver:  "mysql",
	url:     "sqlc:sqlc@/sqlc",
	dialect: sqlc.MySQL,
	filter:  "mysql",
}

var postgres = opts{
	driver:  "postgres",
	url:     "postgres://sqlc:sqlc@localhost/sqlc?sslmode=disable",
	dialect: sqlc.Postgres,
	filter:  "postgres",
}

func main() {
	migrate(sqlite)
	migrate(mysql)
	migrate(postgres)
}

func migrate(o opts) {

	db, err := sql.Open(o.driver, o.url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	prefix := fmt.Sprintf("test/db/%s", o.filter)
	assetNames, _ := test.AssetDir(prefix)

	for i, name := range assetNames {
		assetNames[i] = fmt.Sprintf("%s/%s", prefix, name)
	}

	steps := sqlc.LoadBindata(assetNames, test.Asset)
	err = sqlc.Migrate(db, o.dialect, steps)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
