package sqlc

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	c := NewContext()
	assert.NotNil(t, c)

	foo := Table("foo")
	bar := Varchar("bar")
	baz := Varchar("baz")

	c.Select(bar).From(foo).Where(baz.Eq("quux"))

	sql, err := c.RenderSQL()
	assert.NoError(t, err)

	assert.Equal(t, "SELECT bar FROM foo WHERE baz = ?", sql)

	stmt, placeholders, err := c.Build()
	assert.NoError(t, err)

	assert.Equal(t, "SELECT bar FROM foo WHERE baz = ?", stmt)
	assert.Equal(t, []interface{}{"quux"}, placeholders)

}

func TestSelectStar(t *testing.T) {

	foo := Table("foo")

	c := NewContext()

	c.Select().From(foo)
	sql, err := c.RenderSQL()
	assert.NoError(t, err)
	assert.Equal(t, "SELECT * FROM foo", sql)
}

func TestIntegration(t *testing.T) {

	db, err := sql.Open("sqlite3", "sqlc.db")
	assert.NoError(t, err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS foo (baz TEXT PRIMARY KEY, bar TEXT);`)
	assert.NoError(t, err)

	_, err = db.Exec(`DELETE FROM foo;`)
	assert.NoError(t, err)

	_, err = db.Exec(`INSERT INTO foo (baz, bar) VALUES (?,?)`, "quux", "gorp")
	assert.NoError(t, err)

	foo := Table("foo")
	bar := Varchar("bar")
	baz := Varchar("baz")

	c := NewContext()

	row, err := c.Select(bar).From(foo).Where(baz.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	var barScan string
	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "gorp", barScan)

}
