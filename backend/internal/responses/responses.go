package responses

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
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

type SingleSeriesOverview struct {
	SeriesHeader SeriesHeader `json:"series_header"`

	WinnerTeamId   pgtype.Int8 `json:"winner_team_id"`
	WinnerTeamName pgtype.Text `json:"winner_team_name"`
	FinalStatus    pgtype.Text `json:"final_status"`

	FixtureMatches []MatchInfo `json:"fixture_matches"`
	ResultMatches  []MatchInfo `json:"result_matches"`
}

type SingleSeriesMatches struct {
	SeriesHeader SeriesHeader `json:"series_header"`
	Matches      []MatchInfo  `json:"matches"`
}

type SingleSeriesTeams struct {
	SeriesHeader SeriesHeader `json:"series_header"`

	Teams []struct {
		TeamId       pgtype.Int8 `json:"team_id"`
		TeamName     pgtype.Text `json:"team_name"`
		TeamImageUrl pgtype.Text `json:"team_image_url"`
	} `json:"teams"`
}

type SingleSeriesSquadsList struct {
	SeriesHeader SeriesHeader `json:"series_header"`

	SquadsList []struct {
		SquadId      pgtype.Int8 `json:"squad_id"`
		SquadLabel   pgtype.Text `json:"squad_label"`
		TeamImageUrl pgtype.Text `json:"team_image_url"`
	} `json:"squad_list"`
}

type SingleSeriesSingleSquad struct {
	SeriesHeader SeriesHeader `json:"series_header"`

	SquadsList []struct {
		SquadId    pgtype.Int8 `json:"squad_id"`
		SquadLabel pgtype.Text `json:"squad_label"`
	} `json:"squad_list"`

	Players []struct {
		PlayerId            pgtype.Int8 `json:"player_id"`
		PlayerName          pgtype.Text `json:"player_name"`
		PlayingRole         pgtype.Text `json:"playing_role"`
		DateOfBirth         pgtype.Date `json:"date_of_birth"`
		IsRHB               pgtype.Bool `json:"is_rhb"`
		PrimaryBowlingStyle pgtype.Text `json:"primary_bowling_style"`
		IsCaptain           pgtype.Bool `json:"is_captain"`
		IsViceCaptain       pgtype.Bool `json:"is_vice_captain"`
		IsWk                pgtype.Bool `json:"is_wk"`
	} `json:"players"`
}

type SeriesHeader struct {
	SeriesId       pgtype.Int8 `json:"series_id"`
	SeriesName     pgtype.Text `json:"series_name"`
	Season         pgtype.Text `json:"season"`
	TournamentId   pgtype.Int8 `json:"tournament_id"`
	TournamentName pgtype.Text `json:"tournament_name"`

	TopBatters []struct {
		BatterId       pgtype.Int8 `json:"batter_id"`
		BatterName     pgtype.Text `json:"batter_name"`
		BatterImageUrl pgtype.Text `json:"batter_image_url"`

		InningsBatted pgtype.Int8   `json:"innings_batted"`
		RunsScored    pgtype.Text   `json:"runs_scored"`
		Average       pgtype.Float8 `json:"average"`
	} `json:"top_batters"`

	TopBowlers []struct {
		BowlerId       pgtype.Int8 `json:"bowler_id"`
		BowlerName     pgtype.Text `json:"bowler_name"`
		BowlerImageUrl pgtype.Text `json:"bowler_image_url"`

		InningsBowled pgtype.Int8   `json:"innings_bowled"`
		WicketsTaken  pgtype.Text   `json:"wickets_taken"`
		Average       pgtype.Float8 `json:"average"`
	} `json:"top_bowlers"`
}

// Matches

type MatchInfo struct {
	MatchId               pgtype.Int8 `json:"match_id"`
	PlayingLevel          pgtype.Text `json:"playing_level"`
	PlayingFormat         pgtype.Text `json:"playing_format"`
	MatchType             pgtype.Text `json:"match_type"`
	EventMatchNumber      pgtype.Int8 `json:"event_match_number"`
	MatchState            pgtype.Text `json:"match_state"`
	MatchStateDescription pgtype.Text `json:"match_state_description"`
	FinalResult           pgtype.Text `json:"final_result"` // completed, abandoned, no result

	MatchWinnerId        pgtype.Int8 `json:"match_winner_team_id"`
	MatchLoserId         pgtype.Int8 `json:"match_loser_team_id"`
	IsWonByInnings       pgtype.Bool `json:"is_won_by_innings"`
	IsWonByRuns          pgtype.Bool `json:"is_won_by_runs"`
	WinMargin            pgtype.Int8 `json:"win_margin"`                // runs or wickets
	BallsMargin          pgtype.Int8 `json:"balls_remaining_after_win"` // successful chases
	SuperOverWinnerId    pgtype.Int8 `json:"super_over_winner_id"`
	BowlOutWinnerId      pgtype.Int8 `json:"bowl_out_winner_id"`
	OutcomeSpecialMethod pgtype.Text `json:"outcome_special_method"`
	TossWinnerId         pgtype.Int8 `json:"toss_winner_team_id"`
	TossLoserId          pgtype.Int8 `json:"toss_loser_team_id"`
	IsTossDecisionBat    pgtype.Bool `json:"is_toss_decision_bat"`

	Season           pgtype.Text        `json:"season"`
	StartDate        pgtype.Date        `json:"start_date"`
	EndDate          pgtype.Date        `json:"end_date"`
	StartDateTimeUtc pgtype.Timestamptz `json:"start_datetime_utc"`
	IsDayNight       pgtype.Bool        `json:"is_day_night"`
	GroundId         pgtype.Int8        `json:"ground_id"`
	GroundName       pgtype.Text        `json:"ground_name"`
	MainSeriesId     pgtype.Int8        `json:"main_series_id"`
	MainSeriesName   pgtype.Text        `json:"main_series_name"`

	Team1Id       pgtype.Int8 `json:"team1_id"`
	Team1Name     pgtype.Text `json:"team1_name"`
	Team1ImageUrl pgtype.Text `json:"team1_image_url"`
	Team2Id       pgtype.Int8 `json:"team2_id"`
	Team2Name     pgtype.Text `json:"team2_name"`
	Team2ImageUrl pgtype.Text `json:"team2_image_url"`

	InningsScores []TeamInningsShortInfo `json:"innings_scores"`
}

type TeamInningsShortInfo struct {
	InningsId       pgtype.Int8 `json:"innings_id"`
	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`

	TotalRuns    pgtype.Int8   `json:"total_runs"`
	TotalOvers   pgtype.Float8 `json:"total_overs"`
	TotalWickets pgtype.Int8   `json:"total_wickets"`
	InningsEnd   pgtype.Text   `json:"innings_end"`
	TargetRuns   pgtype.Int8   `json:"target_runs"`
	MaxOvers     pgtype.Int8   `json:"max_overs"`
}

type AllMatchesResponse struct {
	Matches []MatchInfo `json:"matches"`
	Next    bool        `json:"next"`
}

// 	PlayerId  pgtype.Int8 `json:"player_id"`
// 	AwardType pgtype.Text `json:"award_type"`
// }

type MatchHeader struct {
	MatchInfo
	PlayerAwards []PlayerAwardInfo `json:"player_awards"`
}

type PlayerAwardInfo struct {
	PlayerId   pgtype.Int8 `json:"player_id"`
	PlayerName pgtype.Text `json:"player_name"`
	AwardType  pgtype.Text `json:"award_type"`
}

// Match Summary

type MatchSummary struct {
	MatchHeader      MatchHeader             `json:"match_header"`
	ScorecardSummary []ScorecardSummaryEntry `json:"scorecard_summary,omitempty"`
	LatestCommentary []InningsBbbCommentary  `json:"latest_commentary"`
}

type ScorecardSummaryEntry struct {
	InningsId       pgtype.Int8 `json:"innings_id"`
	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`

	TotalRuns    pgtype.Int8   `json:"total_runs"`
	TotalWickets pgtype.Int8   `json:"total_wickets"`
	TotalOvers   pgtype.Float8 `json:"total_overs"`

	TopBatters []struct {
		BatterId    pgtype.Int8 `json:"batter_id"`
		BatterName  pgtype.Text `json:"batter_name"`
		RunsScored  pgtype.Int8 `json:"runs_scored"`
		BallsFaced  pgtype.Int8 `json:"balls_faced"`
		FourScored  pgtype.Int8 `json:"fours_scored"`
		SixesScored pgtype.Int8 `json:"sixes_scored"`
	} `json:"top_batters"`

	TopBowlers []struct {
		BolwerId     pgtype.Int8   `json:"bowler_id"`
		BowlerName   pgtype.Text   `json:"bowler_name"`
		OversBowled  pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers  pgtype.Int8   `json:"maiden_overs"`
		WicketsTaken pgtype.Int8   `json:"wickets_taken"`
		RunsConceded pgtype.Int8   `json:"runs_conceded"`
	} `json:"top_bowlers"`
}

/* Scorecards Page */

type MatchScorecardResponse struct {
	MatchHeader       MatchHeader            `json:"match_header"`
	InningsScorecards []TeamInningsScorecard `json:"innings_scorecards"`
}

type TeamInningsScorecard struct {
	InningsId       pgtype.Int8 `json:"innings_id"`
	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`

	TotalRuns    pgtype.Int8   `json:"total_runs"`
	TotalOvers   pgtype.Float8 `json:"total_overs"`
	TotalWickets pgtype.Int8   `json:"total_wickets"`
	Byes         pgtype.Int8   `json:"byes"`
	LegByes      pgtype.Int8   `json:"leg_byes"`
	Wides        pgtype.Int8   `json:"wides"`
	Noballs      pgtype.Int8   `json:"noballs"`
	Penalty      pgtype.Int8   `json:"penalty"`

	InningsEnd pgtype.Text `json:"innings_end"`
	TargetRuns pgtype.Int8 `json:"target_runs"`
	MaxOvers   pgtype.Int8 `json:"max_overs"`

	BatterScorecardEntries []BatterScorecardEntry `json:"batter_scorecard_entries"`
	BowlerScorecardEntries []BowlerScorecardEntry `json:"bowler_scorecard_entries"`
	FallOfWickets          []FallOfWickets        `json:"fall_of_wickets"`
}

type BatterScorecardEntry struct {
	BatterId        pgtype.Int8 `json:"batter_id"`
	BatterName      pgtype.Text `json:"batter_name"`
	BattingPosition pgtype.Int8 `json:"batting_position"`
	HasBatted       pgtype.Bool `json:"has_batted"`

	RunsScored    pgtype.Int8 `json:"runs_scored"`
	BallsFaced    pgtype.Int8 `json:"balls_faced"`
	MinutesBatted pgtype.Int8 `json:"minutes_batted"`
	FoursScored   pgtype.Int8 `json:"fours_scored"`
	SixesScored   pgtype.Int8 `json:"sixes_scored"`

	DismissalType   *pgtype.Text `json:"dismissal_type,omitempty"`
	DismissedById   *pgtype.Int8 `json:"dismissed_by_id,omitempty"`
	DismissedByName *pgtype.Text `json:"dismissed_by_name,omitempty"`
	Fielder1Id      *pgtype.Int8 `json:"fielder1_id,omitempty"`
	Fielder1Name    *pgtype.Text `json:"fielder1_name,omitempty"`
	Fielder2Id      *pgtype.Int8 `json:"fielder2_id,omitempty"`
	Fielder2Name    *pgtype.Text `json:"fielder2_name,omitempty"`
}

type BowlerScorecardEntry struct {
	BowlerId        pgtype.Int8 `json:"bowler_id"`
	BowlerName      pgtype.Text `json:"bowler_name"`
	BowlingPosition pgtype.Int8 `json:"bowling_position"`

	WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
	RunsConceded    pgtype.Int8   `json:"runs_conceded"`
	OversBowled     pgtype.Float8 `json:"overs_bowled"`
	MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
	FoursConceded   pgtype.Int8   `json:"fours_conceded"`
	SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	WidesConceded   pgtype.Int8   `json:"wides_conceded"`
	NoballsConceded pgtype.Int8   `json:"noballs_conceded"`
}

type FallOfWickets struct {
	BatterId      pgtype.Int8   `json:"batter_id"`
	BatterName    pgtype.Text   `json:"batter_name"`
	BallNumber    pgtype.Float8 `json:"ball_number"`
	TeamRuns      pgtype.Int8   `json:"team_runs"`
	WicketNumber  pgtype.Int8   `json:"wicket_number"`
	DismissalType pgtype.Text   `json:"dismissal_type"`
}

// Match Stats

type MatchStatsResponse struct {
	MatchHeader MatchHeader `json:"match_header"`

	Innings []struct {
		InningsId       pgtype.Int8 `json:"innings_id"`
		InningsNumber   pgtype.Int8 `json:"innings_number"`
		BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
		BattingTeamName pgtype.Text `json:"batting_team_name"`

		Partnerships []struct {
			ForWicket        pgtype.Int8   `json:"for_wicket"`
			IsUnbeaten       pgtype.Bool   `json:"is_unbeaten"`
			StartBallNumber  pgtype.Float8 `json:"start_ball_number"`
			EndBallNumber    pgtype.Float8 `json:"end_ball_number"`
			PartnershipsRuns pgtype.Int8   `json:"partnership_runs"`

			Batter1Id    pgtype.Int8 `json:"batter1_id"`
			Batter1Name  pgtype.Text `json:"batter1_name"`
			Batter1Runs  pgtype.Int8 `json:"batter1_runs"`
			Batter1Balls pgtype.Int8 `json:"batter1_balls"`

			Batter2Id    pgtype.Int8 `json:"batter2_id"`
			Batter2Name  pgtype.Text `json:"batter2_name"`
			Batter2Runs  pgtype.Int8 `json:"batter2_runs"`
			Batter2Balls pgtype.Int8 `json:"batter2_balls"`
		} `json:"partnerships"`

		Overs []struct {
			OverNumber pgtype.Int8 `json:"over_number"`
			Runs       pgtype.Int8 `json:"runs"`
			Balls      pgtype.Int8 `json:"balls"`
			Wickets    pgtype.Int8 `json:"wickets"`
		} `json:"overs"`
	} `json:"innings"`
}

// Match Squad

type MatchSquadResponse struct {
	MatchHeader MatchHeader      `json:"match_header"`
	TeamSquads  []TeamSquadEntry `json:"team_squads"`
}

type TeamSquadEntry struct {
	TeamId   pgtype.Int8        `json:"team_id"`
	TeamName pgtype.Text        `json:"team_name"`
	Players  []PlayerSquadEntry `json:"players"`
}

type PlayerSquadEntry struct {
	PlayerId   pgtype.Int8 `json:"player_id"`
	PlayerName pgtype.Text `json:"player_name"`

	IsCaptain     pgtype.Bool `json:"is_captain"`
	IsWk          pgtype.Bool `json:"is_wk"`
	IsDebut       pgtype.Bool `json:"is_debut"`
	IsViceCaptain pgtype.Bool `json:"is_vice_captain"`
	PlayingStatus pgtype.Text `json:"playing_status"`
}

/* Commentary Page */

// single innings
type MatchCommentaryResponse struct {
	MatchHeader MatchHeader            `json:"match_header"`
	Commentary  []InningsBbbCommentary `json:"commentary"`
}

type InningsBbbCommentary struct {
	InningsId             pgtype.Int8   `json:"innings_id"`
	InningsDeliveryNumber pgtype.Int8   `json:"innings_delivery_number"`
	BallNumber            pgtype.Float8 `json:"ball_number"`
	OverNumber            pgtype.Int8   `json:"over_number"`

	BatterId     *pgtype.Int8 `json:"batter_id,omitempty"`
	BatterName   *pgtype.Text `json:"batter_name,omitempty"`
	BowlerId     *pgtype.Int8 `json:"bowler_id,omitempty"`
	BowlerName   *pgtype.Text `json:"bowler_name,omitempty"`
	Fielder1Id   *pgtype.Int8 `json:"fielder1_id,omitempty"`
	Fielder1Name *pgtype.Text `json:"fielder1_name,omitempty"`
	Fielder2Id   *pgtype.Int8 `json:"fielder2_id,omitempty"`
	Fielder2Name *pgtype.Text `json:"fielder2_name,omitempty"`

	Wides     pgtype.Int8 `json:"wides"`
	Noballs   pgtype.Int8 `json:"noballs"`
	Legbyes   pgtype.Int8 `json:"legbyes"`
	Byes      pgtype.Int8 `json:"byes"`
	TotalRuns pgtype.Int8 `json:"total_runs"`
	IsFour    pgtype.Bool `json:"is_four"`
	IsSix     pgtype.Bool `json:"is_six"`

	Player1DismissedId   *pgtype.Int8 `json:"player1_dismissed_id,omitempty"`
	Player1DismissedName *pgtype.Text `json:"player1_dismissed_name,omitempty"`
	Player1DismissalType *pgtype.Text `json:"player1_dismissal_type,omitempty"`
	Player1DimissedRuns  *pgtype.Int8 `json:"player1_dismissed_runs,omitempty"`
	Player1DimissedBalls *pgtype.Int8 `json:"player1_dismissed_balls,omitempty"`
	Player1DimissedFours *pgtype.Int8 `json:"player1_dismissed_fours,omitempty"`
	Player1DimissedSixes *pgtype.Int8 `json:"player1_dismissed_sixes,omitempty"`

	Player2DismissedId   *pgtype.Int8 `json:"player2_dismissed_id,omitempty"`
	Player2DismissedName *pgtype.Text `json:"player2_dismissed_name,omitempty"`
	Player2DismissalType *pgtype.Text `json:"player2_dismissal_type,omitempty"`
	Player2DimissedRuns  *pgtype.Int8 `json:"player2_dismissed_runs,omitempty"`
	Player2DimissedBalls *pgtype.Int8 `json:"player2_dismissed_balls,omitempty"`
	Player2DimissedFours *pgtype.Int8 `json:"player2_dismissed_fours,omitempty"`
	Player2DimissedSixes *pgtype.Int8 `json:"player2_dismissed_sixes,omitempty"`

	Commentary pgtype.Text `json:"commentary"`
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

type StatsFilters struct {
	PrimaryTeams    []TeamAsForeignField       `json:"primary_teams"`
	OppositionTeams []TeamAsForeignField       `json:"opposition_teams"`
	HostNations     []HostNationAsForeignField `json:"host_nations"`
	Continents      []ContinentAsForeignField  `json:"continents"`
	Grounds         []GroundAsForeignField     `json:"grounds"`
	MinDate         pgtype.Date                `json:"min_date"`
	MaxDate         pgtype.Date                `json:"max_date"`
	Seasons         []pgtype.Text              `json:"seasons"`
	Series          []SeriesAsForeignField     `json:"series"`
	Tournaments     []TournamentAsForeignField `json:"tournaments"`
}
