package sqlc

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type TableMeta struct {
	Name   string
	Fields []FieldMeta
}

type FieldMeta struct {
	Name string
	Type string
}

func Generate(db *sql.DB) error {

	rows, err := db.Query("SELECT name FROM sqlite_master where type = 'table' and name <> 'sqlite_sequence';")
	if err != nil {
		return err
	}

	tables := make([]TableMeta, 0)

	for rows.Next() {
		var t TableMeta
		rows.Scan(&t.Name)
		tables = append(tables, t)
	}

	for i, table := range tables {
		pragma := fmt.Sprintf("PRAGMA table_info(%s);", table.Name)
		rows, err = db.Query(pragma)
		if err != nil {
			return err
		}

		fields := make([]FieldMeta, 0)

		for rows.Next() {
			var notNull, pk sql.NullBool
			var id sql.NullInt64
			var colName, colType, defaultValue sql.NullString
			err = rows.Scan(&id, &colName, &colType, &notNull, &defaultValue, &pk)
			if err != nil {
				return err
			}
			field := FieldMeta{Name: colName.String, Type: colType.String}
			fields = append(fields, field)
		}
		tables[i].Fields = fields
	}

	params := make(map[string]interface{})
	params["Tables"] = tables
	// TODO unhardcode this
	params["Package"] = "test"

	m := template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
	}

	schemaBin, _ := sqlc_tmpl_schema_tmpl()
	t := template.Must(template.New("schema.tmpl").Funcs(m).Parse(string(schemaBin)))

	var b bytes.Buffer
	t.Execute(&b, params)

	// TODO unhardcode this
	if err := ioutil.WriteFile("test/generated_objects.go", b.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("Could not write templated file: %s", err)
		return err
	}

	return nil
}
