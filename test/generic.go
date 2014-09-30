package test

import (
	"database/sql"
	"github.com/relops/sqlc/sqlc"
	. "github.com/relops/sqlc/test/generated/generic"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func runTests(t *testing.T, db *sql.DB) {

	_, err := sqlc.InsertInto(CALL_RECORDS).
		SetString(CALL_RECORDS.IMSI, "230023741299234").
		SetTime(CALL_RECORDS.TIMESTAMP, time.Now()).
		SetInt(CALL_RECORDS.DURATION, 10).
		SetString(CALL_RECORDS.REGION, "quux").
		SetString(CALL_RECORDS.CALLING_NUMBER, "220082769234739").
		SetString(CALL_RECORDS.CALLED_NUMBER, "275617294783934").
		Exec(db)

	assert.NoError(t, err)

	row, err := sqlc.Select(CALL_RECORDS.DURATION).From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	var durationScan int
	err = row.Scan(&durationScan)
	assert.NoError(t, err)

	assert.Equal(t, 10, durationScan)

	_, err = sqlc.Update(CALL_RECORDS).SetInt(CALL_RECORDS.DURATION, 11).Where(CALL_RECORDS.REGION.Eq("quux")).Exec(db)
	assert.NoError(t, err)

	row, err = sqlc.Select(CALL_RECORDS.DURATION).From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	err = row.Scan(&durationScan)
	assert.NoError(t, err)

	assert.Equal(t, 11, durationScan)

	_, err = sqlc.Delete(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).Exec(db)
	assert.NoError(t, err)

	row, err = sqlc.Select(CALL_RECORDS.IMSI).From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(db)
	assert.NoError(t, err)

	err = row.Scan(&durationScan)
	assert.Equal(t, err, sql.ErrNoRows)
}
