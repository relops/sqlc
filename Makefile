test: fields.go
	go test -v .

fields.go: tmpl/fields.tmpl field_generator.go
	go run field_generator.go
