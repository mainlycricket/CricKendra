package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ReadArticleCategories(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::article_category))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	categories, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var category string
		err := row.Scan(&category)
		return category, err
	})

	return categories, err
}

func ReadArticleStatusOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::article_status))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	statusOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var status string
		err := row.Scan(&status)
		return status, err
	})

	return statusOptions, err
}
