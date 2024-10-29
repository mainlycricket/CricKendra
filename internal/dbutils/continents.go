package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertContinent(ctx context.Context, db *pgxpool.Pool, continent *models.Continent) error {
	query := `INSERT INTO continents (name) VALUES($1)`

	cmd, err := db.Exec(ctx, query, continent.Name)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert continent")
	}

	return nil
}
