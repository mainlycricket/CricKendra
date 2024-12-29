WITH match_teams AS (
    SELECT matches.team1_id AS team_id
    FROM matches
    WHERE matches.playing_format = 'ODI'
        AND matches.team1_id IN (1, 8, 10)
        AND matches.team2_id IN (1, 8, 10)
        AND matches.ground_id IN (63, 70, 79, 90, 124)
        AND matches.start_date BETWEEN '2008-08-18' AND '2024-09-27'
        AND matches.season IN (
            '2022/23',
            '2019/20',
            '2017/18',
            '2013/14',
            '2011/12'
        )
    UNION
    SELECT matches.team2_id AS team_id
    FROM matches
    WHERE matches.playing_format = 'ODI'
        AND matches.team1_id IN (1, 8, 10)
        AND matches.team2_id IN (1, 8, 10)
        AND matches.ground_id IN (63, 70, 79, 90, 124)
        AND matches.start_date BETWEEN '2008-08-18' AND '2024-09-27'
        AND matches.season IN (
            '2022/23',
            '2019/20',
            '2017/18',
            '2013/14',
            '2011/12'
        )
)
SELECT matches.ground_id AS ground_id,
    grounds.name AS ground_name,
    MIN(matches.start_date) AS min_date,
    MAX(matches.start_date) AS max_date,
    COUNT(DISTINCT match_teams.team_id) AS teams_count,
    COUNT(DISTINCT matches.id) AS matches_played,
    SUM(
        CASE
            WHEN match_teams.team_id = matches.match_winner_team_id THEN 1
            ELSE 0
        END
    ) AS matches_won,
    SUM(
        CASE
            WHEN match_teams.team_id = matches.match_loser_team_id THEN 1
            ELSE 0
        END
    ) AS matches_lost,
    (
        CASE
            WHEN SUM(
                CASE
                    WHEN match_teams.team_id = matches.match_loser_team_id THEN 1
                    ELSE 0
                END
            ) > 0 THEN SUM(
                CASE
                    WHEN match_teams.team_id = matches.match_winner_team_id THEN 1
                    ELSE 0
                END
            ) * 1.0 / SUM(
                CASE
                    WHEN match_teams.team_id = matches.match_loser_team_id THEN 1
                    ELSE 0
                END
            )
        END
    ) AS win_loss_ratio,
    COUNT(
        CASE
            WHEN matches.final_result = 'drawn' THEN 1
        END
    ) AS matches_drawn,
    COUNT(
        CASE
            WHEN matches.final_result = 'tied' THEN 1
        END
    ) AS matches_tied,
    COUNT(
        CASE
            WHEN matches.final_result = 'no_result' THEN 1
        END
    ) AS matches_no_result,
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
FROM match_teams
    LEFT JOIN innings ON innings.bowling_team_id = match_teams.team_id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN teams ON match_teams.team_id = teams.id
    LEFT JOIN grounds ON matches.ground_id = grounds.id
WHERE matches.playing_format = 'ODI'
    AND matches.team1_id IN (1, 8, 10)
    AND matches.team2_id IN (1, 8, 10)
    AND matches.ground_id IN (63, 70, 79, 90, 124)
    AND matches.start_date BETWEEN '2008-08-18' AND '2024-09-27'
    AND matches.season IN (
        '2022/23',
        '2019/20',
        '2017/18',
        '2013/14',
        '2011/12'
    )
    AND innings.is_super_over = FALSE
GROUP BY matches.ground_id,
    grounds.name
ORDER BY matches_won;