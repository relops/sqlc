package sqlc

import (
	"database/sql"
	log "github.com/cihub/seelog"
)

func LoadBindata(assets []string, r func(string) ([]byte, error)) []string {

	steps := make([]string, len(assets))

	for i, asset := range assets {
		stepBin, _ := r(asset)
		steps[i] = string(stepBin)
	}

	return steps
}

func Migrate(db *sql.DB, steps []string) error {

	var current int
	// TODO Eat our own dogfood
	currentVersion := "SELECT MAX(version) FROM schema_versions"

	if err := db.QueryRow(currentVersion).Scan(&current); err != nil {
		// We assume that this error arises because the table does not exist
		if err := initTable(db); err != nil {
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

		if _, err := txn.Exec(insertVersion, version); err != nil {
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

func initTable(db *sql.DB) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := txn.Exec(versionsTable); err != nil {
		txn.Rollback()
		return err
	}

	version := 0
	if _, err := txn.Exec(insertVersion, version); err != nil {
		txn.Rollback()
		return err
	}

	return txn.Commit()
}

const versionsTable = `
CREATE TABLE schema_versions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    version INT NOT NULL,                
    ts TIMESTAMP DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
);`

const insertVersion = "INSERT INTO schema_versions (version) VALUES (?);"
