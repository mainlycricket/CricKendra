package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertFallOfWicketsEntries(ctx context.Context, db DB_Exec, entries []models.FallOfWicket) error {
	query := `INSERT INTO fall_of_wickets (innings_id, batter_id, team_runs, wicket_number) VALUES($1, $2, $3, $4)`

	batch := pgx.Batch{}

	for _, entry := range entries {
		_ = batch.Queue(query, entry.InningsId, entry.BatterId, entry.TeamRuns, entry.WicketNumber)
	}

	return db.SendBatch(ctx, &batch).Close()
}
