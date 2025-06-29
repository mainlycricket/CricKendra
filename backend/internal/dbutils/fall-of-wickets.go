package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
)

func InsertFallOfWicketsEntries(ctx context.Context, db DB_Exec, entries []models.FallOfWicket) error {
	query := `INSERT INTO fall_of_wickets (innings_id, innings_delivery_number, batter_id, ball_number, team_runs, wicket_number, dismissal_type) VALUES($1, $2, $3, $4, $5, $6, $7)`

	batch := pgx.Batch{}

	for _, entry := range entries {
		_ = batch.Queue(query, entry.InningsId, entry.InningsDeliveryNumber, entry.BatterId, entry.BallNumber, entry.TeamRuns, entry.WicketNumber, entry.DismissalType)
	}

	return db.SendBatch(ctx, &batch).Close()
}
