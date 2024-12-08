SELECT mse.player_id,
    players.name AS player_name,
    COUNT(DISTINCT mse.team_id) AS teams_count,
    MIN(matches.start_date) AS min_date,
    MAX(matches.start_date) AS max_date,
    COUNT(DISTINCT matches.id) AS matches_played,
    COUNT(
        DISTINCT CASE
            WHEN matches.match_winner_team_id = mse.team_id THEN matches.id
        END
    ) AS matches_won,
    COUNT(
        DISTINCT CASE
            WHEN matches.match_loser_team_id = mse.team_id THEN matches.id
        END
    ) AS matches_won,
    (
        CASE
            WHEN COUNT(
                DISTINCT CASE
                    WHEN matches.match_loser_team_id = mse.team_id THEN matches.id
                END
            ) > 0 THEN COUNT(
                DISTINCT CASE
                    WHEN matches.match_winner_team_id = mse.team_id THEN matches.id
                END
            ) * 1.0 / COUNT(
                DISTINCT CASE
                    WHEN matches.match_loser_team_id = mse.team_id THEN matches.id
                END
            )
        END
    ) AS win_loss_ratio,
    COUNT(innings.id) AS innings_count,
    SUM(innings.total_runs) AS total_runs,
    SUM(innings.total_balls) AS total_balls,
    SUM(innings.total_wickets) AS total_wickets,
    (
        CASE
            WHEN SUM(innings.total_wickets) > 0 THEN SUM(innings.total_runs) * 1.0 / SUM(innings.total_wickets)
        END
    ) AS average,
    (
        CASE
            WHEN SUM(innings.total_balls) > 0 THEN SUM(innings.total_runs) * 6.0 / SUM(innings.total_balls)
        END
    ) AS scoring_rate,
    MAX(innings.total_runs) AS highest_score,
    MIN(innings.total_runs) AS lowest_score
FROM match_squad_entries mse
    LEFT JOIN matches ON matches.id = mse.match_id
    LEFT JOIN innings ON innings.match_id = matches.id
    AND innings.bowling_team_id = mse.team_id
    LEFT JOIN players ON mse.player_id = players.id
WHERE matches.playing_format = 'ODI'
    AND matches.team1_id IN (1, 8, 10)
    AND matches.team2_id IN (1, 8, 10)
    AND matches.ground_id IN (63, 70, 79, 90, 124)
    AND matches.start_date >= '2008-08-18'
    AND matches.start_date <= '2024-09-27'
    AND matches.season IN (
        '2022/23',
        '2019/20',
        '2017/18',
        '2013/14',
        '2011/12'
    )
    AND innings.is_super_over = FALSE
GROUP BY mse.player_id,
    players.name
ORDER BY SUM(innings.total_runs) DESC