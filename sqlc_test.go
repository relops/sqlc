package sqlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var foo = Table("foo")
var quux = Table("quux")
var bar = Varchar(foo, "bar")
var baz = Varchar(foo, "baz")
var id = Varchar(quux, "id")

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
		Select(bar).From(foo).Where(baz.Eq("quux")),
		"SELECT foo.bar FROM foo WHERE foo.baz = ?",
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
		Select(bar).From(foo).Join(quux).On(id.IsEq(bar)),
		"SELECT foo.bar FROM foo JOIN quux ON quux.id = foo.bar",
	},
	{
		Select().From(Select(bar).From(foo)),
		"SELECT * FROM (SELECT foo.bar FROM foo)",
	},
	{
		InsertInto(foo).Set(bar, "quux"),
		"INSERT INTO foo (bar) VALUES (?)",
	},
	{
		Update(foo).Set(bar, "quux").Where(baz.Eq("gorp")),
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
		selection{selection: table{name: "foo"}},
	},
	{
		Select(bar, baz).From(foo),
		selection{selection: table{name: "foo"}, projection: []Field{bar, baz}},
	},
	{
		Select(bar).From(foo).Join(quux).On(id.IsEq(bar)),
		selection{
			selection:  table{name: "foo"},
			projection: []Field{bar},
			joinTarget: nil,
			joins: []join{
				join{
					target:   table{name: "quux"},
					joinType: Join,
					conds:    []JoinCondition{id.IsEq(bar)},
				},
			},
		},
	},
	{
		Select(bar).From(foo).GroupBy(bar).OrderBy(bar),
		selection{
			selection:  table{name: "foo"},
			projection: []Field{bar},
			groups:     []Field{bar},
			ordering:   []Field{bar},
		},
	},
	{
		Select().From(Select(bar).From(foo)),
		selection{
			selection: &selection{
				selection: table{name: "foo"}, projection: []Field{bar},
			},
		},
	},
}

var insertTrees = []struct {
	Constructed InsertSetStep
	Expected    insert
}{
	{
		InsertInto(foo).Set(bar, "quux"),
		insert{
			table: table{name: "foo"},
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
		Update(foo).Set(bar, "quux").Where(baz.Eq("gorp")),
		update{
			table: table{name: "foo"},
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
			table: table{name: "foo"},
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
		assert.Equal(t, rendered.Expected, rendered.Constructed.String())
	}
}
