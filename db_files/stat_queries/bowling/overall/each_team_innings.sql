WITH best_innings AS (
    SELECT innings.id AS innings_id,
        MAX(bs.wickets_taken) AS max_wickets
    FROM matches
        LEFT JOIN innings ON innings.match_id = matches.id
        LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
    WHERE matches.playing_format = 'ODI'
        AND matches.ground_id IN (63, 90)
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
    GROUP BY innings.id
)
SELECT matches.id AS match_id,
    innings.innings_number,
    innings.bowling_team_id,
    teams1.name AS bowling_team_name,
    innings.batting_team_id,
    teams2.name AS batting_team_name,
    matches.season,
    cities.name AS city_name,
    matches.start_date,
    COUNT(DISTINCT mse.player_id) AS players_count,
    COUNT(DISTINCT matches.id) AS matches_played,
    COUNT(innings.id) AS innings_count,
    SUM(bs.balls_bowled) / 6 + (SUM(balls_bowled) % 6) * 0.1 AS overs_bowled,
    SUM(bs.runs_conceded) AS runs_conceded,
    SUM(bs.wickets_taken) AS wickets_taken,
    (
        CASE
            WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.runs_conceded) * 1.0 / SUM(bs.wickets_taken)
            ELSE NULL
        END
    ) AS average,
    (
        CASE
            WHEN SUM(bs.wickets_taken) > 0 THEN SUM(bs.balls_bowled) * 1.0 / SUM(bs.wickets_taken)
            ELSE NULL
        END
    ) AS strike_rate,
    (
        CASE
            WHEN SUM(bs.balls_bowled) > 0 THEN SUM(bs.runs_conceded) * 6.0 / SUM(bs.balls_bowled)
            ELSE NULL
        END
    ) AS economy,
    COUNT(
        CASE
            WHEN bs.wickets_taken = 4 THEN 1
        END
    ) AS four_wkt_hauls,
    COUNT(
        CASE
            WHEN bs.wickets_taken >= 5 THEN 1
        END
    ) AS five_wkt_hauls,
    MAX(bs.wickets_taken) AS best_innings_wickets,
    MIN(
        CASE
            WHEN bs.wickets_taken = bi.max_wickets THEN bs.runs_conceded
        END
    ) AS best_innings_runs,
    SUM(bs.fours_conceded) AS fours_conceded,
    SUM(bs.sixes_conceded) AS sixes_conceded
FROM matches
    LEFT JOIN innings ON innings.match_id = matches.id
    LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
    LEFT JOIN best_innings bi ON bi.innings_id = innings.id
    LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
    AND mse.team_id = innings.bowling_team_id
    AND mse.player_id = bs.bowler_id
    AND mse.playing_status IN ('playing_xi')
    LEFT JOIN grounds ON matches.ground_id = grounds.id
    LEFT JOIN cities ON grounds.city_id = cities.id
    LEFT JOIN teams teams1 ON innings.bowling_team_id = teams1.id
    LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
WHERE matches.playing_format = 'ODI'
    AND matches.ground_id IN (63, 90)
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
    innings.id,
    teams1.name,
    teams2.name,
    cities.name
ORDER BY SUM(bs.wickets_taken) DESC;