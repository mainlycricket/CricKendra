package dbutils

import (
	"context"
	"errors"

	"github.com/mainlycricket/CricKendra/backend/internal/models"
)

func InsertBowlingScorecardEntry(ctx context.Context, db DB_Exec, entry *models.BowlingScorecard) error {

	query := `INSERT INTO bowling_scorecards (innings_id, bowler_id, bowling_position, wickets_taken, runs_conceded, balls_bowled, maiden_overs, fours_conceded, sixes_conceded, wides_conceded, noballs_conceded) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	cmd, err := db.Exec(ctx, query, entry.InningsId, entry.BowlerId, entry.BowlingPosition, entry.WicketsTaken.Int64, entry.RunsConceded.Int64, entry.BallsBowled.Int64, entry.MaidenOvers.Int64, entry.FoursConceded.Int64, entry.SixesConceded.Int64, entry.WidesConceded.Int64, entry.NoballsConceded.Int64)

	if cmd.RowsAffected() < 0 {
		return errors.New("failed to insert bowling scorecard entry")
	}

	return err
}
