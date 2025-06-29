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
    matches.season,
    matches.ground_id,
    grounds.name,
    matches.main_series_id,
    series.name,
    matches.team1_id,
    t1.name,
    t1.image_url,
    matches.team2_id,
    t2.name,
    t2.image_url,
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
                    mse.team_id,
                    mse.player_id,
                    players.name,
                    mse.is_captain,
                    mse.is_wk,
                    mse.is_debut,
                    mse.is_vice_captain,
                    mse.playing_status
                )
            )
        FROM
            match_squad_entries mse
            LEFT JOIN players ON mse.player_id = players.id
        WHERE
            mse.match_id = matches.id
    ) AS squad_entries,
    (
        SELECT
            ARRAY_AGG (
                ROW (player_awards.player_id, player_awards.award_type)
            )
        FROM
            player_awards
        WHERE
            player_awards.match_id = matches.id
    ) AS match_awards,
    (
        SELECT
            ARRAY_AGG (
                ROW (
                    innings.id,
                    innings.innings_number,
                    innings.batting_team_id,
                    innings.bowling_team_id,
                    innings.total_runs,
                    innings.total_balls,
                    innings.total_wickets,
                    innings.byes,
                    innings.leg_byes,
                    innings.wides,
                    innings.noballs,
                    innings.penalty,
                    innings.is_super_over,
                    innings.innings_end,
                    innings.target_runs,
                    innings.max_overs,
                    (
                        SELECT
                            ARRAY_AGG (
                                ROW (
                                    bs.batter_id,
                                    bs.batting_position,
                                    bs.runs_scored,
                                    bs.balls_faced,
                                    bs.minutes_batted,
                                    bs.fours_scored,
                                    bs.sixes_scored,
                                    bs.dismissal_type,
                                    bs.dismissed_by_id,
                                    bs.fielder1_id,
                                    bs.fielder2_id
                                )
                            )
                        FROM
                            batting_scorecards bs
                        WHERE
                            bs.innings_id = innings.id
                    ),
                    (
                        SELECT
                            ARRAY_AGG (
                                ROW (
                                    bs.bowler_id,
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
                        WHERE
                            bs.innings_id = innings.id
                    )
                )
            )
        FROM
            innings
        WHERE
            innings.match_id = matches.id
            AND innings.innings_number IS NOT NULL
    ) AS innings
FROM
    matches
    LEFT JOIN teams t1 ON matches.team1_id = t1.id
    LEFT JOIN teams t2 ON matches.team2_id = t2.id
    LEFT JOIN grounds ON matches.ground_id = grounds.id
    LEFT JOIN series ON matches.main_series_id = series.id
WHERE
    matches.id = 1
