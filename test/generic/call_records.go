package generic

import (
	"database/sql"
	"fmt"
	"github.com/shutej/sqlc/sqlc"
	. "github.com/shutej/sqlc/test/generated/generic"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func RunCallRecordGroupTests(t *testing.T, d sqlc.Dialect, db *sql.DB) {

	records := 100

	for i := 0; i < records; i++ {

		var region = "VT"

		if i%2 == 0 {
			region = "NY"
		}

		imsi := 230023741299234 + i
		_, err := sqlc.InsertInto(CALL_RECORDS).
			SetString(CALL_RECORDS.IMSI, fmt.Sprintf("%d", imsi)).
			SetTime(CALL_RECORDS.TIMESTAMP, time.Now()).
			SetInt(CALL_RECORDS.DURATION, i).
			SetString(CALL_RECORDS.REGION, region).
			SetString(CALL_RECORDS.CALLING_NUMBER, "220082769234739").
			SetString(CALL_RECORDS.CALLED_NUMBER, "275617294783934").
			Exec(d, db)
		assert.NoError(t, err)
	}

	row, err := sqlc.SelectCount().From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("VT")).QueryRow(d, db)
	assert.NoError(t, err)

	var count int
	err = row.Scan(&count)
	assert.NoError(t, err)

	assert.Equal(t, records/2, count)

	q1 := sqlc.Select(
		CALL_RECORDS.REGION,
		CALL_RECORDS.DURATION.Min(),
		CALL_RECORDS.DURATION.Max(),
		CALL_RECORDS.DURATION.Avg()).
		From(CALL_RECORDS).GroupBy(CALL_RECORDS.REGION).OrderBy(CALL_RECORDS.REGION)

	row, err = q1.QueryRow(d, db)
	assert.NoError(t, err)

	var regionScan string
	var min, max int
	var avg float32
	err = row.Scan(&regionScan, &min, &max, &avg)
	assert.NoError(t, err)

	assert.Equal(t, "NY", regionScan)
	assert.Equal(t, 0, min)
	assert.Equal(t, 98, max)
	assert.Equal(t, 49.0, avg)

	row, err = sqlc.SelectCount().From(q1).QueryRow(d, db)
	assert.NoError(t, err)

	err = row.Scan(&count)
	assert.NoError(t, err)

	assert.Equal(t, 2, count) // GROUP BY NY and VT

}

func RunCallRecordTests(t *testing.T, d sqlc.Dialect, db *sql.DB) {

	_, err := sqlc.InsertInto(CALL_RECORDS).
		SetString(CALL_RECORDS.IMSI, "230023741299234").
		SetTime(CALL_RECORDS.TIMESTAMP, time.Now()).
		SetInt(CALL_RECORDS.DURATION, 10).
		SetString(CALL_RECORDS.REGION, "quux").
		SetString(CALL_RECORDS.CALLING_NUMBER, "220082769234739").
		SetString(CALL_RECORDS.CALLED_NUMBER, "275617294783934").
		Exec(d, db)

	assert.NoError(t, err)

	row, err := sqlc.SelectCount().From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(d, db)
	assert.NoError(t, err)

	var count int
	err = row.Scan(&count)
	assert.NoError(t, err)

	assert.Equal(t, 1, count)

	row, err = sqlc.Select(CALL_RECORDS.DURATION).From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(d, db)
	assert.NoError(t, err)

	var durationScan int
	err = row.Scan(&durationScan)
	assert.NoError(t, err)

	assert.Equal(t, 10, durationScan)

	_, err = sqlc.Update(CALL_RECORDS).SetInt(CALL_RECORDS.DURATION, 11).Where(CALL_RECORDS.REGION.Eq("quux")).Exec(d, db)
	assert.NoError(t, err)

	row, err = sqlc.Select(CALL_RECORDS.DURATION).From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(d, db)
	assert.NoError(t, err)

	err = row.Scan(&durationScan)
	assert.NoError(t, err)

	assert.Equal(t, 11, durationScan)

	_, err = sqlc.Delete(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).Exec(d, db)
	assert.NoError(t, err)

	row, err = sqlc.Select().From(CALL_RECORDS).Where(CALL_RECORDS.REGION.Eq("quux")).QueryRow(d, db)
	assert.NoError(t, err)

	err = row.Scan(&durationScan)
	assert.Equal(t, err, sql.ErrNoRows)
}
