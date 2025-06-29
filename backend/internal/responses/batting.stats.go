package responses

import "github.com/jackc/pgx/v5/pgtype"

/* Overall Stats */

type Overall_Batting_Summary_Group struct {
	Teams []struct {
		TeamId       pgtype.Int8 `json:"team_id"`
		TeamName     pgtype.Text `json:"team_name"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"teams"`

	Oppositions []struct {
		OppositionTeamId   pgtype.Int8 `json:"opposition_team_id"`
		OppositionTeamName pgtype.Text `json:"opposition_team_name"`
		PlayersCount       pgtype.Int8 `json:"players_count"`
		MinDate            pgtype.Date `json:"min_date"`
		MaxDate            pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"oppositions"`

	HostNations []struct {
		HostNationId   pgtype.Int8 `json:"host_nation_id"`
		HostNationName pgtype.Text `json:"host_nation_name"`
		PlayersCount   pgtype.Int8 `json:"players_count"`
		MinDate        pgtype.Date `json:"min_date"`
		MaxDate        pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"host_nations"`

	Continents []struct {
		ContinentId   pgtype.Int8 `json:"continent_id"`
		ContinentName pgtype.Text `json:"continent_name"`
		PlayersCount  pgtype.Int8 `json:"players_count"`
		MinDate       pgtype.Date `json:"min_date"`
		MaxDate       pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"continents"`

	Years []struct {
		Year         pgtype.Int8 `json:"year"`
		PlayersCount pgtype.Int8 `json:"players_count"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"years"`

	Seasons []struct {
		Season       pgtype.Text `json:"season"`
		PlayersCount pgtype.Int8 `json:"players_count"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"seasons"`

	HomeAway []struct {
		HomeAwayLabel pgtype.Text `json:"home_away_label"`
		PlayersCount  pgtype.Int8 `json:"players_count"`
		MinDate       pgtype.Date `json:"min_date"`
		MaxDate       pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"home_away"`

	TossWonLost []struct {
		TossResult   pgtype.Text `json:"toss_result"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"toss_won_lost"`

	TossDecision []struct {
		TossResult        pgtype.Text `json:"toss_result"`
		IsTossDecisionBat pgtype.Bool `json:"is_toss_decision_bat"`
		PlayersCount      pgtype.Int8 `json:"players_count"`
		MinDate           pgtype.Date `json:"min_date"`
		MaxDate           pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"toss_decision"`

	BatBowlFirst []struct {
		BatBowlFirst pgtype.Text `json:"bat_bowl_first"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"bat_bowl_first"`

	InningsNumber []struct {
		InningsNumber pgtype.Int8 `json:"innings_number"`
		PlayersCount  pgtype.Int8 `json:"players_count"`
		MinDate       pgtype.Date `json:"min_date"`
		MaxDate       pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"innings_number"`

	MatchResult []struct {
		MatchResult  pgtype.Text `json:"match_result"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"match_result"`

	MatchResultBatBowlFirst []struct {
		MatchResult  pgtype.Text `json:"match_result"`
		BatBowlFirst pgtype.Text `json:"bat_bowl_first"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"match_result_bat_bowl_first"`

	SeriesTeamsCount []struct {
		TeamsCount   pgtype.Text `json:"teams_count"`
		PlayersCount pgtype.Int8 `json:"players_count"`
		MinDate      pgtype.Date `json:"min_date"`
		MaxDate      pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"series_teams_count"`

	SeriesEventMatchNumber []struct {
		EventMatchNumber pgtype.Int8 `json:"event_match_number"`
		PlayersCount     pgtype.Int8 `json:"players_count"`
		MinDate          pgtype.Date `json:"min_date"`
		MaxDate          pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"series_event_match_number"`

	Tournaments []struct {
		TournamentId   pgtype.Int8 `json:"tournament_id"`
		TournamentName pgtype.Text `json:"tournament_name"`
		PlayersCount   pgtype.Int8 `json:"players_count"`
		MinDate        pgtype.Date `json:"min_date"`
		MaxDate        pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"tournaments"`

	BattingPositions []struct {
		BattingPosition pgtype.Int8 `json:"batting_position"`
		PlayersCount    pgtype.Int8 `json:"players_count"`
		MinDate         pgtype.Date `json:"min_date"`
		MaxDate         pgtype.Date `json:"max_date"`

		MatchesPlayed      pgtype.Int8   `json:"matches_played"`
		InningsBatted      pgtype.Int8   `json:"innings_batted"`
		RunsScored         pgtype.Int8   `json:"runs_scored"`
		BallsFaced         pgtype.Int8   `json:"balls_faced"`
		NotOuts            pgtype.Int8   `json:"not_outs"`
		Average            pgtype.Float8 `json:"average"`
		StrikeRate         pgtype.Float8 `json:"strike_rate"`
		HighestScore       pgtype.Int8   `json:"highest_score"`
		HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
		Centuries          pgtype.Int8   `json:"centuries"`
		HalfCenturies      pgtype.Int8   `json:"half_centuries"`
		FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
		Ducks              pgtype.Int8   `json:"ducks"`
		FoursScored        pgtype.Int8   `json:"fours_scored"`
		SixesScored        pgtype.Int8   `json:"sixes_scored"`
	} `json:"batting_positions"`
}

type Overall_Batting_Batter_Group struct {
	BatterId         pgtype.Int8   `json:"batter_id"`
	BatterName       pgtype.Text   `json:"batter_name"`
	TeamsRepresented []pgtype.Text `json:"teams_represented"`
	MinDate          pgtype.Date   `json:"min_date"`
	MaxDate          pgtype.Date   `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_TeamInnings_Group struct {
	MatchId         pgtype.Int8 `json:"match_id"`
	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`
	BowlingTeamId   pgtype.Int8 `json:"bowling_team_id"`
	BowlingTeamName pgtype.Text `json:"bowling_team_name"`
	Season          pgtype.Text `json:"season"`
	CityName        pgtype.Text `json:"city_name"`
	StartDate       pgtype.Date `json:"start_date"`
	PlayersCount    pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type Overall_Batting_Match_Group struct {
	MatchId      pgtype.Int8 `json:"match_id"`
	Team1Id      pgtype.Int8 `json:"team1_id"`
	Team1Name    pgtype.Text `json:"team1_name"`
	Team2Id      pgtype.Int8 `json:"team2_id"`
	Team2Name    pgtype.Text `json:"team2_name"`
	Season       pgtype.Text `json:"season"`
	CityName     pgtype.Text `json:"city_name"`
	StartDate    pgtype.Date `json:"start_date"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type Overall_Batting_Team_Group struct {
	TeamId       pgtype.Int8 `json:"team_id"`
	TeamName     pgtype.Text `json:"team_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Opposition_Group struct {
	TeamId       pgtype.Int8 `json:"opposition_team_id"`
	TeamName     pgtype.Text `json:"opposition_team_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Ground_Group struct {
	GroundId     pgtype.Int8 `json:"ground_id"`
	GroundName   pgtype.Text `json:"ground_name"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_HostNation_Group struct {
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	PlayersCount   pgtype.Int8 `json:"players_count"`
	MinDate        pgtype.Date `json:"min_date"`
	MaxDate        pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Continent_Group struct {
	ContinentId   pgtype.Int8 `json:"continent_id"`
	ContinentName pgtype.Text `json:"continent_name"`
	PlayersCount  pgtype.Int8 `json:"players_count"`
	MinDate       pgtype.Date `json:"min_date"`
	MaxDate       pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Series_Group struct {
	SeriesId     pgtype.Int8 `json:"series_id"`
	SeriesName   pgtype.Text `json:"series_name"`
	SeriesSeason pgtype.Text `json:"series_season"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Tournament_Group struct {
	TournamentId   pgtype.Int8 `json:"tournament_id"`
	TournamentName pgtype.Text `json:"tournament_name"`
	PlayersCount   pgtype.Int8 `json:"players_count"`
	MinDate        pgtype.Date `json:"min_date"`
	MaxDate        pgtype.Date `json:"max_date"`
	OverallBattingStats
}

type Overall_Batting_Year_Group struct {
	Year         pgtype.Int8 `json:"year"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type Overall_Batting_Season_Group struct {
	Season       pgtype.Text `json:"season"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type Overall_Batting_Decade_Group struct {
	Decade       pgtype.Int8 `json:"decade"`
	PlayersCount pgtype.Int8 `json:"players_count"`
	OverallBattingStats
}

type Overall_Batting_Aggregate_Group struct {
	PlayersCount pgtype.Int8 `json:"players_count"`
	MinDate      pgtype.Date `json:"min_date"`
	MaxDate      pgtype.Date `json:"max_date"`
	OverallBattingStats
}

/* Individual Stats */

type Individual_Batting_Innings_Group struct {
	MatchId   pgtype.Int8 `json:"match_id"`
	StartDate pgtype.Date `json:"start_date"`
	GroundId  pgtype.Int8 `json:"ground_id"`
	CityName  pgtype.Text `json:"city_name"`

	InningsNumber   pgtype.Int8 `json:"innings_number"`
	BatterId        pgtype.Int8 `json:"batter_id"`
	BatterName      pgtype.Text `json:"batter_name"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`
	BowlingTeamId   pgtype.Int8 `json:"bowling_team_id"`
	BowlingTeamName pgtype.Text `json:"bowling_team_name"`

	RunsScored  pgtype.Int8   `json:"runs_scored"`
	BallsFaced  pgtype.Int8   `json:"balls_faced"`
	IsNotOut    pgtype.Bool   `json:"is_not_out"`
	StrikeRate  pgtype.Float8 `json:"strike_rate"`
	FoursScored pgtype.Int8   `json:"fours_scored"`
	SixesScored pgtype.Int8   `json:"sixes_scored"`
}

type Individual_Batting_MatchTotals_Group struct {
	MatchId   pgtype.Int8 `json:"match_id"`
	StartDate pgtype.Date `json:"start_date"`
	GroundId  pgtype.Int8 `json:"ground_id"`
	CityName  pgtype.Text `json:"city_name"`

	BatterId        pgtype.Int8 `json:"batter_id"`
	BatterName      pgtype.Text `json:"batter_name"`
	BattingTeamId   pgtype.Int8 `json:"batting_team_id"`
	BattingTeamName pgtype.Text `json:"batting_team_name"`
	BowlingTeamId   pgtype.Int8 `json:"bowling_team_id"`
	BowlingTeamName pgtype.Text `json:"bowling_team_name"`

	Innings []struct {
		RunsScored pgtype.Int8 `json:"runs_scored"`
		IsNotOut   pgtype.Bool `json:"is_not_out"`
	} `json:"innings"`

	RunsScored  pgtype.Int8   `json:"runs_scored"`
	BallsFaced  pgtype.Int8   `json:"balls_faced"`
	StrikeRate  pgtype.Float8 `json:"strike_rate"`
	FoursScored pgtype.Int8   `json:"fours_scored"`
	SixesScored pgtype.Int8   `json:"sixes_scored"`
}

type Individual_Batting_Series_Group struct {
	SeriesId     pgtype.Int8 `json:"series_id"`
	SeriesName   pgtype.Text `json:"series_name"`
	SeriesSeason pgtype.Text `json:"series_season"`
	Overall_Batting_Batter_Group
}

type Individual_Batting_Tournaments_Group struct {
	TournamentId   pgtype.Int8 `json:"tournament_id"`
	TournamentName pgtype.Text `json:"tournament_name"`
	Overall_Batting_Batter_Group
}

type Individual_Batting_Ground_Group struct {
	GroundId   pgtype.Int8 `json:"ground_id"`
	GroundName pgtype.Text `json:"ground_name"`
	Overall_Batting_Batter_Group
}

type Individual_Batting_HostNation_Group struct {
	HostNationId   pgtype.Int8 `json:"host_nation_id"`
	HostNationName pgtype.Text `json:"host_nation_name"`
	Overall_Batting_Batter_Group
}

type Individual_Batting_Opposition_Group struct {
	OppositionTeamId   pgtype.Int8 `json:"opposition_team_id"`
	OppositionTeamName pgtype.Text `json:"opposition_team_name"`
	Overall_Batting_Batter_Group
}

type Individual_Batting_Year_Group struct {
	Year             pgtype.Int8   `json:"year"`
	BatterId         pgtype.Int8   `json:"batter_id"`
	BatterName       pgtype.Text   `json:"batter_name"`
	TeamsRepresented []pgtype.Text `json:"teams_represented"`
	OverallBattingStats
}

type Individual_Batting_Season_Group struct {
	Season           pgtype.Text   `json:"season"`
	BatterId         pgtype.Int8   `json:"batter_id"`
	BatterName       pgtype.Text   `json:"batter_name"`
	TeamsRepresented []pgtype.Text `json:"teams_represented"`
	OverallBattingStats
}

// Embedded in other structs
type OverallBattingStats struct {
	MatchesPlayed      pgtype.Int8   `json:"matches_played"`
	InningsBatted      pgtype.Int8   `json:"innings_batted"`
	RunsScored         pgtype.Int8   `json:"runs_scored"`
	BallsFaced         pgtype.Int8   `json:"balls_faced"`
	NotOuts            pgtype.Int8   `json:"not_outs"`
	Average            pgtype.Float8 `json:"average"`
	StrikeRate         pgtype.Float8 `json:"strike_rate"`
	HighestScore       pgtype.Int8   `json:"highest_score"`
	HighestNotOutScore pgtype.Int8   `json:"highest_not_out_score"`
	Centuries          pgtype.Int8   `json:"centuries"`
	HalfCenturies      pgtype.Int8   `json:"half_centuries"`
	FiftyPlusScores    pgtype.Int8   `json:"fifty_plus_scores"`
	Ducks              pgtype.Int8   `json:"ducks"`
	FoursScored        pgtype.Int8   `json:"fours_scored"`
	SixesScored        pgtype.Int8   `json:"sixes_scored"`
}
