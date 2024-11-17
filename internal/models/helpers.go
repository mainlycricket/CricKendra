package models

import (
	"slices"

	"github.com/jackc/pgx/v5/pgtype"
)

func (innings *Innings) SetDefaultScore() {
	innings.TotalRuns = pgtype.Int8{Int64: 0, Valid: true}
	innings.TotalWkts = pgtype.Int8{Int64: 0, Valid: true}
	innings.Byes = pgtype.Int8{Int64: 0, Valid: true}
	innings.Legbyes = pgtype.Int8{Int64: 0, Valid: true}
	innings.Wides = pgtype.Int8{Int64: 0, Valid: true}
	innings.Noballs = pgtype.Int8{Int64: 0, Valid: true}
	innings.Penalty = pgtype.Int8{Int64: 0, Valid: true}
	innings.IsSuperOver = pgtype.Bool{Bool: false, Valid: true}
}

func IsBowlerDismissal(dismissalType string) bool {
	bowlerWickets := []string{"caught", "bowled", "lbw", "stumpted", "hit wicket", "caught and bowled"}
	return slices.Contains(bowlerWickets, dismissalType)
}
