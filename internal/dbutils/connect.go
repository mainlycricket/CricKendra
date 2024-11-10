package dbutils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB_Exec interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
}

func Connect(ctx context.Context, connectionUrl string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return RegisterDataTypes(ctx, conn)
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

func RegisterDataTypes(ctx context.Context, conn *pgx.Conn) error {
	dataTypeNames := []string{
		"career_stats",
	}

	for _, typeName := range dataTypeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(dataType)
	}

	return nil
}
