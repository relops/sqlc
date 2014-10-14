package sqlc

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var foo = Table("foo")
var quux = Table("quux")
var bar = String(foo, "bar")
var baz = String(foo, "baz")
var id = String(quux, "id")
var col = String(quux, "col")

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
		"SELECT foo.bar, foo.baz FROM foo",
	},
	{
		Select(bar.Div(5)).From(foo),
		"SELECT foo.bar / 5 FROM foo",
	},
	{
		Select(bar.Div(1.72)).From(foo),
		"SELECT foo.bar / 1.72 FROM foo",
	},
	{
		Select(Count().Cast("REAL")).From(foo),
		"SELECT CAST(COUNT(*) AS REAL) FROM foo",
	},
	{
		Select(bar.Div(5).As("result")).From(foo),
		"SELECT foo.bar / 5 AS result FROM foo",
	},
	{
		Select(Count().Cast("REAL").Div(20)).From(foo),
		"SELECT CAST(COUNT(*) AS REAL) / 20 FROM foo",
	},
	{
		Select(Count().Cast("REAL").Div(20).Ceil().Cast("INT")).From(foo),
		"SELECT CAST(CEIL(CAST(COUNT(*) AS REAL) / 20) AS INT) FROM foo",
	},
	{
		Select(bar.As("x"), baz.As("y")).From(foo),
		"SELECT foo.bar AS x, foo.baz AS y FROM foo",
	},
	{
		// This is more verbose that it needs to be
		// generally speaking apps would use generated objects, but this example uses the runtime API
		// to create aliased objects
		Select(foo.As("f").StringField("bar").As("x")).From(foo.As("f")),
		"SELECT f.bar AS x FROM foo AS f",
	},
	{
		Select(bar).From(foo).Where(baz.Eq("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz = ?",
	},
	{
		Select(bar).From(foo).Where(baz.Lt("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz < ?",
	},
	{
		Select(bar).From(foo).Where(baz.Le("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz <= ?",
	},
	{
		Select(bar).From(foo).Where(baz.Gt("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz > ?",
	},
	{
		Select(bar).From(foo).Where(baz.Ge("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz >= ?",
	},
	{
		SelectCount().From(foo).Where(baz.Eq("quux")),
		"SELECT COUNT(*) FROM foo WHERE foo.baz = ?",
	},
	{
		Select(bar).From(foo).Where(baz.Eq("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz = ?",
	},
	{
		Select(bar).From(foo).GroupBy(bar).OrderBy(bar),
		"SELECT foo.bar FROM foo GROUP BY foo.bar ORDER BY foo.bar",
	},
	{
		Select(bar, Count()).From(foo).GroupBy(bar),
		"SELECT foo.bar, COUNT(*) FROM foo GROUP BY foo.bar",
	},
	{
		Select(bar).From(foo).Join(quux).On(id.IsEq(bar)),
		"SELECT foo.bar FROM foo JOIN quux ON quux.id = foo.bar",
	},
	{
		// This is more verbose that it needs to be
		// generally speaking apps would use generated objects, but this example uses the runtime API
		// to create aliased objects
		Select(bar).From(foo.As("f")).Join(quux.As("q")).On(quux.As("q").StringField("id").IsEq(foo.As("f").StringField("bar"))),
		"SELECT f.bar FROM foo AS f JOIN quux AS q ON q.id = f.bar",
	},
	{
		Select(bar, col).From(foo).Join(quux).On(id.IsEq(bar), col.IsEq(baz)),
		"SELECT foo.bar, quux.col FROM foo JOIN quux ON (quux.id = foo.bar AND quux.col = foo.baz)",
	},
	{
		Select(bar).From(foo).LeftOuterJoin(quux).On(id.IsEq(bar)),
		"SELECT foo.bar FROM foo LEFT OUTER JOIN quux ON quux.id = foo.bar",
	},
	{
		Select().From(Select(bar).From(foo)),
		"SELECT * FROM (SELECT foo.bar FROM foo)",
	},
	{
		InsertInto(foo).SetString(bar, "quux"),
		"INSERT INTO foo (bar) VALUES (?)",
	},
	{
		Update(foo).SetString(bar, "quux").Where(baz.Eq("gorp")),
		"UPDATE foo SET bar = ? WHERE foo.baz = ?",
	},
	{
		Delete(foo).Where(baz.Eq("gorp")),
		"DELETE FROM foo WHERE foo.baz = ?",
	},
}

var selectTrees = []struct {
	Constructed Selectable
	Expected    selection
}{
	{
		Select().From(foo),
		selection{
			selection: table{name: "foo", fieldMap: make(map[string]Field)},
		},
	},
	{
		Select(bar, baz).From(foo),
		selection{
			selection: table{
				name:     "foo",
				fieldMap: make(map[string]Field),
			}, projection: []Field{bar, baz},
		},
	},
	{
		Select(bar).From(foo).Join(quux).On(id.IsEq(bar)),
		selection{
			selection:  table{name: "foo", fieldMap: make(map[string]Field)},
			projection: []Field{bar},
			joinTarget: nil,
			joinType:   NotJoined,
			joins: []join{
				join{
					target:   table{name: "quux", fieldMap: make(map[string]Field)},
					joinType: Join,
					conds:    []JoinCondition{id.IsEq(bar)},
				},
			},
		},
	},
	{
		Select(bar).From(foo).GroupBy(bar).OrderBy(bar),
		selection{
			selection:  table{name: "foo", fieldMap: make(map[string]Field)},
			projection: []Field{bar},
			groups:     []Field{bar},
			ordering:   []Field{bar},
		},
	},
	{
		Select().From(Select(bar).From(foo)),
		selection{
			selection: &selection{
				selection:  table{name: "foo", fieldMap: make(map[string]Field)},
				projection: []Field{bar},
			},
		},
	},
}

var insertTrees = []struct {
	Constructed InsertSetStep
	Expected    insert
}{
	{
		InsertInto(foo).SetString(bar, "quux"),
		insert{
			table: table{name: "foo", fieldMap: make(map[string]Field)},
			bindings: []TableFieldBinding{
				TableFieldBinding{
					Field: bar,
					Value: "quux",
				},
			},
		},
	},
}

var updateTrees = []struct {
	Constructed Executable
	Expected    update
}{
	{
		Update(foo).SetString(bar, "quux").Where(baz.Eq("gorp")),
		update{
			table: table{name: "foo", fieldMap: make(map[string]Field)},
			bindings: []TableFieldBinding{
				TableFieldBinding{
					Field: bar,
					Value: "quux",
				},
			},
			predicate: []Condition{
				Condition{
					Binding: FieldBinding{
						Field: baz,
						Value: "gorp",
					},
					Predicate: EqPredicate,
				},
			},
		},
	},
}

var deleteTrees = []struct {
	Constructed Executable
	Expected    deletion
}{
	{
		Delete(foo).Where(baz.Eq("gorp")),
		deletion{
			table: table{name: "foo", fieldMap: make(map[string]Field)},
			predicate: []Condition{
				Condition{
					Binding: FieldBinding{
						Field: baz,
						Value: "gorp",
					},
					Predicate: EqPredicate,
				},
			},
		},
	},
}

func TestDeleteTrees(t *testing.T) {
	for _, tree := range deleteTrees {
		assert.Equal(t, &tree.Expected, tree.Constructed)
	}
}

func TestUpdateTrees(t *testing.T) {
	for _, tree := range updateTrees {
		assert.Equal(t, &tree.Expected, tree.Constructed)
	}
}

func TestInsertTrees(t *testing.T) {
	for _, tree := range insertTrees {
		assert.Equal(t, &tree.Expected, tree.Constructed)
	}
}

func TestSelectTrees(t *testing.T) {
	for _, tree := range selectTrees {
		assert.Equal(t, &tree.Expected, tree.Constructed)
	}
}

func TestRendered(t *testing.T) {
	for _, rendered := range rendered {
		// TODO This does a substring match because of the potential random alias name,
		// should probably figure out a way to strip out the alias
		r := rendered.Constructed.String(Sqlite)
		contains := strings.Contains(r, rendered.Expected)
		if !contains {
			assert.Equal(t, rendered.Expected, r)
		}

	}
}
