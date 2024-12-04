WITH match_stats AS (
    SELECT matches.id AS match_id,
        COUNT(DISTINCT matches.id) AS matches_count,
        COUNT(DISTINCT player_id) AS players_count
    FROM match_squad_entries mse
        LEFT JOIN matches ON mse.match_id = matches.id
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
    GROUP BY matches.id
),
bowling_performance AS (
    SELECT matches.id AS match_id,
        MAX(bs.wickets_taken) AS best_innings_fig_wkts,
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
    GROUP BY matches.id
),
best_innings_fig_runs AS (
    SELECT matches.id AS match_id,
        MIN(bs.runs_conceded) AS best_innings_fig_runs
    FROM bowling_scorecards bs
        LEFT JOIN innings ON bs.innings_id = innings.id
        LEFT JOIN matches ON innings.match_id = matches.id
        LEFT JOIN bowling_performance bp ON matches.id = bp.match_id
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
    GROUP BY matches.id
)
SELECT matches.id AS match_id,
    matches.start_date,
    matches.team1_id,
    teams.name AS team1_name,
    matches.team2_id,
    teams2.name AS team2_name,
    matches.season,
    grounds.name AS ground_name,
    ms.players_count,
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
    LEFT JOIN teams ON matches.team1_id = teams.id
    LEFT JOIN teams teams2 ON matches.team2_id = teams2.id
    LEFT JOIN grounds ON matches.ground_id = grounds.id
    LEFT JOIN bowling_performance bp ON matches.id = bp.match_id
    LEFT JOIN match_stats ms ON matches.id = ms.match_id
    LEFT JOIN best_innings_fig_runs bifr ON matches.id = bifr.match_id
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
GROUP BY matches.id,
    teams.name,
    teams2.name,
    grounds.name,
    ms.matches_count,
    ms.players_count,
    bifr.best_innings_fig_runs,
    bp.best_innings_fig_wkts,
    bp.four_wkt_hauls,
    bp.five_wkt_hauls
ORDER BY SUM(bs.wickets_taken) DESC;