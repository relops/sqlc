sqlc
----

[![Build Status](https://travis-ci.org/relops/sqlc.png?branch=master)](https://travis-ci.org/relops/sqlc)
[![GoDoc](http://godoc.org/_?status.png)](http://godoc.org/github.com/relops/sqlc)

A composable type safe fluent API to generate nested and complex SQL queries

Taking heavy inspiration from [JOOQ][], `sqlc` generates SQL queries for you:
	
	var FOO sqlc.TableLike // auto-generated by sqlc by introspecting your DB schema
	
	var db *db.DB // For integration with database/sql
	var d Dialect // Either sqlite, mysql or postgres

	row, err := Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).QueryRow(d, db)

If you don't want to use `database/sql`, you don't have to. `String(Dialect)` is an API call to just produce the SQL string that you use in any way that you want to:

	// Renders `SELECT foo.bar FROM foo WHERE foo.baz = ?`
	sql := Select(FOO.BAR).From(FOO).Where(FOO.BAZ.Eq("quux")).String(d)

Code Generation
---------------

Install the `sqlc` command line tool:

	$ go get github.com/relops/sqlc

Make sure `sqlc` is on your PATH (usually $GOPATH/bin).

Then point `sqlc` at your sqlite DB file:

	$ sqlc -h
	Usage:
	  sqlc [OPTIONS]

	Application Options:
	  -f, --file=    The path to the sqlite file
	  -u, --url=     The DB URL
	  -o, --output=  The path to save the generated objects to
	  -p, --package= The package to put the generated objects into

	Help Options:
	  -h, --help     Show this help message

Now you can use the generated objects in your app.

Database Support
----------------

* Sqlite
* MySQL
* Postgres

Features
--------

* SELECT ... FROM ... WHERE
* GROUP BY
* ORDER BY
* INSERTs
* UPDATEs
* DELETEs
* INNER JOINS
* Sub queries
* Data types:
  * VARCHAR
  * INTEGER
  * TIMESTAMP
* Functions:
  * AVG
  * MAX
  * MIN
* Statement rendering
* Querying via database/sql
* Code generation of table and field objects from an exising DB schema

Pre-requisites
--------------

The integration tests depend on [go-bindata](https://github.com/jteeuwen/go-bindata).

Status
------

Experimental - this is work in progress. Basically I'm trying to port [JOOQ][] to Go, but I don't know yet whether it will work.

[jooq]: http://jooq.org
