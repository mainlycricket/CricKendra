package dbutils

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
)

func InsertDelivery(ctx context.Context, db DB_Exec, delivery *models.Delivery) error {
	query := `INSERT INTO deliveries (innings_id, innings_delivery_number, ball_number, over_number, batter_id, bowler_id, non_striker_id, batter_runs, wides, noballs, legbyes, byes, penalty, total_extras, total_runs, bowler_runs, is_four, is_six, player1_dismissed_id, player1_dismissal_type, player2_dismissed_id, player2_dismissal_type, is_pace, bowling_style, is_batter_rhb, is_non_striker_rhb, line, length, ball_type, ball_speed, misc, ww_region, foot_type, shot_type, fielder1_id, fielder2_id, commentary, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39)`

	cmd, err := db.Exec(ctx, query, delivery.InningsId, delivery.InningsDeliveryNumber, delivery.BallNumber, delivery.OverNumber, delivery.BatterId, delivery.BowlerId, delivery.NonStrikerId, delivery.BatterRuns, delivery.Wides, delivery.Noballs, delivery.Legbyes, delivery.Byes, delivery.Penalty, delivery.TotalExtras, delivery.TotalRuns, delivery.BowlerRuns, delivery.IsFour, delivery.IsSix, delivery.Player1DismissedId, delivery.Player1DismissalType, delivery.Player2DismissedId, delivery.Player2DismissalType, delivery.IsPace, delivery.BowlingStyle, delivery.IsBatterRHB, delivery.IsNonStrikerRHB, delivery.Line, delivery.Length, delivery.BallType, delivery.BallSpeed, delivery.Misc, delivery.WwRegion, delivery.FootType, delivery.ShotType, delivery.Fielder1Id, delivery.Fielder2Id, delivery.Commentary, delivery.CreatedAt, delivery.UpdatedAt)

	if cmd.RowsAffected() == 0 {
		return errors.New("failed to insert delivery")
	}

	return err
}

// will also update batter & bowler scorecard entries, innings scores
func InsertDeliveryWithScoringData(ctx context.Context, db DB_Exec, input *models.DeliveryScoringInput) error {
	batch := pgx.Batch{}

	// insert delivery
	batch.Queue(`INSERT INTO deliveries (innings_id, innings_delivery_number, ball_number, over_number, batter_id, bowler_id, non_striker_id, batter_runs, wides, noballs, legbyes, byes, penalty, total_extras, total_runs, bowler_runs, is_four, is_six, player1_dismissed_id, player1_dismissal_type, fielder1_id, fielder2_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)`, input.InningsId, input.InningsDeliveryNumber, input.BallNumber, input.OverNumber, input.BatterId, input.BowlerId, input.NonStrikerId, input.BatterRuns, input.Wides, input.Noballs, input.Legbyes, input.Byes, input.Penalty, input.TotalExtras, input.TotalRuns, input.BowlerRuns, input.IsFour, input.IsSix, input.Player1DismissedId, input.Player1DismissalType, input.Fielder1Id, input.Fielder2Id)

	var isContinuousDelivery bool
	db.QueryRow(ctx,
		`SELECT max(innings_delivery_number) + 1 = $2 FROM deliveries WHERE innings_id = $1`,
		input.InningsId, input.InningsDeliveryNumber,
	).Scan(&isContinuousDelivery)

	// ensure partnership
	if isContinuousDelivery {
		query, args := getEnsurePartnershipTriggerBatch(input)
		_ = batch.Queue(query, args...)
	}

	// update batter entry
	if input.Wides.Int64 == 0 {
		query, args := getDeliveryBatterTriggerBatch(input, false)
		_ = batch.Queue(query, args...)
	}

	// update dismissed1 entry
	if input.Player1DismissedId.Valid {
		query, args := getDeliveryPlayer1DismissedTriggerBatch(input, false)
		_ = batch.Queue(query, args...)
	}

	// update bowler entry
	bowlerQuery, bowlerArgs := getDeliveryBowlerTriggerBatch(input, false)
	_ = batch.Queue(bowlerQuery, bowlerArgs...)

	// update innings entry
	inningsQuery, inningsArgs := getDeliveryInningsTriggerBatch(input, false)
	_ = batch.Queue(inningsQuery, inningsArgs...)

	// update FoW & Partnership
	if isContinuousDelivery {
		if input.Player1DismissedId.Valid {
			query, args := getDeliveryFallofWktTriggerBatch(input, false)
			_ = batch.Queue(query, args...)
		}

		query, args := getUpdatePartnershipTriggerBatch(input)
		_ = batch.Queue(query, args...)
	} else {
		_ = batch.Queue(`SELECT sync_fow_partnership($1)`, input.InningsId)
	}

	result := db.SendBatch(ctx, &batch)
	return result.Close()
}

// will also undo previous batter & bowler scorecard entries, innings scores and then reapply new input
func UpdateDeliveryWithScoringData(ctx context.Context, db DB_Exec, newInput *models.DeliveryScoringInput) error {
	existingInput := models.DeliveryScoringInput{
		InningsId:             newInput.InningsId,
		InningsDeliveryNumber: newInput.InningsDeliveryNumber,
	}

	query := `
	SELECT
		ball_number, over_number, batter_id, bowler_id, non_striker_id, batter_runs, wides,
		noballs, legbyes, byes, penalty, total_extras, total_runs, bowler_runs, is_four, is_six,
		player1_dismissed_id, player1_dismissal_type, fielder1_id, fielder2_id
	FROM deliveries
	WHERE innings_id = $1 AND innings_delivery_number = $2`

	err := db.QueryRow(ctx, query, newInput.InningsId, newInput.InningsDeliveryNumber).Scan(&existingInput.BallNumber, &existingInput.OverNumber, &existingInput.BatterId, &existingInput.BowlerId, &existingInput.NonStrikerId, &existingInput.BatterRuns, &existingInput.Wides, &existingInput.Noballs, &existingInput.Legbyes, &existingInput.Byes, &existingInput.Penalty, &existingInput.TotalExtras, &existingInput.TotalRuns, &existingInput.BowlerRuns, &existingInput.IsFour, &existingInput.IsSix, &existingInput.Player1DismissedId, &existingInput.Player1DismissalType, &existingInput.Fielder1Id, &existingInput.Fielder2Id)

	if err != nil {
		return err
	}

	batch := pgx.Batch{}

	// undo batter entry
	if existingInput.Wides.Int64 == 0 {
		query, args := getDeliveryBatterTriggerBatch(&existingInput, true)
		_ = batch.Queue(query, args...)
	}

	// undo dismissed1 entry
	if existingInput.Player1DismissedId.Valid {
		query, args := getDeliveryPlayer1DismissedTriggerBatch(&existingInput, true)
		_ = batch.Queue(query, args...)
	}

	// undo bowler entry
	bowlerQuery, bowlerArgs := getDeliveryBowlerTriggerBatch(&existingInput, true)
	_ = batch.Queue(bowlerQuery, bowlerArgs...)

	// undo innings entry
	inningsQuery, inningsArgs := getDeliveryInningsTriggerBatch(&existingInput, true)
	_ = batch.Queue(inningsQuery, inningsArgs...)

	// update delivery
	batch.Queue(`
		UPDATE deliveries SET
			ball_number = $1, over_number = $2, batter_id = $3, bowler_id = $4, non_striker_id = $5,
			batter_runs = $6, wides = $7, noballs = $8, legbyes = $9, byes = $10, penalty = $11,
			total_extras = $12, total_runs = $13, bowler_runs = $14, is_four = $15, is_six = $16,
			player1_dismissed_id = $17, player1_dismissal_type = $18, fielder1_id = $19, fielder2_id = $20
		WHERE innings_id = $21 AND innings_delivery_number = $22`,
		newInput.BallNumber, newInput.OverNumber, newInput.BatterId, newInput.BowlerId, newInput.NonStrikerId, newInput.BatterRuns, newInput.Wides, newInput.Noballs, newInput.Legbyes, newInput.Byes, newInput.Penalty, newInput.TotalExtras, newInput.TotalRuns, newInput.BowlerRuns, newInput.IsFour, newInput.IsSix, newInput.Player1DismissedId, newInput.Player1DismissalType, newInput.Fielder1Id, newInput.Fielder2Id, newInput.InningsId, newInput.InningsDeliveryNumber)

	// update batter entry
	if newInput.Wides.Int64 == 0 {
		query, args := getDeliveryBatterTriggerBatch(newInput, false)
		_ = batch.Queue(query, args...)
	}

	// update dismissed1 entry
	if newInput.Player1DismissedId.Valid {
		query, args := getDeliveryPlayer1DismissedTriggerBatch(newInput, false)
		_ = batch.Queue(query, args...)
	}

	// update bowler entry
	bowlerQuery1, bowlerArgs1 := getDeliveryBowlerTriggerBatch(newInput, false)
	_ = batch.Queue(bowlerQuery1, bowlerArgs1...)

	// update innings entry
	inningsQuery1, inningsArgs1 := getDeliveryInningsTriggerBatch(newInput, false)
	_ = batch.Queue(inningsQuery1, inningsArgs1...)

	_ = batch.Queue(`SELECT sync_fow_partnership($1)`, newInput.InningsId)

	result := db.SendBatch(ctx, &batch)

	return result.Close()
}

func UpdateDeliveryPlayer2Dimissal(ctx context.Context, db DB_Exec, newInput *models.DeliveryPlayer2DismissedInput) error {
	existingInput := models.Delivery{
		InningsId:             newInput.InningsId,
		InningsDeliveryNumber: newInput.InningsDeliveryNumber,
	}

	query := `
		SELECT ball_number, player2_dismissed_id, player2_dismissal_type
		FROM deliveries
		WHERE innings_id = $1 AND innings_delivery_number = $2`

	err := db.QueryRow(ctx, query, existingInput.InningsId, existingInput.InningsDeliveryNumber).Scan(&existingInput.BallNumber, &existingInput.Player2DismissedId, &existingInput.Player2DismissalType)
	if err != nil {
		return err
	}

	batch := pgx.Batch{}

	query = `
		UPDATE deliveries SET
			player2_dismissed_id = $1, player2_dismissal_type = $2
		WHERE innings_id = $3 AND innings_delivery_number = $4
	`

	_ = batch.Queue(query, newInput.Player2DismissedId, newInput.Player2DismissalType, newInput.InningsId, newInput.InningsDeliveryNumber)

	if existingInput.Player2DismissedId.Valid {
		query = `UPDATE batting_scorecards SET
					dismissal_type = NULL, dismissal_ball_number = NULL
				WHERE innings_id = $1 AND batter_id = $2`

		_ = batch.Queue(query, existingInput.InningsId, existingInput.Player2DismissedId)

		if models.IsTeamDismissal(existingInput.Player2DismissalType.String) {
			query = `UPDATE innings SET total_wickets = total_wickets - 1 WHERE id = $1`
			_ = batch.Queue(query, newInput.InningsId)
		}
	}

	if newInput.Player2DismissedId.Valid {
		query = `UPDATE batting_scorecards SET
					dismissal_type = $1, dismissal_ball_number = $2
				WHERE innings_id = $3 AND batter_id = $4`

		_ = batch.Queue(query, newInput.Player2DismissalType, newInput.InningsDeliveryNumber, newInput.InningsId, newInput.Player2DismissedId)

		if models.IsTeamDismissal(newInput.Player2DismissalType.String) {
			query = `UPDATE innings SET total_wickets = total_wickets + 1 WHERE id = $1`
			_ = batch.Queue(query, newInput.InningsId)
		}
	}

	_ = batch.Queue(`SELECT sync_fow_partnership($1)`, newInput.InningsId)

	return db.SendBatch(ctx, &batch).Close()
}

func UpdateDeliveryCommentary(ctx context.Context, db DB_Exec, input *models.DeliveryCommentaryInput) error {
	query := `
		UPDATE deliveries
			SET commentary = $1, updated_at = now()
		WHERE innings_id = $2 AND innings_delivery_number = $3`

	cmd, err := db.Exec(ctx, query, input.Commentary, input.InningsId, input.InningsDeliveryNumber)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return errors.New("failed to update commentary")
	}

	return nil
}

func UpdateDeliveryAdvanceInfo(ctx context.Context, db DB_Exec, input *models.DeliveryAdvanceInfoInput) error {
	query := `
		UPDATE deliveries
			SET is_pace = $1, bowling_style = $2, is_batter_rhb = $3, is_non_striker_rhb = $4,
			line = $5, length = $6, ball_type = $7, ball_speed = $8, misc = $9, ww_region = $10,
			foot_type = $11, shot_type = $12, updated_at = now()
		WHERE innings_id = $13 AND innings_delivery_number = $14`

	cmd, err := db.Exec(ctx, query, input.IsPace, input.BowlingStyle, input.IsBatterRHB, input.IsNonStrikerRHB, input.Line, input.Length, input.BallType, input.BallSpeed, input.Misc, input.WwRegion, input.FootType, input.ShotType, input.InningsId, input.InningsDeliveryNumber)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("failed to insert delivery")
	}

	return nil
}

func ReadDeliveriesByMatchInnings(ctx context.Context, db DB_Exec, match_id int64, innings_id int64) (responses.MatchCommentaryResponse, error) {
	var response responses.MatchCommentaryResponse
	matchHeader := &response.MatchHeader

	query := fmt.Sprintf(`
		WITH innings_commentary AS (
			SELECT
				deliveries.innings_id, deliveries.innings_delivery_number, deliveries.ball_number, deliveries.over_number,
		
				deliveries.batter_id, batter.name, deliveries.bowler_id, bowler.name, deliveries.fielder1_id, fielder1.name, deliveries.fielder2_Id, fielder2.name,
	
				deliveries.wides, deliveries.noballs, deliveries.legbyes, deliveries.byes, deliveries.total_runs, deliveries.is_four, deliveries.is_six,
			
				deliveries.player1_dismissed_id, player1_dismissed.name, deliveries.player1_dismissal_type, bs1.runs_scored, bs1.balls_faced, bs1.fours_scored, bs1.sixes_scored,

				deliveries.player2_dismissed_id, player2_dismissed.name, deliveries.player2_dismissal_type, bs2.runs_scored, bs2.balls_faced, bs2.fours_scored, bs2.sixes_scored,

				deliveries.commentary

			FROM deliveries

			LEFT JOIN innings ON deliveries.innings_id = innings.id
				AND innings.is_super_over = FALSE
				AND innings.id = $2

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

	rows := db.QueryRow(ctx, query, match_id, innings_id)

	err := rows.Scan(
		&matchHeader.MatchId, &matchHeader.PlayingLevel, &matchHeader.PlayingFormat, &matchHeader.MatchType, &matchHeader.EventMatchNumber,
		&matchHeader.MatchState, &matchHeader.MatchStateDescription, &matchHeader.FinalResult,

		&matchHeader.MatchWinnerId, &matchHeader.MatchLoserId, &matchHeader.IsWonByInnings, &matchHeader.IsWonByRuns,
		&matchHeader.WinMargin, &matchHeader.BallsMargin, &matchHeader.SuperOverWinnerId, &matchHeader.BowlOutWinnerId, &matchHeader.OutcomeSpecialMethod, &matchHeader.TossWinnerId, &matchHeader.TossLoserId, &matchHeader.IsTossDecisionBat,

		&matchHeader.Season, &matchHeader.StartDate, &matchHeader.EndDate, &matchHeader.StartDateTimeUtc, &matchHeader.IsDayNight, &matchHeader.GroundId, &matchHeader.GroundName, &matchHeader.MainSeriesId, &matchHeader.MainSeriesName,

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

// helpers

func getDeliveryBatterTriggerBatch(input *models.DeliveryScoringInput, undoFlag bool) (string, []any) {
	var fourInc, sixInc int

	if input.IsFour.Valid && input.IsFour.Bool {
		fourInc++
	} else if input.IsSix.Valid && input.IsSix.Bool {
		sixInc++
	}

	sign := '+'
	if undoFlag {
		sign = '-'
	}

	query := fmt.Sprintf(`
	UPDATE batting_scorecards SET
		runs_scored = runs_scored %c $1, balls_faced = balls_faced %c 1,
		fours_scored = fours_scored %c $2, sixes_scored = sixes_scored %c $3
	WHERE innings_id = $4 AND batter_id = $5
	`, sign, sign, sign, sign)

	args := []any{input.BatterRuns, fourInc, sixInc, input.InningsId, input.BatterId}

	return query, args
}

func getDeliveryPlayer1DismissedTriggerBatch(input *models.DeliveryScoringInput, undoFlag bool) (string, []any) {
	var (
		query         string
		args          []any
		dismissedById pgtype.Int8
	)

	if input.Player1DismissedId.Valid && models.IsBowlerDismissal(input.Player1DismissalType.String) {
		dismissedById = input.BowlerId
	}

	if undoFlag {
		query = `
			UPDATE batting_scorecards SET
				dismissed_by_id = NULL, dismissal_type = NULL, dismissal_ball_number = NULL,
				fielder1_id = NULL, fielder2_id = NULL
			WHERE innings_id = $1 AND batter_id = $2`

		args = []any{input.InningsId, input.Player1DismissedId}
	} else {
		query = `
			UPDATE batting_scorecards SET
				dismissed_by_id = $1, dismissal_type = $2, dismissal_ball_number = $3,
				fielder1_id = $4, fielder2_id = $5
			WHERE innings_id = $6 AND batter_id = $7`

		args = []any{
			dismissedById, input.Player1DismissalType, input.InningsDeliveryNumber, input.Fielder1Id, input.Fielder2Id, input.InningsId, input.Player1DismissedId,
		}
	}

	return query, args
}

func getDeliveryBowlerTriggerBatch(input *models.DeliveryScoringInput, undoFlag bool) (string, []any) {
	var bowlerBallInc, bowlerWicketInc, fourInc, sixInc int

	if input.IsFour.Valid && input.IsFour.Bool {
		fourInc++
	} else if input.IsSix.Valid && input.IsSix.Bool {
		sixInc++
	}

	if input.Wides.Int64 == 0 && input.Noballs.Int64 == 0 {
		bowlerBallInc++
	}

	if input.Player1DismissedId.Valid && models.IsBowlerDismissal(input.Player1DismissalType.String) {
		bowlerWicketInc++
	}

	sign := '+'
	if undoFlag {
		sign = '-'
	}

	query := fmt.Sprintf(`
		WITH match AS (
			SELECT balls_per_over
			FROM matches
			WHERE id = $1
		),
		maiden AS (
			SELECT
				(CASE 
					WHEN COUNT(DISTINCT bowler_id) = 1
					AND SUM(bowler_runs) = 0
					AND COUNT(ball_number) - SUM(wides) - SUM(noballs) = (SELECT balls_per_over FROM match)
					THEN 1 ELSE 0
				END) AS increase
			FROM deliveries
			WHERE innings_id = $2 AND over_number = $3
		)
		UPDATE bowling_scorecards
		SET wickets_taken = wickets_taken %c $4, runs_conceded = runs_conceded %c $5,
			balls_bowled = balls_bowled %c $6, fours_conceded = fours_conceded %c $7,
			sixes_conceded = sixes_conceded %c $8, wides_conceded = wides_conceded %c $9,
			noballs_conceded = noballs_conceded %c $10, maiden_overs = maiden_overs %c (SELECT increase FROM maiden)
		WHERE innings_id = $2 AND bowler_id = $11
	`, sign, sign, sign, sign, sign, sign, sign, sign)

	args := []any{input.MatchId, input.InningsId, input.OverNumber, bowlerWicketInc, input.BowlerRuns, bowlerBallInc, fourInc, sixInc, input.Wides, input.Noballs, input.BowlerId}

	return query, args
}

func getDeliveryInningsTriggerBatch(input *models.DeliveryScoringInput, undoFlag bool) (string, []any) {
	var bowlerBallInc, teamWicketInc int

	if input.Wides.Int64 == 0 && input.Noballs.Int64 == 0 {
		bowlerBallInc++
	}

	if input.Player1DismissedId.Valid && models.IsTeamDismissal(input.Player1DismissalType.String) {
		teamWicketInc++
	}

	sign := '+'
	if undoFlag {
		sign = '-'
	}

	query := fmt.Sprintf(`
		UPDATE innings SET
			total_runs = total_runs %c $1, total_balls = total_balls %c $2,
			total_wickets = total_wickets %c $3, byes = byes %c $4, leg_byes = leg_byes %c $5, 
			wides = wides %c $6, noballs = noballs %c $7, penalty = penalty %c $8,
			striker_id = $9, non_striker_id = $10, bowler1_id = $11, bowler2_id = $12
		WHERE id = $13
	`, sign, sign, sign, sign, sign, sign, sign, sign)

	args := []any{input.TotalRuns, bowlerBallInc, teamWicketInc, input.Byes, input.Legbyes, input.Wides, input.Noballs, input.Penalty, input.NewStrikerId, input.NonStrikerId, input.NewBowler1Id, input.NewBowler2Id, input.InningsId}

	return query, args
}

func getDeliveryFallofWktTriggerBatch(input *models.DeliveryScoringInput, undoFlag bool) (string, []any) {
	var (
		query string
		args  []any
	)

	if undoFlag {
		query = "DELETE FROM fall_of_wickets WHERE innings_id = $1 AND innings_delivery_number = $2 AND batter_id = $3"
		args = []any{input.InningsId, input.InningsDeliveryNumber, input.Player1DismissedId}
	} else {
		query = `
			WITH current_innings AS (
				SELECT total_runs, total_wickets
				FROM innings
				WHERE innings.id = $1
			)
			INSERT INTO fall_of_wickets (innings_id, innings_delivery_number, batter_id, ball_number, dismissal_type, team_runs, wicket_number)
			SELECT $1, $2, $3, $4, $5, current_innings.total_runs, current_innings.total_wickets
			FROM current_innings
		`

		args = []any{input.InningsId, input.InningsDeliveryNumber, input.Player1DismissedId, input.BallNumber, input.Player1DismissalType}
	}

	return query, args
}

func getEnsurePartnershipTriggerBatch(input *models.DeliveryScoringInput) (string, []any) {
	query := `
		WITH innings_data AS (
			SELECT total_runs, total_wickets FROM innings WHERE id = $1
		)
		INSERT INTO batting_partnerships (
			innings_id, wicket_number, start_innings_delivery_number, end_innings_delivery_number,
			batter1_id, batter1_runs, batter1_balls, batter2_id, batter2_runs, batter2_balls,
			start_team_runs, end_team_runs, start_ball_number, end_ball_number, is_unbeaten
		)
		SELECT $1, innings_data.total_wickets + 1, $2, $3,
			$4, 0, 0, $5, 0, 0,
			innings_data.total_runs, innings_data.total_runs, $6, $7, true
		FROM innings_data
		WHERE NOT EXISTS (
			SELECT 1 FROM batting_partnerships WHERE innings_id = $1 AND is_unbeaten = TRUE
		);
	`

	args := []any{input.InningsId, input.InningsDeliveryNumber, input.InningsDeliveryNumber, input.BatterId, input.NonStrikerId, input.BallNumber, input.BallNumber}

	return query, args
}

func getUpdatePartnershipTriggerBatch(input *models.DeliveryScoringInput) (string, []any) {
	query := `
		WITH innings_data AS (
			SELECT total_runs FROM innings WHERE id = $1
		), partnership_data AS (
			SELECT batter1_id, batter2_id FROM batting_partnerships WHERE innings_id = $1 AND is_unbeaten = TRUE
		), delivery_data AS (
			SELECT
				CASE WHEN batter1_id = batter_id THEN batter_runs ELSE 0 END AS batter1_runs,
				CASE WHEN batter1_id = batter_id AND wides = 0 THEN 1 ELSE 0 END AS batter1_balls,
				CASE WHEN batter2_id = batter_id THEN batter_runs ELSE 0 END AS batter2_runs,
				CASE WHEN batter2_id = batter_id AND wides = 0 THEN 1 ELSE 0 END AS batter2_balls
			FROM deliveries
			CROSS JOIN partnership_data
			WHERE innings_id = $1 AND innings_delivery_number = $2
		)
		UPDATE batting_partnerships SET
			end_innings_delivery_number = $2, end_team_runs = innings_data.total_runs,
			end_ball_number = $3, is_unbeaten = $4,
			batter1_runs = batting_partnerships.batter1_runs + delivery_data.batter1_runs,
			batter1_balls = batting_partnerships.batter1_balls + delivery_data.batter1_balls,
			batter2_runs = batting_partnerships.batter2_runs + delivery_data.batter2_runs,
			batter2_balls = batting_partnerships.batter2_balls + delivery_data.batter2_balls
		FROM delivery_data
		CROSS JOIN innings_data
		WHERE innings_id = $1 AND is_unbeaten = TRUE
	`

	args := []any{input.InningsId, input.InningsDeliveryNumber, input.BallNumber, !input.Player1DismissedId.Valid}

	return query, args
}
