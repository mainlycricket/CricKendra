package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

// Function Names are in Query_Overall_Batting_x format, x represents grouping

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
