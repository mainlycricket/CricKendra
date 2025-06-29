package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	"github.com/mainlycricket/CricKendra/backend/pkg/pgxutils"
)

func InsertInnings(ctx context.Context, db DB_Exec, innings *models.Innings) (int64, error) {
	var id int64

	query := `INSERT INTO innings (match_id, innings_number, batting_team_id, bowling_team_id, total_runs, total_balls, total_wickets, byes, leg_byes, wides, noballs, penalty, is_super_over, innings_end, target_runs, max_overs) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`

	err := db.QueryRow(ctx, query, innings.MatchId, innings.InningsNumber, innings.BattingTeamId, innings.BowlingTeamId, innings.TotalRuns.Int64, innings.TotalBalls.Int64, innings.TotalWkts.Int64, innings.Byes.Int64, innings.Legbyes.Int64, innings.Wides.Int64, innings.Noballs.Int64, innings.Penalty.Int64, innings.IsSuperOver, innings.InningsEnd, &innings.TargetRuns, &innings.MaxOvers).Scan(&id)

	return id, err
}

// all-out, declare, target reached etc
// Will also mark strikers, non-strikers, current bowlers as NULL
func UpdateInningsEnd(ctx context.Context, db DB_Exec, input *models.InningsEndInput) error {
	query := `
		UPDATE innings
			SET innings_end = $1,
				striker_id = NULL, non_striker_id = NULL, bowler1_id = NULL, bowler2_id = NULL
		WHERE id = $2
	`

	cmd, err := db.Exec(ctx, query, input.InningsEnd, input.InningsId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return errors.New("failed to update innings end")
	}

	return nil
}

func UpdateInningsCurrentBatters(ctx context.Context, db DB_Exec, input *models.InningsCurrentBattersInput) error {
	query := `
		UPDATE innings SET
			striker_id = $1, non_striker_id = $2
		WHERE id = $3
	`

	cmd, err := db.Exec(ctx, query, input.StrikerId, input.NonStrikerId, input.InningsId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return errors.New("failed to update innings current batters")
	}

	return nil
}

func UpdateInningsCurrentBowlers(ctx context.Context, db DB_Exec, input *models.InningsCurrentBowlersInput) error {
	query := `
		UPDATE innings SET
			bowler1_id = $1, bowler2_id = $2
		WHERE id = $3
	`

	cmd, err := db.Exec(ctx, query, input.Bowler1Id, input.Bowler2Id, input.InningsId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return errors.New("failed to update innings current bowlers")
	}

	return nil
}

func ReadInnings(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllInningsResponse, error) {
	var response responses.AllInningsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "innings",
		DefaultLimit: 20,
		DefaultSort:  []string{"-id"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Innings](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT id, match_id, innings_number FROM innings %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	inningsList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllInnings, error) {
		var innings responses.AllInnings

		err := rows.Scan(&innings.Id, &innings.MatchId, &innings.InningsNumber)

		return innings, err
	})

	if len(inningsList) > queryInfoOutput.RecordsCount {
		response.Innings = inningsList[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Innings = inningsList
		response.Next = false
	}

	return response, err
}

func UpdateInnings(ctx context.Context, db DB_Exec, innings *models.Innings) error {
	query := `UPDATE innings SET match_id = $1, innings_number = $2, batting_team_id = $3, bowling_team_id = $4, total_runs = $5, total_balls = $6, total_wickets = $7, byes = $8, leg_byes = $9, wides = $10, noballs = $11, penalty = $12, is_super_over = $13, innings_end = $14, target_runs = $15, max_overs = $16 WHERE id = $17`

	cmd, err := db.Exec(ctx, query, innings.MatchId, innings.InningsNumber, innings.BattingTeamId, innings.BowlingTeamId, innings.TotalRuns.Int64, innings.TotalBalls.Int64, innings.TotalWkts.Int64, innings.Byes.Int64, innings.Legbyes.Int64, innings.Wides.Int64, innings.Noballs.Int64, innings.Penalty.Int64, innings.IsSuperOver, innings.InningsEnd, innings.TargetRuns, innings.MaxOvers, innings.Id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to update innings")
	}

	return nil
}
