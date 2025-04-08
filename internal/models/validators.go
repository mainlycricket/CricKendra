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

func (in *DeliveryPlayer2DismissedInput) Validate(ctx context.Context, db *pgxpool.Pool) error {
	if in.Player2DismissedId.Valid {
		if !IsValidDismissal2(in.Player2DismissalType.String) {
			return errors.New("invalid dismissal 2 type")
		}

		var player1DismissedId pgtype.Int8
		query := `SELECT player1_dismissed_id FROM deliveries WHERE innings_id = $1 AND innings_delivery_number = $2`

		err := db.QueryRow(ctx, query, in.InningsId, in.InningsDeliveryNumber).Scan(&player1DismissedId)
		if err != nil {
			return err
		}

		if !player1DismissedId.Valid {
			return errors.New("no player1 wicket found for this delivery")
		}
	}

	return nil
}
