package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

// Function Names are in Query_Overall_Team_x format, x represents grouping

func Query_Overall_Team_Teams(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY team_id,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Players(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("mse.team_id")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT mse.player_id,
			players.name AS player_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			COUNT(DISTINCT mse.team_id) AS teams_count,
			%s
		FROM matches
		LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		LEFT JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			AND innings.%s = mse.team_id
		LEFT JOIN players ON mse.player_id = players.id
		%s
		GROUP BY mse.player_id,
			players.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, team_total_for, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureGround()
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for, team_total_against := "batting_team_id", "bowling_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
		team_total_against = "batting_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), match_teams AS (
			SELECT combined_teams.match_id, COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.id AS match_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.id AS match_id, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.match_id
		)
		SELECT matches.id AS match_id,
			innings.%s,
			team1.name AS team1_name,
			innings.%s,
			team2.name AS team2_name,
			matches.ground_id,
			matches.city_name,
			matches.season,
			matches.start_date,
			match_teams.teams_count,
			%s
		FROM matches
		LEFT JOIN match_teams ON match_teams.match_id = matches.id
		LEFT JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		LEFT JOIN teams team1 ON innings.%s = team1.id
		LEFT JOIN teams team2 ON innings.%s = team2.id
		%s
		GROUP BY matches.id,
			innings.batting_team_id,
			innings.bowling_team_id,
			team1.name,
			team2.name,
			matches.ground_id,
			matches.city_name,
			matches.season,
			matches.start_date,
			match_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_total_against, team_numbers_query, team_total_for, team_total_against, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), series_teams AS (
			SELECT combined_teams.series_id,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.series_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.series_id, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.series_id
		)
		SELECT matches.series_id,
			matches.series_name,
			matches.series_season,
			series_teams.teams_count,
			%s
		FROM matches
		JOIN series_teams ON series_teams.series_id = matches.series_id
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season,
			series_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), tournament_teams AS (
			SELECT combined_teams.tournament_id,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.tournament_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.tournament_id, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.tournament_id
		)
		SELECT matches.tournament_id,
			matches.tournament_name,
			tournament_teams.teams_count,
			%s
		FROM matches
		JOIN tournament_teams ON tournament_teams.tournament_id = matches.tournament_id
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name,
			tournament_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), ground_teams AS (
			SELECT combined_teams.ground_id,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.ground_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.ground_id, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.ground_id
		)
		SELECT matches.ground_id,
			matches.ground_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			ground_teams.teams_count,
			%s
		FROM matches
		JOIN ground_teams ON ground_teams.ground_id = matches.ground_id
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY matches.ground_id,
			matches.ground_name,
			ground_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), host_nation_teams AS (
			SELECT combined_teams.host_nation_id,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.host_nation_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.host_nation_id, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.host_nation_id
		)
		SELECT matches.host_nation_id,
			matches.host_nation_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			host_nation_teams.teams_count,
			%s
		FROM matches
		JOIN host_nation_teams ON host_nation_teams.host_nation_id = matches.host_nation_id
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY matches.host_nation_id,
			matches.host_nation_name,
			host_nation_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Continents(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureContinent()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), continent_teams AS (
			SELECT combined_teams.continent_id,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.continent_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.continent_id, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.continent_id
		)
		SELECT matches.continent_id,
			matches.continent_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			continent_teams.teams_count,
			%s
		FROM matches
		JOIN continent_teams ON continent_teams.continent_id = matches.continent_id
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY matches.continent_id,
			matches.continent_name,
			continent_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	fmt.Println(query)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), year_teams AS (
			SELECT combined_teams.year,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT date_part('year', matches.start_date)::integer AS year, 
					matches.team1_id AS team_id
				FROM matches
				UNION
				SELECT date_part('year', matches.start_date)::integer AS year,
					matches.team2_id AS team_id
				FROM matches
			) combined_teams
			GROUP BY combined_teams.year
		)
		SELECT date_part('year', matches.start_date)::integer AS match_year,
			year_teams.teams_count,
			%s
		FROM matches
		JOIN year_teams ON year_teams.year = date_part('year', matches.start_date)::integer
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY match_year,
			year_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), season_teams AS (
			SELECT combined_teams.season,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.season, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.season, matches.team2_id AS team_id FROM matches
			) combined_teams
			GROUP BY combined_teams.season
		)
		SELECT matches.season,
			season_teams.teams_count,
			%s
		FROM matches
		JOIN season_teams ON season_teams.season = matches.season
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY matches.season,
			season_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Decades(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), decade_teams AS (
			SELECT combined_teams.decade,
				COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT 10 * date_part('decade', matches.start_date)::integer AS decade, 
					matches.team1_id AS team_id
				FROM matches
				UNION
				SELECT 10 * date_part('decade', matches.start_date)::integer AS decade,
					matches.team2_id AS team_id
				FROM matches
			) combined_teams
			GROUP BY combined_teams.decade
		)
		SELECT 10 * date_part('decade', matches.start_date)::integer AS match_decade,
			decade_teams.teams_count,
			%s
		FROM matches
		JOIN decade_teams ON decade_teams.decade = 10 * date_part('decade', matches.start_date)::integer
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY match_decade,
			decade_teams.teams_count
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_numbers_query, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Team_Aggregate(params *url.Values) (string, []any, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), aggregate_teams AS (
			SELECT COUNT(DISTINCT combined_teams.team_id) AS teams_count
			FROM (
				SELECT matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.team2_id AS team_id FROM matches
			) combined_teams
		)
		SELECT aggregate_teams.teams_count,
			%s
		FROM matches
		LEFT JOIN aggregate_teams ON TRUE
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		%s
		GROUP BY aggregate_teams.teams_count
		ORDER BY matches_won DESC;
		`, matchQuery, team_numbers_query, inningsCondition)

	return query, sqlWhere.args, nil
}

// Function Names are in Query_Individual_Team_x format, x represents grouping

func Query_Individual_Team_Innings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureGround()
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for, team_total_against := "batting_team_id", "bowling_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for, team_total_against = "bowling_team_id", "batting_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.id,
			innings.%s AS team_id,
			teams.name AS team_name,
			innings.%s AS opposition_id,
			teams2.name AS opposition_name,
			matches.ground_id,
			matches.city_name,
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
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		LEFT JOIN teams ON innings.%s = teams.id
		LEFT JOIN teams teams2 ON innings.%s = teams2.id
		%s
		GROUP BY matches.id,
			innings.batting_team_id,
			teams.name,
			innings.bowling_team_id,
			teams2.name,
			matches.ground_id,
			matches.city_name,
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
		`, matchQuery, team_total_for, team_total_against, team_total_for, team_total_against, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.fields = append(sqlWhere.matchQuery.fields, "matches.win_margin", "matches.balls_remaining_after_win", "matches.is_won_by_runs", "matches.is_won_by_innings")
	sqlWhere.matchQuery.ensureGround()
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for, team_total_against := "batting_team_id", "bowling_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for, team_total_against = "bowling_team_id", "batting_team_id"
	}

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		), match_teams AS (
			SELECT combined_teams.match_id, combined_teams.team_id
			FROM ( 
				SELECT matches.id AS match_id, matches.team1_id AS team_id FROM matches
				UNION
				SELECT matches.id AS match_id, matches.team2_id AS team_id FROM matches
			) combined_teams
		)
		SELECT matches.id,
			innings.%s AS team1_id,
			teams.name AS team1_name,
			innings.%s AS team2_id,
			teams2.name AS team2_name,
			matches.ground_id,
			matches.city_name,
			matches.start_date,
			matches.final_result,
			matches.match_winner_team_id,
			matches.toss_winner_team_id,
			innings.innings_number,
			matches.win_margin,
			matches.balls_remaining_after_win,
			matches.is_won_by_runs,
			matches.is_won_by_innings
		FROM matches
		JOIN match_teams ON matches.id = match_teams.match_id
		JOIN innings ON innings.match_id = matches.id
		JOIN teams ON innings.%s = teams.id
		JOIN teams teams2 ON innings.%s = teams2.id
		%s
		GROUP BY matches.id,
			innings.batting_team_id,
			innings.bowling_team_id,
			teams.name,
			teams2.name,
			matches.ground_id,
			matches.city_name,
			matches.start_date,
			matches.final_result,
			matches.match_winner_team_id,
			matches.toss_winner_team_id,
			innings.innings_number,
			matches.win_margin,
			matches.balls_remaining_after_win,
			matches.is_won_by_runs,
			matches.is_won_by_innings
		ORDER BY matches.start_date DESC
		%s;
		`, matchQuery, team_total_for, team_total_against, team_total_for, team_total_against, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			matches.series_id,
			matches.series_name,
			matches.series_season,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season,
			innings.%s,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, team_total_for, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			matches.tournament_id,
			matches.tournament_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name,
			innings.%s,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, team_total_for, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			matches.ground_id,
			matches.ground_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY matches.ground_id,
			matches.ground_name,
			innings.%s,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, team_total_for, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			matches.host_nation_id,
			matches.host_nation_name,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY matches.host_nation_id,
			matches.host_nation_name,
			innings.%s,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, team_total_for, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			date_part('year', matches.start_date) AS match_year,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY date_part('year', matches.start_date),
			innings.%s,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, team_total_for, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Team_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, team_stats)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	team_total_for := "batting_team_id"
	if params.Get("team_total_for") == "bowling" {
		team_total_for = "bowling_team_id"
	}

	team_numbers_query := getTeamNumbersQuery("innings." + team_total_for)

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.%s AS team_id,
			teams.name AS team_name,
			matches.season,
			%s
		FROM matches
		JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		JOIN teams ON innings.%s = teams.id
		%s
		GROUP BY matches.season,
			innings.%s,
			teams.name
		ORDER BY matches_won DESC
		%s;
		`, matchQuery, team_total_for, team_numbers_query, team_total_for, inningsCondition, team_total_for, pagination)

	return query, sqlWhere.args, limit, nil
}

// helpers

func getTeamNumbersQuery(teamField string) string {
	query := fmt.Sprintf(`
		COUNT(DISTINCT matches.id) AS matches_played,
		COUNT(CASE WHEN %s = matches.match_winner_team_id THEN 1 END) AS matches_won,
		COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END) AS matches_lost,
		(
			CASE
				WHEN COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END) > 0
				THEN 
					COUNT(CASE WHEN %s = matches.match_winner_team_id THEN 1 END) * 1.0 / COUNT(CASE WHEN %s = matches.match_loser_team_id THEN 1 END)
			END
		) AS win_loss_ratio,
		COUNT(CASE WHEN matches.final_result = 'draw' THEN 1  END) AS matches_drawn,
		COUNT(CASE WHEN matches.final_result = 'tie' THEN 1 END) AS matches_tied,
		COUNT(CASE WHEN matches.final_result = 'no result' THEN 1 END) AS matches_no_result,
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
		MIN(
			CASE WHEN innings.innings_end = 'all_out' THEN innings.total_runs END
		) AS lowest_score
		`, teamField, teamField, teamField, teamField, teamField)

	return query
}
