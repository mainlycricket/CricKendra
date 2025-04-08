WITH
	match_wickets AS (
		SELECT
			bs.bowler_id,
			matches.id AS match_id,
			matches.playing_format,
			SUM(bs.wickets_taken) AS total_wickets,
			SUM(bs.runs_conceded) AS total_runs
		FROM
			matches
			LEFT JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE
			matches.id = 1
		GROUP BY
			matches.id,
			matches.playing_format,
			bs.bowler_id
	),
	ten_wicket_hauls AS (
		SELECT
			mw.bowler_id,
			mw.playing_format,
			COUNT(*) AS hauls_count
		FROM
			match_wickets mw
		WHERE
			mw.total_wickets >= 10
		GROUP BY
			mw.bowler_id,
			mw.playing_format
	),
	best_bowling_match AS (
		SELECT DISTINCT
			ON (mw.bowler_id, mw.playing_format) mw.bowler_id,
			mw.playing_format,
			mw.total_wickets AS wickets,
			mw.total_runs AS runs
		FROM
			match_wickets mw
		ORDER BY
			mw.bowler_id,
			mw.playing_format,
			mw.total_wickets DESC,
			mw.total_runs ASC
	),
	best_bowling_innings AS (
		SELECT
			bs.bowler_id,
			matches.playing_format,
			MAX(bs.wickets_taken) AS wickets
		FROM
			matches
			LEFT JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
		WHERE
			matches.id = 1
		GROUP BY
			bs.bowler_id,
			matches.playing_format
	),
	best_bowling_figures AS (
		SELECT
			bbm.bowler_id,
			bbm.playing_format,
			twh.hauls_count AS ten_wicket_hauls,
			bbm.wickets AS best_match_wickets,
			bbm.runs AS best_match_runs,
			bbi.wickets AS best_innings_wickets
		FROM
			best_bowling_match bbm
			LEFT JOIN ten_wicket_hauls twh ON twh.bowler_id = bbm.bowler_id
			LEFT JOIN best_bowling_innings bbi ON bbi.bowler_id = bbm.bowler_id
			AND bbi.playing_format = bbm.playing_format
	),
	match_stats AS (
		-- Get batting and bowling stats per player per match
		SELECT
			m.playing_format,
			mse.player_id,
			-- stats
			ROW (
				COUNT(DISTINCT m.id),
				COUNT(
					DISTINCT CASE
						WHEN bs.has_batted THEN i.id
					END
				),
				COALESCE(SUM(bs.runs_scored), 0),
				COUNT(
					CASE
						WHEN bs.dismissal_type IS NULL
						OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
					END
				),
				COALESCE(SUM(bs.balls_faced), 0),
				COALESCE(SUM(bs.fours_scored), 0),
				COALESCE(SUM(bs.sixes_scored), 0),
				COUNT(
					CASE
						WHEN bs.runs_scored >= 100 THEN 1
					END
				),
				COUNT(
					CASE
						WHEN bs.runs_scored >= 50
						AND bs.runs_scored < 100 THEN 1
					END
				),
				MAX(bs.runs_scored),
				CASE
					WHEN MAX(bs.runs_scored) = MAX(
						CASE
							WHEN bs.dismissal_type IS NULL
							OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN bs.runs_scored
						END
					) THEN TRUE
					ELSE FALSE
				END,
				COUNT(
					DISTINCT CASE
						WHEN bws.bowling_position IS NOT NULL THEN i.id
					END
				),
				COALESCE(SUM(bws.runs_conceded), 0),
				COALESCE(SUM(bws.wickets_taken), 0),
				COALESCE(SUM(bws.balls_bowled), 0),
				COALESCE(SUM(bws.fours_conceded), 0),
				COALESCE(SUM(bws.sixes_conceded), 0),
				COUNT(
					CASE
						WHEN bws.wickets_taken >= 4
						AND bws.wickets_taken < 5 THEN 1
					END
				),
				COUNT(
					CASE
						WHEN bws.wickets_taken >= 5 THEN 1
					END
				),
				bbf.ten_wicket_hauls,
				MIN(
					CASE
						WHEN bbf.best_innings_wickets = bws.wickets_taken THEN bws.runs_conceded
					END
				),
				bbf.best_innings_wickets,
				bbf.best_match_runs,
				bbf.best_match_wickets
			)::career_stats AS stats
		FROM
			matches m
			JOIN match_squad_entries mse ON mse.match_id = m.id
			JOIN innings i ON i.match_id = m.id
			LEFT JOIN batting_scorecards bs ON bs.innings_id = i.id
			AND mse.player_id = bs.batter_id
			LEFT JOIN bowling_scorecards bws ON bws.innings_id = i.id
			AND mse.player_id = bws.bowler_id
			LEFT JOIN best_bowling_figures bbf ON bbf.bowler_id = mse.player_id
			AND bbf.playing_format = m.playing_format
		WHERE
			m.id = 1
			AND i.is_super_over = FALSE
			AND i.innings_number IS NOT NULL
		GROUP BY
			m.playing_format,
			mse.player_id,
			bbf.ten_wicket_hauls,
			bbf.best_innings_wickets,
			bbf.best_match_wickets,
			bbf.best_match_runs
	)
UPDATE players p
SET -- Update individual format stats
	db_test_stats = (
		SELECT
			combine_career_stats (
				(
					SELECT
						ms.stats
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
						AND ms.playing_format = 'Test'
				),
				db_test_stats
			)
	),
	db_odi_stats = (
		SELECT
			combine_career_stats (
				(
					SELECT
						ms.stats
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
						AND ms.playing_format = 'ODI'
				),
				db_odi_stats
			)
	),
	db_t20i_stats = (
		SELECT
			combine_career_stats (
				(
					SELECT
						ms.stats
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
						AND ms.playing_format = 'T20I'
				),
				db_t20i_stats
			)
	),
	db_fc_stats = (
		SELECT
			combine_career_stats (
				(
					SELECT
						ms.stats
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
						AND (
							ms.playing_format = 'first_class'
							OR ms.playing_format = 'Test'
						)
				),
				db_fc_stats
			)
	),
	db_lista_stats = (
		SELECT
			combine_career_stats (
				(
					SELECT
						ms.stats
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
						AND (
							ms.playing_format = 'list_a'
							OR ms.playing_format = 'ODI'
						)
				),
				db_lista_stats
			)
	),
	db_t20_stats = (
		SELECT
			combine_career_stats (
				(
					SELECT
						ms.stats
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
						AND (
							ms.playing_format = 'T20'
							OR ms.playing_format = 'T20I'
						)
				),
				db_t20_stats
			)
	)
WHERE
	EXISTS (
		SELECT
			1
		FROM
			match_stats ms
		WHERE
			ms.player_id = p.id
	);