package responses

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/models"
)

// Players

type AllPlayersResponse struct {
	Players []AllPlayers `json:"players"`
	Next    bool         `json:"next"`
}

type AllPlayers struct {
	Id                  pgtype.Int8 `json:"id"`
	Name                pgtype.Text `json:"name"`
	PlayingRole         pgtype.Text `json:"playing_role"`
	Nationality         pgtype.Text `json:"nationality"`
	IsMale              pgtype.Bool `json:"is_male"`
	DateOfBirth         pgtype.Date `json:"date_of_birth"`
	IsRHB               pgtype.Bool `json:"is_rhb"`
	PrimaryBowlingStyle pgtype.Text `json:"primary_bowling_style"`
}

type SinglePlayer struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	FullName    pgtype.Text `json:"full_name"`
	PlayingRole pgtype.Text `json:"playing_role"`
	Nationality pgtype.Text `json:"nationality"`
	IsMale      pgtype.Bool `json:"is_male"`
	DateOfBirth pgtype.Date `json:"date_of_birth"`
	ImageURL    pgtype.Text `json:"image_url"`
	Biography   pgtype.Text `json:"biography"`

	IsRHB               pgtype.Bool          `json:"is_rhb"`
	BowlingStyles       []pgtype.Text        `json:"bowling_styles"`
	PrimaryBowlingStyle pgtype.Text          `json:"primary_bowling_style"`
	TeamsRepresented    []TeamAsForeignField `json:"teams_represented"`

	TestStats  *models.CareerStats `json:"test_stats"`
	OdiStats   *models.CareerStats `json:"odi_stats"`
	T20iStats  *models.CareerStats `json:"t20i_stats"`
	FcStats    *models.CareerStats `json:"fc_stats"`
	ListAStats *models.CareerStats `json:"lista_stats"`
	T20Stats   *models.CareerStats `json:"t20_stats"`

	CricsheetId pgtype.Text `json:"cricsheet_id"`
	CricinfoId  pgtype.Text `json:"cricinfo_id"`
	CricbuzzId  pgtype.Text `json:"cricbuzz_id"`
}

// Teams

type AllTeamsResponse struct {
	Teams []AllTeams `json:"teams"`
	Next  bool       `json:"next"`
}
type AllTeams struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	IsMale       pgtype.Bool `json:"is_male"`
	ImageURL     pgtype.Text `json:"image_url"`
	PlayingLevel pgtype.Text `json:"playing_level"`
	ShortName    pgtype.Text `json:"short_name"`
}

// Seasons

type AllSeasonsResponse struct {
	Seasons []string `json:"seasons"`
	Next    bool     `json:"next"`
}

// Continents

type AllContinentsResponse struct {
	Continents []AllContinents `json:"continents"`
	Next       bool            `json:"next"`
}

type AllContinents struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

// HostNations

type AllHostNationsResponse struct {
	HostNations []AllHostNations `json:"host_nations"`
	Next        bool             `json:"next"`
}

type AllHostNations struct {
	Id            pgtype.Int8 `json:"id"`
	Name          pgtype.Text `json:"name"`
	ContinetId    pgtype.Int8 `json:"continent_id"`
	ContinentName pgtype.Text `json:"continent_name"`
}

// Cities

type AllCitiesResponse struct {
	Cities []AllCities `json:"cities"`
	Next   bool        `json:"next"`
}

type AllCities struct {
	Id             pgtype.Int8 `json:"id"`
	Name           pgtype.Text `json:"name"`
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	ContinetId     pgtype.Int8 `json:"continent_id"`
	ContinentName  pgtype.Text `json:"continent_name"`
}

// Grounds

type AllGroundsResponse struct {
	Grounds []AllGrounds `json:"grounds"`
	Next    bool         `json:"next"`
}

type AllGrounds struct {
	Id             pgtype.Int8 `json:"id"`
	Name           pgtype.Text `json:"name"`
	CityId         pgtype.Int8 `json:"city_id"`
	CityName       pgtype.Text `json:"city_name"`
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	ContinetId     pgtype.Int8 `json:"continent_id"`
	ContinentName  pgtype.Text `json:"continent_name"`
}

// Tournaments

type AllTournamentsResponse struct {
	Tournaments []AllTournaments `json:"tournaments"`
	Next        bool             `json:"next"`
}
type AllTournaments struct {
	Id            pgtype.Int8 `json:"id"`
	Name          pgtype.Text `json:"name"`
	IsMale        pgtype.Bool `json:"is_male"`
	PlayingLevel  pgtype.Text `json:"playing_level"`
	PlayingFormat pgtype.Text `json:"playing_format"`
}

// Cricsheet People

type CricsheetPeople struct {
	Identifier pgtype.Text `json:"cricsheet_id"`
	Name       pgtype.Text `json:"name"`
	UniqueName pgtype.Text `json:"unique_name"`
	CricinfoId pgtype.Text `json:"cricinfo_id"`
	CricbuzzId pgtype.Text `json:"cricbuzz_id"`
}

// Series

type AllSeriesResponse struct {
	Series []AllSeries `json:"series"`
	Next   bool        `json:"next"`
}

type AllSeries struct {
	Id            pgtype.Int8          `json:"id"`
	Name          pgtype.Text          `json:"name"`
	IsMale        pgtype.Bool          `json:"is_male"`
	PlayingLevel  pgtype.Text          `json:"playing_level"`
	PlayingFormat pgtype.Text          `json:"playing_format"`
	Season        pgtype.Text          `json:"season"`
	Teams         []TeamAsForeignField `json:"teams"`
	StartDate     pgtype.Date          `json:"start_date"`
	EndDate       pgtype.Date          `json:"end_date"`
	WinnerTeamId  pgtype.Int8          `json:"winner_team_id"`
	FinalStatus   pgtype.Text          `json:"final_status"`
	TourFlag      pgtype.Text          `json:"tour_flag"`
}

type SingleSeries struct {
	AllSeries
	TournamentId   pgtype.Int8  `json:"tournament_id"`
	TournamentName pgtype.Text  `json:"tournament_name"`
	Matches        []AllMatches `json:"matches"`
}

// Matches

type AllMatchesResponse struct {
	Matches []AllMatches `json:"matches"`
	Next    bool         `json:"next"`
}

type AllMatches struct {
	MatchInfo
	Innings []TeamInningsShortInfo `json:"innings"`
}

type MatchInfo struct {
	Id               pgtype.Int8        `json:"id"`
	PlayingLevel     pgtype.Text        `json:"playing_level"`
	PlayingFormat    pgtype.Text        `json:"playing_format"`
	MatchType        pgtype.Text        `json:"match_type"`
	CricsheetId      pgtype.Text        `json:"cricsheet_id"`
	BallsPerOver     pgtype.Int8        `json:"balls_per_over"`
	EventMatchNumber pgtype.Int8        `json:"event_match_number"`
	StartDate        pgtype.Date        `json:"start_date"`
	EndDate          pgtype.Date        `json:"end_date"`
	StartTime        pgtype.Timestamptz `json:"start_time"`
	IsDayNight       pgtype.Bool        `json:"is_day_night"`
	GroundId         pgtype.Int8        `json:"ground_id"`
	GroundName       pgtype.Text        `json:"ground_name"`

	Team1Id       pgtype.Int8 `json:"team1_id"`
	Team1Name     pgtype.Text `json:"team1_name"`
	Team1ImageUrl pgtype.Text `json:"team1_image_url"`
	Team2Id       pgtype.Int8 `json:"team2_id"`
	Team2Name     pgtype.Text `json:"team2_name"`
	Team2ImageUrl pgtype.Text `json:"team2_image_url"`

	Season         pgtype.Text `json:"season"`
	MainSeriesId   pgtype.Int8 `json:"main_series_id"`
	MainSeriesName pgtype.Text `json:"main_series_name"`

	CurrentStatus        pgtype.Text `json:"current_status"`
	FinalResult          pgtype.Text `json:"final_result"`
	OutcomeSpecialMethod pgtype.Text `json:"outcome_special_method"`
	MatchWinnerId        pgtype.Int8 `json:"match_winner_team_id"`
	BowlOutWinnerId      pgtype.Int8 `json:"bowl_out_winner_id"`
	SuperOverWinnerId    pgtype.Int8 `json:"super_over_winner_id"`
	IsWonByInnings       pgtype.Bool `json:"is_won_by_innings"`
	IsWonByRuns          pgtype.Bool `json:"is_won_by_runs"`
	WinMargin            pgtype.Int8 `json:"win_margin"`
	BallsMargin          pgtype.Int8 `json:"balls_remaining_after_win"`
	TossWinnerId         pgtype.Int8 `json:"toss_winner_team_id"`
	IsTossDecisionBat    pgtype.Bool `json:"is_toss_decision_bat"`
}

type SingleMatchResponse struct {
	MatchInfo

	SquadEntries []MatchSquad             `json:"squad_entries"`
	MatchAwards  []MatchAwards            `json:"match_awards"`
	Innings      []InningsScorecardRecord `json:"innings"`
}

type MatchAwards struct {
	PlayerId  pgtype.Int8 `json:"player_id"`
	AwardType pgtype.Text `json:"award_type"`
}

// Squads

type MatchSquad struct {
	TeamId        pgtype.Int8 `json:"team_id"`
	PlayerId      pgtype.Int8 `json:"player_id"`
	PlayerName    pgtype.Text `json:"player_name"`
	IsCaptain     pgtype.Bool `json:"is_captain"`
	IsWk          pgtype.Bool `json:"is_wk"`
	IsDebut       pgtype.Bool `json:"is_debut"`
	IsViceCaptain pgtype.Bool `json:"is_vice_captain"`
	PlayingStatus pgtype.Text `json:"playing_status"`
}

// Scorecards

type TeamInningsShortInfo struct {
	InningsNumber pgtype.Int8 `json:"innings_number"`
	BattingTeamId pgtype.Int8 `json:"batting_team_id"`
	TotalRuns     pgtype.Int8 `json:"total_runs"`
	TotalBalls    pgtype.Int8 `json:"total_balls"`
	TotalWickets  pgtype.Int8 `json:"total_wickets"`
	InningsEnd    pgtype.Text `json:"innings_end"`
	TargetRuns    pgtype.Int8 `json:"target_runs"`
	MaxOvers      pgtype.Int8 `json:"max_overs"`
}

type InningsScorecardRecord struct {
	InningsId               pgtype.Int8              `json:"innings_id"`
	InningsNumber           pgtype.Int8              `json:"innings_number"`
	BattingTeamId           pgtype.Int8              `json:"batting_team_id"`
	BowlingTeamId           pgtype.Int8              `json:"bowling_team_id"`
	TotalRuns               pgtype.Int8              `json:"total_runs"`
	TotalBalls              pgtype.Int8              `json:"total_balls"`
	TotalWickets            pgtype.Int8              `json:"total_wickets"`
	Byes                    pgtype.Int8              `json:"byes"`
	LegByes                 pgtype.Int8              `json:"leg_byes"`
	Wides                   pgtype.Int8              `json:"wides"`
	Noballs                 pgtype.Int8              `json:"noballs"`
	Penalty                 pgtype.Int8              `json:"penalty"`
	IsSuperOver             pgtype.Bool              `json:"is_super_over"`
	InningsEnd              pgtype.Text              `json:"innings_end"`
	TargetRuns              pgtype.Int8              `json:"target_runs"`
	MaxOvers                pgtype.Int8              `json:"max_overs"`
	BattingScorecardEntries []BattingScorecardRecord `json:"batting_scorecard_entries"`
	BowlingScorecardEntries []BowlingScorecardRecord `json:"bowling_scorecard_entries"`
}

type BattingScorecardRecord struct {
	BatterId        pgtype.Int8 `json:"batter_id"`
	BattingPosition pgtype.Int8 `json:"batting_position"`
	HasBatted       pgtype.Bool `json:"has_batted"`
	RunsScored      pgtype.Int8 `json:"runs_scored"`
	BallsFaced      pgtype.Int8 `json:"balls_faced"`
	MinutesBatted   pgtype.Int8 `json:"minutes_batted"`
	FoursScored     pgtype.Int8 `json:"fours_scored"`
	SixesScored     pgtype.Int8 `json:"sixes_scored"`
	DismissalType   pgtype.Text `json:"dismissal_type"`
	DismissedById   pgtype.Int8 `json:"dismissed_by_id"`
	Fielder1Id      pgtype.Int8 `json:"fielder1_id"`
	Fielder2Id      pgtype.Int8 `json:"fielder2_id"`
}

type BowlingScorecardRecord struct {
	BowlerId        pgtype.Int8 `json:"bowler_id"`
	BowlingPosition pgtype.Int8 `json:"bowling_position"`
	WicketsTaken    pgtype.Int8 `json:"wickets_taken"`
	RunsConceded    pgtype.Int8 `json:"runs_conceded"`
	BallsBowled     pgtype.Int8 `json:"balls_bowled"`
	MaidenOvers     pgtype.Int8 `json:"maiden_overs"`
	FoursConceded   pgtype.Int8 `json:"fours_conceded"`
	SixesConceded   pgtype.Int8 `json:"sixes_conceded"`
	WidesConceded   pgtype.Int8 `json:"wides_conceded"`
	NoballsConceded pgtype.Int8 `json:"noballs_conceded"`
}

// Series Squads

type AllSeriesSquadResponse struct {
	Squads []AllSeriesSquads `json:"squads"`
	Next   bool              `json:"next"`
}

type AllSeriesSquads struct {
	Id         pgtype.Int8 `json:"id"`
	SeriesId   pgtype.Int8 `json:"series_id"`
	TeamId     pgtype.Int8 `json:"team_id"`
	SquadLabel pgtype.Text `json:"squad_label"`
}

// Innings

type AllInningsResponse struct {
	Innings []AllInnings `json:"innings"`
	Next    bool         `json:"next"`
}

type AllInnings struct {
	Id            pgtype.Int8 `json:"id"`
	MatchId       pgtype.Int8 `json:"match_id"`
	InningsNumber pgtype.Int8 `json:"innings_number"`
}

// Stats
type StatsResponse[T any] struct {
	Stats []T  `json:"stats"`
	Next  bool `json:"next"`
}
