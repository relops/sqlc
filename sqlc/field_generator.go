// +build ignore

package main

import (
	"bytes"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/relops/sqlc/meta"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var logConfig = `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Date(2006-02-01 03:04:05.000) - %Msg%n"/>
	</formats>
</seelog>`

func init() {
	logger, err := log.LoggerFromConfigAsString(logConfig)

	if err != nil {
		fmt.Printf("Could not load seelog configuration: %s\n", err)
		return
	}

	log.ReplaceLogger(logger)
}

type PredicateInfo struct {
	Predicate     string
	FieldFunction string
	JoinFunction  string
}

var preds = []PredicateInfo{
	PredicateInfo{Predicate: "EqPredicate", FieldFunction: "Eq", JoinFunction: "IsEq"},
}

func main() {
	params := make(map[string]interface{})
	params["types"] = meta.Types
	params["predicates"] = preds

	m := template.FuncMap{
		"toLower": strings.ToLower,
	}

	t, err := template.New("fields.tmpl").Funcs(m).ParseFiles("sqlc/tmpl/fields.tmpl")
	if err != nil {
		log.Errorf("Could not open template: %s", err)
		return
	}

	var b bytes.Buffer
	t.Execute(&b, params)

	if err := ioutil.WriteFile("sqlc/fields.go", b.Bytes(), os.ModePerm); err != nil {
		log.Errorf("Could not write templated file: %s", err)
		return
	}

	log.Info("Regenerated fields")
}
