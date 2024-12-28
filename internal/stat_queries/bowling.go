package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

const bowling_numbers_query string = `
	COUNT(DISTINCT matches.id) AS matches_played,
	COUNT(innings.id) AS innings_count,
	SUM(bs.balls_bowled) / 6 + (SUM(balls_bowled) % 6) * 0.1 AS overs_bowled,
	SUM(bs.runs_conceded) AS runs_conceded,
	SUM(bs.wickets_taken) AS wickets_taken,
	(
		CASE
			WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.runs_conceded) * 1.0 / SUM(bs.wickets_taken)
			ELSE NULL
		END
	) AS average,
	(
		CASE
			WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.balls_bowled) * 1.0 / SUM(bs.wickets_taken)
			ELSE NULL
		END
	) AS strike_rate,
	(
		CASE
			WHEN SUM(bs.balls_bowled) > 0 THEN SUM(bs.runs_conceded) * 6.0 / SUM(bs.balls_bowled)
			ELSE NULL
		END
	) AS economy,
	COUNT(
		CASE
			WHEN bs.wickets_taken = 4 THEN 1
		END
	) AS four_wkt_hauls,
	COUNT(
		CASE
			WHEN bs.wickets_taken >= 5 THEN 1
		END
	) AS five_wkt_hauls,
	MAX(bs.wickets_taken) AS best_innings_wickets,
	MIN(
		CASE
			WHEN bs.wickets_taken = bi.max_wickets THEN bs.runs_conceded
		END
	) AS best_innings_runs,
	SUM(bs.fours_conceded) AS fours_conceded,
	SUM(bs.sixes_conceded) AS sixes_conceded
`

const bowling_common_joins string = `
	LEFT JOIN innings ON innings.match_id = matches.id
	LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
	LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.bowling_team_id
		AND mse.player_id = bs.bowler_id
		AND mse.playing_status IN ('playing_xi')
`

// Function Names are in Query_Overall_Bowling_x format, x represents grouping

func Query_Overall_Bowling_Bowlers(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT bs.bowler_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY bs.bowler_id
	)
	SELECT bs.bowler_id,
		players.name AS bowler_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.bowler_id = bs.bowler_id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY bs.bowler_id,
		players.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_TeamInnings(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT innings.id AS innings_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY innings.id
	)
	SELECT matches.id AS match_id,
		innings.innings_number,
		innings.bowling_team_id,
		teams1.name AS bowling_team_name,
		innings.batting_team_id,
		teams2.name AS batting_team_name,
		matches.season,
		cities.name AS city_name,
		matches.start_date,
		COUNT(DISTINCT mse.player_id) AS players_count,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.innings_id = innings.id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN cities ON grounds.city_id = cities.id
		LEFT JOIN teams teams1 ON innings.bowling_team_id = teams1.id
		LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY matches.id,
		innings.id,
		teams1.name,
		teams2.name,
		cities.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT matches.id AS match_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY matches.id
	)
	SELECT matches.id AS match_id,
		matches.team1_id,
		teams1.name AS team1_name,
		matches.team2_id,
		teams2.name AS team2_name,
		matches.season,
		cities.name AS city_name,
		matches.start_date,
		COUNT(DISTINCT mse.player_id) AS players_count,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.match_id = matches.id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN cities ON grounds.city_id = cities.id
		LEFT JOIN teams teams1 ON matches.team1_id = teams1.id
		LEFT JOIN teams teams2 ON matches.team2_id = teams2.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY matches.id,
		teams1.name,
		teams2.name,
		cities.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Teams(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT innings.bowling_team_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY innings.bowling_team_id
	)
	SELECT innings.bowling_team_id,
		teams.name AS bowling_team_name,
		COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.bowling_team_id = innings.bowling_team_id
		LEFT JOIN teams ON innings.bowling_team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY innings.bowling_team_id,
		teams.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT innings.batting_team_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY innings.batting_team_id
	)
	SELECT innings.batting_team_id AS opposition_team_id,
		teams.name AS opposition_team_name,
		COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.batting_team_id = innings.batting_team_id
		LEFT JOIN teams ON innings.batting_team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY innings.batting_team_id,
		teams.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT matches.ground_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY matches.ground_id
	)
	SELECT matches.ground_id,
		grounds.name AS ground_name,
		COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.ground_id = matches.ground_id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY matches.ground_id,
		grounds.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
	    SELECT cities.host_nation_id,
	        MAX(bs.wickets_taken) AS max_wickets
	    FROM matches
	        LEFT JOIN innings ON innings.match_id = matches.id
	        LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
	        LEFT JOIN grounds ON matches.ground_id = grounds.id
	        LEFT JOIN cities ON grounds.city_id = cities.id
	    WHERE innings.is_super_over = FALSE
	        %s
	    GROUP BY cities.host_nation_id
	)
	SELECT host_nations.id AS host_nation_id,
	    host_nations.name AS host_nation_name,
	    COUNT(DISTINCT mse.player_id) AS players_count,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    %s
	FROM matches
		%s
	    LEFT JOIN grounds ON matches.ground_id = grounds.id
	    LEFT JOIN cities ON grounds.city_id = cities.id
	    LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
	    LEFT JOIN best_innings bi ON bi.host_nation_id = host_nations.id
	WHERE innings.is_super_over = FALSE
	    %s
	GROUP BY host_nations.id,
	    host_nations.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Continents(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT host_nations.continent_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
			LEFT JOIN grounds ON matches.ground_id = grounds.id
			LEFT JOIN cities ON grounds.city_id = cities.id
			LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY host_nations.continent_id
	)
	SELECT continents.id AS continents_id,
		continents.name AS continents_name,
		COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN cities ON grounds.city_id = cities.id
		LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
		LEFT JOIN continents ON host_nations.continent_id = continents.id
		LEFT JOIN best_innings bi ON bi.continent_id = continents.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY continents.id,
		continents.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT date_part('year', matches.start_date)::integer AS match_year,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY date_part('year', matches.start_date)::integer
	)
	SELECT date_part('year', matches.start_date)::integer AS match_year,
		COUNT(DISTINCT mse.player_id) AS players_count,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.match_year = date_part('year', matches.start_date)
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY date_part('year', matches.start_date)::integer
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT matches.season,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY matches.season
	)
	SELECT matches.season,
		COUNT(DISTINCT mse.player_id) AS players_count,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.season = matches.season
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY matches.season
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Bowling_Aggregate(params *url.Values) (string, []any, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
	)
	SELECT COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON TRUE
	WHERE innings.is_super_over = FALSE
		%s
	ORDER BY SUM(bs.wickets_taken) DESC;
	`, condition, bowling_numbers_query, bowling_common_joins, condition)

	return query, sqlWhere.args, nil
}

// Function Names are in Query_Individual_Bowling_x format, x represents grouping

func Query_Individual_Bowling_Innings(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT matches.id AS match_id,
		matches.start_date,
		matches.ground_id,
		cities.name AS city_name,
		innings.innings_number,
		bs.bowler_id,
		players.name AS bowler_name,
		innings.batting_team_id AS bowling_team_id,
		teams.short_name as bowling_team_name,
		innings.bowling_team_id AS batting_team_id,
		teams2.name AS batting_team_name,
		SUM(bs.balls_bowled) / 6 + (SUM(balls_bowled) %% 6) * 0.1 AS overs_bowled,
		SUM(bs.runs_conceded) AS runs_conceded,
		SUM(bs.wickets_taken) AS wickets_taken,
		(
			CASE
				WHEN SUM(bs.balls_bowled) > 0 THEN SUM(bs.runs_conceded) * 6.0 / SUM(bs.balls_bowled)
				ELSE NULL
			END
		) AS economy,
		SUM(bs.fours_conceded) AS fours_conceded,
		SUM(bs.sixes_conceded) AS sixes_conceded
	FROM bowling_scorecards bs
		LEFT JOIN innings ON innings.id = bs.innings_id
		LEFT JOIN teams ON innings.bowling_team_id = teams.id
		LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
		LEFT JOIN matches ON innings.match_id = matches.id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN cities ON grounds.city_id = cities.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY bs.innings_id,
		bs.bowler_id,
		players.name,
		matches.ground_id,
		cities.name,
		innings.batting_team_id,
		innings.bowling_team_id,
		teams.short_name,
		teams2.name,
		innings.innings_number,
		matches.start_date,
		matches.id,
		bs.runs_conceded,
		bs.balls_bowled,
		bs.wickets_taken
	ORDER BY bs.wickets_taken DESC
	%s;
	`, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT bs.bowler_id,
			matches.ground_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY bs.bowler_id,
			matches.ground_id
	)
	SELECT matches.ground_id,
		grounds.name AS ground_name,
		bs.bowler_id,
		players.name AS bowler_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.ground_id = matches.ground_id
			AND bi.bowler_id = bs.bowler_id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY matches.ground_id,
		grounds.name,
		bs.bowler_id,
		players.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT bs.bowler_id,
			cities.host_nation_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
			LEFT JOIN grounds ON matches.ground_id = grounds.id
			LEFT JOIN cities ON grounds.city_id = cities.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY bs.bowler_id,
			cities.host_nation_id
	)
	SELECT cities.host_nation_id,
		host_nations.name AS host_nation_name,
		bs.bowler_id,
		players.name AS bowler_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN cities ON grounds.city_id = cities.id
		LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
		LEFT JOIN best_innings bi ON bi.host_nation_id = cities.host_nation_id
			AND bi.bowler_id = bs.bowler_id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY cities.host_nation_id,
		host_nations.name,
		bs.bowler_id,
		players.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT bs.bowler_id,
			innings.batting_team_id,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY bs.bowler_id,
			innings.batting_team_id
	)
	SELECT innings.batting_team_id AS opposition_team_id,
		teams2.name AS opposition_team_name,
		bs.bowler_id,
		players.name AS bowler_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.batting_team_id = innings.batting_team_id
			AND bi.bowler_id = bs.bowler_id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY innings.batting_team_id,
		teams2.name,
		bs.bowler_id,
		players.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT bs.bowler_id,
			date_part('year', matches.start_date)::integer AS match_year,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY bs.bowler_id,
			date_part('year', matches.start_date)::integer
	)
	SELECT date_part('year', matches.start_date)::integer AS match_year,
		bs.bowler_id,
		players.name AS bowler_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.match_year = date_part('year', matches.start_date)::integer
			AND bi.bowler_id = bs.bowler_id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY date_part('year', matches.start_date)::integer,
		bs.bowler_id,
		players.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Bowling_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, bowling_stats)
	condition := sqlWhere.getConditionString("AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH best_innings AS (
		SELECT bs.bowler_id,
			matches.season,
			MAX(bs.wickets_taken) AS max_wickets
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE innings.is_super_over = FALSE
			%s
		GROUP BY bs.bowler_id,
			matches.season
	)
	SELECT matches.season,
		bs.bowler_id,
		players.name AS bowler_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		%s
	FROM matches
		%s
		LEFT JOIN best_innings bi ON bi.season = matches.season
			AND bi.bowler_id = bs.bowler_id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN players ON bs.bowler_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY matches.season,
		bs.bowler_id,
		players.name
	ORDER BY SUM(bs.wickets_taken) DESC
	%s;
	`, condition, bowling_numbers_query, bowling_common_joins, condition, pagination)

	return query, sqlWhere.args, limit, nil
}
