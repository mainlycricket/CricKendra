SELECT continents.id AS continent_id,
    continents.name AS continent_name,
    COUNT(DISTINCT mse.player_id) AS players_count,
    MIN(matches.start_date) AS min_date,
    MAX(matches.start_date) AS max_date,
    COUNT(DISTINCT matches.id) AS matches_played,
    COUNT(innings.id) AS innings_count,
    SUM(bs.runs_scored) AS runs_scored,
    SUM(bs.balls_faced) AS balls_faced,
    COUNT(
        CASE
            WHEN bs.dismissal_type IS NULL
            OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
        END
    ) AS not_outs,
    (
        CASE
            WHEN COUNT(
                CASE
                    WHEN bs.dismissal_type IS NOT NULL
                    AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
                END
            ) > 0 THEN SUM(bs.runs_scored) * 1.0 / COUNT(
                CASE
                    WHEN bs.dismissal_type IS NOT NULL
                    AND bs.dismissal_type NOT IN ('retired hurt', 'retired not out') THEN 1
                END
            )
        END
    ) AS average,
    (
        CASE
            WHEN SUM(bs.balls_faced) > 0 THEN SUM(bs.runs_scored) * 100.0 / SUM(bs.balls_faced)
        END
    ) AS strike_rate,
    MAX(bs.runs_scored) AS highest_score,
    MAX(
        CASE
            WHEN bs.dismissal_type IS NULL
            OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN runs_scored
            ELSE 0
        END
    ) as highest_not_out_score,
    COUNT(
        CASE
            WHEN bs.runs_scored >= 100 THEN 1
        END
    ) AS centuries,
    COUNT(
        CASE
            WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
        END
    ) AS half_centuries,
    COUNT(
        CASE
            WHEN bs.runs_scored >= 100 THEN 1
        END
    ) + COUNT(
        CASE
            WHEN bs.runs_scored BETWEEN 50 AND 99 THEN 1
        END
    ) AS fifty_plus_scores,
    COUNT(
        CASE
            WHEN bs.runs_scored = 0 THEN 1
        END
    ) AS ducks,
    SUM(bs.fours_scored) AS fours_scored,
    SUM(bs.sixes_scored) AS sixes_scored
FROM matches
    LEFT JOIN innings ON innings.match_id = matches.id
    LEFT JOIN batting_scorecards bs ON bs.innings_id = innings.id
    LEFT JOIN match_squad_entries mse ON mse.match_id = matches.id
    AND mse.team_id = innings.batting_team_id
    AND mse.player_id = bs.batter_id
    AND mse.playing_status IN ('playing_xi')
    LEFT JOIN grounds ON matches.ground_id = grounds.id
    LEFT JOIN cities ON grounds.city_id = cities.id
    LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id
    LEFT JOIN continents ON host_nations.continent_id = continents.id
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
GROUP BY continents.id,
    continents.name
ORDER BY runs_scored DESC;