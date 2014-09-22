test: columns.go
	go test -v .

columns.go: tmpl/columns.tmpl column_generator.go
	go run column_generator.go
