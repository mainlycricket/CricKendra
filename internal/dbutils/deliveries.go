package dbutils

import (
	"context"
	"errors"
	"fmt"

	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func InsertDelivery(ctx context.Context, db DB_Exec, delivery *models.Delivery) error {
	query := `INSERT INTO deliveries (innings_id, innings_delivery_number, ball_number, over_number, batter_id, bowler_id, non_striker_id, batter_runs, wides, noballs, legbyes, byes, penalty, total_extras, total_runs, bowler_runs, is_four, is_six, player1_dismissed_id, player1_dismissal_type, player2_dismissed_id, player2_dismissal_type, is_pace, bowling_style, is_batter_rhb, is_non_striker_rhb, line, length, ball_type, ball_speed, misc, ww_region, foot_type, shot_type, fielder1_id, fielder2_id, commentary, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39)`

	cmd, err := db.Exec(ctx, query, delivery.InningsId, delivery.InningsDeliveryNumber, delivery.BallNumber, delivery.OverNumber, delivery.BatterId, delivery.BowlerId, delivery.NonStrikerId, delivery.BatterRuns, delivery.Wides, delivery.Noballs, delivery.Legbyes, delivery.Byes, delivery.Penalty, delivery.TotalExtras, delivery.TotalRuns, delivery.BowlerRuns, delivery.IsFour, delivery.IsSix, delivery.Player1DismissedId, delivery.Player1DismissalType, delivery.Player2DismissedId, delivery.Player2DismissalType, delivery.IsPace, delivery.BowlingStyle, delivery.IsBatterRHB, delivery.IsNonStrikerRHB, delivery.Line, delivery.Length, delivery.BallType, delivery.BallSpeed, delivery.Misc, delivery.WwRegion, delivery.FootType, delivery.ShotType, delivery.Fielder1Id, delivery.Fielder2Id, delivery.Commentary, delivery.CreatedAt, delivery.UpdatedAt)

	if cmd.RowsAffected() == 0 {
		return errors.New("failed to insert delivery")
	}

	return err
}

func ReadDeliveriesByMatchInnings(ctx context.Context, db DB_Exec, match_id int64, innings_number int) (responses.MatchCommentaryResponse, error) {
	var response responses.MatchCommentaryResponse
	matchHeader := &response.MatchHeader

	query := fmt.Sprintf(`
		WITH innings_commentary AS (
			SELECT
				deliveries.innings_delivery_number, deliveries.ball_number, deliveries.over_number,
		
				deliveries.batter_id, batter.name, deliveries.bowler_id, bowler.name, deliveries.fielder1_id, fielder1.name, deliveries.fielder2_Id, fielder2.name,
	
				deliveries.wides, deliveries.noballs, deliveries.legbyes, deliveries.byes, deliveries.total_runs, deliveries.is_four, deliveries.is_six,
			
				deliveries.player1_dismissed_id, player1_dismissed.name, deliveries.player1_dismissal_type, bs1.runs_scored, bs1.balls_faced, bs1.fours_scored, bs1.sixes_scored,

				deliveries.player2_dismissed_id, player2_dismissed.name, deliveries.player2_dismissal_type, bs2.runs_scored, bs2.balls_faced, bs2.fours_scored, bs2.sixes_scored,

				deliveries.commentary

			FROM deliveries

			LEFT JOIN innings ON deliveries.innings_id = innings.id
				AND innings.is_super_over = FALSE
				AND innings.innings_number = $2

			LEFT JOIN matches ON innings.match_id = matches.id

			LEFT JOIN players batter ON deliveries.batter_id = batter.id
			LEFT JOIN players bowler ON deliveries.bowler_id = bowler.id
			LEFT JOIN players fielder1 ON deliveries.fielder1_id = fielder1.id
			LEFT JOIN players fielder2 ON deliveries.fielder2_id = fielder2.id

			LEFT JOIN players player1_dismissed ON deliveries.player1_dismissed_id = player1_dismissed.id
			LEFT JOIN players player2_dismissed ON deliveries.player2_dismissed_id = player2_dismissed.id

			LEFT JOIN batting_scorecards bs1 ON bs1.innings_id = deliveries.innings_id AND bs1.batter_id = deliveries.player1_dismissed_id
			LEFT JOIN batting_scorecards bs2 ON bs2.innings_id = deliveries.innings_id AND bs2.batter_id = deliveries.player2_dismissed_id

			WHERE matches.id = $1

			ORDER BY deliveries.innings_delivery_number
		)

		SELECT
			%s,
			(
				SELECT ARRAY_AGG(innings_commentary.*)
				FROM innings_commentary
			)
		FROM matches
		%s
		WHERE matches.id = $1
		GROUP BY %s
	`,
		matchHeaderQuery.selectFields,
		matchHeaderQuery.joins,
		matchHeaderQuery.groupByFields,
	)

	rows := db.QueryRow(ctx, query, match_id, innings_number)

	err := rows.Scan(
		&matchHeader.MatchId, &matchHeader.PlayingLevel, &matchHeader.PlayingFormat, &matchHeader.MatchType, &matchHeader.EventMatchNumber,

		&matchHeader.Season, &matchHeader.StartDate, &matchHeader.EndDate, &matchHeader.StartTime, &matchHeader.IsDayNight, &matchHeader.GroundId, &matchHeader.GroundName, &matchHeader.MainSeriesId, &matchHeader.MainSeriesName,

		&matchHeader.Team1Id, &matchHeader.Team1Name, &matchHeader.Team1ImageUrl, &matchHeader.Team2Id, &matchHeader.Team2Name, &matchHeader.Team2ImageUrl,

		&matchHeader.InningsScores,
		&matchHeader.PlayerAwards,

		&response.Commentary,
	)

	if err != nil {
		return response, err
	}

	return response, nil
}
