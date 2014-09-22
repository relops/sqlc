package sqlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	c := NewContext()
	assert.NotNil(t, c)

	foo := Table("foo")
	bar := Field("bar")

	c.Select(bar).From(foo)

	sql, err := c.Render()
	assert.NoError(t, err)

	assert.Equal(t, "SELECT bar FROM foo", sql)
}
