package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
)

func InsertBattingPartnershipEntries(ctx context.Context, db DB_Exec, entries []models.BattingPartnership) error {
	query := `INSERT INTO batting_partnerships (innings_id, wicket_number, start_innings_delivery_number, end_innings_delivery_number, batter1_id, batter1_runs, batter1_balls, batter2_id, batter2_runs, batter2_balls, start_team_runs, end_team_runs, start_ball_number, end_ball_number, is_unbeaten) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	batch := pgx.Batch{}

	for _, entry := range entries {
		_ = batch.Queue(query, entry.InningsId, entry.WicketNumber, entry.StartInningsDeliveryNumber, entry.EndInningsDeliveryNumber, entry.Batter1Id, entry.Batter1Runs, entry.Batter1Balls, entry.Batter2Id, entry.Batter2Runs, entry.Batter2Balls, entry.StartTeamRuns, entry.EndTeamRuns, entry.StartBallNumber, entry.EndBallNumber, entry.IsUnbeaten)
	}

	return db.SendBatch(ctx, &batch).Close()
}
