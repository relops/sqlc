MIGRATION_DIR := test/db
MIGRATION_SCRIPTS := $(foreach dir, $(MIGRATION_DIR), $(wildcard $(dir)/*))

tests:  test/test.db \
		test/generated/generic/objects.go \
		test/generated/sqlite/objects.go \
		test/generated/mysql/objects.go \
		test/generated/postgres/objects.go \
		sqlc/fields.go \
		sqlc/schema.go
	go run test/migrate_db.go
	go test -v ./...

test/test.db: test/migration_steps.go
	go run test/migrate_db.go

test/generated/mysql:
	mkdir -p $@

test/generated/sqlite:
	mkdir -p $@

test/generated/generic:
	mkdir -p $@

test/generated/postgres:
	mkdir -p $@

test/generated/generic/objects.go: test/test.db test/migration_steps.go test/generated/generic main.go
	go run main.go -p generic -o $@ -f test/test.db -t sqlite

test/generated/sqlite/objects.go: test/test.db test/migration_steps.go test/generated/sqlite main.go
	go run main.go -p sqlite -o $@ -f test/test.db -t sqlite

test/generated/mysql/objects.go: test/migration_steps.go test/generated/mysql main.go
	go run main.go -p mysql -o $@ -u "sqlc:sqlc@/sqlc" -t mysql

test/generated/postgres/objects.go: test/migration_steps.go test/generated/postgres main.go
	go run main.go -p postgres -o $@ -u "postgres://sqlc:sqlc@localhost/sqlc?sslmode=disable" -t postgres

sqlc/fields.go: sqlc/tmpl/fields.tmpl sqlc/field_generator.go
	go run sqlc/field_generator.go

sqlc/schema.go: sqlc/fields.go sqlc/tmpl/schema.tmpl
	go-bindata -pkg=sqlc -o=$@ sqlc/tmpl

test/migration_steps.go: $(MIGRATION_SCRIPTS)
	go-bindata -pkg=test -o=$@ test/db/mysql test/db/sqlite test/db/postgres
