#!/bin/sh
set -e

psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "/schema.sql"

TABLES="cricsheet_people continents host_nations cities grounds tournaments"

for table_name in $TABLES; do
    file_name="/seed_data/$table_name.csv"
    psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c \
        "\COPY $table_name FROM '$file_name' WITH CSV HEADER"
    if [ "$table_name" != "cricsheet_people" ]; then
        psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c \
        "SELECT setval('${table_name}_id_seq', (SELECT MAX(id) FROM ${table_name}) + 1);";
    fi
done

echo "Database setup completed!"

