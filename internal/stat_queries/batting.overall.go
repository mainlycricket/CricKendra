package statqueries

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

// Function Names are in Query_Overall_Batting_x_Stats format, x represents grouping

func Query_Overall_Batting_Batters(params *url.Values) (string, []any, int, error) {
	var filters []string
	var args []any

	HandlePlayingFormat(params, &filters, &args)
	HandleIsMale(params, &filters, &args)
	HandleMinStartDate(params, &filters, &args)
	HandleMaxStartDate(params, &filters, &args)
	HandleSeasons(params, &filters, &args)
	HandleBattingTeam(params, &filters, &args)
	HandleBowlingTeam(params, &filters, &args)
	HandleGround(params, &filters, &args)

	condition := fmt.Sprintf(` AND %s`, strings.Join(filters, " AND "))

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT bs.batter_id,
		players.name AS batter_name,
		ARRAY_AGG(DISTINCT teams.short_name) AS teams,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		COUNT(DISTINCT matches.id) AS matches_played,
		COUNT(innings.id) AS innings_count,
		SUM(bs.runs_scored) AS runs_scored,
		SUM(bs.balls_faced) AS balls_faced,
		COUNT(
			CASE
				WHEN bs.dismissal_type IS NULL
				OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
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
				WHEN bs.runs_scored >= 100 THEN 1
			END
		) + COUNT(
			CASE
				WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
			END
		) AS fifty_plus_scores,
		COUNT(
			CASE
				WHEN bs.runs_scored = 0 THEN 1
			END
		) AS ducks,
		SUM(bs.fours_scored) AS fours_scored,
		SUM(bs.sixes_scored) AS sixes_scored
	FROM matches
		LEFT JOIN innings ON innings.match_id = matches.id
		LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = bs.batter_id
		LEFT JOIN teams ON mse.team_id = teams.id
	WHERE innings.is_super_over = FALSE
	%s
	GROUP BY bs.batter_id,
		players.name
	ORDER BY runs_scored DESC
	%s;
	`, condition, pagination)

	return query, args, limit, nil
}

func Query_Overall_Batting_Teams(params *url.Values) (string, []any, int, error) {
	var filters []string
	var args []any

	HandlePlayingFormat(params, &filters, &args)
	HandleIsMale(params, &filters, &args)
	HandleMinStartDate(params, &filters, &args)
	HandleMaxStartDate(params, &filters, &args)
	HandleSeasons(params, &filters, &args)
	HandleBattingTeam(params, &filters, &args)
	HandleBowlingTeam(params, &filters, &args)
	HandleGround(params, &filters, &args)

	condition := fmt.Sprintf(` AND %s`, strings.Join(filters, " AND "))

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT innings.batting_team_id AS team_id,
		teams.name AS team_name,
		COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		COUNT(DISTINCT matches.id) AS matches_played,
		COUNT(innings.id) AS innings_count,
		SUM(bs.runs_scored) AS runs_scored,
		SUM(bs.balls_faced) AS balls_faced,
		COUNT(
			CASE
				WHEN bs.dismissal_type IS NULL
				OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
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
				WHEN bs.runs_scored >= 100 THEN 1
			END
		) + COUNT(
			CASE
				WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
			END
		) AS fifty_plus_scores,
		COUNT(
			CASE
				WHEN bs.runs_scored = 0 THEN 1
			END
		) AS ducks,
		SUM(bs.fours_scored) AS fours_scored,
		SUM(bs.sixes_scored) AS sixes_scored
	FROM matches
		LEFT JOIN innings ON innings.match_id = matches.id
		LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
		LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = bs.batter_id
		LEFT JOIN teams ON innings.batting_team_id = teams.id
	WHERE innings.is_super_over = FALSE
	%s
	GROUP BY innings.batting_team_id,
		teams.name
	ORDER BY runs_scored DESC
	%s;
	`, condition, pagination)

	return query, args, limit, nil
}

func Query_Overall_Batting_Oppositions(params *url.Values) (string, []any, int, error) {
	var filters []string
	var args []any

	HandlePlayingFormat(params, &filters, &args)
	HandleIsMale(params, &filters, &args)
	HandleMinStartDate(params, &filters, &args)
	HandleMaxStartDate(params, &filters, &args)
	HandleSeasons(params, &filters, &args)
	HandleBattingTeam(params, &filters, &args)
	HandleBowlingTeam(params, &filters, &args)
	HandleGround(params, &filters, &args)

	condition := fmt.Sprintf(` AND %s`, strings.Join(filters, " AND "))

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT innings.bowling_team_id AS team_id,
		teams.name AS team_name,
		COUNT(DISTINCT mse.player_id) AS players_count,
		MIN(matches.start_date) AS min_date,
		MAX(matches.start_date) AS max_date,
		COUNT(DISTINCT matches.id) AS matches_played,
		COUNT(innings.id) AS innings_count,
		SUM(bs.runs_scored) AS runs_scored,
		SUM(bs.balls_faced) AS balls_faced,
		COUNT(
			CASE
				WHEN bs.dismissal_type IS NULL
				OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
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
				WHEN bs.runs_scored >= 100 THEN 1
			END
		) + COUNT(
			CASE
				WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
			END
		) AS fifty_plus_scores,
		COUNT(
			CASE
				WHEN bs.runs_scored = 0 THEN 1
			END
		) AS ducks,
		SUM(bs.fours_scored) AS fours_scored,
		SUM(bs.sixes_scored) AS sixes_scored
	FROM matches
		LEFT JOIN innings ON innings.match_id = matches.id
		LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
		LEFT JOIN players ON bs.batter_id = players.id
		LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = bs.batter_id
		LEFT JOIN teams ON innings.bowling_team_id  = teams.id
	WHERE innings.is_super_over = FALSE
	%s
	GROUP BY innings.bowling_team_id,
		teams.name
	ORDER BY runs_scored DESC
	%s;
	`, condition, pagination)

	return query, args, limit, nil
}

func Query_Overall_Batting_Seasons(params *url.Values) (string, []any, int, error) {
	var filters []string
	var args []any

	HandlePlayingFormat(params, &filters, &args)
	HandleIsMale(params, &filters, &args)
	HandleMinStartDate(params, &filters, &args)
	HandleMaxStartDate(params, &filters, &args)
	HandleSeasons(params, &filters, &args)
	HandleBattingTeam(params, &filters, &args)
	HandleBowlingTeam(params, &filters, &args)
	HandleGround(params, &filters, &args)

	condition := fmt.Sprintf(` AND %s`, strings.Join(filters, " AND "))

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT matches.season,
		COUNT(DISTINCT mse.player_id) AS players_count,
		COUNT(DISTINCT matches.id) AS matches_played,
		COUNT(innings.id) AS innings_count,
		SUM(bs.runs_scored) AS runs_scored,
		SUM(bs.balls_faced) AS balls_faced,
		COUNT(
			CASE
				WHEN bs.dismissal_type IS NULL
				OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
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
				WHEN bs.runs_scored >= 100 THEN 1
			END
		) + COUNT(
			CASE
				WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
			END
		) AS fifty_plus_scores,
		COUNT(
			CASE
				WHEN bs.runs_scored = 0 THEN 1
			END
		) AS ducks,
		SUM(bs.fours_scored) AS fours_scored,
		SUM(bs.sixes_scored) AS sixes_scored
	FROM matches
		LEFT JOIN innings ON innings.match_id = matches.id
		LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
		LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = bs.batter_id
	WHERE innings.is_super_over = FALSE
	%s
	GROUP BY matches.season
	ORDER BY runs_scored DESC
	%s;
	`, condition, pagination)

	return query, args, limit, nil
}

func Query_Overall_Batting_Years(params *url.Values) (string, []any, int, error) {
	var filters []string
	var args []any

	HandlePlayingFormat(params, &filters, &args)
	HandleIsMale(params, &filters, &args)
	HandleMinStartDate(params, &filters, &args)
	HandleMaxStartDate(params, &filters, &args)
	HandleSeasons(params, &filters, &args)
	HandleBattingTeam(params, &filters, &args)
	HandleBowlingTeam(params, &filters, &args)
	HandleGround(params, &filters, &args)

	condition := fmt.Sprintf(` AND %s`, strings.Join(filters, " AND "))

	skip, limit := pgxutils.GetPaginationParams(params)
	pagination := fmt.Sprintf(`OFFSET %d ROWS FETCH FIRST %d ROWS ONLY`, skip, (limit + 1))

	query := fmt.Sprintf(`SELECT date_part('year', matches.start_date)::int AS match_year,
		COUNT(DISTINCT mse.player_id) AS players_count,
		COUNT(DISTINCT matches.id) AS matches_played,
		COUNT(innings.id) AS innings_count,
		SUM(bs.runs_scored) AS runs_scored,
		SUM(bs.balls_faced) AS balls_faced,
		COUNT(
			CASE
				WHEN bs.dismissal_type IS NULL
				OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
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
				WHEN bs.runs_scored >= 100 THEN 1
			END
		) + COUNT(
			CASE
				WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
			END
		) AS fifty_plus_scores,
		COUNT(
			CASE
				WHEN bs.runs_scored = 0 THEN 1
			END
		) AS ducks,
		SUM(bs.fours_scored) AS fours_scored,
		SUM(bs.sixes_scored) AS sixes_scored
	FROM matches
		LEFT JOIN innings ON innings.match_id = matches.id
		LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
		LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
		AND mse.team_id = innings.batting_team_id
		AND mse.player_id = bs.batter_id
	WHERE innings.is_super_over = FALSE
	%s
	GROUP BY date_part('year', matches.start_date)::int
	ORDER BY runs_scored DESC
	%s;
	`, condition, pagination)

	return query, args, limit, nil
}
