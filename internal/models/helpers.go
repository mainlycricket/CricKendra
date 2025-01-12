package models

import (
	"slices"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewInnings(matchId, battingTeamId, bowlingTeamId int64) *Innings {
	innings := &Innings{}

	innings.MatchId = pgtype.Int8{Int64: matchId, Valid: true}
	innings.BattingTeamId = pgtype.Int8{Int64: battingTeamId, Valid: true}
	innings.BowlingTeamId = pgtype.Int8{Int64: bowlingTeamId, Valid: true}
	innings.TotalRuns = pgtype.Int8{Int64: 0, Valid: true}
	innings.TotalBalls = pgtype.Int8{Int64: 0, Valid: true}
	innings.TotalWkts = pgtype.Int8{Int64: 0, Valid: true}
	innings.Byes = pgtype.Int8{Int64: 0, Valid: true}
	innings.Legbyes = pgtype.Int8{Int64: 0, Valid: true}
	innings.Wides = pgtype.Int8{Int64: 0, Valid: true}
	innings.Noballs = pgtype.Int8{Int64: 0, Valid: true}
	innings.Penalty = pgtype.Int8{Int64: 0, Valid: true}
	innings.IsSuperOver = pgtype.Bool{Bool: false, Valid: true}

	return innings
}

func IsBowlerDismissal(dismissalType string) bool {
	bowlerWickets := []string{"caught", "bowled", "lbw", "stumpted", "hit wicket", "caught and bowled"}
	return slices.Contains(bowlerWickets, dismissalType)
}

func IsTeamDismissal(dismissalType string) bool {
	teamWickets := []string{"caught", "bowled", "lbw", "stumpted", "hit wicket", "handled the ball", "obstructing the field", "timed out", "hit the ball twice", "caught and bowled", "retired out"}
	return slices.Contains(teamWickets, dismissalType)
}
