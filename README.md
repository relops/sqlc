sqlc
----

`sqlc` generates SQL for you:

	c := NewContext()	

	foo := Table("foo")
	bar := Field("bar")

	sql, _ := c.Select(bar).From(foo).Render() // Renders `SELECT bar FROM foo`
