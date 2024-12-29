package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

const team_numbers_query string = `
	COUNT(DISTINCT matches.id) AS matches_played,
	SUM(
	    CASE
	        WHEN match_teams.team_id = matches.match_winner_team_id THEN 1
	        ELSE 0
	    END
	) AS matches_won,
	SUM(
	    CASE
	        WHEN match_teams.team_id = matches.match_loser_team_id THEN 1
	        ELSE 0
	    END
	) AS matches_lost,
	(
	    CASE
	        WHEN SUM(
	            CASE
	                WHEN match_teams.team_id = matches.match_loser_team_id THEN 1
	                ELSE 0
	            END
	        ) > 0 THEN SUM(
	            CASE
	                WHEN match_teams.team_id = matches.match_winner_team_id THEN 1
	                ELSE 0
	            END
	        ) * 1.0 / SUM(
	            CASE
	                WHEN match_teams.team_id = matches.match_loser_team_id THEN 1
	                ELSE 0
	            END
	        )
	    END
	) AS win_loss_ratio,
	COUNT(
	    CASE
	        WHEN matches.final_result = 'drawn' THEN 1
	    END
	) AS matches_drawn,
	COUNT(
	    CASE
	        WHEN matches.final_result = 'tied' THEN 1
	    END
	) AS matches_tied,
	COUNT(
	    CASE
	        WHEN matches.final_result = 'no_result' THEN 1
	    END
	) AS matches_no_result,
	COUNT(innings.id) AS innings_count,
	SUM(innings.total_runs) AS total_runs,
	SUM(innings.total_balls) AS total_balls,
	SUM(innings.total_wickets) AS total_wickets,
	(
	    CASE
	        WHEN SUM(innings.total_wickets) > 0 THEN SUM(innings.total_runs) * 1.0 / SUM(innings.total_wickets)
	    END
	) AS average,
	(
	    CASE
	        WHEN SUM(innings.total_balls) > 0 THEN SUM(innings.total_runs) * 6.0 / SUM(innings.total_balls)
	    END
	) AS scoring_rate,
	MAX(innings.total_runs) AS highest_score,
	MIN(innings.total_runs) AS lowest_score
`

const team_common_joins string = `
	LEFT JOIN matches ON innings.match_id = matches.id
	LEFT JOIN teams ON match_teams.team_id = teams.id
`

// Function Names are in Query_Overall_Team_x format, x represents grouping

func Query_Overall_Team_Teams(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT match_teams.team_id,
	    teams.name AS team_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY match_teams.team_id,
    	teams.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Players(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT mse.player_id,
	    players.name AS player_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    COUNT(DISTINCT mse.team_id) AS teams_count,
	    COUNT(DISTINCT matches.id) AS matches_played,
	    COUNT(
	        DISTINCT CASE
	            WHEN matches.match_winner_team_id = mse.team_id THEN matches.id
	        END
	    ) AS matches_won,
	    COUNT(
	        DISTINCT CASE
	            WHEN matches.match_loser_team_id = mse.team_id THEN matches.id
	        END
	    ) AS matches_lost,
	    (
	        CASE
	            WHEN COUNT(
	                DISTINCT CASE
	                    WHEN matches.match_loser_team_id = mse.team_id THEN matches.id
	                END
	            ) > 0 THEN COUNT(
	                DISTINCT CASE
	                    WHEN matches.match_winner_team_id = mse.team_id THEN matches.id
	                END
	            ) * 1.0 / COUNT(
	                DISTINCT CASE
	                    WHEN matches.match_loser_team_id = mse.team_id THEN matches.id
	                END
	            )
	        END
	    ) AS win_loss_ratio,
		COUNT(
		    CASE
		        WHEN matches.final_result = 'drawn' THEN 1
		    END
		) AS matches_drawn,
		COUNT(
		    CASE
		        WHEN matches.final_result = 'tied' THEN 1
		    END
		) AS matches_tied,
		COUNT(
		    CASE
		        WHEN matches.final_result = 'no_result' THEN 1
		    END
		) AS matches_no_result,
	    COUNT(innings.id) AS innings_count,
	    SUM(innings.total_runs) AS total_runs,
	    SUM(innings.total_balls) AS total_balls,
	    SUM(innings.total_wickets) AS total_wickets,
	    (
	        CASE
	            WHEN SUM(innings.total_wickets) > 0 THEN SUM(innings.total_runs) * 1.0 / SUM(innings.total_wickets)
	        END
	    ) AS average,
	    (
	        CASE
	            WHEN SUM(innings.total_balls) > 0 THEN SUM(innings.total_runs) * 6.0 / SUM(innings.total_balls)
	        END
	    ) AS scoring_rate,
    MAX(innings.total_runs) AS highest_score,
    MIN(innings.total_runs) AS lowest_score
	FROM match_squad_entries mse
	    LEFT JOIN matches ON matches.id = mse.match_id
	    LEFT JOIN innings ON innings.match_id = matches.id
	    	AND innings.%s = mse.team_id
	    LEFT JOIN players ON mse.player_id = players.id
	WHERE innings.is_super_over = FALSE
		%s
	GROUP BY mse.player_id,
	    players.name
	ORDER BY matches_won DESC
    %s;
	`, team_total_for, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT matches.id AS match_id,
	    matches.team1_id,
	    team1.name AS team1_name,
	    matches.team2_id,
	    team2.name AS team2_name,
	    cities.name AS city_name,
	    matches.season,
	    matches.start_date,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
        LEFT JOIN teams team1 ON matches.team1_id = team1.id
        LEFT JOIN teams team2 ON matches.team2_id = team2.id
        LEFT JOIN grounds ON matches.ground_id = grounds.id
        LEFT JOIN cities ON cities.id = grounds.city_id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY matches.id,
        team1.name,
        team2.name,
        cities.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT matches.ground_id AS ground_id,
	    grounds.name AS ground_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
        LEFT JOIN grounds ON matches.ground_id = grounds.id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY matches.ground_id,
        grounds.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT cities.host_nation_id,
	    host_nations.name AS host_nation_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
        LEFT JOIN grounds ON matches.ground_id = grounds.id
        LEFT JOIN cities ON grounds.city_id = cities.id
        LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY cities.host_nation_id,
        host_nations.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Continents(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT host_nations.continent_id,
	    continents.name AS continent_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
        LEFT JOIN grounds ON matches.ground_id = grounds.id
        LEFT JOIN cities ON grounds.city_id = cities.id
        LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
        LEFT JOIN continents ON host_nations.continent_id = continents.id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY host_nations.continent_id,
        continents.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT date_part('year', matches.start_date)::integer AS match_year,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY date_part('year', matches.start_date)::integer
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT matches.season,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY matches.season
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Decades(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT date_part('decade', matches.start_date)::integer * 10 AS decade,
	    COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY date_part('decade', matches.start_date)::integer * 10
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Aggregate(params *url.Values) (string, []any, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT COUNT(DISTINCT match_teams.team_id) AS teams_count,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    ORDER BY matches_won DESC;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition)

	return query, sqlWhere.args, nil
}

// Function Names are in Query_Individual_Team_x format, x represents grouping

func Query_Individual_Team_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	condition := sqlWhere.getConditionString("WHERE ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	team1_matches_query := getIndividualMatchesQuery(condition, false)
	team2_matches_query := getIndividualMatchesQuery(condition, true)

	query := fmt.Sprintf(`WITH matches_list AS (
		%s
		UNION
		%s
	)
	SELECT matches.*
	FROM matches_list matches
    ORDER BY matches.start_date DESC
    %s;
	`, team1_matches_query, team2_matches_query, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Innings(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	main_condition := sqlWhere.getConditionString("AND ")

	team_field, opposition_field := "batting_team_id", "bowling_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_field, opposition_field = "bowling_team_id", "batting_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT matches.id,
		innings.%s AS team_id,
		teams.name AS team_name,
		innings.%s AS opposition_id,
		teams2.name AS opposition_name,
		matches.ground_id,
  		cities.name AS city_name,
	    matches.start_date,
		matches.final_result,
		matches.match_winner_team_id,
		innings.id AS innings_id,
		innings.innings_number,
		innings.innings_end,
		innings.total_runs,
		innings.total_wickets,
		innings.total_balls / 6 + (innings.total_balls %% 6) * 0.1 AS overs,
		(
		    CASE
		        WHEN innings.total_balls > 0 THEN innings.total_runs * 6.0 / innings.total_balls
		    END
		) AS scoring_rate
    FROM innings
        LEFT JOIN matches ON innings.match_id = matches.id
        LEFT JOIN teams ON innings.%s = teams.id
        LEFT JOIN teams teams2 ON innings.%s = teams2.id
        LEFT JOIN grounds ON matches.ground_id = grounds.id
        LEFT JOIN cities ON grounds.city_id = cities.id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY matches.id,
		innings.batting_team_id,
		teams.name,
		innings.bowling_team_id,
		teams2.name,
		matches.ground_id,
    	cities.name ,
	    matches.start_date,
		matches.final_result,
		matches.match_winner_team_id,
		innings.id,
		innings.innings_number,
		innings.innings_end,
		innings.total_runs,
		innings.total_wickets
    ORDER BY matches.start_date DESC
    %s;
	`, team_field, opposition_field, team_field, opposition_field, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT match_teams.team_id,
		teams.name AS team_name,
		matches.ground_id,
	    grounds.name AS ground_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
        LEFT JOIN grounds ON matches.ground_id = grounds.id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY match_teams.team_id,
    	teams.name,	
    	matches.ground_id,
        grounds.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT match_teams.team_id,
		teams.name AS team_name,
		cities.host_nation_id,
	    host_nations.name AS host_nation_name,
	    MIN(matches.start_date) AS min_date,
	    MAX(matches.start_date) AS max_date,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
        LEFT JOIN grounds ON matches.ground_id = grounds.id
        LEFT JOIN cities ON grounds.city_id = cities.id
        LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY match_teams.team_id,
    	teams.name,	
    	cities.host_nation_id,
        host_nations.name
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT match_teams.team_id,
		teams.name AS team_name,
		date_part('year', matches.start_date)::integer AS match_year,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY match_teams.team_id,
    	teams.name,	
    	date_part('year', matches.start_date)::integer
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := &sqlWhere{}
	sqlWhere.applyFilters(params, team_stats)
	teams_conditon := sqlWhere.getConditionString("WHERE ")
	main_condition := sqlWhere.getConditionString("AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`WITH match_teams AS (
	    SELECT matches.team1_id AS team_id FROM matches %s
	    UNION
	    SELECT matches.team2_id AS team_id FROM matches %s
	)
	SELECT match_teams.team_id,
		teams.name AS team_name,
		matches.season,
	    %s
    FROM match_teams
        LEFT JOIN innings ON innings.%s = match_teams.team_id
        %s
    WHERE innings.is_super_over = FALSE
    	%s
    GROUP BY match_teams.team_id,
    	teams.name,	
    	matches.season
    ORDER BY matches_won DESC
    %s;
	`, teams_conditon, teams_conditon, team_numbers_query, team_total_for, team_common_joins, main_condition, pagination)

	return query, sqlWhere.args, limit, nil
}

// helpers

func getIndividualMatchesQuery(condition string, isTeamSwitch bool) string {
	const common_fields string = `
		matches.ground_id,
	    cities.name,
	    matches.start_date,
	    matches.final_result,
		matches.match_winner_team_id,
		matches.toss_winner_team_id,
		matches.is_toss_decision_bat,
		matches.win_margin,
		matches.balls_remaining_after_win,
		matches.is_won_by_runs,
		matches.is_won_by_innings
	`

	team_fields := `
		matches.team1_id AS team_id,
		teams.name AS team_name,
		matches.team2_id AS opposition_id,
		teams2.name AS opposition_name
	`

	if isTeamSwitch {
		team_fields = `
			matches.team1_id AS opposition_id,
			teams.name AS opposition_name,
			matches.team2_id AS team_id,
			teams2.name AS team_name
		`
	}

	query := fmt.Sprintf(`
		SELECT matches.id,
			%s,
			%s
	    FROM matches
 			LEFT JOIN teams ON matches.team1_id = teams.id
		 	LEFT JOIN teams teams2 ON matches.team2_id = teams2.id
		    LEFT JOIN grounds ON matches.ground_id = grounds.id
		    LEFT JOIN cities ON grounds.city_id = cities.id
			%s
	    GROUP BY matches.id,
			matches.team1_id,
			teams.name,
			matches.team2_id,
			teams2.name,
			%s
	`, team_fields, common_fields, condition, common_fields)

	return query
}
