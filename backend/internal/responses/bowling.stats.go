package responses

import "github.com/jackc/pgx/v5/pgtype"

/* Overall Stats */

type Overall_Bowling_Summary_Group struct {
	Teams []struct {
		TeamId       pgtype.Int8 `json:"team_id"`
		TeamName     pgtype.Text `json:"team_name"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"teams"`

	Oppositions []struct {
		OppositionTeamId   pgtype.Int8 `json:"opposition_team_id"`
		OppositionTeamName pgtype.Text `json:"opposition_team_name"`
		PlayersCount       pgtype.Int8 `json:"players_count"`
		MinDate            pgtype.Date `json:"min_date"`
		MaxDate            pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"oppositions"`

	HostNations []struct {
		HostNationId   pgtype.Int8 `json:"host_nation_id"`
		HostNationName pgtype.Text `json:"host_nation_name"`
		PlayersCount   pgtype.Int8 `json:"players_count"`
		MinDate        pgtype.Date `json:"min_date"`
		MaxDate        pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"host_nations"`

	Continents []struct {
		ContinentId   pgtype.Int8 `json:"continent_id"`
		ContinentName pgtype.Text `json:"continent_name"`
		PlayersCount  pgtype.Int8 `json:"players_count"`
		MinDate       pgtype.Date `json:"min_date"`
		MaxDate       pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"continents"`

	Years []struct {
		Year         pgtype.Int8 `json:"year"`
		PlayersCount pgtype.Int8 `json:"players_count"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"years"`

	Seasons []struct {
		Season       pgtype.Text `json:"season"`
		PlayersCount pgtype.Int8 `json:"players_count"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"seasons"`

	HomeAway []struct {
		HomeAwayLabel pgtype.Text `json:"home_away_label"`
		PlayersCount  pgtype.Int8 `json:"players_count"`
		MinDate       pgtype.Date `json:"min_date"`
		MaxDate       pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"home_away"`

	TossWonLost []struct {
		TossResult   pgtype.Text `json:"toss_result"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"toss_won_lost"`

	TossDecision []struct {
		TossResult        pgtype.Text `json:"toss_result"`
		IsTossDecisionBat pgtype.Bool `json:"is_toss_decision_bat"`
		PlayersCount      pgtype.Int8 `json:"players_count"`
		MinDate           pgtype.Date `json:"min_date"`
		MaxDate           pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"toss_decision"`

	BatBowlFirst []struct {
		BatBowlFirst pgtype.Text `json:"bat_bowl_first"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"bat_bowl_first"`

	InningsNumber []struct {
		InningsNumber pgtype.Int8 `json:"innings_number"`
		PlayersCount  pgtype.Int8 `json:"players_count"`
		MinDate       pgtype.Date `json:"min_date"`
		MaxDate       pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"innings_number"`

	MatchResult []struct {
		MatchResult  pgtype.Text `json:"match_result"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"match_result"`

	MatchResultBatBowlFirst []struct {
		MatchResult  pgtype.Text `json:"match_result"`
		BatBowlFirst pgtype.Text `json:"bat_bowl_first"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"match_result_bat_bowl_first"`

	SeriesTeamsCount []struct {
		TeamsCount   pgtype.Text `json:"teams_count"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"series_teams_count"`

	SeriesEventMatchNumber []struct {
		EventMatchNumber pgtype.Int8 `json:"event_match_number"`
		PlayersCount     pgtype.Int8 `json:"players_count"`
		MinDate          pgtype.Date `json:"min_date"`
		MaxDate          pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"series_event_match_number"`

	Tournaments []struct {
		TournamentId   pgtype.Int8 `json:"tournament_id"`
		TournamentName pgtype.Text `json:"tournament_name"`
		PlayersCount   pgtype.Int8 `json:"players_count"`
		MinDate        pgtype.Date `json:"min_date"`
		MaxDate        pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"tournaments"`

	BowlingPositions []struct {
		BowlingPosition pgtype.Int8 `json:"bowling_position"`
		PlayersCount    pgtype.Int8 `json:"players_count"`
		MinDate         pgtype.Date `json:"min_date"`
		MaxDate         pgtype.Date `json:"max_date"`

		MatchesPlayed   pgtype.Int8   `json:"matches_played"`
		InningsBowled   pgtype.Int8   `json:"innings_bowled"`
		OversBowled     pgtype.Float8 `json:"overs_bowled"`
		MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
		RunsConceded    pgtype.Int8   `json:"runs_conceded"`
		WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
		Average         pgtype.Float8 `json:"average"`
		StrikeRate      pgtype.Float8 `json:"strike_rate"`
		Economy         pgtype.Float8 `json:"economy"`
		FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
		FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
		TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
		BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
		BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
		BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
		BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
		FoursConceded   pgtype.Int8   `json:"fours_conceded"`
		SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
	} `json:"bowling_positions"`
}

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
	TeamId       pgtype.Int8 `json:"team_id"`
	TeamName     pgtype.Text `json:"team_name"`
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

type Overall_Bowling_Series_Group struct {
	SeriesId     pgtype.Int8 `json:"series_id"`
	SeriesName   pgtype.Text `json:"series_name"`
	SeriesSeason pgtype.Text `json:"series_season"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBowlingStats
}

type Overall_Bowling_Tournament_Group struct {
	TournamentId   pgtype.Int8 `json:"tournament_id"`
	TournamentName pgtype.Text `json:"tournament_name"`
	PlayersCount   pgtype.Int8 `json:"players_count"`
	MinDate        pgtype.Date `json:"min_date"`
	MaxDate        pgtype.Date `json:"max_date"`
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

type Overall_Bowling_Decade_Group struct {
	Decade       pgtype.Int8 `json:"decade"`
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
	MaidenOvers   pgtype.Float8 `json:"maiden_overs"`
	RunsConceded  pgtype.Int8   `json:"runs_conceded"`
	WicketsTaken  pgtype.Int8   `json:"wickets_taken"`
	Economy       pgtype.Float8 `json:"economy"`
	FoursConceded pgtype.Int8   `json:"fours_conceded"`
	SixesConceded pgtype.Int8   `json:"sixes_conceded"`
}

type Individual_Bowling_MatchTotals_Group struct {
	MatchId   pgtype.Int8 `json:"match_id"`
	StartDate pgtype.Date `json:"start_date"`
	GroundId  pgtype.Int8 `json:"ground_id"`
	CityName  pgtype.Text `json:"city_name"`

	BowlerId        pgtype.Int8 `json:"bowler_id"`
	BowlerName      pgtype.Text `json:"bowler_name"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`
	BowlingTeamId   pgtype.Int8 `json:"bowling_team_id"`
	BowlingTeamName pgtype.Text `json:"bowling_team_name"`

	OversBowled   pgtype.Float8 `json:"overs_bowled"`
	MaidenOvers   pgtype.Float8 `json:"maiden_overs"`
	RunsConceded  pgtype.Int8   `json:"runs_conceded"`
	WicketsTaken  pgtype.Int8   `json:"wickets_taken"`
	Average       pgtype.Float8 `json:"average"`
	Economy       pgtype.Float8 `json:"economy"`
	StrikeRate    pgtype.Float8 `json:"strike_rate"`
	FoursConceded pgtype.Int8   `json:"fours_conceded"`
	SixesConceded pgtype.Int8   `json:"sixes_conceded"`
}

type Individual_Bowling_Ground_Group struct {
	GroundId   pgtype.Int8 `json:"ground_id"`
	GroundName pgtype.Text `json:"ground_name"`
	Overall_Bowling_Bowler_Group
}

type Individual_Bowling_Series_Group struct {
	SeriesId     pgtype.Int8 `json:"series_id"`
	SeriesName   pgtype.Text `json:"series_name"`
	SeriesSeason pgtype.Text `json:"series_season"`
	Overall_Bowling_Bowler_Group
}

type Individual_Bowling_Tournament_Group struct {
	TournamentId   pgtype.Int8 `json:"tournament_id"`
	TournamentName pgtype.Text `json:"tournament_name"`
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
	MaidenOvers     pgtype.Int8   `json:"maiden_overs"`
	RunsConceded    pgtype.Int8   `json:"runs_conceded"`
	WicketsTaken    pgtype.Int8   `json:"wickets_taken"`
	Average         pgtype.Float8 `json:"average"`
	StrikeRate      pgtype.Float8 `json:"strike_rate"`
	Economy         pgtype.Float8 `json:"economy"`
	FourWktHauls    pgtype.Int8   `json:"four_wicket_hauls"`
	FiveWktHauls    pgtype.Int8   `json:"five_wicket_hauls"`
	TenWktHauls     pgtype.Int8   `json:"ten_wicket_hauls"`
	BestMatchWkts   pgtype.Int8   `json:"best_match_wickets"`
	BestMatchRuns   pgtype.Int8   `json:"best_match_runs"`
	BestInningsWkts pgtype.Int8   `json:"best_innings_wickets"`
	BestInningsRuns pgtype.Int8   `json:"best_innings_runs"`
	FoursConceded   pgtype.Int8   `json:"fours_conceded"`
	SixesConceded   pgtype.Int8   `json:"sixes_conceded"`
}
