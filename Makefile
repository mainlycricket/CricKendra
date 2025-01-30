include secrets.sh

.PHONY: build-main	
build-main:
	go build -o bin/ ./cmd/main-server/

.PHONY: build-live
build-live:
	go build -o bin/ ./cmd/live-server/

.PHONY: build-csv-parser
build-csv-parser:
	go build -o bin/ ./cmd/csv-parser/

.PHONY: run-main
run-main: build-main
	./bin/main-server

.PHONY: run-live
run-live: build-live
	./bin/live-server

.PHONY: run-csv-parser
run-csv-parser: build-csv-parser
	./bin/csv-parser
	$(MAKE) sync-db-stats

.PHONY: time-csv-parser
time-csv-parser: reset-db build-csv-parser 
	time ./bin/csv-parser
	$(MAKE) sync-db-stats

.PHONY: install-tools
install-tools:
	go install github.com/cespare/reflex@latest

.PHONY: watch-main
watch-main:
	reflex -r '\.go$$' -s -- sh -c 'make build-main && ./bin/main-server'

.PHONY: watch-live
watch-live:
	reflex -r '\.go$$' -s -- sh -c './bin/live-server'
	
DB_HOST ?= $(POSTGRES_HOST)
DB_USER ?= $(POSTGRES_USER)
DB_NAME ?= $(POSTGRES_DB)
SEED_FILES := cricsheet_people continents host_nations cities grounds tournaments

.PHONY: drop-db
drop-db:
	psql -h $(DB_HOST) -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

.PHONY: create-db
create-db:
	psql -h $(DB_HOST) -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME);"

.PHONY: seed-db
seed-db:
	@for file in $(SEED_FILES); do \
		if [ -f "db_files/seed_data/$$file.csv" ]; then \
			psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "\COPY $$file FROM 'db_files/seed_data/$$file.csv' WITH (FORMAT csv, HEADER true);" || exit 1; \
			if [ $$file != cricsheet_people ]; then \
			    psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "SELECT setval('$${file}_id_seq', (SELECT MAX(id) FROM $$file)+1);"; \
			fi; \
		else \
			echo "Warning: $$file.csv not found in db_files/seed_data/"; \
		fi \
	done	
	
.PHONY: list-seed-files
list-seed-files:
	@echo "Available seed files:"
	@ls -1 db_files/seed_data/*.csv | sed 's|.*/||' | sed 's|.csv$$||'

.PHONY: setup-db
setup-db: drop-db create-db
	psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f ./db_files/full_db.sql
	
.PHONY: backup-db
backup-db:
	pg_dump -h $(DB_HOST) -U $(DB_USER) $(DB_NAME) > ./db_files/full_db.sql
	pg_dump -h $(DB_HOST) -U $(DB_USER) --schema-only $(DB_NAME) > ./db_files/schema.sql
	
.PHONY: reset-db
reset-db: drop-db create-db
	psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f ./db_files/schema.sql
	$(MAKE) seed-db

.PHONY: sync-db-stats
sync-db-stats: 
	psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f ./db_files/sync_player_db_stats.sql

.PHONY: test-postman
test-postman:
	postman collection run $(POSTMAN_COLLECTION_ID)

.PHONY: test-code
test-code:
	$(MAKE) backup-db
	psql -h $(DB_HOST) -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME)_test;"
	psql -h $(DB_HOST) -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)_test;"
	psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME)_test -f ./db_files/full_db.sql
	go test ./...
	psql -h $(DB_HOST) -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME)_test;"
