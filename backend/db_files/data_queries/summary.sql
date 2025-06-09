WITH
	top_batters AS (
		SELECT
			bs.innings_id,
			bs.batter_id,
			batter.name AS batter_name,
			bs.runs_scored,
			bs.balls_faced,
			ROW_NUMBER() OVER (
				PARTITION BY
					bs.innings_id
				ORDER BY
					bs.runs_scored DESC,
					bs.balls_faced ASC,
					bs.dismissal_type IS NULL,
					bs.batting_position ASC
			) AS rn
		FROM
			batting_scorecards bs
			LEFT JOIN innings ON bs.innings_id = innings.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			LEFT JOIN players batter ON bs.batter_id = batter.id
		WHERE
			innings.match_id = 1
	),
	aggregated_top_batters AS (
		SELECT
			top_batters.innings_id,
			ARRAY_AGG (
				ROW (
					top_batters.batter_id,
					top_batters.batter_name,
					top_batters.runs_scored,
					top_batters.balls_faced
				)
			) AS batters
		FROM
			top_batters
		WHERE
			top_batters.rn <= 2
		GROUP BY
			top_batters.innings_id
	),
	top_bowlers AS (
		SELECT
			bs.innings_id,
			bs.bowler_id,
			bowler.name AS bowler_name,
			bs.balls_bowled / 6 + bs.balls_bowled % 6 * 0.1 AS overs_bowled,
			bs.runs_conceded,
			bs.wickets_taken,
			ROW_NUMBER() OVER (
				PARTITION BY
					bs.innings_id
				ORDER BY
					bs.wickets_taken DESC,
					bs.runs_conceded ASC,
					bs.balls_bowled / 6 + bs.balls_bowled % 6 * 0.1 DESC,
					bs.bowling_position ASC
			) AS rn
		FROM
			bowling_scorecards bs
			LEFT JOIN innings ON bs.innings_id = innings.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
			LEFT JOIN players bowler ON bs.bowler_id = bowler.id
		WHERE
			innings.match_id = 1
	),
	aggregated_top_bowlers AS (
		SELECT
			top_bowlers.innings_id,
			ARRAY_AGG (
				ROW (
					top_bowlers.bowler_id,
					top_bowlers.bowler_name,
					top_bowlers.overs_bowled,
					top_bowlers.wickets_taken,
					top_bowlers.runs_conceded
				)
			) AS bowlers
		FROM
			top_bowlers
		WHERE
			top_bowlers.rn <= 2
		GROUP BY
			top_bowlers.innings_id
	),
	scorecard_summary AS (
		SELECT
			innings.innings_number,
			innings.batting_team_id,
			batting_team.name,
			innings.total_runs,
			innings.total_wickets,
			innings.total_balls,
			ARRAY_AGG (aggregated_top_batters.batters),
			ARRAY_AGG (aggregated_top_bowlers.bowlers)
		FROM
			innings
			LEFT JOIN teams batting_team ON innings.batting_team_id = batting_team.id
			LEFT JOIN aggregated_top_batters ON aggregated_top_batters.innings_id = innings.id
			LEFT JOIN aggregated_top_bowlers ON aggregated_top_bowlers.innings_id = innings.id
		WHERE
			innings.match_id = 1
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE
		GROUP BY
			innings.innings_number,
			innings.batting_team_id,
			batting_team.name,
			innings.total_runs,
			innings.total_wickets,
			innings.total_balls
	),
	latest_commentary AS (
		SELECT
			deliveries.innings_delivery_number,
			deliveries.ball_number,
			deliveries.over_number,
			deliveries.batter_id,
			batter.name,
			deliveries.bowler_id,
			bowler.name,
			deliveries.fielder1_id,
			fielder1.name,
			deliveries.fielder2_Id,
			fielder2.name,
			deliveries.wides,
			deliveries.noballs,
			deliveries.legbyes,
			deliveries.byes,
			deliveries.total_runs,
			deliveries.is_four,
			deliveries.is_six,
			deliveries.player1_dismissed_id,
			player1_dismissed.name,
			deliveries.player1_dismissal_type,
			bs1.runs_scored,
			bs1.balls_faced,
			bs1.fours_scored,
			bs1.sixes_scored,
			deliveries.player2_dismissed_id,
			player2_dismissed.name,
			deliveries.player2_dismissal_type,
			bs2.runs_scored,
			bs2.balls_faced,
			bs2.fours_scored,
			bs2.sixes_scored,
			deliveries.commentary
		FROM
			deliveries
			LEFT JOIN innings ON deliveries.innings_id = innings.id
			AND innings.is_super_over = FALSE
			LEFT JOIN matches ON innings.match_id = matches.id
			LEFT JOIN players batter ON deliveries.batter_id = batter.id
			LEFT JOIN players bowler ON deliveries.bowler_id = bowler.id
			LEFT JOIN players fielder1 ON deliveries.fielder1_id = fielder1.id
			LEFT JOIN players fielder2 ON deliveries.fielder2_id = fielder2.id
			LEFT JOIN players player1_dismissed ON deliveries.player1_dismissed_id = player1_dismissed.id
			LEFT JOIN players player2_dismissed ON deliveries.player2_dismissed_id = player2_dismissed.id
			LEFT JOIN batting_scorecards bs1 ON bs1.innings_id = deliveries.innings_id
			AND bs1.batter_id = deliveries.player1_dismissed_id
			LEFT JOIN batting_scorecards bs2 ON bs2.innings_id = deliveries.innings_id
			AND bs2.batter_id = deliveries.player2_dismissed_id
		WHERE
			matches.id = 1
		ORDER BY
			innings.innings_number DESC,
			deliveries.innings_delivery_number DESC
		FETCH FIRST
			50 ROWS ONLY
	)
SELECT
	matches.id,
	matches.playing_level,
	matches.playing_format,
	matches.match_type,
	matches.event_match_number,
	-- Day 1, 2, etc - Test / FC
	-- Stumps, Innings Break, Tea/Lunch/Dinner, Stopped
	-- Need 50 runs, won by 5 wkts, trail/lead by 8 runs, won the toss and chose to bat, match starts in
	matches.season,
	matches.start_date,
	matches.end_date,
	matches.start_time,
	matches.is_day_night,
	matches.ground_id,
	grounds.name,
	matches.main_series_id,
	main_series.name,
	matches.team1_id,
	team1.name,
	team1.image_url,
	matches.team2_id,
	team2.name,
	team2.image_url,
	(
		SELECT
			ARRAY_AGG (
				-- order is necessary for struct scanning
				ROW (
					player_awards.player_id,
					players.name,
					player_awards.award_type
				)
			)
		FROM
			player_awards
			LEFT JOIN players ON player_awards.player_id = players.id
		WHERE
			player_awards.match_id = matches.id
	) AS match_awards,
	(
		SELECT
			ARRAY_AGG (
				-- order is necessary for struct scanning
				ROW (
					innings.innings_number,
					innings.batting_team_id,
					batting_team.name,
					innings.total_runs,
					innings.total_balls,
					innings.total_wickets,
					innings.innings_end,
					innings.target_runs,
					innings.max_overs
				)
			)
	) AS team_innings_short_info,
	(
		SELECT
			ARRAY_AGG (scorecard_summary.*)
		FROM
			scorecard_summary
	) AS scorecard_summary,
	(
		SELECT
			ARRAY_AGG (latest_commentary.*)
		FROM
			latest_commentary
	) AS latest_commentary
FROM
	matches
	LEFT JOIN innings ON innings.match_id = matches.id
	AND innings.innings_number IS NOT NULL
	AND innings.is_super_over = FALSE
	LEFT JOIN teams batting_team ON innings.batting_team_id = batting_team.id
	LEFT JOIN teams team1 ON matches.team1_id = team1.id
	LEFT JOIN teams team2 ON matches.team2_id = team2.id
	LEFT JOIN grounds ON matches.ground_id = grounds.id
	LEFT JOIN series main_series ON matches.main_series_id = main_series.id
WHERE
	matches.id = 1
GROUP BY
	matches.id,
	grounds.name,
	main_series.name,
	team1.name,
	team1.image_url,
	team2.name,
	team2.image_url;