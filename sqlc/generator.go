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

type Options struct {
	File    string `short:"f" long:"file" description:"The path to the sqlite file" required:"true"`
	Output  string `short:"o" long:"output" description:"The path to save the generated objects to" required:"true"`
	Package string `short:"p" long:"package" description:"The package to put the generated objects into" required:"true"`
}

func Generate(db *sql.DB, opts *Options) error {

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
	params["Package"] = opts.Package

	m := template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
	}

	schemaBin, _ := sqlc_tmpl_schema_tmpl()
	t := template.Must(template.New("schema.tmpl").Funcs(m).Parse(string(schemaBin)))

	var b bytes.Buffer
	t.Execute(&b, params)

	// TODO unhardcode this
	if err := ioutil.WriteFile(opts.Output, b.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("Could not write templated file: %s", err)
		return err
	}

	return nil
}
