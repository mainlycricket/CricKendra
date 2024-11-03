package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertInnings(ctx context.Context, db *pgxpool.Pool, innings *models.Innings) error {
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
