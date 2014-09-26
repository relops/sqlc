package sqlc

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegration(t *testing.T) {

	db, err := sql.Open("sqlite3", "sqlc.db")
	assert.NoError(t, err)

	err = Migrate(db)
	assert.NoError(t, err)

	row, err := Select(bar).From(foo).Where(baz.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	var barScan string
	err = row.Scan(&barScan)
	assert.NoError(t, err)

	assert.Equal(t, "gorp", barScan)

}
