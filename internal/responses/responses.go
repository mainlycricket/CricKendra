package responses

import "github.com/jackc/pgx/v5/pgtype"

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

	TestStats  CareerStats `json:"test_stats"`
	OdiStats   CareerStats `json:"odi_stats"`
	T20iStats  CareerStats `json:"t20i_stats"`
	FcStats    CareerStats `json:"fc_stats"`
	ListAStats CareerStats `json:"lista_stats"`
	T20Stats   CareerStats `json:"t20_stats"`

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
}

type SingleSeries struct {
	AllSeries
	TournamentId       pgtype.Int8            `json:"tournament_id"`
	TournamentName     pgtype.Text            `json:"tournament_name"`
	ParentSeriesId     pgtype.Int8            `json:"parent_series_id"`
	ParentSeriesName   pgtype.Text            `json:"parent_series_name"`
	PlayersOfTheSeries []PlayerAsForeignField `json:"players_of_the_series"`
	Matches            []MatchShortInfo       `json:"matches"`
}

// Tours

type AllToursResponse struct {
	Tours []AllTours `json:"tours"`
	Next  bool       `json:"next"`
}

type AllTours struct {
	Id              pgtype.Int8   `json:"id"`
	TouringTeamId   pgtype.Int8   `json:"touring_team_id"`
	TouringTeamName pgtype.Text   `json:"touring_team_name"`
	HostNationsId   []pgtype.Int8 `json:"host_nations_id"`
	HostNationsName []pgtype.Text `json:"host_nations_name"`
	Season          pgtype.Text   `json:"season"`
	StartDate       pgtype.Date   `json:"start_date"`
	EndDate         pgtype.Date   `json:"end_date"`
}

// Matches

type AllMatchesResponse struct {
	Matches []AllMatches `json:"matches"`
	Next    bool         `json:"next"`
}

type AllMatches struct {
	MatchShortInfo
	SeriesId   pgtype.Int8 `json:"series_id"`
	SeriesName pgtype.Text `json:"series_name"`
}

type MatchShortInfo struct {
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

	Innings []TeamInningsShortInfo `json:"innings"`
}

type SingleMatchResponse struct {
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

	Team1Id       pgtype.Int8  `json:"team1_id"`
	Team1Name     pgtype.Text  `json:"team1_name"`
	Team1ImageUrl pgtype.Text  `json:"team1_image_url"`
	Team2Id       pgtype.Int8  `json:"team2_id"`
	Team2Name     pgtype.Text  `json:"team2_name"`
	Team2ImageUrl pgtype.Text  `json:"team2_image_url"`
	SquadEntries  []MatchSquad `json:"squad_entries"`

	Season     pgtype.Text `json:"season"`
	SeriesId   pgtype.Int8 `json:"series_id"`
	SeriesName pgtype.Text `json:"series_name"`

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

	MatchAwards []MatchAwards            `json:"match_awards"`
	Innings     []InningsScorecardRecord `json:"innings"`
}

type MatchAwards struct {
	PlayerId   pgtype.Int8 `json:"player_id"`
	PlayerName pgtype.Text `json:"player_name"`
	AwardType  pgtype.Text `json:"award_type"`
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
}

// Scorecards

type TeamInningsShortInfo struct {
	InnigsNumber  pgtype.Int8 `json:"innings_number"`
	BattingTeamId pgtype.Int8 `json:"batting_team_id"`
	TotalRuns     pgtype.Int8 `json:"total_runs"`
	TotalBalls    pgtype.Int8 `json:"total_balls"`
	TotalWickets  pgtype.Int8 `json:"total_wickets"`
	InningsEnd    pgtype.Text `json:"innings_end"`
	TargetRuns    pgtype.Int8 `json:"target_runs"`
	TargetBalls   pgtype.Int8 `json:"target_balls"`
}

type InningsScorecardRecord struct {
	InnigsNumber            pgtype.Int8              `json:"innings_number"`
	BattingTeamId           pgtype.Int8              `json:"batting_team_id"`
	BattingTeamName         pgtype.Text              `json:"batting_team_name"`
	BowlingTeamId           pgtype.Int8              `json:"bowling_team_id"`
	BowlingTeamName         pgtype.Text              `json:"bowling_team_name"`
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
	TargetBalls             pgtype.Int8              `json:"target_balls"`
	BattingScorecardEntries []BattingScorecardRecord `json:"batting_scorecard_entries"`
	BowlingScorecardEntries []BowlingScorecardRecord `json:"bowling_scorecard_entries"`
}

type BattingScorecardRecord struct {
	BatterId        pgtype.Int8 `json:"batter_id"`
	BatterName      pgtype.Text `json:"batter_name"`
	BattingPosition pgtype.Int8 `json:"batting_position"`
	RunsScored      pgtype.Int8 `json:"runs_scored"`
	BallsFaced      pgtype.Int8 `json:"balls_faced"`
	FoursScored     pgtype.Int8 `json:"fours_scored"`
	SixesScored     pgtype.Int8 `json:"sixes_scored"`
	DismissedById   pgtype.Int8 `json:"dismissed_by_id"`
	DismissedByName pgtype.Text `json:"dismissed_by_name"`
	Fielder1Id      pgtype.Int8 `json:"fielder1_id"`
	Fielder1Name    pgtype.Text `json:"fielder1_name"`
	Fielder2Id      pgtype.Int8 `json:"fielder2_id"`
	Fielder2Name    pgtype.Text `json:"fielder2_name"`
	DismissalType   pgtype.Text `json:"dismissal_type"`
}

type BowlingScorecardRecord struct {
	BowlerId        pgtype.Int8 `json:"bowler_id"`
	BowlerName      pgtype.Text `json:"bowler_name"`
	BowlingPosition pgtype.Int8 `json:"bowling_position"`
	WicketsTaken    pgtype.Int8 `json:"wickets_taken"`
	RunsConceded    pgtype.Int8 `json:"runs_conceded"`
	BallsBowled     pgtype.Int8 `json:"balls_bowled"`
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

// Extras

type CareerStats struct {
	MatchesPlayed pgtype.Int8 `json:"matches_played"`

	InningsBatted     pgtype.Int8 `json:"innings_batted"`
	RunsScored        pgtype.Int8 `json:"runs_scored"`
	BattingDismissals pgtype.Int8 `json:"batting_dismissals"`
	BallsFaced        pgtype.Int8 `json:"balls_faced"`
	FoursScored       pgtype.Int8 `json:"fours_scored"`
	SixesScored       pgtype.Int8 `json:"sixes_scored"`
	CenturiesScored   pgtype.Int8 `json:"centuries_scored"`
	FiftiesScored     pgtype.Int8 `json:"fifties_scored"`
	HighestScore      pgtype.Int8 `json:"highest_score"`
	IsHighestNotOut   pgtype.Bool `json:"is_highest_not_out"`

	InningsBowled    pgtype.Int8 `json:"innings_bowled"`
	RunsConceded     pgtype.Int8 `json:"runs_conceded"`
	WicketsTaken     pgtype.Int8 `json:"wickets_taken"`
	BallsBowled      pgtype.Int8 `json:"balls_bowled"`
	FoursConceded    pgtype.Int8 `json:"fours_conceded"`
	SixesConceded    pgtype.Int8 `json:"sixes_conceded"`
	FourWktHauls     pgtype.Int8 `json:"four_wkt_hauls"`
	FiveWktHauls     pgtype.Int8 `json:"five_wkt_hauls"`
	TenWktHauls      pgtype.Int8 `json:"ten_wkt_hauls"`
	BestInnFigRuns   pgtype.Int8 `json:"best_inn_fig_runs"`
	BestInnFigWkts   pgtype.Int8 `json:"best_inn_fig_wkts"`
	BestMatchFigRuns pgtype.Int8 `json:"best_match_fig_runs"`
	BestMatchFigWkts pgtype.Int8 `json:"best_match_fig_wkts"`
}
