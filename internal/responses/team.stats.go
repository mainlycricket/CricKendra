package responses

import "github.com/jackc/pgx/v5/pgtype"

/* Overall Stats */

type Overall_Team_Teams_Group struct {
	TeamId       pgtype.Int8 `json:"team_id"`
	TeamName     pgtype.Text `json:"team_name"`
	MinStartDate pgtype.Date `json:"min_start_date"`
	MaxStartDate pgtype.Date `json:"max_start_date"`
	OverallTeamStats
}

type Overall_Team_Players_Group struct {
	PlayerId     pgtype.Int8 `json:"player_id"`
	PlayerName   pgtype.Text `json:"player_name"`
	MinStartDate pgtype.Date `json:"min_start_date"`
	MaxStartDate pgtype.Date `json:"max_start_date"`
	TeamsCount   pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Matches_Group struct {
	MatchId    pgtype.Int8 `json:"match_id"`
	Team1Id    pgtype.Int8 `json:"team1_id"`
	Team1Name  pgtype.Text `json:"team1_name"`
	Team2Id    pgtype.Int8 `json:"team2_id"`
	Team2Name  pgtype.Text `json:"team2_name"`
	City       pgtype.Text `json:"city_name"`
	Season     pgtype.Text `json:"season"`
	StartDate  pgtype.Date `json:"start_date"`
	TeamsCount pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Grounds_Group struct {
	GroundId     pgtype.Int8 `json:"ground_id"`
	GroundName   pgtype.Text `json:"ground_name"`
	MinStartDate pgtype.Date `json:"min_start_date"`
	MaxStartDate pgtype.Date `json:"max_start_date"`
	TeamsCount   pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_HostNations_Group struct {
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	MinStartDate   pgtype.Date `json:"min_start_date"`
	MaxStartDate   pgtype.Date `json:"max_start_date"`
	TeamsCount     pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Continents_Group struct {
	ContinentId   pgtype.Int8 `json:"continent_id"`
	ContinentName pgtype.Text `json:"continent_name"`
	MinStartDate  pgtype.Date `json:"min_start_date"`
	MaxStartDate  pgtype.Date `json:"max_start_date"`
	TeamsCount    pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Years_Group struct {
	Year       pgtype.Int8 `json:"year"`
	TeamsCount pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Seasons_Group struct {
	Season     pgtype.Text `json:"season"`
	TeamsCount pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Decades_Group struct {
	Decade     pgtype.Int8 `json:"decade"`
	TeamsCount pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

type Overall_Team_Aggregate_Group struct {
	TeamsCount pgtype.Int8 `json:"teams_count"`
	OverallTeamStats
}

// Embedded in other structs
type OverallTeamStats struct {
	MatchesPlayed   pgtype.Int8   `json:"matches_played"`
	MatchesWon      pgtype.Int8   `json:"matches_won"`
	MatchesLost     pgtype.Int8   `json:"matches_lost"`
	WinLossRatio    pgtype.Float8 `json:"win_loss_ratio"`
	MatchesDrawn    pgtype.Int8   `json:"matches_drawn"`
	MatchesTied     pgtype.Int8   `json:"matches_tied"`
	MatchesNoResult pgtype.Int8   `json:"matches_no_result"`

	InningsCount pgtype.Int8   `json:"innings_count"`
	TotalRuns    pgtype.Int8   `json:"total_runs"`
	TotalBalls   pgtype.Int8   `json:"total_balls"`
	TotalWickets pgtype.Int8   `json:"total_wickets"`
	Average      pgtype.Float8 `json:"average"`
	ScoringRate  pgtype.Float8 `json:"scoring_rate"`
	HighestScore pgtype.Int8   `json:"highest_score"`
	LowestScore  pgtype.Int8   `json:"lowest_score"`
}
