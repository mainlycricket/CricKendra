package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertInnings(ctx context.Context, db DB_Exec, innings *models.Innings) error {
	query := `INSERT INTO innings (match_id, innings_number, batting_team_id, bowling_team_id, total_runs, total_wickets, byes, leg_byes, wides, noballs, penalty, is_super_over, innings_end) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	cmd, err := db.Exec(ctx, query, innings.MatchId, innings.InningsNumber, innings.BattingTeamId, innings.BowlingTeamId, innings.TotalRuns, innings.TotalWkts, innings.Byes, innings.Legbyes, innings.Wides, innings.Noballs, innings.Penalty, innings.IsSuperOver, innings.InningsEnd)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert innings")
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
	query := `UPDATE innings SET match_id = $1, innings_number = $2, batting_team_id = $3, bowling_team_id = $4, total_runs = $5, total_wickets = $6, byes = $7, leg_byes = $8, wides = $9, noballs = $10, penalty = $11, is_super_over = $12, innings_end = $13 WHERE id = $14`

	cmd, err := db.Exec(ctx, query, innings.MatchId, innings.InningsNumber, innings.BattingTeamId, innings.BowlingTeamId, innings.TotalRuns, innings.TotalWkts, innings.Byes, innings.Legbyes, innings.Wides, innings.Noballs, innings.Penalty, innings.IsSuperOver, innings.InningsEnd, innings.Id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to update innings")
	}

	return nil
}

func ReadInningsEndOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::innings_end))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	inningsEndOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var inningsEnd string
		err := row.Scan(&inningsEnd)
		return inningsEnd, err
	})

	return inningsEndOptions, err
}
