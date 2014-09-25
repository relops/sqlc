package sqlc

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

var foo = Table("foo")
var bar = Varchar("bar")
var baz = Varchar("baz")

var rendered = []struct {
	Constructed Renderable
	Expected    string
}{
	{
		Select().From(foo),
		"SELECT * FROM foo",
	},
	{
		Select(bar, baz).From(foo),
		"SELECT bar, baz FROM foo",
	},
	{
		Select(bar).From(foo).Where(baz.Eq("quux")),
		"SELECT bar FROM foo WHERE baz = ?",
	},
	{
		Select().From(Select(bar).From(foo)),
		"SELECT * FROM (SELECT bar FROM foo)",
	},
}

var trees = []struct {
	Constructed Selectable
	Expected    selection
}{
	{
		Select().From(foo),
		selection{selection: table{name: "foo"}},
	},
	{
		Select(bar, baz).From(foo),
		selection{selection: table{name: "foo"}, projection: []Column{bar, baz}},
	},
	{
		Select().From(Select(bar).From(foo)),
		selection{
			selection: &selection{
				selection: table{name: "foo"}, projection: []Column{bar},
			},
		},
	},
}

func TestTrees(t *testing.T) {
	for _, tree := range trees {
		assert.Equal(t, &tree.Expected, tree.Constructed)
	}
}

func TestRendered(t *testing.T) {
	for _, rendered := range rendered {
		assert.Equal(t, rendered.Expected, rendered.Constructed.String())
	}
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

	row, err := Select(bar).From(foo).Where(baz.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	var barScan string
	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "gorp", barScan)

}
