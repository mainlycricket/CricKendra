WITH batting_performance AS (
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
        AND matches.playing_format = 'ODI'
        AND matches.start_date >= '2008-08-18'
    GROUP BY batter_id
)
SELECT bs.batter_id,
    players.name AS batter_name,
    COUNT(innings.id) AS innings_count,
    SUM(bs.runs_scored) AS runs_scored,
    SUM(bs.balls_faced) AS balls_faced,
    bp.not_outs,
    bp.highest_score,
    bp.centuries,
    bp.half_centuries,
    bp.ducks,
    SUM(bs.fours_scored) AS fours_scored,
    SUM(bs.sixes_scored) AS sixes_scored
FROM batting_scorecards bs
    LEFT JOIN players ON bs.batter_id = players.id
    LEFT JOIN innings ON bs.innings_id = innings.id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN batting_performance bp ON bs.batter_id = bp.batter_id
WHERE innings.is_super_over = FALSE
    AND matches.playing_format = 'ODI'
    AND matches.start_date >= '2008-08-18'
GROUP BY bs.batter_id,
    players.name,
    bp.highest_score,
    bp.not_outs,
    bp.centuries,
    bp.half_centuries,
    bp.ducks
ORDER BY runs_scored DESC;