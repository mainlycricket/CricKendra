package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertCity(ctx context.Context, db *pgxpool.Pool, city *models.City) error {
	query := `INSERT INTO cities (name, host_nation_id) VALUES($1, $2)`

	cmd, err := db.Exec(ctx, query, city.Name, city.HostNationId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert city")
	}

	return nil
}
