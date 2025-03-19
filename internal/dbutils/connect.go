package dbutils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB_Exec interface {
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

func Connect(ctx context.Context, connectionUrl string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return registerDataTypes(ctx, conn)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return pool, nil
}

func registerDataTypes(ctx context.Context, conn *pgx.Conn) error {
	dataTypeNames := []string{
		"career_stats",
		"dismissal_type",
		"innings_end",
	}

	for _, typeName := range dataTypeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			return fmt.Errorf("failed to register custom type: %s", typeName)
		}
		conn.TypeMap().RegisterType(dataType)
	}

	return nil
}
