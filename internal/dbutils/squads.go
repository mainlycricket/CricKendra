package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertSquadEntry(ctx context.Context, db *pgxpool.Pool, entry *models.Squad) error {
	query := `INSERT INTO squads (player_id, series_id, match_id, is_captain, is_wk, is_debut, playing_status) VALUES($1, $2, $3, $4, $5, $6, $7)`

	cmd, err := db.Exec(ctx, query, entry.PlayerId, entry.SeriesId, entry.MatchId, entry.IsCaptain, entry.IsWk, entry.IsDebut, entry.PlayingStatus)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert squad entry")
	}

	return nil
}
