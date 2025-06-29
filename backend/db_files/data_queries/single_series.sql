WITH
    matches_data AS (
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
            grounds.name AS ground_name,
            matches.team1_id,
            t1.name AS team1_name,
            t1.image_url AS team1_image_url,
            matches.team2_id,
            t2.name AS team2_name,
            t2.image_url AS team2_image_url,
            matches.season,
            matches.main_series_id,
            series.name AS main_series_name,
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
            ) AS innings
        FROM
            matches
            LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
            LEFT JOIN series ON matches.main_series_id = series.id
            LEFT JOIN teams t1 ON matches.team1_id = t1.id
            LEFT JOIN teams t2 ON matches.team2_id = t2.id
            LEFT JOIN grounds ON matches.ground_id = grounds.id
            LEFT JOIN innings ON innings.match_id = matches.id
            AND innings.is_super_over = FALSE
            AND innings.innings_number IS NOT NULL
        WHERE
            mse.series_id = 1
        GROUP BY
            matches.id,
            grounds.name,
            t1.name,
            t2.name,
            t1.image_url,
            t2.image_url,
            series.name
    )
SELECT
    s.id,
    s.name,
    s.is_male,
    s.playing_level,
    s.playing_format,
    s.season,
    ARRAY_AGG (DISTINCT ROW (ste.team_id, t.name)),
    s.start_date,
    s.end_date,
    s.winner_team_id,
    s.final_status,
    s.tour_flag,
    s.tournament_id,
    tournaments.name,
    ARRAY_AGG (DISTINCT ROW (matches_data))
FROM
    series s
    LEFT JOIN matches_data ON TRUE
    LEFT JOIN tournaments ON s.tournament_id = tournaments.id
    LEFT JOIN series_team_entries ste ON s.id = ste.series_id
    LEFT JOIN teams t ON ste.team_id = t.id
WHERE
    s.id = 1
GROUP BY
    s.id,
    tournaments.name
