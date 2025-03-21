package models

import "github.com/jackc/pgx/v5/pgtype"

type TossDecisionInput struct {
	TossWinnerId      pgtype.Int8 `json:"toss_winner_team_id"`
	TossLoserId       pgtype.Int8 `json:"toss_loser_team_id"`
	IsTossDecisionBat pgtype.Bool `json:"is_toss_decision_bat"`
}
