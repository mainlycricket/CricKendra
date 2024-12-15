package responses

import "github.com/jackc/pgx/v5/pgtype"

/* Overall Stats */

type Overall_Bowling_Bowler_Group struct {
	BowlerId         pgtype.Int8   `json:"bowler_id"`
	BowlerName       pgtype.Text   `json:"bowler_name"`
	TeamsRepresented []pgtype.Text `json:"teams_represented"`
	MinDate          pgtype.Date   `json:"min_date"`
	MaxDate          pgtype.Date   `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_TeamInnings_Group struct {
	MatchId         pgtype.Int8 `json:"match_id"`
	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BowlingTeamId   pgtype.Int8 `json:"bowling_team_id"`
	BowlingTeamName pgtype.Text `json:"bowling_team_name"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`
	Season          pgtype.Text `json:"season"`
	CityName        pgtype.Text `json:"city_name"`
	StartDate       pgtype.Date `json:"start_date"`
	PlayersCount    pgtype.Int8 `json:"players_count"`
	OverallBowlingStats
}

type Overall_Bowling_Match_Group struct {
	MatchId      pgtype.Int8 `json:"match_id"`
	Team1Id      pgtype.Int8 `json:"team1_id"`
	Team1Name    pgtype.Text `json:"team1_name"`
	Team2Id      pgtype.Int8 `json:"team2_id"`
	Team2Name    pgtype.Text `json:"team2_name"`
	Season       pgtype.Text `json:"season"`
	CityName     pgtype.Text `json:"city_name"`
	StartDate    pgtype.Date `json:"start_date"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBowlingStats
}

type Overall_Bowling_Team_Group struct {
	TeamId       pgtype.Int8 `json:"bowling_team_id"`
	TeamName     pgtype.Text `json:"bowling_team_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_Opposition_Group struct {
	OppositionId   pgtype.Int8 `json:"opposition_team_id"`
	OppositionName pgtype.Text `json:"opposition_team_name"`
	PlayersCount   pgtype.Int8 `json:"players_count"`
	MinDate        pgtype.Date `json:"min_date"`
	MaxDate        pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_Ground_Group struct {
	GroundId     pgtype.Int8 `json:"ground_id"`
	GroundName   pgtype.Text `json:"ground_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_HostNation_Group struct {
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	PlayersCount   pgtype.Int8 `json:"players_count"`
	MinDate        pgtype.Date `json:"min_date"`
	MaxDate        pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_Continent_Group struct {
	ContinentId   pgtype.Int8 `json:"continent_id"`
	ContinentName pgtype.Text `json:"continent_name"`
	PlayersCount  pgtype.Int8 `json:"players_count"`
	MinDate       pgtype.Date `json:"min_date"`
	MaxDate       pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_Year_Group struct {
	Year         pgtype.Int8 `json:"year"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBowlingStats
}

type Overall_Bowling_Season_Group struct {
	Season       pgtype.Text `json:"season"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBowlingStats
}

type Overall_Bowling_Aggregate_Group struct {
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

/* Individual Stats */

type Individual_Bowling_Innings_Group struct {
	MatchId   pgtype.Int8 `json:"match_id"`
	StartDate pgtype.Date `json:"start_date"`
	GroundId  pgtype.Int8 `json:"ground_id"`
	CityName  pgtype.Text `json:"city_name"`

	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BowlerId        pgtype.Int8 `json:"bowler_id"`
	BowlerName      pgtype.Text `json:"bowler_name"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`
	BowlingTeamId   pgtype.Int8 `json:"bowling_team_id"`
	BowlingTeamName pgtype.Text `json:"bowling_team_name"`

	OversBowled   pgtype.Float8 `json:"overs_bowled"`
	RunsConceded  pgtype.Int8   `json:"runs_conceded"`
	WicketsTaken  pgtype.Int8   `json:"wickets_taken"`
	Economy       pgtype.Float8 `json:"economy"`
	FoursConceded pgtype.Int8   `json:"fours_conceded"`
	SixesConceded pgtype.Int8   `json:"sixes_conceded"`
}

type Individual_Bowling_Ground_Group struct {
	GroundId   pgtype.Int8 `json:"ground_id"`
	GroundName pgtype.Text `json:"ground_name"`
	Overall_Bowling_Bowler_Group
}

type Individual_Bowling_HostNation_Group struct {
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	Overall_Bowling_Bowler_Group
}

type Individual_Bowling_Opposition_Group struct {
	OppositionTeamId   pgtype.Int8 `json:"opposition_team_id"`
	OppositionTeamName pgtype.Text `json:"opposition_team_name"`
	Overall_Bowling_Bowler_Group
}

type Individual_Bowling_Year_Group struct {
	Year pgtype.Int8 `json:"year"`
	Overall_Bowling_Bowler_Group
}

type Individual_Bowling_Season_Group struct {
	Season pgtype.Text `json:"season"`
	Overall_Bowling_Bowler_Group
}

// Embedded in other structs
type OverallBowlingStats struct {
	MatchesPlayed   pgtype.Int8   `json:"matches_played"`
	InningsBowled   pgtype.Int8   `json:"innings_bowled"`
	OversBowled     pgtype.Float8 `json:"overs_bowled"`
	RunsConceded    pgtype.Int8   `json:"runs_conceded"`
	WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
	Average         pgtype.Float8 `json:"average"`
	StrikeRate      pgtype.Float8 `json:"strike_rate"`
	Economy         pgtype.Float8 `json:"economy"`
	FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
	FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
	BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
	BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
	FoursConceded   pgtype.Int8   `json:"fours_conceded"`
	SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
}
