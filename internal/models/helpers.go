package models

import "github.com/jackc/pgx/v5/pgtype"

func (innings *Innings) SetDefaultScore() {
	innings.TotalRuns = pgtype.Int8{Int64: 0, Valid: true}
	innings.TotalWkts = pgtype.Int8{Int64: 0, Valid: true}
	innings.Byes = pgtype.Int8{Int64: 0, Valid: true}
	innings.Legbyes = pgtype.Int8{Int64: 0, Valid: true}
	innings.Wides = pgtype.Int8{Int64: 0, Valid: true}
	innings.Noballs = pgtype.Int8{Int64: 0, Valid: true}
	innings.Penalty = pgtype.Int8{Int64: 0, Valid: true}
}
