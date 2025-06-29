package statqueries

import "fmt"

const (
	batting_stats = iota
	bowling_stats
	team_stats
)

const (
	overall_players_group = iota
	overall_teamInnings_group
	overall_matches_group
	overall_teams_group
	overall_oppositions_group
	overall_grounds_group
	overall_hostNations_group
	overall_continents_group
	overall_series_group
	overall_tournaments_group
	overall_years_group
	overall_seasons_group
	overall_decade_group
	overall_aggregate_group

	inidividual_innings_group
	inidividual_matchTotals_group
	inidividual_matchResults_group
	inidividual_series_group
	inidividual_tournaments_group
	inidividual_grounds_group
	inidividual_hostNations_group
	inidividual_oppositions_group
	inidividual_years_group
	inidividual_seasons_group
)

const matches_count_query = `COUNT(DISTINCT matches.id)`

/* BATTING */

const (
	batting_innings_count  string = `COUNT(CASE WHEN batting_scorecards.has_batted THEN innings.id END)`
	batting_runs_scored    string = `SUM(batting_scorecards.runs_scored)`
	batting_balls_faced    string = `SUM(batting_scorecards.balls_faced)`
	batting_not_outs       string = `COUNT(CASE WHEN batting_scorecards.has_batted AND (batting_scorecards.dismissal_type IS NULL OR batting_scorecards.dismissal_type IN ('retired hurt', 'retired not out')) THEN 1 END)`
	batting_highest_score  string = `MAX(batting_scorecards.runs_scored)`
	batting_hs_notout      string = `MAX(CASE WHEN batting_scorecards.dismissal_type IS NULL OR batting_scorecards.dismissal_type IN ('retired hurt', 'retired not out') THEN batting_scorecards.runs_scored ELSE 0 END)`
	batting_average        string = `(CASE WHEN COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) > 0 THEN SUM(batting_scorecards.runs_scored) * 1.0 / COUNT(CASE WHEN batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END) ELSE NULL END)`
	batting_strike_rate    string = `(CASE WHEN SUM(batting_scorecards.balls_faced) > 0 THEN SUM(batting_scorecards.runs_scored) * 100.0 / SUM(batting_scorecards.balls_faced) ELSE 0 END)`
	batting_centuries      string = `COUNT(CASE WHEN batting_scorecards.runs_scored >= 100 THEN 1 END)`
	batting_half_centuries string = `COUNT(CASE WHEN batting_scorecards.runs_scored BETWEEN 50 AND 99 THEN 1 END)`
	batting_fifty_plus     string = `COUNT(CASE WHEN batting_scorecards.runs_scored >= 50 THEN 1 END)`
	batting_ducks          string = `COUNT(CASE WHEN batting_scorecards.has_batted AND batting_scorecards.runs_scored = 0 AND batting_scorecards.dismissal_type IS NOT NULL AND batting_scorecards.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1 END)`
	batting_fours_scored   string = `SUM(batting_scorecards.fours_scored)`
	batting_sixes_scored   string = `SUM(batting_scorecards.sixes_scored)`
)

const batting_numbers_query string = matches_count_query + ", " +
	batting_innings_count + ", " +
	batting_runs_scored + " AS runs_scored, " +
	batting_balls_faced + ", " +
	batting_not_outs + ", " +
	batting_average + ", " +
	batting_strike_rate + ", " +
	batting_highest_score + ", " +
	batting_hs_notout + ", " +
	batting_centuries + ", " +
	batting_half_centuries + ", " +
	batting_fifty_plus + ", " +
	batting_ducks + ", " +
	batting_fours_scored + ", " +
	batting_sixes_scored

const batting_common_joins string = `
	LEFT JOIN innings ON innings.match_id = matches.id
		AND innings.innings_number IS NOT NULL
		AND innings.is_super_over = FALSE
	LEFT JOIN batting_scorecards ON batting_scorecards.innings_id = innings.id
	LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = batting_scorecards.batter_id
		AND mse.playing_status IN ('playing_xi')
`

/* BOWLING */

const (
	inningsBowled_query     string = `COUNT(CASE WHEN bowling_scorecards.bowling_position IS NOT NULL THEN innings.id END)`
	oversBowled_query       string = `SUM(bowling_scorecards.balls_bowled) / 6 + (SUM(balls_bowled) % 6) * 0.1`
	maidenOvers_query       string = `SUM(bowling_scorecards.maiden_overs)`
	runsConceded_query      string = `SUM(bowling_scorecards.runs_conceded)`
	wicketsTaken_query      string = `SUM(bowling_scorecards.wickets_taken)`
	bowlingAverage_query    string = `(CASE WHEN SUM(bowling_scorecards.wickets_taken) > 0 THEN SUM(bowling_scorecards.runs_conceded) * 1.0 / SUM(bowling_scorecards.wickets_taken) ELSE NULL END)`
	bowlingStrikeRate_query string = `(CASE WHEN SUM(bowling_scorecards.wickets_taken) > 0 THEN SUM(bowling_scorecards.balls_bowled) * 1.0 / SUM(bowling_scorecards.wickets_taken) ELSE NULL END)`
	bowlingEconomy_query    string = `(CASE WHEN SUM(bowling_scorecards.balls_bowled) > 0 THEN SUM(bowling_scorecards.runs_conceded) * 6.0 / SUM(bowling_scorecards.balls_bowled) ELSE NULL END)`
	fourWktHauls_query      string = `COUNT(CASE WHEN bowling_scorecards.wickets_taken = 4 THEN 1 END)`
	fiveWktHauls_query      string = `COUNT(CASE WHEN bowling_scorecards.wickets_taken >= 5 THEN 1  END)`
	tenWktHauls_query       string = `bbf.ten_wicket_hauls`
	bestMatchWkts_query     string = `bbf.best_match_wickets`
	bestMatchRuns_query     string = `bbf.best_match_runs`
	bestInningsWkts_query   string = `MAX(bowling_scorecards.wickets_taken)`
	bestInningsRuns_query   string = `MIN(CASE WHEN bowling_scorecards.wickets_taken = bbf.best_innings_wickets THEN bowling_scorecards.runs_conceded END)`
	foursConceded_query     string = `SUM(bowling_scorecards.fours_conceded)`
	sixesConceded_query     string = `SUM(bowling_scorecards.sixes_conceded)`
)

const bowling_numbers_query string = matches_count_query + "," +
	inningsBowled_query + "," +
	oversBowled_query + "," +
	maidenOvers_query + "," +
	runsConceded_query + "," +
	wicketsTaken_query + "," +
	bowlingAverage_query + "," +
	bowlingStrikeRate_query + "," +
	bowlingEconomy_query + "," +
	fourWktHauls_query + "," +
	fiveWktHauls_query + "," +
	tenWktHauls_query + "," +
	bestMatchWkts_query + "," +
	bestMatchRuns_query + "," +
	bestInningsWkts_query + "," +
	bestInningsRuns_query + "," +
	foursConceded_query + "," +
	sixesConceded_query

const bowling_common_joins string = `
	LEFT JOIN innings ON innings.match_id = matches.id
		AND innings.innings_number IS NOT NULL
		AND innings.is_super_over = FALSE
	LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
	LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.bowling_team_id
		AND mse.player_id = bowling_scorecards.bowler_id
		AND mse.playing_status IN ('playing_xi')
`

/* TEAM */

const (
	matchesDrawn_query     string = `COUNT(CASE WHEN matches.final_result = 'draw' THEN 1  END)`
	matchesTied_query      string = `COUNT(CASE WHEN matches.final_result = 'tie' THEN 1 END)`
	matchesNoResult_query  string = `COUNT(CASE WHEN matches.final_result = 'no result' THEN 1 END)`
	teamInningsCount_query string = `COUNT(CASE WHEN innings.innings_number > 0 THEN innings.id END)`
	teamTotalRuns_query    string = `SUM(innings.total_runs)`
	teamTotalBalls_query   string = `SUM(innings.total_balls)`
	teamTotalWkts_query    string = `SUM(innings.total_wickets)`
	teamAverage_query      string = `(CASE WHEN SUM(innings.total_wickets) > 0 THEN SUM (innings.total_runs) * 1.0 / SUM(innings.total_wickets) END)`
	teamScoringRate_query  string = `(CASE WHEN SUM(innings.total_balls) > 0 THEN SUM (innings.total_runs) * 6.0 / SUM(innings.total_balls) END)`
	teamHighestScore_query string = `MAX(innings.total_runs)`
	team_LowestScore_query string = `MIN(CASE WHEN innings.innings_end = 'all_out' THEN innings.total_runs END)`
)

func getTeamNumbersQuery(teamField string) string {
	matchesWon_query := getMatchesWonQuery(teamField)
	matchesLost_query := getMatchesLostQuery(teamField)
	winLossRatio_query := getWinLossRatioNullQuery(teamField)

	query := matches_count_query + "," +
		matchesWon_query + " AS matches_won, " +
		matchesLost_query + " AS matches_lost, " +
		winLossRatio_query + " AS win_loss_ratio, " +
		matchesDrawn_query + "," +
		matchesTied_query + "," +
		matchesNoResult_query + "," +
		teamInningsCount_query + "," +
		teamTotalRuns_query + "," +
		teamTotalBalls_query + "," +
		teamTotalWkts_query + "," +
		teamAverage_query + "," +
		teamScoringRate_query + "," +
		teamHighestScore_query + "," +
		team_LowestScore_query

	return query
}

func getMatchesWonQuery(teamField string) string {
	return fmt.Sprintf(`COUNT(CASE WHEN %s = matches.match_winner_team_id THEN 1 END)`, teamField)
}

func getMatchesLostQuery(teamField string) string {
	return fmt.Sprintf(`COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END)`, teamField)
}

func getWinLossRatioNullQuery(teamField string) string {
	return fmt.Sprintf(`(CASE
				WHEN COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END) > 0
				THEN 
					COUNT(CASE WHEN %s = matches.match_winner_team_id THEN 1 END) * 1.0 / COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END)
				ELSE NULL
				END)`, teamField, teamField, teamField)
}

func getWinLossRatioInfiniteQuery(teamField string) string {
	return fmt.Sprintf(`(CASE
				WHEN COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END) > 0
				THEN 
					COUNT(CASE WHEN %s = matches.match_winner_team_id THEN 1 END) * 1.0 / COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END)
				ELSE '+infinity'
				END)`, teamField, teamField, teamField)
}
