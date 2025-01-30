package statqueries

import (
	"fmt"
	"net/url"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

// Function Names are in Query_Overall_Bowling_x format, x represents grouping

func Query_Overall_Bowling_Bowlers(params *url.Values) (string, []any, int, error) {
	sqlWhere := newSqlWhere(bowling_stats, overall_players_group)

	sqlWhere.applyFilters(params)

	matchQuery := sqlWhere.matchQuery.prepareQuery()
	inningsCondition := sqlWhere.inningsFilters.getClause()
	qualificationClause := sqlWhere.qualifications.getClause()

	best_bowling_figures := prepareBestBowlingFigures("bowling_scorecards.bowler_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = bowling_scorecards.bowler_id
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

	best_bowling_figures := prepareBestBowlingFigures("innings.id", sqlWhere.inningsFilters.conditions)

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
	LEFT JOIN best_bowling_figures bbf ON bbf.group_field = innings.id
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

	best_bowling_figures := prepareBestBowlingFigures("matches.id", sqlWhere.inningsFilters.conditions)

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
	LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.id
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

	best_bowling_figures := prepareBestBowlingFigures("innings.bowling_team_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = innings.bowling_team_id
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

	best_bowling_figures := prepareBestBowlingFigures("innings.batting_team_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = innings.batting_team_id
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

	best_bowling_figures := prepareBestBowlingFigures("matches.ground_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.ground_id
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

	best_bowling_figures := prepareBestBowlingFigures("matches.host_nation_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.host_nation_id
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

	best_bowling_figures := prepareBestBowlingFigures("matches.continent_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.continent_id
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

	best_bowling_figures := prepareBestBowlingFigures("matches.series_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.series_id
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

	best_bowling_figures := prepareBestBowlingFigures("matches.tournament_id", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.tournament_id
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

	best_bowling_figures := prepareBestBowlingFigures("date_part('year', matches.start_date)::integer", sqlWhere.inningsFilters.conditions)

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
			LEFT JOIN best_bowling_figures bbf ON bbf.group_field = date_part('year', matches.start_date)::integer
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

	best_bowling_figures := prepareBestBowlingFigures("matches.season", sqlWhere.inningsFilters.conditions)

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
		LEFT JOIN best_bowling_figures bbf ON bbf.group_field = matches.season
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

	best_bowling_figures := prepareBestBowlingFigures("10 * date_part('decade', matches.start_date)::integer", sqlWhere.inningsFilters.conditions)

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
			LEFT JOIN best_bowling_figures bbf ON bbf.group_field = 10 * date_part('decade', matches.start_date)::integer
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

	best_bowling_figures := prepareBestBowlingFigures("TRUE", sqlWhere.inningsFilters.conditions)

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

func prepareBestBowlingFigures(groupField string, inningsConditions []string) string {
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
			LEFT JOIN bowling_scorecards bowling_scorecards ON bowling_scorecards.innings_id = innings.id
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
			SELECT DISTINCT ON (mw.group_field)
				mw.group_field,
				mw.total_wickets AS wickets,
				mw.total_runs AS runs
			FROM match_wickets mw
			ORDER BY mw.group_field,
				mw.total_wickets DESC,
				mw.total_runs ASC
		),
		best_bowling_innings AS (
			SELECT %s AS group_field,
				MAX(bowling_scorecards.wickets_taken) AS wickets
			FROM matches
			LEFT JOIN innings ON innings.match_id = matches.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards ON bowling_scorecards.innings_id = innings.id
			%s
			GROUP BY group_field
		),
		best_bowling_figures AS (
			SELECT bbm.group_field,
				twh.hauls_count AS ten_wicket_hauls,
				bbm.wickets AS best_match_wickets,
				bbm.runs AS best_match_runs,
				bbi.wickets AS best_innings_wickets
			FROM best_bowling_match bbm
			LEFT JOIN ten_wicket_hauls twh ON twh.group_field = bbm.group_field
			LEFT JOIN best_bowling_innings bbi ON bbi.group_field = bbm.group_field
		)
	`, groupField, condition, groupField, condition)

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
