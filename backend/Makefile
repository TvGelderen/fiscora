include .env

build:
	@go build -o bin/main

run: build
	@./bin/main

seed: build
	@./bin/main -seed

seed-me: build
	@./bin/main -seed-me

create-migration:
	@GOOSE_MIGRATION_DIR=./database/migrations goose postgres ${DB_CONNECTION_STRING} create $(name) sql

migrate:
	@GOOSE_MIGRATION_DIR=./database/migrations goose postgres ${DB_CONNECTION_STRING} up

migrate-down:
	@GOOSE_MIGRATION_DIR=./database/migrations goose postgres ${DB_CONNECTION_STRING} down

redo:
	@GOOSE_MIGRATION_DIR=./database/migrations goose postgres ${DB_CONNECTION_STRING} redo
