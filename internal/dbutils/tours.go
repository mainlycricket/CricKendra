package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertTour(ctx context.Context, db *pgxpool.Pool, tour *models.Tour) error {
	query := `INSERT INTO tours (touring_team_id, host_nations_id, season) VALUES($1, $2, $3)`

	cmd, err := db.Exec(ctx, query, tour.TouringTeamId, tour.HostNationsId, tour.Season)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert tour")
	}

	return nil
}
