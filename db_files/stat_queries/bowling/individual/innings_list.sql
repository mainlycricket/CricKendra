SELECT matches.id AS match_id,
    matches.start_date,
    matches.ground_id,
    cities.name AS city_name,
    innings.innings_number,
    bs.bowler_id,
    players.name AS bowler_name,
    innings.batting_team_id AS bowling_team_id,
    teams.short_name as bowling_team_name,
    innings.bowling_team_id AS batting_team_id,
    teams2.name AS batting_team_name,
    SUM(bs.balls_bowled) / 6 + (SUM(balls_bowled) % 6) * 0.1 AS overs_bowled,
    SUM(bs.runs_conceded) AS runs_conceded,
    SUM(bs.wickets_taken) AS wickets_taken,
    (
        CASE
            WHEN SUM(bs.balls_bowled) > 0 THEN SUM(bs.runs_conceded) * 6.0 / SUM(bs.balls_bowled)
            ELSE NULL
        END
    ) AS economy,
    SUM(bs.fours_conceded) AS fours_conceded,
    SUM(bs.sixes_conceded) AS sixes_conceded
FROM bowling_scorecards bs
    LEFT JOIN innings ON innings.id = bs.innings_id
    LEFT JOIN teams ON innings.bowling_team_id = teams.id
    LEFT JOIN teams teams2 ON innings.batting_team_id = teams2.id
    LEFT JOIN matches ON innings.match_id = matches.id
    LEFT JOIN players ON bs.bowler_id = players.id
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
    bs.bowler_id,
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
    bs.runs_conceded,
    bs.balls_bowled,
    bs.wickets_taken
ORDER BY bs.wickets_taken DESC,
    bs.runs_conceded ASC;