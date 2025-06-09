package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/backend/pkg/pgxutils"
)

// Function Names are in Query_Overall_Batting_x format, x represents grouping

func Query_Overall_Batting_Summary(params *url.Values) (string, []any, error) {
	commonSqlWhere := newSqlWhere(batting_stats, -1)
	commonSqlWhere.applyFilters(params)
	commonSqlWhere.matchQuery.ensureHostNation()
	commonSqlWhere.matchQuery.ensureContinent()
	commonMatchQuery := commonSqlWhere.matchQuery.prepareQuery()
	commonInningsCondition := commonSqlWhere.inningsFilters.getClause()

	seriesSqlWhere := newSqlWhere(batting_stats, -1)
	seriesSqlWhere.applyFilters(params)
	seriesSqlWhere.matchQuery.ensureSeries()
	seriesMatchQuery := seriesSqlWhere.matchQuery.prepareQuery()

	tournamentSqlWhere := newSqlWhere(batting_stats, -1)
	tournamentSqlWhere.applyFilters(params)
	tournamentSqlWhere.matchQuery.ensureTournament()
	tournamentMatchQuery := tournamentSqlWhere.matchQuery.prepareQuery()

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
		
		teams_summary AS (
			SELECT innings.batting_team_id AS team_id,
			teams.name AS team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN teams ON innings.batting_team_id = teams.id
			%s
			GROUP BY innings.batting_team_id,
				teams.name
		),

		opposition_summary AS (
			SELECT innings.bowling_team_id AS team_id,
			teams.name AS team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			LEFT JOIN players ON batting_scorecards.batter_id = players.id
			LEFT JOIN teams ON innings.bowling_team_id  = teams.id
			%s
			GROUP BY innings.bowling_team_id,
				teams.name
		),
		
		host_nations_summary AS (
			SELECT
				matches.host_nation_id,
				matches.host_nation_name,
				COUNT(DISTINCT mse.player_id),
				MIN(matches.start_date),
				MAX(matches.start_date),
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY matches.host_nation_id,
			matches.host_nation_name
		),

		continents_summary AS (
			SELECT matches.continent_id,
			matches.continent_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
			FROM common_matches matches
			%s
			%s
			GROUP BY matches.continent_id,
				matches.continent_name
		),

		years_summary AS (
			SELECT date_part('year', matches.start_date)::int AS match_year,
			COUNT(DISTINCT mse.player_id) AS players_count,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY date_part('year', matches.start_date)::int
		),

		seasons_summary AS (
			SELECT matches.season,
			COUNT(DISTINCT mse.player_id) AS players_count,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY matches.season
		),

		home_away_summary AS (
			SELECT (CASE 
				WHEN matches.is_neutral_venue THEN 'neutral'
				WHEN innings.batting_team_id = matches.home_team_id THEN 'home'
				WHEN innings.batting_team_id = matches.away_team_id THEN 'away'
				ELSE 'unknown' END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY (CASE 
				WHEN matches.is_neutral_venue THEN 'neutral'
				WHEN innings.batting_team_id = matches.home_team_id THEN 'home'
				WHEN innings.batting_team_id = matches.away_team_id THEN 'away'
				ELSE 'unknown' END
			)
		),

		toss_won_lost_summary AS (
			SELECT (CASE 
				WHEN innings.batting_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY (CASE 
				WHEN innings.batting_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			)
		),

		toss_decision_summary AS (
			SELECT (CASE 
				WHEN innings.batting_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			), matches.is_toss_decision_bat,
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY (CASE 
				WHEN innings.batting_team_id = matches.toss_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.toss_loser_team_id THEN 'lost'
				END
			), matches.is_toss_decision_bat
		),

		bat_bowl_first_summary AS (
			SELECT (CASE 
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY (CASE 
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			)
		),

		innings_number_summary AS (
			SELECT innings.innings_number,
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY innings.innings_number
		),

		match_result_summary AS (
			SELECT (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.batting_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.match_loser_team_id THEN 'lost'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.batting_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.match_loser_team_id THEN 'lost'
				END
			)
		),

		match_result_bat_bowl_first_summary AS (
			SELECT (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.batting_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.match_loser_team_id THEN 'lost'
				END
			),
			(CASE 
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			),
			COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY (CASE 
				WHEN matches.final_result = 'tie' THEN 'tied'
				WHEN matches.final_result = 'draw' THEN 'drawn'
				WHEN matches.final_result = 'no result' THEN 'no result'
				WHEN innings.batting_team_id = matches.match_winner_team_id THEN 'won'
				WHEN innings.batting_team_id = matches.match_loser_team_id THEN 'lost'
				END
			),
			(CASE 
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat = FALSE)
					THEN 'bat'
				WHEN (innings.batting_team_id = matches.toss_winner_team_id AND matches.is_toss_decision_bat = FALSE
					OR innings.batting_team_id = matches.toss_loser_team_id AND matches.is_toss_decision_bat)
					THEN 'bowl'
				END
			)
		),

		series_teams_count_summary AS (
			SELECT series_teams.teams_count,
				COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM series_matches matches
			%s
			JOIN series_teams ON series_teams.series_id = matches.series_id
			%s
			GROUP BY series_teams.teams_count
		),

		series_match_number_summary AS (
			SELECT matches.event_match_number,
				COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM series_matches matches
			%s
			JOIN series_teams ON series_teams.series_id = matches.series_id AND series_teams.teams_count = '2'
			%s
			GROUP BY matches.event_match_number
		),

		tournament_summary AS (
			SELECT matches.tournament_id, matches.tournament_name,
				COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM tournament_matches matches
			%s
			%s
			GROUP BY matches.tournament_id, matches.tournament_name
		),

		batting_position_summary AS (
			SELECT batting_scorecards.batting_position,
				COUNT(DISTINCT mse.player_id) AS players_count,
				MIN(matches.start_date) AS min_date,
				MAX(matches.start_date) AS max_date,
				%s
			FROM common_matches matches
			%s
			%s
			GROUP BY batting_scorecards.batting_position
		)

		SELECT
			(SELECT ARRAY_AGG(teams_summary.*) FROM teams_summary),
			(SELECT ARRAY_AGG(opposition_summary.*) FROM opposition_summary),
			(SELECT ARRAY_AGG(host_nations_summary.*) FROM host_nations_summary),
			(SELECT ARRAY_AGG(continents_summary.*) FROM continents_summary),
			(SELECT ARRAY_AGG(years_summary.*) FROM years_summary),
			(SELECT ARRAY_AGG(seasons_summary.*) FROM seasons_summary),
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
			(SELECT ARRAY_AGG(batting_position_summary.*) FROM batting_position_summary)
		;
		`, commonMatchQuery, seriesMatchQuery, tournamentMatchQuery,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
		batting_numbers_query, batting_common_joins, commonInningsCondition,
	)

	return query, commonSqlWhere.args, nil
}

func Query_Overall_Batting_Batters(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_players_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY batting_scorecards.batter_id,
			players.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_TeamInnings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_teamInnings_group)
	sqlWhere.applyFilters(params)
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
			innings.innings_number,
			innings.batting_team_id,
			teams1.name AS batting_team_name,
			innings.bowling_team_id,
			teams2.name AS bowling_team_name,
			matches.season,
			matches.city_name,
			matches.start_date,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
		%s
		LEFT JOIN teams teams1 ON innings.batting_team_id = teams1.id
		LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
		%s
		GROUP BY matches.id,
			matches.start_date,
			matches.season,	
			matches.city_name,
			innings.id,
			teams1.name,
			teams2.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_matches_group)
	sqlWhere.applyFilters(params)
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
		SELECT 
			matches.id AS match_id,
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
			teams2.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Teams(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_teams_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.batting_team_id AS team_id,
			teams.name AS team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN teams ON innings.batting_team_id = teams.id
		%s
		GROUP BY innings.batting_team_id,
			teams.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_oppositions_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.bowling_team_id AS team_id,
			teams.name AS team_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON innings.bowling_team_id  = teams.id
		%s
		GROUP BY innings.bowling_team_id,
			teams.name
		%s
		%s
			%s;
			`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_grounds_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.ground_id,
			matches.ground_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		%s
		GROUP BY matches.ground_id,
			matches.ground_name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_hostNations_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.host_nation_id,
			matches.host_nation_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		%s
		GROUP BY matches.host_nation_id,
			matches.host_nation_name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Continents(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_continents_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureContinent()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.continent_id,
			matches.continent_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		%s
		GROUP BY matches.continent_id,
			matches.continent_name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_series_group)
	sqlWhere.applyFilters(params)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.series_id,
			matches.series_name,
			matches.series_season,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_tournaments_group)
	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.tournament_id,
			matches.tournament_name,
			COUNT(DISTINCT mse.player_id) AS players_count,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_years_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT date_part('year', matches.start_date)::int AS match_year,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
		%s
		%s
		GROUP BY date_part('year', matches.start_date)::int
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_seasons_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.season,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
		%s
		%s
		GROUP BY matches.season
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Decades(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_decade_group)
	sqlWhere.applyFilters(params)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT 10 * date_part('decade', matches.start_date)::int AS match_decade,
			COUNT(DISTINCT mse.player_id) AS players_count,
			%s
		FROM matches
		%s
		%s
		GROUP BY 10 * date_part('decade', matches.start_date)::int
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Aggregate(params *url.Values) (string, []any, error) {
	sqlWhere := newSqlWhere(batting_stats, overall_aggregate_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
		FROM matches
		%s
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause)

	return query, sqlWhere.args, nil
}

// Function Names are in Query_Individual_Batting_x format, x represents grouping

func Query_Individual_Batting_Innings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_innings_group)

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
			batting_scorecards.batter_id,
			players.name AS batter_name,
			innings.batting_team_id,
			teams.short_name as batting_team_name,
			innings.bowling_team_id,
			teams2.name AS bowling_team_name,
			batting_scorecards.runs_scored,
			batting_scorecards.balls_faced,
			(
				CASE
					WHEN batting_scorecards.dismissal_type IS NULL
					OR batting_scorecards.dismissal_type IN ('retired hurt', 'retired not out') THEN TRUE
					ELSE FALSE
				END
			) AS is_not_out,
			(
				CASE
					WHEN batting_scorecards.balls_faced > 0 THEN batting_scorecards.runs_scored * 100.0 / batting_scorecards.balls_faced
					ELSE 0
				END
			) AS strike_rate,
			batting_scorecards.fours_scored,
			batting_scorecards.sixes_scored
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN batting_scorecards ON batting_scorecards.innings_id = innings.id
			LEFT JOIN teams ON innings.batting_team_id = teams.id
			LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
			LEFT JOIN players ON batting_scorecards.batter_id = players.id
		%s
		GROUP BY batting_scorecards.innings_id,
			batting_scorecards.batter_id,
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
			batting_scorecards.runs_scored,
			batting_scorecards.balls_faced,
			batting_scorecards.dismissal_type
		%s
		%s
		%s;
		`, matchQuery, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_MatchTotals(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_matchTotals_group)

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
			batting_scorecards.batter_id,
			players.name AS batter_name,
			innings.batting_team_id,
			teams.short_name as batting_team_name,
			innings.bowling_team_id,
			teams2.name AS bowling_team_name,
			ARRAY_AGG(ROW(
				batting_scorecards.runs_scored, (
				CASE
					WHEN batting_scorecards.dismissal_type IS NULL
					OR batting_scorecards.dismissal_type IN ('retired hurt', 'retired not out') THEN TRUE
					ELSE FALSE
				END
			))),
			%s,
			%s,
			%s,
			%s,
			%s
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN batting_scorecards ON batting_scorecards.innings_id = innings.id
			LEFT JOIN teams ON innings.batting_team_id = teams.id
			LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
			LEFT JOIN players ON batting_scorecards.batter_id = players.id
		%s
		GROUP BY batting_scorecards.innings_id,
			batting_scorecards.batter_id,
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
		`, matchQuery, batting_runs_scored, batting_balls_faced, batting_strike_rate, batting_fours_scored, batting_sixes_scored, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_series_group)

	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureSeries()

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.series_id,
			matches.series_name, 
			matches.series_season,
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season,
			batting_scorecards.batter_id,
			players.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_tournaments_group)

	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureTournament()

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.tournament_id,
			matches.tournament_name,
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name,
			batting_scorecards.batter_id,
			players.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_grounds_group)

	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureGround()

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.ground_id,
			matches.ground_name, 
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.ground_id,
			matches.ground_name,
			batting_scorecards.batter_id,
			players.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_hostNations_group)

	sqlWhere.applyFilters(params)
	sqlWhere.matchQuery.ensureHostNation()

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.host_nation_id,
			matches.host_nation_name,
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY batting_scorecards.batter_id,
			players.name,
			matches.host_nation_id,
			matches.host_nation_name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_oppositions_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.bowling_team_id AS opposition_team_id,
			teams2.name AS opposition_team_name,
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
		%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
		%s
		GROUP BY batting_scorecards.batter_id,
			players.name,
			innings.bowling_team_id,
			teams2.name
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_years_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT date_part('year', matches.start_date)::integer AS match_year,
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY batting_scorecards.batter_id,
			players.name,
			date_part('year', matches.start_date)::integer
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(batting_stats, inidividual_seasons_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.season,
			batting_scorecards.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			%s
		FROM matches
		%s
		LEFT JOIN players ON batting_scorecards.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY batting_scorecards.batter_id,
			players.name,
			matches.season
		%s
		%s
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, qualificationClause, sqlWhere.sortingClause, pagination)

	return query, sqlWhere.args, limit, nil
}
