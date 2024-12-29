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

/* Individual Stats */

type Individual_Team_Matches_Group struct {
	IndividualMatchInfo

	TossWinnerId      pgtype.Int8 `json:"toss_winner_id"`
	IsTossDecisionBat pgtype.Bool `json:"is_toss_decison_bat"`
	WinMargin         pgtype.Int8 `json:"win_margin"`
	BallsMargin       pgtype.Int8 `json:"balls_remaining_after_win"`
	IsWonByRuns       pgtype.Bool `json:"is_won_by_runs"`
	IsWonByInnings    pgtype.Bool `json:"is_won_by_innings"`
}

type Individual_Team_Innings_Group struct {
	IndividualMatchInfo

	InningsId     pgtype.Int8   `json:"innings_id"`
	InningsNumber pgtype.Int8   `json:"innings_number"`
	InningsEnd    pgtype.Text   `json:"innings_end"`
	TotalRuns     pgtype.Int8   `json:"total_runs"`
	TotalWickets  pgtype.Int8   `json:"total_wickets"`
	TotalOvers    pgtype.Float8 `json:"total_overs"`
	ScoringRate   pgtype.Float8 `json:"scoring_rate"`
}

type Individual_Team_Grounds_Group struct {
	TeamId       pgtype.Int8 `json:"team_id"`
	TeamName     pgtype.Text `json:"team_name"`
	GroundId     pgtype.Int8 `json:"ground_id"`
	GroundName   pgtype.Text `json:"ground_name"`
	MinStartDate pgtype.Date `json:"min_start_date"`
	MaxStartDate pgtype.Date `json:"max_start_date"`
	OverallTeamStats
}

type Individual_Team_HostNations_Group struct {
	TeamId         pgtype.Int8 `json:"team_id"`
	TeamName       pgtype.Text `json:"team_name"`
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	MinStartDate   pgtype.Date `json:"min_start_date"`
	MaxStartDate   pgtype.Date `json:"max_start_date"`
	OverallTeamStats
}

type Individual_Team_Years_Group struct {
	TeamId   pgtype.Int8 `json:"team_id"`
	TeamName pgtype.Text `json:"team_name"`
	Year     pgtype.Int8 `json:"year"`
	OverallTeamStats
}

type Individual_Team_Seasons_Group struct {
	TeamId   pgtype.Int8 `json:"team_id"`
	TeamName pgtype.Text `json:"team_name"`
	Season   pgtype.Text `json:"season"`
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

type IndividualMatchInfo struct {
	MatchId        pgtype.Int8 `json:"match_id"`
	TeamId         pgtype.Int8 `json:"team_id"`
	TeamName       pgtype.Text `json:"team_name"`
	OppositionId   pgtype.Int8 `json:"opposition_id"`
	OppositionName pgtype.Text `json:"opposition_name"`
	GroundId       pgtype.Int8 `json:"ground_id"`
	CityName       pgtype.Text `json:"city_name"`
	StartDate      pgtype.Date `json:"start_date"`
	FinalResult    pgtype.Text `json:"final_result"`
	MatchWinnerId  pgtype.Int8 `json:"match_winner_id"`
}
