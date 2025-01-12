-- fix domestic format stats
-- fix bowling match stats (ten_wkt_haul, best_match_fig_runs, best_match_fig_wkts)
WITH best_bowling_innings AS (
    -- Get best bowling figures per player per match format
    SELECT m.playing_format,
        mse.player_id,
        MAX(bs.wickets_taken) AS max_wickets
    FROM matches m
        JOIN match_squad_entries mse ON mse.match_id = m.id
        JOIN innings i ON i.match_id = m.id
        JOIN bowling_scorecards bs ON bs.innings_id = i.id
        AND bs.bowler_id = mse.player_id
    GROUP BY m.playing_format,
        mse.player_id
),
match_stats AS (
    -- Get batting and bowling stats per player per match
    SELECT m.playing_format,
        mse.player_id,
        -- stats
        ROW(
            COUNT(DISTINCT m.id),
            COUNT(
                DISTINCT CASE
                    WHEN bs.has_batted THEN i.id
                END
            ),
            COALESCE(SUM(bs.runs_scored), 0),
            COUNT(
                CASE
                    WHEN bs.dismissal_type IS NULL
                    OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN 1
                END
            ),
            COALESCE(SUM(bs.balls_faced), 0),
            COALESCE(SUM(bs.fours_scored), 0),
            COALESCE(SUM(bs.sixes_scored), 0),
            COUNT(
                CASE
                    WHEN bs.runs_scored >= 100 THEN 1
                END
            ),
            COUNT(
                CASE
                    WHEN bs.runs_scored >= 50
                    AND bs.runs_scored < 100 THEN 1
                END
            ),
            MAX(bs.runs_scored),
            CASE
                WHEN MAX(bs.runs_scored) = MAX(
                    CASE
                        WHEN bs.dismissal_type IS NULL
                        OR bs.dismissal_type IN ('retired hurt', 'retired not out') THEN bs.runs_scored
                    END
                ) THEN TRUE
                ELSE FALSE
            END,
            COUNT(
                DISTINCT CASE
                    WHEN bws.bowling_position IS NOT NULL THEN i.id
                END
            ),
            COALESCE(SUM(bws.runs_conceded), 0),
            COALESCE(SUM(bws.wickets_taken), 0),
            COALESCE(SUM(bws.balls_bowled), 0),
            COALESCE(SUM(bws.fours_conceded), 0),
            COALESCE(SUM(bws.sixes_conceded), 0),
            COUNT(
                CASE
                    WHEN bws.wickets_taken >= 4
                    AND bws.wickets_taken < 5 THEN 1
                END
            ),
            COUNT(
                CASE
                    WHEN bws.wickets_taken >= 5 THEN 1
                END
            ),
            COUNT(
                CASE
                    WHEN bws.wickets_taken >= 10 THEN 1
                END
            ),
            -- fix this should consider entire match
            MIN(
                CASE
                    WHEN bws.wickets_taken = bbi.max_wickets THEN bws.runs_conceded
                END
            ),
            MAX(bws.wickets_taken),
            MIN(
                CASE
                    WHEN bws.wickets_taken = bbi.max_wickets THEN bws.runs_conceded
                END
            ),
            MAX(bws.wickets_taken)
        )::career_stats AS stats
    FROM matches m
        JOIN match_squad_entries mse ON mse.match_id = m.id
        JOIN innings i ON i.match_id = m.id
        LEFT JOIN batting_scorecards bs ON bs.innings_id = i.id
        AND mse.player_id = bs.batter_id
        LEFT JOIN bowling_scorecards bws ON bws.innings_id = i.id
        AND mse.player_id = bws.bowler_id
        LEFT JOIN best_bowling_innings bbi ON bbi.player_id = mse.player_id
        AND bbi.playing_format = m.playing_format
    WHERE i.is_super_over = FALSE
        AND i.innings_number IS NOT NULL
    GROUP BY m.playing_format,
        mse.player_id
)
UPDATE players p
SET -- Update individual format stats
    db_odi_stats = (
        SELECT ms.stats
        FROM match_stats ms
        WHERE ms.player_id = p.id
            AND ms.playing_format = 'ODI'
    ),
    db_t20i_stats = (
        SELECT ms.stats
        FROM match_stats ms
        WHERE ms.player_id = p.id
            AND ms.playing_format = 'T20I'
    )
WHERE EXISTS (
        SELECT 1
        FROM match_stats ms
        WHERE ms.player_id = p.id
    );