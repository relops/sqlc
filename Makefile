MIGRATION_DIR := test/db
MIGRATION_SCRIPTS := $(foreach dir, $(MIGRATION_DIR), $(wildcard $(dir)/*))

tests: test/migration_steps.go fields.go
	go test -v ./...

fields.go: tmpl/fields.tmpl field_generator.go
	go run field_generator.go

test/migration_steps.go: $(MIGRATION_SCRIPTS)
	go-bindata -o=$@ $(MIGRATION_DIR)
