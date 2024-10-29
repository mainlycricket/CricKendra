.PHONY: build-main build-live run-main run-live watch-main watch-live install-tools backup-db setup-db

build-main:
	go build -o bin/ ./cmd/main-server/

build-live:
	go build -o bin/ ./cmd/live-server/

run-main: build-main
	./bin/main-server

run-live: build-live
	./bin/live-server

install-tools:
	go install github.com/cespare/reflex@latest

watch-main:
	reflex -r '\.go$$' -s -- sh -c 'make build-main && ./bin/main-server'

watch-live:
	reflex -r '\.go$$' -s -- sh -c './bin/live-server'

backup-db:
	pg_dump -h localhost -U postgres crickendra > db.sql

setup-db:
	psql -h localhost -U postgres -c "DROP DATABASE IF EXISTS crickendra;"
	psql -h localhost -U postgres -c "CREATE DATABASE crickendra;"
	psql -h localhost -U postgres -d crickendra -f db.sql
