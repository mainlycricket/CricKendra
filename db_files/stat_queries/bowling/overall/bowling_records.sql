WITH bowling_performance AS (
    SELECT bowler_id,
        COUNT(
            CASE
                WHEN wickets_taken = 4 THEN 1
            END
        ) AS four_wkt_hauls,
        COUNT(
            CASE
                WHEN wickets_taken >= 5 THEN 1
            END
        ) AS five_wkt_hauls
    FROM bowling_scorecards bs
        LEFT JOIN innings ON bs.innings_id = innings.id
        LEFT JOIN matches ON innings.match_id = matches.id
    WHERE innings.is_super_over = FALSE
        AND matches.playing_format = 'ODI'
        AND matches.start_date >= '2018-01-01'
    GROUP BY bowler_id
)
SELECT bs.bowler_id,
    players.name AS bowler_name,
    COUNT(innings.id) AS innings_count,
    SUM(bs.runs_conceded) AS runs_conceded,
    SUM(bs.balls_bowled) AS balls_bowled,
    SUM(bs.wickets_taken) AS wickets_taken,
    bp.four_wkt_hauls,
    bp.five_wkt_hauls,
    SUM(bs.fours_conceded) AS fours_conceded,
    SUM(bs.sixes_conceded) AS sixes_conceded
FROM bowling_scorecards bs
    LEFT JOIN players ON bs.bowler_id = players.id
    LEFT JOIN innings ON bs.innings_id = innings.id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN bowling_performance bp ON bs.bowler_id = bp.bowler_id
WHERE innings.is_super_over = FALSE
    AND matches.playing_format = 'ODI'
    AND matches.start_date >= '2018-01-01'
GROUP BY bs.bowler_id,
    players.name,
    bp.four_wkt_hauls,
    bp.five_wkt_hauls
ORDER BY wickets_taken DESC;