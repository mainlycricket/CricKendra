package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CareerStats struct {
	MatchesPlayed pgtype.Int8 `json:"matches_played"`
	DebutMatchId  pgtype.Int8 `json:"debut_match_id"`
	LastMatchId   pgtype.Int8 `json:"last_match_id"`

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

type Team struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	IsMale       pgtype.Bool `json:"is_male"`
	ImageURL     pgtype.Text `json:"image_url"`
	PlayingLevel pgtype.Text `json:"playing_level"`
	ShortName    pgtype.Text `json:"short_name"`
}

type Player struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	FullName    pgtype.Text `json:"full_name"`
	PlayingRole pgtype.Text `json:"playing_role"`
	Nationality pgtype.Text `json:"nationality"`
	IsMale      pgtype.Bool `json:"is_male"`
	DateOfBirth pgtype.Date `json:"date_of_birth"`
	ImageURL    pgtype.Text `json:"image_url"`
	Biography   pgtype.Text `json:"biography"`

	IsRHB               pgtype.Bool   `json:"is_rhb"`
	BowlingStyles       []pgtype.Text `json:"bowling_styles"`
	PrimaryBowlingStyle pgtype.Text   `json:"primary_bowling_style"`
	TeamsRepresentedId  []pgtype.Int8 `json:"teams_represented_id"`

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

type Continent struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type HostNation struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	ContinentId pgtype.Int8 `json:"continent_id"`
}

type City struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	HostNationId pgtype.Int8 `json:"host_nation_id"`
}

type Ground struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	HostNationId pgtype.Int8 `json:"host_nation_id"`
	CityId       pgtype.Int8 `json:"city_id"`
}

type Tournament struct {
	Id            pgtype.Int8 `json:"id"`
	Name          pgtype.Text `json:"name"`
	IsMale        pgtype.Bool `json:"is_male"`
	PlayingLevel  pgtype.Text `json:"playing_level"`
	PlayingFormat pgtype.Text `json:"playing_format"`
}

type Season struct {
	Season pgtype.Text `json:"season"`
}

type Tour struct {
	Id            pgtype.Int8   `json:"id"`
	TouringTeamId pgtype.Int8   `json:"touring_team_id"`
	HostNationsId []pgtype.Int8 `json:"host_nations_id"`
	Season        pgtype.Text   `json:"season"`
}

type Series struct {
	Id             pgtype.Int8   `json:"id"`
	Name           pgtype.Text   `json:"name"`
	IsMale         pgtype.Bool   `json:"is_male"`
	PlayingLevel   pgtype.Text   `json:"playing_level"`
	PlayingFormat  pgtype.Text   `json:"playing_format"`
	Season         pgtype.Text   `json:"season"`
	TeamsId        []pgtype.Int8 `json:"teams_id"`
	HostNationsId  []pgtype.Int8 `json:"host_nations_id"`
	TournamentId   pgtype.Int8   `json:"tournament_id"`
	ParentSeriesId pgtype.Int8   `json:"parent_series_id"`
	TourId         pgtype.Int8   `json:"tour_id"`
}

type Squad struct {
	PlayerId      pgtype.Int8 `json:"player_id"`
	SeriesId      pgtype.Int8 `json:"series_id"`
	MatchId       pgtype.Int8 `json:"match_id"`
	IsCaptain     pgtype.Bool `json:"is_captain"`
	IsWk          pgtype.Bool `json:"is_wk"`
	IsDebut       pgtype.Bool `json:"is_debut"`
	PlayingStatus pgtype.Text `json:"playing_status"` // withdrawn, playing XI, bench, substitute
}

type Match struct {
	Id                pgtype.Int8        `json:"id"`
	StartDateTime     pgtype.Timestamptz `json:"start_datetime"`
	Team1Id           pgtype.Int8        `json:"team1_id"`
	Team2Id           pgtype.Int8        `json:"team2_id"`
	IsMale            pgtype.Bool        `json:"is_male"`
	TournamentId      pgtype.Int8        `json:"tournament_id"`
	SeriesId          pgtype.Int8        `json:"series_id"`
	ParentSeriesId    pgtype.Int8        `json:"parent_series_id"`
	TourId            pgtype.Int8        `json:"tour_id"`
	HostNationId      pgtype.Int8        `json:"host_nation_id"`
	ContinentId       pgtype.Int8        `json:"continent_id"`
	GroundId          pgtype.Int8        `json:"ground_id"`
	CurrentStatus     pgtype.Text        `json:"current_status"`
	FinalResult       pgtype.Text        `json:"final_result"` // completed, abandoned, no result
	HomeTeamId        pgtype.Int8        `json:"home_team_id"`
	AwayTeamId        pgtype.Int8        `json:"away_team_id"`
	MatchType         pgtype.Text        `json:"match_type"` // preliminary, semifinal, final
	PlayingLevel      pgtype.Text        `json:"playing_level"`
	PlayingFormat     pgtype.Text        `json:"playing_format"`
	Season            pgtype.Text        `json:"season"`
	IsDayNight        pgtype.Bool        `json:"is_day_night"`
	IsDLS             pgtype.Bool        `json:"is_dls"`
	TossWinnerId      pgtype.Int8        `json:"toss_winner_team_id"`
	TossLoserId       pgtype.Int8        `json:"toss_loser_team_id"`
	IsTossDecisionBat pgtype.Bool        `json:"is_toss_decision_bat"`
	MatchWinnerId     pgtype.Int8        `json:"match_winner_team_id"`
	MatchLoserId      pgtype.Int8        `json:"match_loser_team_id"`
	IsWonByRuns       pgtype.Bool        `json:"is_won_by_runs"`
	WinMargin         pgtype.Int8        `json:"win_margin"`                // runs or wickets
	BallsMargin       pgtype.Int8        `json:"balls_remaining_after_win"` // successful chases
	PotmId            pgtype.Int8        `json:"potm_id"`
	BallsPerOver      pgtype.Int8        `json:"balls_per_over"`

	ScorersId      []pgtype.Int8 `json:"scorers_id"`
	CommentatorsId []pgtype.Int8 `json:"commentators_id"`
}

type Innings struct {
	Id            pgtype.Int8 `json:"id"`
	MatchId       pgtype.Int8 `json:"match_id"`
	InningsNumber pgtype.Int8 `json:"innings_number"`
	BattingTeamId pgtype.Int8 `json:"batting_team_id"`
	BowlingTeamId pgtype.Int8 `json:"bowling_team_id"`
	TotalRuns     pgtype.Int8 `json:"total_runs"`
	TotalWkts     pgtype.Int8 `json:"total_wickets"`
	Byes          pgtype.Int8 `json:"byes"`
	Legbyes       pgtype.Int8 `json:"leg_byes"`
	Wides         pgtype.Int8 `json:"wides"`
	Noballs       pgtype.Int8 `json:"noballs"`
	Penalty       pgtype.Int8 `json:"penalty"`
	IsSuperOver   pgtype.Bool `json:"is_super_over"`
	Status        pgtype.Text `json:"status"` // completed, declared, fortfeited, all out
}

type BattingScorecard struct {
	Id              pgtype.Int8 `json:"id"`
	InningsId       pgtype.Int8 `json:"innings_id"`
	BatterId        pgtype.Int8 `json:"batter_id"`
	BattingPosition pgtype.Int8 `json:"batting_position"`
	RunsScored      pgtype.Int8 `json:"runs_scored"`
	BallsFaced      pgtype.Int8 `json:"balls_faced"`
	MinutesBatted   pgtype.Int8 `json:"minutes_batted"`
	FoursScored     pgtype.Int8 `json:"fours_scored"`
	SixesScored     pgtype.Int8 `json:"sixes_scored"`
	DismissedById   pgtype.Int8 `json:"dismissed_by_id"`
	DismissalType   pgtype.Text `json:"dismissal_type"`
	DismissalBallId pgtype.Int8 `json:"dismissal_ball_id"`
	Fielder1Id      pgtype.Int8 `json:"fielder1_id"`
	Fielder2Id      pgtype.Int8 `json:"fielder2_id"`
}

type BowlingScorecard struct {
	Id              pgtype.Int8 `json:"id"`
	InningsId       pgtype.Int8 `json:"innings_id"`
	BowlerId        pgtype.Int8 `json:"bowler_id"`
	BowlingPosition pgtype.Int8 `json:"bowling_position"`
	WicketsTaken    pgtype.Int8 `json:"wickets_taken"`
	RunsConceded    pgtype.Int8 `json:"runs_conceded"`
	BallsBowled     pgtype.Int8 `json:"balls_bowled"`
	FoursConceded   pgtype.Int8 `json:"fours_conceded"`
	SixesConceded   pgtype.Int8 `json:"sixes_conceded"`
	WidesConceded   pgtype.Int8 `json:"wides_conceded"`
	NoballsConceded pgtype.Int8 `json:"noballs_conceded"`
}

type Deliveries struct {
	Id                   pgtype.Int8        `json:"id"`
	InningsId            pgtype.Int8        `json:"innings_id"`
	BallNumber           pgtype.Float8      `json:"ball_number"`
	OverNumber           pgtype.Int8        `json:"over_number"`
	BatterId             pgtype.Int8        `json:"batter_id"`
	BowlerId             pgtype.Int8        `json:"bowler_id"`
	NonStrikerId         pgtype.Int8        `json:"non_striker_id"`
	BattingTeamId        pgtype.Int8        `json:"batting_team_id"`
	BowlingTeamId        pgtype.Int8        `json:"bowling_team_id"`
	BatterRuns           pgtype.Int8        `json:"batter_runs"`
	Wides                pgtype.Int8        `json:"wides"`
	Noballs              pgtype.Int8        `json:"noballs"`
	Legbyes              pgtype.Int8        `json:"legbyes"`
	Byes                 pgtype.Int8        `json:"byes"`
	Penalty              pgtype.Int8        `json:"penalty"`
	TotalExtras          pgtype.Int8        `json:"total_extras"`
	TotalRuns            pgtype.Int8        `json:"total_runs"`
	BowlerRuns           pgtype.Int8        `json:"bowler_runs"`
	IsFour               pgtype.Bool        `json:"is_four"`
	IsSix                pgtype.Bool        `json:"is_six"`
	Player1DismissedId   pgtype.Int8        `json:"player1_dismissed_id"`
	Player1DismissalType pgtype.Text        `json:"player1_dismissal_type"`
	Player2DismissedId   pgtype.Int8        `json:"player2_dismissed_id"`
	Player2DismissalType pgtype.Text        `json:"player2_dismissal_type"`
	IsPace               pgtype.Bool        `json:"is_pace"`            // true if pacer, false if spin
	BowlingStyle         pgtype.Text        `json:"bowling_style"`      // RAFM, LAFM, LAF etc
	IsBatterRHB          pgtype.Bool        `json:"is_batter_rhb"`      // true if batter is RHB, false if LHB
	IsNonStrikerRHB      pgtype.Bool        `json:"is_non_striker_rhb"` // true if non-striker is RHB, false if LHB
	Line                 pgtype.Text        `json:"line"`
	Length               pgtype.Text        `json:"length"`
	BallType             pgtype.Text        `json:"ball_type"`  // inswinger, googly
	BallSpeed            pgtype.Float8      `json:"ball_speed"` // kph
	Misc                 pgtype.Text        `json:"misc"`       // edged, missed
	WwRegion             pgtype.Text        `json:"ww_region"`  // cover, mid-wkt
	FootType             pgtype.Text        `json:"foot_type"`  // front foot, back foot, step out
	ShotType             pgtype.Text        `json:"shot_type"`  // straight drive, pull shot
	Fielder1Id           pgtype.Int8        `json:"fielder1_id"`
	Fielder2Id           pgtype.Int8        `json:"fielder2_id"`
	Commentary           pgtype.Text        `json:"commentary"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
}

type BlogArticles struct {
	Id             pgtype.Int8        `json:"id"`
	Title          pgtype.Text        `json:"title"`
	Content        pgtype.Text        `json:"content"`
	AuthorId       pgtype.Int8        `json:"author_id"`
	Category       pgtype.Int8        `json:"category"`
	Status         pgtype.Text        `json:"status"`
	PlayerTags     []pgtype.Int8      `json:"player_tags"`
	TeamTags       []pgtype.Int8      `json:"team_tags"`
	SeriesTags     []pgtype.Int8      `json:"series_tags"`
	TournamentTags []pgtype.Int8      `json:"tournament_tags"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	Id       pgtype.Int8 `json:"id"`
	Name     pgtype.Text `json:"name"`
	Email    pgtype.Text `json:"email"`
	Password pgtype.Text `json:"password"`
	Role     pgtype.Text `json:"role"`
}
