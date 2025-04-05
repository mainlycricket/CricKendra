package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (in *MatchStateInput) Validate(ctx context.Context, db *pgxpool.Pool) error {
	var currentState, finalResult pgtype.Text

	query := `SELECT match_state, final_result FROM matches WHERE id = $1`

	err := db.QueryRow(ctx, query, in.MatchId).Scan(&currentState, &finalResult)
	if err != nil {
		return fmt.Errorf(`error while reading current match state: %v`, err)
	}

	if currentState.String == "completed" {
		return errors.New(`match state is already completed`)
	}

	if in.State.String == "completed" && !finalResult.Valid {
		return errors.New("match result should be declared to set the match state to complete")
	}

	return nil
}
