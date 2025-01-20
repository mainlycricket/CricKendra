package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

const batting_numbers_query string = `
	COUNT(DISTINCT matches.id) AS matches_played,
	COUNT(CASE WHEN bs.has_batted THEN innings.id END) AS innings_count,
	SUM(bs.runs_scored) AS runs_scored,
	SUM(bs.balls_faced) AS balls_faced,
	COUNT(
		CASE
			WHEN bs.has_batted AND (
				bs.dismissal_type IS NULL
				OR bs.dismissal_type IN ('retired hurt', 'retired not out')) THEN 1
		END
	) AS not_outs,
	(
		CASE
			WHEN COUNT(
				CASE
					WHEN bs.dismissal_type IS NOT NULL
					AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
				END
			) > 0 THEN SUM(bs.runs_scored) * 1.0 / COUNT(
				CASE
					WHEN bs.dismissal_type IS NOT NULL
					AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
				END
			)
		END
	) AS average,
	(
		CASE
			WHEN SUM(bs.balls_faced) > 0 THEN SUM(bs.runs_scored) * 100.0 / SUM(bs.balls_faced)
		END
	) AS strike_rate,
	MAX(bs.runs_scored) AS highest_score,
	MAX(
		CASE
			WHEN bs.dismissal_type IS NULL
			OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN runs_scored
			ELSE 0
		END
	) as highest_not_out_score,		
	COUNT(
		CASE
			WHEN bs.runs_scored >= 100 THEN 1
		END
	) AS centuries,
	COUNT(
		CASE
			WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
		END
	) AS half_centuries,
	COUNT(
		CASE
			WHEN bs.runs_scored >= 50 THEN 1
		END
	) AS fifty_plus_scores,		
	COUNT(
		CASE
			WHEN bs.has_batted AND
				bs.runs_scored = 0 AND
				bs.dismissal_type IS NOT NULL AND
				bs.dismissal_type NOT IN ('retired hurt', 'retired not out')
			THEN 1
			END
	) AS ducks,
	SUM(bs.fours_scored) AS fours_scored,
	SUM(bs.sixes_scored) AS sixes_scored
`

const batting_common_joins string = `
	LEFT JOIN innings ON innings.match_id = matches.id
		AND innings.innings_number IS NOT NULL
		AND innings.is_super_over = FALSE
	LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
	LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = bs.batter_id
		AND mse.playing_status IN ('playing_xi')
`

// Function Names are in Query_Overall_Batting_x format, x represents grouping

func Query_Overall_Batting_Batters(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY bs.batter_id,
			players.name
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_TeamInnings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Matches(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Teams(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON innings.bowling_team_id  = teams.id
		%s
		GROUP BY innings.bowling_team_id,
			teams.name
		ORDER BY runs_scored DESC
			%s;
			`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Continents(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureContinent()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Overall_Batting_Aggregate(params *url.Values) (string, []any, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
		ORDER BY runs_scored DESC;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition)

	return query, sqlWhere.args, nil
}

// Function Names are in Query_Individual_Batting_x format, x represents grouping

func Query_Individual_Batting_Innings(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureGround()
	sqlWhere.matchQuery.ensureCity()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

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
			bs.batter_id,
			players.name AS batter_name,
			innings.batting_team_id,
			teams.short_name as batting_team_name,
			innings.bowling_team_id,
			teams2.name AS bowling_team_name,
			bs.runs_scored,
			bs.balls_faced,
			(
				CASE
					WHEN bs.dismissal_type IS NULL
					OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN TRUE
					ELSE FALSE
				END
			) AS is_not_out,
			(
				CASE
					WHEN bs.balls_faced > 0 THEN bs.runs_scored * 100.0 / bs.balls_faced
					ELSE 0
				END
			) AS strike_rate,
			bs.fours_scored,
			bs.sixes_scored
		FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
			LEFT JOIN teams ON innings.batting_team_id = teams.id
			LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
			LEFT JOIN players ON bs.batter_id = players.id
		%s
		GROUP BY bs.innings_id,
			bs.batter_id,
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
			bs.runs_scored,
			bs.balls_faced,
			bs.dismissal_type
		ORDER BY bs.runs_scored DESC
		%s;
		`, matchQuery, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Series(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)

	sqlWhere.matchQuery.ensureSeries()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.series_id,
			matches.series_name, 
			matches.series_season,
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.series_id,
			matches.series_name,
			matches.series_season,
			bs.batter_id,
			players.name
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Tournaments(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)

	sqlWhere.matchQuery.ensureTournament()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.tournament_id,
			matches.tournament_name,
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.tournament_id,
			matches.tournament_name,
			bs.batter_id,
			players.name
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Grounds(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureGround()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.ground_id,
			matches.ground_name, 
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY matches.ground_id,
			matches.ground_name,
			bs.batter_id,
			players.name
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_HostNations(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	sqlWhere.matchQuery.ensureHostNation()
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.host_nation_id,
			matches.host_nation_name,
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY bs.batter_id,
			players.name,
			matches.host_nation_id,
			matches.host_nation_name
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Oppositions(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT innings.bowling_team_id AS opposition_team_id,
			teams2.name AS opposition_team_name,
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			MIN(matches.start_date) AS min_date,
			MAX(matches.start_date) AS max_date,
		%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
		%s
		GROUP BY bs.batter_id,
			players.name,
			innings.bowling_team_id,
			teams2.name
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Years(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT date_part('year', matches.start_date)::integer AS match_year,
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY bs.batter_id,
			players.name,
			date_part('year', matches.start_date)::integer
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}

func Query_Individual_Batting_Seasons(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere()
	sqlWhere.applyFilters(params, batting_stats)
	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := prefixJoin(sqlWhere.inningsFilters.conditions, "WHERE", " AND ")

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`
		WITH matches AS (
			%s
		)
		SELECT matches.season,
			bs.batter_id,
			players.name AS batter_name,
			ARRAY_AGG(DISTINCT teams.short_name) AS teams_represented,
			%s
		FROM matches
		%s
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN teams ON mse.team_id = teams.id
		%s
		GROUP BY bs.batter_id,
			players.name,
			matches.season
		ORDER BY runs_scored DESC
		%s;
		`, matchQuery, batting_numbers_query, batting_common_joins, inningsCondition, pagination)

	return query, sqlWhere.args, limit, nil
}
