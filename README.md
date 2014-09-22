sqlc
----

[![Build Status](https://travis-ci.org/relops/sqlc.png?branch=master)](https://travis-ci.org/relops/sqlc)

`sqlc` generates SQL for you:

	c := NewContext()	

	foo := Table("foo")
	bar := Varchar("bar")
	baz := Varchar("baz")

	// Renders `SELECT bar FROM foo WHERE baz = ?`
	sql, _ := c.Select(bar).From(foo).Where(baz.Eq("quux")).Render()
