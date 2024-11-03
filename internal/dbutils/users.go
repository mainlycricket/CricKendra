package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ReadUserRoleOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::user_role))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	userRoleOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var role string
		err := row.Scan(&role)
		return role, err
	})

	return userRoleOptions, err
}
