package responses

import "github.com/jackc/pgx/v5/pgtype"

// Players

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

type AllTeams struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	IsMale       pgtype.Bool `json:"is_male"`
	ImageURL     pgtype.Text `json:"image_url"`
	PlayingLevel pgtype.Text `json:"playing_level"`
	ShortName    pgtype.Text `json:"short_name"`
}

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
