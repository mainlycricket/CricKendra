package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertGround(ctx context.Context, db *pgxpool.Pool, ground *models.Ground) error {
	query := `INSERT INTO grounds (name, host_nation_id, city_id) VALUES($1, $2, $3)`

	cmd, err := db.Exec(ctx, query, ground.Name, ground.HostNationId, ground.CityId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert ground")
	}

	return nil
}
