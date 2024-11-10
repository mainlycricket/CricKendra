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
	Id            pgtype.Int8 `json:"id"`
	Name          pgtype.Text `json:"name"`
	IsMale        pgtype.Bool `json:"is_male"`
	PlayingLevel  pgtype.Text `json:"playing_level"`
	PlayingFormat pgtype.Text `json:"playing_format"`
	Season        pgtype.Text `json:"season"`
}

// Tours

type AllToursResponse struct {
	Tours []AllTours `json:"tours"`
	Next  bool       `json:"next"`
}

type AllTours struct {
	Id              pgtype.Int8                `json:"id"`
	TouringTeamId   pgtype.Int8                `json:"touring_team_id"`
	TouringTeamName pgtype.Text                `json:"touring_team_name"`
	HostNations     []HostNationAsForeignField `json:"host_nations"`
	Season          pgtype.Text                `json:"season"`
}

// Innings

type AllMatchesResponse struct {
	Matches []AllMatches `json:"matches"`
	Next    bool         `json:"next"`
}

type AllMatches struct {
	Id          pgtype.Int8 `json:"id"`
	CricsheetId pgtype.Text `json:"cricsheet_id"`
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

// Extras

type CareerStats struct {
	MatchesPlayed pgtype.Int8         `json:"matches_played"`
	DebutMatchId  MatchAsForeignField `json:"debut_match"`
	LastMatchId   MatchAsForeignField `json:"last_match"`

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
