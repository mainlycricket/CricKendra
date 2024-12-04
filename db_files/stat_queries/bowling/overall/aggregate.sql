WITH match_stats AS (
    SELECT COUNT(DISTINCT matches.id) AS matches_count,
        COUNT(DISTINCT player_id) AS players_count
    FROM match_squad_entries mse
        LEFT JOIN matches ON mse.match_id = matches.id
        LEFT JOIN innings ON innings.match_id = matches.id
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
        AND mse.team_id IN (1, 8, 10)
        AND (
            CASE
                WHEN mse.team_id = matches.team1_id THEN matches.team2_id
                ELSE matches.team1_id
            END
        ) IN (1, 8, 10)
),
bowling_performance AS (
    SELECT MAX(bs.wickets_taken) AS best_innings_fig_wkts,
        COUNT(
            CASE
                WHEN bs.wickets_taken = 4 THEN 1
            END
        ) AS four_wkt_hauls,
        COUNT(
            CASE
                WHEN bs.wickets_taken >= 5 THEN 1
            END
        ) AS five_wkt_hauls
    FROM bowling_scorecards bs
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
        AND innings.batting_team_id IN (1, 8, 10)
        AND innings.bowling_team_id IN (1, 8, 10)
),
best_innings_fig_runs AS (
    SELECT MIN(bs.runs_conceded) AS best_innings_fig_runs
    FROM bowling_scorecards bs
        LEFT JOIN innings ON bs.innings_id = innings.id
        LEFT JOIN matches ON innings.match_id = matches.id
        LEFT JOIN bowling_performance bp ON TRUE
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
        AND bs.wickets_taken = bp.best_innings_fig_wkts
)
SELECT ms.players_count,
    MIN(matches.start_date) AS min_date,
    MAX(matches.start_date) AS max_date,
    ms.matches_count,
    COUNT(innings.id) AS innings_count,
    SUM(bs.balls_bowled) / 6 + (SUM(balls_bowled) % 6) * 0.1 AS overs_bowled,
    SUM(bs.runs_conceded) AS runs_conceded,
    SUM(bs.wickets_taken) AS wickets_taken,
    (
        CASE
            WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.runs_conceded) * 1.0 / SUM(bs.wickets_taken)
            ELSE '+infinity'
        END
    ) AS average,
    (
        CASE
            WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.balls_bowled) * 1.0 / SUM(bs.wickets_taken)
            ELSE '+infinity'
        END
    ) AS strike_rate,
    (
        CASE
            WHEN SUM(bs.balls_bowled) > 0 THEN SUM(bs.runs_conceded) * 6.0 / SUM(bs.balls_bowled)
            ELSE '+infinity'
        END
    ) AS economy,
    bp.best_innings_fig_wkts,
    bifr.best_innings_fig_runs,
    bp.five_wkt_hauls,
    bp.four_wkt_hauls,
    SUM(bs.fours_conceded) AS fours_conceded,
    SUM(bs.sixes_conceded) AS sixes_conceded
FROM bowling_scorecards bs
    LEFT JOIN innings ON bs.innings_id = innings.id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN bowling_performance bp ON TRUE
    LEFT JOIN match_stats ms ON TRUE
    LEFT JOIN best_innings_fig_runs bifr ON TRUE
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
GROUP BY ms.matches_count,
    ms.players_count,
    bifr.best_innings_fig_runs,
    bp.best_innings_fig_wkts,
    bp.four_wkt_hauls,
    bp.five_wkt_hauls
ORDER BY SUM(bs.wickets_taken) DESC;