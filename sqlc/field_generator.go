// +build ignore

package main

import (
	"bytes"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/shutej/sqlc/meta"
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
	PredicateInfo{Predicate: "GtPredicate", FieldFunction: "Gt", JoinFunction: "IsGt"},
	PredicateInfo{Predicate: "GePredicate", FieldFunction: "Ge", JoinFunction: "IsGe"},
	PredicateInfo{Predicate: "LtPredicate", FieldFunction: "Lt", JoinFunction: "IsLt"},
	PredicateInfo{Predicate: "LePredicate", FieldFunction: "Le", JoinFunction: "IsLe"},
}

func argifier(arg meta.FunctionInfo) string {
	argN := strings.Count(arg.Expr, "%")
	if argN == 1 {
		return ""
	} else {
		argsToSub := argN - 1
		args := make([]string, argsToSub)
		for i := 0; i < argsToSub; i++ {
			args[i] = fmt.Sprintf("_%d", i)
		}
		return strings.Join(args, ",")
	}
}

func signifier(arg meta.FunctionInfo) string {
	args := argifier(arg)
	if args == "" {
		return ""
	} else {
		return fmt.Sprintf("%s interface{}", args)
	}
}

func injectifier(arg meta.FunctionInfo) string {
	args := argifier(arg)
	if args == "" {
		return ""
	} else {
		return fmt.Sprintf(", %s", args)
	}
}

func main() {
	params := make(map[string]interface{})
	params["types"] = meta.Types
	params["predicates"] = preds
	params["functions"] = meta.Funcs

	m := template.FuncMap{
		"toLower":     strings.ToLower,
		"signifier":   signifier,
		"injectifier": injectifier,
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
