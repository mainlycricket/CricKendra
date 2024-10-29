package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertSeries(ctx context.Context, db *pgxpool.Pool, series *models.Series) error {
	query := `INSERT INTO series (name, is_male, playing_level, playing_format, season, teams_id, host_nations_id, tournament_id, parent_series_id, tour_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	cmd, err := db.Exec(ctx, query, series.Name, series.IsMale, series.PlayingLevel, series.PlayingFormat, series.Season, series.TeamsId, series.HostNationsId, series.TournamentId, series.ParentSeriesId, series.TourId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert series")
	}

	return nil
}
