package sqlc

import (
	"database/sql"
	"fmt"
	log "github.com/cihub/seelog"
)

func FilterBindata(filter string, r func(string) ([]string, error)) []string {
	assetNames, _ := r(filter)

	for i, name := range assetNames {
		assetNames[i] = fmt.Sprintf("%s/%s", filter, name)
	}

	return assetNames
}

func LoadBindata(assets []string, r func(string) ([]byte, error)) []string {

	steps := make([]string, len(assets))

	for i, asset := range assets {
		stepBin, _ := r(asset)
		steps[i] = string(stepBin)
	}

	return steps
}

func Migrate(db *sql.DB, d Dialect, steps []string) error {

	var current int
	// TODO Eat our own dogfood
	currentVersion := "SELECT MAX(version) FROM schema_versions"

	if err := db.QueryRow(currentVersion).Scan(&current); err != nil {
		// We assume that this error arises because the table does not exist
		if err := initTable(db, d); err != nil {
			log.Errorf("Could not initialize version metadata table: %s", err)
			return err
		} else {
			if err := db.QueryRow(currentVersion).Scan(&current); err != nil {
				log.Errorf("Could not establish current schema version: %s", err)
				return err
			}
		}
	}

	log.Infof("Current DB version: %d", current)

	for i, stmt := range steps {

		version := i + 1

		if version <= int(current) {
			continue
		}

		txn, err := db.Begin()
		if err != nil {
			return err
		}

		log.Infof("Step %d: Applying statement:\n %s", version, stmt)

		_, err = db.Exec(stmt)
		if err != nil {
			log.Error(err)
			return txn.Rollback()
		}

		if _, err := txn.Exec(d.insertVersionSQL(), version); err != nil {
			txn.Rollback()
			return err
		}

		err = txn.Commit()
		if err != nil {
			return err
		}

		log.Infof("Successfully migrated DB to version %d", version)

	}

	return nil
}

func initTable(db *sql.DB, d Dialect) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}

	var versionsTable string
	switch d {
	case Sqlite:
		versionsTable = sqliteVersionsTable
	case MySQL:
		versionsTable = mysqlVersionsTable
	case Postgres:
		versionsTable = postgresVersionsTable
	}

	if _, err := txn.Exec(versionsTable); err != nil {
		txn.Rollback()
		return err
	}

	version := 0

	if _, err := txn.Exec(d.insertVersionSQL(), version); err != nil {
		txn.Rollback()
		return err
	}

	return txn.Commit()
}

func (d Dialect) insertVersionSQL() string {
	switch d {
	case Postgres:
		return insertVersionPostgres
	default:
		return insertVersion
	}
}

const sqliteVersionsTable = `
CREATE TABLE schema_versions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    version INT NOT NULL,                
    ts TIMESTAMP DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
);`

const mysqlVersionsTable = `
CREATE TABLE schema_versions (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
    version INT NOT NULL,                
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

const postgresVersionsTable = `
CREATE TABLE schema_versions (
	id SERIAL PRIMARY KEY,
    version INT NOT NULL,                
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

const insertVersion = "INSERT INTO schema_versions (version) VALUES (?);"
const insertVersionPostgres = "INSERT INTO schema_versions (version) VALUES ($1);"
