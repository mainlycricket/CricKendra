WITH match_stats AS (
    SELECT player_id,
        COUNT(DISTINCT matches.id) AS matches_count,
        MIN(matches.start_date) AS min_date,
        MAX(matches.start_date) AS max_date,
        ARRAY_AGG(DISTINCT(teams.short_name)) AS teams
    FROM match_squad_entries mse
        LEFT JOIN matches ON mse.match_id = matches.id
        LEFT JOIN teams ON teams.id = mse.team_id
    WHERE matches.playing_format = 'ODI'
        AND matches.start_date >= '2008-08-18'
        AND matches.start_date <= '2024-09-27'
        AND matches.ground_id IN (63, 70, 79, 90, 124)
        AND matches.season IN (
            '2022/23',
            '2019/20',
            '2017/18',
            '2013/14',
            '2011/12'
        )
        AND mse.playing_status IN ('playing_xi')
        AND mse.team_id IN (1, 10)
        AND (
            CASE
                WHEN mse.team_id = matches.team1_id THEN matches.team2_id
                ELSE matches.team1_id
            END
        ) IN (1, 8, 10)
    GROUP BY player_id
),
batting_performance AS (
    SELECT batter_id,
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
        AND innings.batting_team_id IN (1, 10)
        AND innings.bowling_team_id IN (1, 8, 10)
    GROUP BY batter_id
)
SELECT bs.batter_id,
    players.name AS batter_name,
    ms.matches_count,
    ms.min_date,
    ms.max_date,
    ms.teams,
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
    LEFT JOIN players ON bs.batter_id = players.id
    LEFT JOIN innings ON bs.innings_id = innings.id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN match_stats ms ON ms.player_id = bs.batter_id
    LEFT JOIN batting_performance bp ON bs.batter_id = bp.batter_id
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
    AND innings.batting_team_id IN (1, 10)
    AND innings.bowling_team_id IN (1, 8, 10)
GROUP BY bs.batter_id,
    players.name,
    ms.matches_count,
    ms.teams,
    ms.min_date,
    ms.max_date,
    bp.highest_score,
    bp.highest_not_out_score,
    bp.not_outs,
    bp.centuries,
    bp.half_centuries,
    bp.ducks
ORDER BY runs_scored DESC;