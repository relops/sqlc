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

Status
------

Experimental - this is work in progress. Basically I'm trying to port [JOOQ][] to Go, but I don't know yet whether it will work.

[jooq]: http://jooq.org