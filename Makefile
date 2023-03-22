-include .env

MIGRATIONS_DIR   = ./sql/migrations/

.PHONY: start
start:
	go run ./cmd/app/main.go

migrate:
	migrate -path "$(MIGRATIONS_DIR)" -database "$(POSTGRES_DSN)" $(filter-out $@,$(MAKECMDGOALS))

db-create-migration:
	migrate create -ext sql -dir "$(MIGRATIONS_DIR)" $(filter-out $@,$(MAKECMDGOALS))

gensql:
	pgxgen crud
	pgxgen sqlc generate

install-tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/tkcrm/pgxgen/cmd/pgxgen@latest

%:
	@:
