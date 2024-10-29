package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertHostNation(ctx context.Context, db *pgxpool.Pool, host_nation *models.HostNation) error {
	query := `INSERT INTO host_nations (name, continent_id) VALUES($1, $2)`

	cmd, err := db.Exec(ctx, query, host_nation.Name, host_nation.ContinentId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert host nation")
	}

	return nil
}
