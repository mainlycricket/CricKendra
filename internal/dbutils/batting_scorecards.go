package dbutils

import (
	"context"
	"errors"

	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertBattingScorecardEntry(ctx context.Context, db DB_Exec, entry *models.BattingScorecard) error {
	query := `INSERT INTO batting_scorecards (innings_id, batter_id, batting_position, has_batted, runs_scored, balls_faced, minutes_batted, fours_scored, sixes_scored, dismissed_by_id, dismissal_type, dismissal_ball_number, fielder1_id, fielder2_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	cmd, err := db.Exec(ctx, query, entry.InningsId, entry.BatterId, entry.BattingPosition, entry.HasBatted, entry.RunsScored, entry.BallsFaced, entry.MinutesBatted, entry.FoursScored, entry.SixesScored, entry.DismissedById, entry.DismissalType, entry.DismissalBallNumber, entry.Fielder1Id, entry.Fielder2Id)

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert batting socrecard entry")
	}

	return err
}
