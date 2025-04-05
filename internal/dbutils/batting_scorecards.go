package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertBattingScorecardEntry(ctx context.Context, db DB_Exec, entry *models.BattingScorecard) error {
	query := `INSERT INTO batting_scorecards (innings_id, batter_id, batting_position, has_batted, runs_scored, balls_faced, minutes_batted, fours_scored, sixes_scored, dismissed_by_id, dismissal_type, dismissal_ball_number, fielder1_id, fielder2_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	cmd, err := db.Exec(ctx, query, entry.InningsId, entry.BatterId, entry.BattingPosition, entry.HasBatted, entry.RunsScored.Int64, entry.BallsFaced.Int64, entry.MinutesBatted.Int64, entry.FoursScored.Int64, entry.SixesScored.Int64, entry.DismissedById, entry.DismissalType, entry.DismissalBallNumber, entry.Fielder1Id, entry.Fielder2Id)

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert batting socrecard entry")
	}

	return err
}

func InsertBattingScorecardEntries(ctx context.Context, db DB_Exec, entries []models.BattingScorecard) error {
	query := `INSERT INTO batting_scorecards (innings_id, batter_id, batting_position, has_batted, runs_scored, balls_faced, minutes_batted, fours_scored, sixes_scored, dismissed_by_id, dismissal_type, dismissal_ball_number, fielder1_id, fielder2_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	batch := pgx.Batch{}

	for _, entry := range entries {
		_ = batch.Queue(query, entry.InningsId, entry.BatterId, entry.BattingPosition, entry.HasBatted, entry.RunsScored.Int64, entry.BallsFaced.Int64, entry.MinutesBatted.Int64, entry.FoursScored.Int64, entry.SixesScored.Int64, entry.DismissedById, entry.DismissalType, entry.DismissalBallNumber, entry.Fielder1Id, entry.Fielder2Id)
	}

	return db.SendBatch(ctx, &batch).Close()
}

func UpdateBatterPositionByInningsId(ctx context.Context, db DB_Exec, input *models.BatterPositionInput) error {
	query := `UPDATE batting_scorecards
				SET batting_position = $1, has_batted = $2
				WHERE innings_id = $3 AND batter_id = $4`

	cmd, err := db.Exec(ctx, query, input.BattingPosition, input.HasBatted, input.InningsId, input.BatterId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to update batting position")
	}

	return nil
}
