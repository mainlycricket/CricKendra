package statqueries

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mainlycricket/CricKendra/internal/utils"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

// Function Names are in Query_Overall_Bowling_x format, x represents grouping

func Query_Overall_Bowling_Summary(params *url.Values) (string, []any, error) {
	commonSqlWhere := newSqlWhere(bowling_stats, -1)
	commonSqlWhere.applyFilters(params)
	commonSqlWhere.matchQuery.ensureHostNation()
	commonSqlWhere.matchQuery.ensureContinent()
	commonMatchQuery := commonSqlWhere.matchQuery.prepareQuery()
	commonInningsCondition := commonSqlWhere.inningsFilters.getClause()

	seriesSqlWhere := newSqlWhere(bowling_stats, -1)
	seriesSqlWhere.applyFilters(params)
	seriesSqlWhere.matchQuery.ensureSeries()
	seriesMatchQuery := seriesSqlWhere.matchQuery.prepareQuery()

	tournamentSqlWhere := newSqlWhere(bowling_stats, -1)
	tournamentSqlWhere.applyFilters(params)
	tournamentSqlWhere.matchQuery.ensureTournament()
	tournamentMatchQuery := tournamentSqlWhere.matchQuery.prepareQuery()

	/* Best Bowling Figures Start */

	teamBestBowlingFigures := prepareBestBowlingFigures("team_", "common_matches", []string{"innings.bowling_team_id"}, nil, commonSqlWhere.inningsFilters.conditions)
	oppositionBestBowlingFigures := prepareBestBowlingFigures("opposition_", "common_matches", []string{"innings.batting_team_id"}, nil, commonSqlWhere.inningsFilters.conditions)
	hostNationBestBowlingFigures := prepareBestBowlingFigures("hostnation_", "common_matches", []string{"matches.host_nation_id"}, nil, commonSqlWhere.inningsFilters.conditions)
	continentBestBowlingFigures := prepareBestBowlingFigures("continent_", "common_matches", []string{"matches.continent_id"}, nil, commonSqlWhere.inningsFilters.conditions)
	yearBestBowlingFigures := prepareBestBowlingFigures("year_", "common_matches", []string{"date_part('year', matches.start_date)::integer"}, nil, commonSqlWhere.inningsFilters.conditions)
	seasonBestBowlingFigures := prepareBestBowlingFigures("season_", "common_matches", []string{"matches.season"}, nil, commonSqlWhere.inningsFilters.conditions)

	homeAwayBestBowlingFigures := prepareBestBowlingFigures("home_away_", "common_matches", []string{`
		(CASE 
			WHEN matches.is_neutral_venue THEN 'neutral'
			WHEN innings.bowling_team_id = matches.home_team_id THEN 'home'
			WHEN innings.bowling_team_id = matches.away_team_id THEN 'away'
			ELSE 'unknown' END
		)`}, nil, commonSqlWhere.inningsFilters.conditions)
	tossWonLostBestBowlingFigures := prepareBestBowlingFigures("toss_won_lost_", "common_matches", []string{`
		(CASE 
			WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
			WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
			END
		)`}, nil, commonSqlWhere.inningsFilters.conditions)
	tossDecisionBestBowlingFigures := prepareBestBowlingFigures("toss_decision_", "common_matches", []string{`
		(CASE 
			WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
			WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
			END
		)`, `matches.is_toss_decision_bat`}, nil, commonSqlWhere.inningsFilters.conditions)
	batBowlFirstBestBowlingFigures := prepareBestBowlingFigures("bat_bowl_first_", "common_matches", []string{`
		(CASE 
			WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
				OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
				THEN 'bat'
			WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
				OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
				THEN 'bowl'
			END
		)`}, nil, commonSqlWhere.inningsFilters.conditions)

	inningsNumberBestBowlingFigures := prepareBestBowlingFigures("innings_number_", "common_matches", []string{"innings.innings_number"}, nil, commonSqlWhere.inningsFilters.conditions)

	matchResultBestBowlingFigures := prepareBestBowlingFigures("match_result_", "common_matches", []string{`
		(CASE 
			WHEN matches.final_result = 'tie' THEN 'tied'
			WHEN matches.final_result = 'draw' THEN 'drawn'
			WHEN matches.final_result = 'no result' THEN 'no result'
			WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
			WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
			END
		)`}, nil, commonSqlWhere.inningsFilters.conditions)

	matchResultBatBowlFirstBestBowlingFigures := prepareBestBowlingFigures("match_result_bat_bowl_first_", "common_matches", []string{`
		(CASE 
			WHEN matches.final_result = 'tie' THEN 'tied'
			WHEN matches.final_result = 'draw' THEN 'drawn'
			WHEN matches.final_result = 'no result' THEN 'no result'
			WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
			WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
			END
		)`, `(CASE 
			WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
				OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
				THEN 'bat'
			WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
				OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
				THEN 'bowl'
			END
		)`}, nil, commonSqlWhere.inningsFilters.conditions)

	seriesTeamsCountBestBowlingFigures := prepareBestBowlingFigures("series_teams_count_", "series_matches", []string{"series_teams.teams_count"}, []string{"JOIN series_teams ON matches.series_id = series_teams.series_id"}, commonSqlWhere.inningsFilters.conditions)
	eventMatchNumberBestBowlingFigures := prepareBestBowlingFigures("event_match_number_", "series_matches", []string{"matches.event_match_number"}, []string{"JOIN series_teams ON matches.series_id = series_teams.series_id AND series_teams.teams_count = '2'"}, commonSqlWhere.inningsFilters.conditions)

	tournamentBestBowlingFigures := prepareBestBowlingFigures("tournament_", "tournament_matches", []string{"matches.tournament_id"}, nil, tournamentSqlWhere.inningsFilters.conditions)
	bowlingPositionBestBowlingFigures := prepareBestBowlingFigures("bowling_position_", "common_matches", []string{"bowling_scorecards.bowling_position"}, nil, commonSqlWhere.inningsFilters.conditions)

	/* Best Bowling Figures End */

	query := fmt.Sprintf(`
		WITH common_matches AS (
			%s
		),
		
		series_matches AS (
			%s
		), series_teams AS (
			SELECT series_matches.series_id, 
				(CASE
					WHEN COUNT(DISTINCT series_team_entries.team_id) = 2 THEN '2'
					WHEN COUNT(DISTINCT series_team_entries.team_id) > 4 THEN '5+'
					ELSE '3-4'
				END) AS teams_count
			FROM series_matches
			JOIN series_team_entries ON series_matches.series_id = series_team_entries.series_id
			GROUP BY series_matches.series_id
		),
		
		tournament_matches AS (
			%s
		),
		
		%s,
		teams_summary AS (
			SELECT innings.bowling_team_id AS team_id,
			teams.name AS team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN team_best_bowling_figures bbf ON bbf.group_field1 = innings.bowling_team_id
			LEFT JOIN teams ON innings.bowling_team_id = teams.id
			%s
			GROUP BY innings.bowling_team_id, teams.name, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		opposition_summary AS (
			SELECT innings.batting_team_id AS opposition_team_id,
			teams.name AS opposition_team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN opposition_best_bowling_figures bbf ON bbf.group_field1 = innings.batting_team_id
			LEFT JOIN teams ON innings.batting_team_id = teams.id
			%s
			GROUP BY innings.batting_team_id, teams.name, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),


		%s,
		host_nation_summary AS (
			SELECT matches.host_nation_id, matches.host_nation_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN hostnation_best_bowling_figures bbf ON bbf.group_field1 = matches.host_nation_id
			%s
			GROUP BY matches.host_nation_id, matches.host_nation_name, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		continent_summary AS (
			SELECT matches.continent_id, matches.continent_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN continent_best_bowling_figures bbf ON bbf.group_field1 = matches.continent_id
			%s
			GROUP BY matches.continent_id, matches.continent_name, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		year_summary AS (
			SELECT date_part('year', matches.start_date)::integer AS match_year,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN year_best_bowling_figures bbf ON bbf.group_field1 = date_part('year', matches.start_date)::integer
			%s
			GROUP BY date_part('year', matches.start_date)::integer, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		season_summary AS (
			SELECT matches.season,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN season_best_bowling_figures bbf ON bbf.group_field1 = matches.season
			%s
			GROUP BY matches.season, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		home_away_summary AS (
			SELECT (CASE 
				WHEN matches.is_neutral_venue THEN 'neutral'
				WHEN innings.bowling_team_id = matches.home_team_id THEN 'home'
				WHEN innings.bowling_team_id = matches.away_team_id THEN 'away'
				ELSE 'unknown' END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			LEFT JOIN home_away_best_bowling_figures bbf ON bbf.group_field1 = (CASE 
				WHEN matches.is_neutral_venue THEN 'neutral'
				WHEN innings.bowling_team_id = matches.home_team_id THEN 'home'
				WHEN innings.bowling_team_id = matches.away_team_id THEN 'away'
				ELSE 'unknown' END
			)
			%s
			GROUP BY (CASE 
				WHEN matches.is_neutral_venue THEN 'neutral'
				WHEN innings.bowling_team_id = matches.home_team_id THEN 'home'
				WHEN innings.bowling_team_id = matches.away_team_id THEN 'away'
				ELSE 'unknown' END
			), bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		toss_won_lost_summary AS (
			SELECT (CASE 
				WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			LEFT JOIN toss_won_lost_best_bowling_figures bbf ON bbf.group_field1 = (CASE 
				WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			)
			%s
			GROUP BY (CASE 
				WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			), bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),


		%s,
		toss_decision_summary AS (
			SELECT (CASE 
				WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			), matches.is_toss_decision_bat,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			LEFT JOIN toss_decision_best_bowling_figures bbf ON bbf.group_field1 = (CASE 
				WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			) AND bbf.group_field2 = matches.is_toss_decision_bat
			%s
			GROUP BY (CASE 
				WHEN innings.bowling_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			), matches.is_toss_decision_bat,
			bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		bat_bowl_first_summary AS (
			SELECT (CASE 
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			LEFT JOIN bat_bowl_first_best_bowling_figures bbf ON bbf.group_field1 = (CASE 
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			)
			%s
			GROUP BY (CASE 
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			),
			bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		innings_number_summary AS (
			SELECT innings.innings_number,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN innings_number_best_bowling_figures bbf ON bbf.group_field1 = innings.innings_number
			%s
			GROUP BY innings.innings_number, bbf.ten_wicket_hauls, 
			bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		match_result_summary AS (
			SELECT (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			LEFT JOIN match_result_best_bowling_figures bbf ON bbf.group_field1 = (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
				END
			)
			%s
			GROUP BY (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
				END
			),
			bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		match_result_bat_bowl_first_summary AS (
			SELECT (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
				END
			),
			(CASE 
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			LEFT JOIN match_result_bat_bowl_first_best_bowling_figures bbf ON bbf.group_field1 = (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
				END
			) AND bbf.group_field2 = (CASE 
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			)
			%s
			GROUP BY (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.bowling_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.bowling_team_id = matches.match_loser_team_id THEN 'lost'
				END
			), (CASE 
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.bowling_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.bowling_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			), bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		series_teams_count_summary AS (
			SELECT series_teams.teams_count,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM series_matches matches
			%s
			LEFT JOIN series_teams ON matches.series_id = series_teams.series_id
			LEFT JOIN series_teams_count_best_bowling_figures bbf ON bbf.group_field1 = series_teams.teams_count
			%s
			GROUP BY series_teams.teams_count, bbf.ten_wicket_hauls,
				bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		series_match_number_summary AS (
			SELECT matches.event_match_number,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM series_matches matches
			%s
			JOIN series_teams ON matches.series_id = series_teams.series_id AND series_teams.teams_count = '2'
			LEFT JOIN event_match_number_best_bowling_figures bbf ON bbf.group_field1 = matches.event_match_number
			%s
			GROUP BY matches.event_match_number, bbf.ten_wicket_hauls,
				bbf.best_match_wickets, bbf.best_match_runs
		),

		%s,
		tournament_summary AS (
			SELECT matches.tournament_id, matches.tournament_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
				%s
			FROM tournament_matches matches
			%s
			LEFT JOIN tournament_best_bowling_figures bbf ON bbf.group_field1 = matches.tournament_id
			%s
			GROUP BY matches.tournament_id, matches.tournament_name,
				bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		),
			
		%s,
		bowling_position_summary AS (
			SELECT bowling_scorecards.bowling_position,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN bowling_position_best_bowling_figures bbf ON bbf.group_field1 = bowling_scorecards.bowling_position
			%s
			GROUP BY bowling_scorecards.bowling_position,
				bbf.ten_wicket_hauls, bbf.best_match_wickets, bbf.best_match_runs
		)

		SELECT
			(SELECT ARRAY_AGG(teams_summary.*) FROM teams_summary),
			(SELECT ARRAY_AGG(opposition_summary.*) FROM opposition_summary),
			(SELECT ARRAY_AGG(host_nation_summary.*) FROM host_nation_summary),
			(SELECT ARRAY_AGG(continent_summary.*) FROM continent_summary),
			(SELECT ARRAY_AGG(year_summary.*) FROM year_summary),
			(SELECT ARRAY_AGG(season_summary.*) FROM season_summary),
			(SELECT ARRAY_AGG(home_away_summary.*) FROM home_away_summary),
			(SELECT ARRAY_AGG(toss_won_lost_summary.*) FROM toss_won_lost_summary),
			(SELECT ARRAY_AGG(toss_decision_summary.*) FROM toss_decision_summary),
			(SELECT ARRAY_AGG(bat_bowl_first_summary.*) FROM bat_bowl_first_summary),
			(SELECT ARRAY_AGG(innings_number_summary.*) FROM innings_number_summary),
			(SELECT ARRAY_AGG(match_result_summary.*) FROM match_result_summary),
			(SELECT ARRAY_AGG(match_result_bat_bowl_first_summary.*) FROM match_result_bat_bowl_first_summary),
			(SELECT ARRAY_AGG(series_teams_count_summary.*) FROM series_teams_count_summary),
			(SELECT ARRAY_AGG(series_match_number_summary.*) FROM series_match_number_summary),
			(SELECT ARRAY_AGG(tournament_summary.*) FROM tournament_summary),
			(SELECT ARRAY_AGG(bowling_position_summary.*) FROM bowling_position_summary)
		;
		`, commonMatchQuery, seriesMatchQuery, tournamentMatchQuery,
		teamBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		oppositionBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		hostNationBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		continentBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		yearBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		seasonBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		homeAwayBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		tossWonLostBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		tossDecisionBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		batBowlFirstBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		inningsNumberBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		matchResultBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		matchResultBatBowlFirstBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		seriesTeamsCountBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		eventMatchNumberBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		tournamentBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
		bowlingPositionBestBowlingFigures, bowling_numbers_query, bowling_common_joins, commonInningsCondition,
	)

	return query, commonSqlWhere.args, nil
}

func Query_Overall_Bowling_Bowlers(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_players_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"bowling_scorecards.bowler_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_TeamInnings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_teamInnings_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"innings.id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
	SELECT matches.id AS match_id,
		innings.innings_number,
		innings.bowling_team_id,
		teams1.name AS bowling_team_name,
		innings.batting_team_id,
		teams2.name AS batting_team_name,
		matches.season,
		matches.city_name,
		matches.start_date,
		COUNT(DISTINCT mse.player_id) AS players_count,
		%s
	FROM matches
	%s
	LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = innings.id
	LEFT JOIN teams teams1 ON innings.bowling_team_id = teams1.id
	LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
	%s
	GROUP BY matches.id,
		matches.start_date,
		matches.season,	
		matches.city_name,
		innings.id,
		teams1.name,
		teams2.name,
		bbf.ten_wicket_hauls,
		bbf.best_match_wickets,
		bbf.best_match_runs
	%s
		%s
	%s;
	`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_matches_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
	SELECT matches.id AS match_id,
		matches.team1_id,
		teams1.name AS team1_name,
		matches.team2_id,
		teams2.name AS team2_name,
		matches.season,
		matches.city_name,
		matches.start_date,
		COUNT(DISTINCT mse.player_id) AS players_count,
		%s
	FROM matches
	%s
	LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.id
	LEFT JOIN teams teams1 ON matches.team1_id = teams1.id
	LEFT JOIN teams teams2 ON matches.team2_id = teams2.id
	%s
	GROUP BY matches.id,
		matches.season,
		matches.start_date,
		matches.city_name,
		matches.team1_id,
		matches.team2_id,
		teams1.name,
		teams2.name,
		bbf.ten_wicket_hauls,
		bbf.best_match_wickets,
		bbf.best_match_runs
	%s
		%s
	%s;
	`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Teams(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_teams_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"innings.bowling_team_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT innings.bowling_team_id,
			teams.name AS bowling_team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = innings.bowling_team_id
		LEFT JOIN teams ON innings.bowling_team_id = teams.id
		%s
		GROUP BY innings.bowling_team_id,
			teams.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_oppositions_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"innings.batting_team_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT innings.batting_team_id AS opposition_team_id,
			teams.name AS opposition_team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN teams ON innings.batting_team_id = teams.id
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = innings.batting_team_id
		%s
		GROUP BY innings.batting_team_id,
			teams.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_grounds_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.ground_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.ground_id,
			matches.ground_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.ground_id
		%s
		GROUP BY matches.ground_id,
			matches.ground_name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_hostNations_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.host_nation_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.host_nation_id,
			matches.host_nation_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.host_nation_id
		%s
		GROUP BY matches.host_nation_id,
			matches.host_nation_name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Continents(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_continents_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureContinent()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.continent_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.continent_id,
			matches.continent_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.continent_id
		%s
		GROUP BY matches.continent_id,
			matches.continent_name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_series_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.series_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.series_id,
			matches.series_name,
			matches.series_season,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.series_id
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_tournaments_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.tournament_id"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.tournament_id,
			matches.tournament_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.tournament_id
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_years_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"date_part('year', matches.start_date)::integer"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT date_part('year', matches.start_date)::integer AS match_year,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
			%s
			LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = date_part('year', matches.start_date)::integer
		%s
		GROUP BY date_part('year', matches.start_date)::integer,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_seasons_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"matches.season"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.season AS matches_season,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = matches.season
		%s
		GROUP BY matches.season,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Decades(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_decade_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"10 * date_part('decade', matches.start_date)::integer"}, nil, sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT 10 * date_part('decade', matches.start_date)::integer AS match_decade,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
			%s
			LEFT JOIN best_bowling_figures bbf ON bbf.group_field1 = 10 * date_part('decade', matches.start_date)::integer
		%s
		GROUP BY 10 * date_part('decade', matches.start_date)::integer,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Aggregate(params *url.Values) (string, []any, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_aggregate_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("", "matches", []string{"TRUE"}, nil, sqlWhere.inningsFilters.conditions)

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON TRUE
		%s
		GROUP BY bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause)

	return query, sqlWhere.args, nil
}

// Function Names are in Query_Individual_Bowling_x format, x represents grouping

func Query_Individual_Bowling_Innings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_innings_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureGround()
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.id AS match_id,
			matches.start_date,
			matches.ground_id,
			matches.city_name,
			innings.innings_number,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			innings.batting_team_id AS bowling_team_id,
			teams.short_name as bowling_team_name,
			innings.bowling_team_id AS batting_team_id,
			teams2.name AS batting_team_name,
			bowling_scorecards.balls_bowled / 6 + balls_bowled %% 6 * 0.1 AS overs_bowled,
			bowling_scorecards.maiden_overs,
			bowling_scorecards.runs_conceded,
			bowling_scorecards.wickets_taken,
			(
				CASE
					WHEN bowling_scorecards.balls_bowled > 0 THEN bowling_scorecards.runs_conceded * 6.0 / bowling_scorecards.balls_bowled
					ELSE NULL
				END
			) AS economy,
			bowling_scorecards.fours_conceded,
			bowling_scorecards.sixes_conceded
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			LEFT JOIN teams ON innings.bowling_team_id = teams.id
			LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
			LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		%s
		GROUP BY bowling_scorecards.innings_id,
			bowling_scorecards.bowler_id,
			players.name,
			matches.ground_id,
			matches.city_name,
			innings.batting_team_id,
			innings.bowling_team_id,
			teams.short_name,
			teams2.name,
			innings.innings_number,
			matches.start_date,
			matches.id,
			bowling_scorecards.runs_conceded,
			bowling_scorecards.balls_bowled,
			bowling_scorecards.wickets_taken
		%s
		%s
		%s;
		`, matchQuery, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_MatchTotals(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_matchTotals_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureGround()
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.id AS match_id,
			matches.start_date,
			matches.ground_id,
			matches.city_name,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			innings.batting_team_id AS bowling_team_id,
			teams.short_name as bowling_team_name,
			innings.bowling_team_id AS batting_team_id,
			teams2.name AS batting_team_name,
			%s,
			%s,
			%s,
			%s,
			%s,
			%s,
			%s,
			%s,
			%s
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			LEFT JOIN teams ON innings.bowling_team_id = teams.id
			LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
			LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		%s
		GROUP BY bowling_scorecards.innings_id,
			bowling_scorecards.bowler_id,
			players.name,
			matches.ground_id,
			matches.city_name,
			innings.batting_team_id,
			innings.bowling_team_id,
			teams.short_name,
			teams2.name,
			matches.start_date,
			matches.id
		%s
		%s
		%s;
		`, matchQuery, oversBowled_query, maidenOvers_query, runsConceded_query, wicketsTaken_query, bowlingAverage_query, bowlingEconomy_query, bowlingStrikeRate_query, foursConceded_query, sixesConceded_query, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_series_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("matches.series_id", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.series_id,
			matches.series_name,
			matches.series_season,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.series_id
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_tournaments_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("matches.tournament_id", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.tournament_id,
			matches.tournament_name,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.tournament_id
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_grounds_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("matches.ground_id", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.ground_id,
			matches.ground_name,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.ground_id
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.ground_id,
			matches.ground_name,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_hostNations_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("matches.host_nation_id", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.host_nation_id,
			matches.host_nation_name,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.host_nation_id
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.host_nation_id,
			matches.host_nation_name,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_oppositions_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("innings.batting_team_id", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT innings.batting_team_id AS opposition_team_id,
			teams2.name AS opposition_team_name,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = innings.batting_team_id
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
		%s
		GROUP BY innings.batting_team_id,
			teams2.name,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_years_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("date_part('year', matches.start_date)::integer", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT date_part('year', matches.start_date)::integer AS match_year,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = date_part('year', matches.start_date)
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY date_part('year', matches.start_date)::integer,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, inidividual_seasons_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures_Individual("matches.season", sqlWhere.inningsFilters.conditions)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		),
		%s
		SELECT matches.season,
			bowling_scorecards.bowler_id,
			players.name AS bowler_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.season
			AND bbf.bowler_id = bowling_scorecards.bowler_id
		LEFT JOIN players ON bowling_scorecards.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.season,
			bowling_scorecards.bowler_id,
			players.name,
			bbf.ten_wicket_hauls,
			bbf.best_match_wickets,
			bbf.best_match_runs
		%s
		%s
		%s;
		`, matchQuery, best_bowling_figures, bowling_numbers_query, bowling_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

// HELPERS

func prepareBestBowlingFigures(prefix, from string, groupFields, extraJoins, inningsConditions []string) string {
	condition := prefixJoin(inningsConditions, "WHERE", " AND ")
	extraJoinsStr := strings.Join(extraJoins, "\n")

	nums := utils.NumbersAsStrings(1, len(groupFields))
	groupFieldsAlias := utils.RepeatWord("group_field", len(groupFields))
	groupFieldsAlias = utils.AddPrefixSuffix(groupFieldsAlias, nil, nums)
	groupFieldsStr := strings.Join(groupFieldsAlias, ", ")

	sqlGroupFields := utils.RepeatWord(" AS ", len(groupFields))
	sqlGroupFields = utils.AddPrefixSuffix(sqlGroupFields, nil, groupFieldsAlias)
	sqlGroupFields = utils.AddPrefixSuffix(groupFields, nil, sqlGroupFields)
	sqlGroupFieldsStr := strings.Join(sqlGroupFields, ", ")

	mwGroupFieldsAlias := utils.RepeatWord("mw.", len(groupFields))
	mwGroupFieldsAlias = utils.AddPrefixSuffix(groupFieldsAlias, mwGroupFieldsAlias, nil)
	mwGroupFieldsStr := strings.Join(mwGroupFieldsAlias, ", ")

	bbmGroupFieldsAlias := utils.RepeatWord("bbm.", len(groupFields))
	bbmGroupFieldsAlias = utils.AddPrefixSuffix(groupFieldsAlias, bbmGroupFieldsAlias, nil)
	bbmGroupFieldsStr := strings.Join(bbmGroupFieldsAlias, ", ")

	eqArr := utils.RepeatWord(" = ", len(groupFields))

	twhGroupFieldsAlias := utils.RepeatWord("twh.", len(groupFields))
	twhGroupFieldsAlias = utils.AddPrefixSuffix(groupFieldsAlias, twhGroupFieldsAlias, nil)
	twhJoinConditions := utils.AddPrefixSuffix(eqArr, twhGroupFieldsAlias, bbmGroupFieldsAlias)
	twhJoinConditionsStr := strings.Join(twhJoinConditions, " AND ")

	bbiGroupFieldsAlias := utils.RepeatWord("bbi.", len(groupFields))
	bbiGroupFieldsAlias = utils.AddPrefixSuffix(groupFieldsAlias, bbiGroupFieldsAlias, nil)
	bbiJoinConditions := utils.AddPrefixSuffix(eqArr, bbiGroupFieldsAlias, bbmGroupFieldsAlias)
	bbiJoinConditionsStr := strings.Join(bbiJoinConditions, " AND ")

	query := fmt.Sprintf(`
		%smatch_wickets AS (
			SELECT bowling_scorecards.bowler_id,
				%s,
				matches.id AS match_id,
				SUM(bowling_scorecards.wickets_taken) AS total_wickets,
				SUM(bowling_scorecards.runs_conceded) AS total_runs
			FROM %s matches
			%s
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			%s
			GROUP BY matches.id,
				%s,
				bowling_scorecards.bowler_id
		),
		%sten_wicket_hauls AS (
			SELECT %s,
				mw.bowler_id,
				COUNT(*) AS hauls_count
			FROM %smatch_wickets mw
			WHERE mw.total_wickets >= 10
			GROUP BY %s,
				mw.bowler_id
		),
		%sbest_bowling_match AS (
			SELECT DISTINCT ON (%s)
				%s,
				mw.total_wickets AS wickets,
				mw.total_runs AS runs
			FROM %smatch_wickets mw
			ORDER BY %s,
				mw.total_wickets DESC,
				mw.total_runs ASC
		),
		%sbest_bowling_innings AS (
			SELECT %s,
				MAX(bowling_scorecards.wickets_taken) AS wickets
			FROM %s matches
			%s
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			%s
			GROUP BY %s
		),
		%sbest_bowling_figures AS (
			SELECT %s,
				twh.hauls_count AS ten_wicket_hauls,
				bbm.wickets AS best_match_wickets,
				bbm.runs AS best_match_runs,
				bbi.wickets AS best_innings_wickets
			FROM %sbest_bowling_match bbm
			LEFT JOIN %sten_wicket_hauls twh ON %s
			LEFT JOIN %sbest_bowling_innings bbi ON %s
		)
	`, prefix, sqlGroupFieldsStr, from, extraJoinsStr, condition, groupFieldsStr,
		prefix, mwGroupFieldsStr, prefix, mwGroupFieldsStr,
		prefix, mwGroupFieldsStr, mwGroupFieldsStr, prefix, mwGroupFieldsStr,
		prefix, sqlGroupFieldsStr, from, extraJoinsStr, condition, groupFieldsStr,
		prefix, bbmGroupFieldsStr, prefix, prefix, twhJoinConditionsStr, prefix, bbiJoinConditionsStr,
	)
	return query
}

func prepareBestBowlingFigures_Individual(groupField string, inningsConditions []string) string {
	condition := prefixJoin(inningsConditions, "WHERE", " AND ")

	query := fmt.Sprintf(`
		match_wickets AS (
			SELECT bowling_scorecards.bowler_id,
				%s AS group_field,
				matches.id AS match_id,
				SUM(bowling_scorecards.wickets_taken) AS total_wickets,
				SUM(bowling_scorecards.runs_conceded) AS total_runs
			FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			%s
			GROUP BY matches.id,
				group_field,
				bowling_scorecards.bowler_id
		),
		ten_wicket_hauls AS (
			SELECT mw.group_field,
				mw.bowler_id,
				COUNT(*) AS hauls_count
			FROM match_wickets mw
			WHERE mw.total_wickets >= 10
			GROUP BY mw.group_field,
				mw.bowler_id
		),
		best_bowling_match AS (
			SELECT DISTINCT ON (mw.bowler_id, mw.group_field)
				mw.bowler_id,
				mw.group_field,
				mw.total_wickets AS wickets,
				mw.total_runs AS runs
			FROM match_wickets mw
			ORDER BY mw.group_field,
				mw.bowler_id,
				mw.total_wickets DESC,
				mw.total_runs ASC
		),
		best_bowling_innings AS (
			SELECT %s AS group_field,
				bowling_scorecards.bowler_id,
				MAX(bowling_scorecards.wickets_taken) AS wickets
			FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			%s
			GROUP BY group_field,
				bowling_scorecards.bowler_id
		),
		best_bowling_figures AS (
			SELECT bbm.group_field,
				bbm.bowler_id,
				twh.hauls_count AS ten_wicket_hauls,
				bbm.wickets AS best_match_wickets,
				bbm.runs AS best_match_runs,
				bbi.wickets AS best_innings_wickets
			FROM best_bowling_match bbm
			LEFT JOIN ten_wicket_hauls twh ON twh.group_field = bbm.group_field
        		AND twh.bowler_id = bbm.bowler_id
			LEFT JOIN best_bowling_innings bbi ON bbi.group_field = bbm.group_field
				AND bbi.bowler_id = bbm.bowler_id
		)
	`, groupField, condition, groupField, condition)

	return query
}
