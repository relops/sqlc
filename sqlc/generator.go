package sqlc

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var integer = regexp.MustCompile("INT|INTEGER")
var varchar = regexp.MustCompile("VARCHAR")
var ts = regexp.MustCompile("TIMESTAMP|DATETIME")

type TableMeta struct {
	Name   string
	Fields []FieldMeta
}

type FieldMeta struct {
	Name string
	Type string
}

type Options struct {
	File    string `short:"f" long:"file" description:"The path to the sqlite file"`
	Url     string `short:"u" long:"url" description:"The DB URL"`
	Output  string `short:"o" long:"output" description:"The path to save the generated objects to" required:"true"`
	Package string `short:"p" long:"package" description:"The package to put the generated objects into" required:"true"`
	Dialect Dialect
}

func (o *Options) Validate() error {
	if o.File == "" && o.Url == "" {
		return errors.New("Must specify EITHER file path for sqlite OR url to DB")
	}

	if o.File != "" && o.Url != "" {
		return errors.New("Cannot specify BOTH file path for sqlite AND url to DB")
	}
	return nil
}

func Generate(db *sql.DB, opts *Options) error {

	tables, err := opts.Dialect.metadata(db)
	if err != nil {
		return err
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

	if err := ioutil.WriteFile(opts.Output, b.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("Could not write templated file: %s", err)
		return err
	}

	return nil
}

func (d Dialect) metadata(db *sql.DB) ([]TableMeta, error) {
	switch d {
	case Sqlite:
		return sqlite(db)
	case MySQL:
		return mysql(db)
	default:
		return nil, errors.New("Unsupported dialect")
	}
}

func mysql(db *sql.DB) ([]TableMeta, error) {
	schema := "sqlc"
	rows, err := db.Query(mysqlTables, schema)
	if err != nil {
		return nil, err
	}

	tables := make([]TableMeta, 0)

	for rows.Next() {
		var t TableMeta
		rows.Scan(&t.Name)
		tables = append(tables, t)
	}

	for i, table := range tables {

		rows, err = db.Query(mysqlColumns, schema, table.Name)
		if err != nil {
			return nil, err
		}

		fields := make([]FieldMeta, 0)

		for rows.Next() {
			var colName, colType sql.NullString
			err = rows.Scan(&colName, &colType)
			if err != nil {
				return nil, err
			}

			var fieldType string

			if integer.MatchString(colType.String) {
				fieldType = "Int"
			} else if varchar.MatchString(colType.String) {
				fieldType = "String"
			} else if ts.MatchString(colType.String) {
				fieldType = "Time"
			}

			field := FieldMeta{Name: colName.String, Type: fieldType}
			fields = append(fields, field)
		}
		tables[i].Fields = fields
	}

	return tables, nil
}

func sqlite(db *sql.DB) ([]TableMeta, error) {
	rows, err := db.Query("SELECT name FROM sqlite_master where type = 'table' and name NOT IN ('sqlite_sequence','schema_versions');")
	if err != nil {
		return nil, err
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
			return nil, err
		}

		fields := make([]FieldMeta, 0)

		for rows.Next() {
			var notNull sql.NullBool
			var id, pk sql.NullInt64
			var colName, colType, defaultValue sql.NullString
			err = rows.Scan(&id, &colName, &colType, &notNull, &defaultValue, &pk)
			if err != nil {
				return nil, err
			}

			var fieldType string

			if integer.MatchString(colType.String) {
				fieldType = "Int"
			} else if varchar.MatchString(colType.String) {
				fieldType = "String"
			} else if ts.MatchString(colType.String) {
				fieldType = "Time"
			}

			field := FieldMeta{Name: colName.String, Type: fieldType}
			//fmt.Printf("Field type: %s\n", fieldType)
			fields = append(fields, field)
		}
		tables[i].Fields = fields
	}

	return tables, nil
}

const mysqlTables = `
	select table_name
	from information_schema.tables
	where table_schema = ? AND table_name != 'schema_versions';
`

const mysqlColumns = `
	SELECT column_name, UPPER(data_type)
	FROM information_schema.columns 
	WHERE table_schema = ? and table_name = ?;
`
