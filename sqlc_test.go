package sqlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	c := NewContext()
	assert.NotNil(t, c)

	foo := Table("foo")
	bar := VarcharField("bar")
	baz := VarcharField("baz")

	c.Select(bar).From(foo).Where(baz.Eq("quux"))

	sql, err := c.Render()
	assert.NoError(t, err)

	assert.Equal(t, "SELECT bar FROM foo WHERE baz = ?", sql)
}
