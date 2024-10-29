package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertSeason(ctx context.Context, db *pgxpool.Pool, season *models.Season) error {
	query := `INSERT INTO seasons (season) VALUES($1)`

	cmd, err := db.Exec(ctx, query, season.Season)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert season")
	}

	return nil
}
