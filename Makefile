MIGRATION_DIR := test/db
MIGRATION_SCRIPTS := $(foreach dir, $(MIGRATION_DIR), $(wildcard $(dir)/*))

test/test.db: test/migration_steps.go
	go run test/migrate_db.go

tests: test/generated/generated_objects.go sqlc/fields.go sqlc/schema.go
	go test -v ./...

test/generated:
	mkdir - $@

test/generated/generated_objects.go: test/generated test/test.db test/object_generator.go
	go run test/object_generator.go

sqlc/fields.go: sqlc/tmpl/fields.tmpl sqlc/field_generator.go
	go run sqlc/field_generator.go

sqlc/schema.go: sqlc/fields.go sqlc/tmpl/schema.tmpl
	go-bindata -pkg=sqlc -o=$@ sqlc/tmpl

test/migration_steps.go: $(MIGRATION_SCRIPTS)
	go-bindata -pkg=test -o=$@ $(MIGRATION_DIR)
