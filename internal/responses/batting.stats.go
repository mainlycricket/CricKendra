package responses

import "github.com/jackc/pgx/v5/pgtype"

type Overall_Batting_Batter_Group struct {
	BatterId         pgtype.Int8   `json:"batter_id"`
	BatterName       pgtype.Text   `json:"batter_name"`
	TeamsRepresented []pgtype.Text `json:"teams_represented"`
	MinDate          pgtype.Date   `json:"min_date"`
	MaxDate          pgtype.Date   `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Team_Group struct {
	TeamId       pgtype.Int8 `json:"batting_team_id"`
	TeamName     pgtype.Text `json:"batting_team_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Opposition_Group struct {
	TeamId       pgtype.Int8 `json:"bowling_team_id"`
	TeamName     pgtype.Text `json:"bowling_team_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Season_Group struct {
	Season       pgtype.Text `json:"season"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type Overall_Batting_Year_Group struct {
	Year         pgtype.Int8 `json:"year"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type OverallBattingStats struct {
	MatchesPlayed      pgtype.Int8   `json:"matches_played"`
	InningsBatted      pgtype.Int8   `json:"innings_batted"`
	RunsScored         pgtype.Int8   `json:"runs_scored"`
	BallsFaced         pgtype.Int8   `json:"balls_faced"`
	NotOuts            pgtype.Int8   `json:"not_outs"`
	Average            pgtype.Float8 `json:"average"`
	StrikeRate         pgtype.Float8 `json:"strike_rate"`
	HighestScore       pgtype.Int8   `json:"highest_score"`
	HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
	Centuries          pgtype.Int8   `json:"centuries"`
	HalfCenturies      pgtype.Int8   `json:"half_centuries"`
	FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
	Ducks              pgtype.Int8   `json:"ducks"`
	FoursScored        pgtype.Int8   `json:"fours_scored"`
	SixesScored        pgtype.Int8   `json:"sixes_scored"`
}
