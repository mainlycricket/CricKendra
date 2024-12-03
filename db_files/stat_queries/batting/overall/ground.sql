WITH match_stats AS (
    SELECT matches.ground_id,
        COUNT(DISTINCT matches.id) AS match_count,
        COUNT(DISTINCT mse.player_id) AS players_count
    FROM match_squad_entries mse
        LEFT JOIN matches ON mse.match_id = matches.id
        LEFT JOIN innings ON innings.match_id = matches.id
    WHERE matches.playing_format = 'ODI'
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
        AND innings.batting_team_id IN (1, 8, 10)
        AND innings.bowling_team_id IN (1, 8, 10)
        AND mse.playing_status IN ('playing_xi')
    GROUP BY matches.ground_id
),
batting_performance AS (
    SELECT matches.ground_id,
        COUNT(
            CASE
                WHEN dismissal_type IS NULL
                OR dismissal_type IN ('retired hurt', 'retired not out') THEN 1
            END
        ) AS not_outs,
        COUNT(
            CASE
                WHEN runs_scored >= 100 THEN 1
            END
        ) AS centuries,
        MAX(runs_scored) AS highest_score,
        MAX(
            CASE
                WHEN dismissal_type IS NULL
                OR dismissal_type IN ('retired hurt', 'retired not out') THEN runs_scored
                ELSE 0
            END
        ) as highest_not_out_score,
        COUNT(
            CASE
                WHEN runs_scored BETWEEN 50 AND 99 THEN 1
            END
        ) AS half_centuries,
        COUNT(
            CASE
                WHEN runs_scored = 0 THEN 1
            END
        ) AS ducks
    FROM batting_scorecards bs
        LEFT JOIN innings ON bs.innings_id = innings.id
        LEFT JOIN matches ON innings.match_id = matches.id
    WHERE innings.is_super_over = FALSE
        AND innings.batting_team_id IN (1, 8, 10)
        AND innings.bowling_team_id IN (1, 8, 10)
        AND matches.playing_format = 'ODI'
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
    GROUP BY matches.ground_id
)
SELECT ms.match_count,
    ms.players_count AS players_count,
    matches.ground_id,
    grounds.name AS ground_name,
    MIN(matches.start_date) AS min_date,
    MAX(matches.end_date) AS max_date,
    COUNT(innings.id) AS innings_count,
    SUM(bs.runs_scored) AS runs_scored,
    SUM(bs.balls_faced) AS balls_faced,
    (
        CASE
            WHEN (COUNT(innings.id) - bp.not_outs) > 0 THEN SUM(bs.runs_scored) * 1.0 / (COUNT(innings.id) - bp.not_outs)
            ELSE sum(bs.runs_scored)
        END
    ) AS average,
    (
        CASE
            WHEN SUM(bs.balls_faced) > 0 THEN SUM(bs.runs_scored) * 100.0 / SUM(bs.balls_faced)
            ELSE 0
        END
    ) AS strike_rate,
    bp.not_outs,
    bp.highest_score,
    (
        CASE
            WHEN bp.highest_score = bp.highest_not_out_score THEN TRUE
            ELSE FALSE
        END
    ) AS is_highest_not_out,
    bp.centuries,
    bp.half_centuries,
    bp.centuries + bp.half_centuries AS fifty_plus_scores,
    bp.ducks,
    SUM(bs.fours_scored) AS fours_scored,
    SUM(bs.sixes_scored) AS sixes_scored
FROM batting_scorecards bs
    LEFT JOIN innings ON innings.id = bs.innings_id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN grounds ON matches.ground_id = grounds.id
    LEFT JOIN batting_performance bp ON bp.ground_id = matches.ground_id
    LEFT JOIN match_stats ms ON ms.ground_id = matches.ground_id
WHERE innings.is_super_over = FALSE
    AND matches.playing_format = 'ODI'
    AND innings.batting_team_id IN (1, 8, 10)
    AND innings.bowling_team_id IN (1, 8, 10)
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
GROUP BY matches.ground_id,
    grounds.name,
    bp.not_outs,
    bp.centuries,
    bp.half_centuries,
    bp.ducks,
    bp.highest_score,
    bp.highest_not_out_score,
    ms.players_count,
    ms.match_count
ORDER BY grounds.name DESC;