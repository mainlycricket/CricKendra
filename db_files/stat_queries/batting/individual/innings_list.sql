SELECT matches.id AS match_id,
    matches.start_date,
    matches.ground_id,
    cities.name AS city_name,
    innings.innings_number,
    bs.batter_id,
    players.name AS batter_name,
    innings.batting_team_id,
    teams.short_name as batting_team_name,
    innings.bowling_team_id,
    teams2.name AS bowling_team_name,
    bs.runs_scored,
    bs.balls_faced,
    (
        CASE
            WHEN bs.dismissal_type IS NULL
            OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN TRUE
            ELSE FALSE
        END
    ) AS is_not_out,
    (
        CASE
            WHEN SUM(bs.balls_faced) > 0 THEN SUM(bs.runs_scored) * 100.0 / SUM(bs.balls_faced)
            ELSE 0
        END
    ) AS strike_rate,
    SUM(bs.fours_scored) AS fours_scored,
    SUM(bs.sixes_scored) AS sixes_scored
FROM batting_scorecards bs
    LEFT JOIN innings ON innings.id = bs.innings_id
    LEFT JOIN teams ON innings.batting_team_id = teams.id
    LEFT JOIN teams teams2 ON innings.bowling_team_id = teams2.id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN players ON bs.batter_id = players.id
    LEFT JOIN grounds ON matches.ground_id = grounds.id
    LEFT JOIN cities ON grounds.city_id = cities.id
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
GROUP BY bs.innings_id,
    bs.batter_id,
    players.name,
    matches.ground_id,
    cities.name,
    innings.batting_team_id,
    innings.bowling_team_id,
    teams.short_name,
    teams2.name,
    innings.innings_number,
    matches.start_date,
    matches.id,
    bs.runs_scored,
    bs.balls_faced,
    bs.dismissal_type
ORDER BY bs.runs_scored DESC;