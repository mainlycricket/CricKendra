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
			ARRAY_AGG (
				-- order is necessary for struct scanning
				ROW (
					innings.innings_number,
					innings.batting_team_id,
					batting_team.name,
					innings.total_runs,
					innings.total_balls,
					innings.total_wickets,
					innings.byes,
					innings.leg_byes,
					innings.wides,
					innings.noballs,
					innings.penalty,
					innings.innings_end,
					innings.target_runs,
					innings.max_overs,
					(
						SELECT
							ARRAY_AGG (
								-- order is necessary for struct scanning
								ROW (
									bs.batter_id,
									batter.name,
									bs.batting_position,
									bs.has_batted,
									bs.runs_scored,
									bs.balls_faced,
									bs.minutes_batted,
									bs.fours_scored,
									bs.sixes_scored,
									bs.dismissal_type,
									bs.dismissed_by_id,
									dismissed_by.name,
									bs.fielder1_id,
									fielder1.name,
									bs.fielder2_id,
									fielder2.name
								)
							)
						FROM
							batting_scorecards bs
							LEFT JOIN players batter ON bs.batter_id = batter.id
							LEFT JOIN players dismissed_by ON bs.dismissed_by_id = dismissed_by.id
							LEFT JOIN players fielder1 ON bs.fielder1_id = fielder1.id
							LEFT JOIN players fielder2 ON bs.fielder2_id = fielder2.id
						WHERE
							bs.innings_id = innings.id
					),
					(
						SELECT
							ARRAY_AGG (
								-- order is necessary for struct scanning
								ROW (
									bs.bowler_id,
									bowler.name,
									bs.bowling_position,
									bs.wickets_taken,
									bs.runs_conceded,
									bs.balls_bowled,
									bs.maiden_overs,
									bs.fours_conceded,
									bs.sixes_conceded,
									bs.wides_conceded,
									bs.noballs_conceded
								)
							)
						FROM
							bowling_scorecards bs
							LEFT JOIN players bowler ON bs.bowler_id = bowler.id
						WHERE
							bs.innings_id = innings.id
							AND bs.bowling_position IS NOT NULL
					)
				)
			)
	) AS match_innings
FROM
	matches
	LEFT JOIN teams team1 ON matches.team1_id = team1.id
	LEFT JOIN teams team2 ON matches.team2_id = team2.id
	LEFT JOIN grounds ON matches.ground_id = grounds.id
	LEFT JOIN series main_series ON matches.main_series_id = main_series.id
	LEFT JOIN innings ON innings.match_id = matches.id
	AND innings.innings_number IS NOT NULL
	AND innings.is_super_over = FALSE
	LEFT JOIN teams batting_team ON innings.batting_team_id = batting_team.id
WHERE
	matches.id = 1
GROUP BY
	matches.id,
	grounds.name,
	main_series.name,
	team1.name,
	team1.image_url,
	team2.name,
	team2.image_url