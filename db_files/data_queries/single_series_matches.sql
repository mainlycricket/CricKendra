SELECT
    matches.id,
    matches.playing_level,
    matches.playing_format,
    matches.match_type,
    matches.cricsheet_id,
    matches.balls_per_over,
    matches.event_match_number,
    matches.start_date,
    matches.end_date,
    matches.start_time,
    matches.is_day_night,
    matches.ground_id,
    grounds.name,
    matches.team1_id,
    t1.name,
    t1.image_url,
    matches.team2_id,
    t2.name,
    t2.image_url,
    matches.season,
    matches.main_series_id,
    ms.name,
    matches.current_status,
    matches.final_result,
    matches.outcome_special_method,
    matches.match_winner_team_id,
    matches.bowl_out_winner_id,
    matches.super_over_winner_id,
    matches.is_won_by_innings,
    matches.is_won_by_runs,
    matches.win_margin,
    matches.balls_remaining_after_win,
    matches.toss_winner_team_id,
    matches.is_toss_decision_bat,
    (
        SELECT
            ARRAY_AGG (
                ROW (
                    innings.innings_number,
                    innings.batting_team_id,
                    innings.total_runs,
                    innings.total_balls,
                    innings.total_wickets,
                    innings.innings_end,
                    innings.target_runs,
                    innings.max_overs
                )
            )
        FROM
            innings
        WHERE
            innings.match_id = matches.id
            AND innings.is_super_over = FALSE
            AND innings.innings_number IS NOT NULL
    )
FROM
    matches
    LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
    LEFT JOIN series ON mse.series_id = series.id
    LEFT JOIN series ms ON matches.main_series_id = ms.id
    LEFT JOIN teams t1 ON matches.team1_id = t1.id
    LEFT JOIN teams t2 ON matches.team2_id = t2.id
    LEFT JOIN grounds ON matches.ground_id = grounds.id
WHERE
    series.id = 1
