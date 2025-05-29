package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertMatch(ctx context.Context, db *pgxpool.Pool, match *models.Match) (int64, error) {
	var matchId int64

	tx, err := db.Begin(ctx)
	if err != nil {
		return matchId, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	query := `INSERT INTO matches (start_date, start_datetime_utc, team1_id, team2_id, is_male, main_series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number, end_date, match_state, match_state_description) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36) RETURNING id`

	err = db.QueryRow(ctx, query, match.StartDate, match.StartDateTimeUtc, match.Team1Id, match.Team2Id, match.IsMale, match.MainSeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate, match.MatchState, match.MatchStateDescription).Scan(&matchId)

	if len(match.SeriesListId) > 0 {
		if err = UpsertMatchSeriesEntries(ctx, tx, matchId, match.SeriesListId); err != nil {
			return matchId, err
		}
	}

	return matchId, err
}

// used for insert and update match both
func UpsertMatchSeriesEntries(ctx context.Context, db DB_Exec, matchId int64, seriesListId []pgtype.Int8) error {
	query := `INSERT INTO match_series_entries (match_id, series_id) VALUES ($1, $2) ON CONFLICT (match_id, series_id) DO NOTHING`

	batch := &pgx.Batch{}
	_ = batch.Queue(`DELETE FROM match_series_entries WHERE match_id = $1`, matchId)

	for _, seriesId := range seriesListId {
		batch.Queue(query, matchId, seriesId)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
}

func UpdateMatchTossDecisionById(ctx context.Context, db DB_Exec, input *models.TossDecisionInput) error {
	query := `
		UPDATE matches SET
			toss_winner_team_id = $1, toss_loser_team_id = $2, is_toss_decision_bat = $3
		WHERE matches.id = $4`

	cmd, err := db.Exec(ctx, query, input.TossWinnerId, input.TossLoserId, input.IsTossDecisionBat, input.MatchId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return errors.New("failed to update toss decision")
	}

	return nil
}

func UpdateMatchResultById(ctx context.Context, db DB_Exec, input *models.MatchResultInput) error {
	query := `
		UPDATE matches SET
			final_result = $1, match_winner_team_id = $2, match_loser_team_id = $3,
			bowl_out_winner_id = $4, super_over_winner_id = $5, is_won_by_innings = $6,
			is_won_by_runs = $7, win_margin = $8, balls_remaining_after_win = $9,
			outcome_special_method = $10
		WHERE id = $11`

	cmd, err := db.Exec(ctx, query, input.FinalResult, input.MatchWinnerId, input.MatchLoserId, input.BowlOutWinnerId, input.SuperOverWinnerId, input.IsWonByInnings, input.IsWonByRuns, input.WinMargin, input.BallsMargin, input.OutcomeSpecialMethod, input.MatchId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return errors.New("failed to update match result")
	}

	return nil
}

func UpdateMatchStateById(ctx context.Context, db DB_Exec, input *models.MatchStateInput) error {
	batch := pgx.Batch{}
	_ = batch.Queue(`UPDATE matches SET match_state = $1 WHERE id = $2`, input.State, input.MatchId)

	// set match bbb done & player career stat updates
	if input.State.String == "completed" {
		_ = batch.Queue(`UPDATE matches SET is_bbb_done = TRUE WHERE id = $1`, input.MatchId)

		statsUpdateQuery := `
			WITH
				match_wickets AS (
					SELECT
						bs.bowler_id,
						matches.id AS match_id,
						matches.playing_format,
						SUM(bs.wickets_taken) AS total_wickets,
						SUM(bs.runs_conceded) AS total_runs
					FROM
						matches
						LEFT JOIN innings ON innings.match_id = matches.id
						AND innings.innings_number IS NOT NULL
						AND innings.is_super_over = FALSE
						LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
					WHERE
						matches.id = $1
					GROUP BY
						matches.id,
						matches.playing_format,
						bs.bowler_id
				),
				ten_wicket_hauls AS (
					SELECT
						mw.bowler_id,
						mw.playing_format,
						COUNT(*) AS hauls_count
					FROM
						match_wickets mw
					WHERE
						mw.total_wickets >= 10
					GROUP BY
						mw.bowler_id,
						mw.playing_format
				),
				best_bowling_match AS (
					SELECT DISTINCT
						ON (mw.bowler_id, mw.playing_format) mw.bowler_id,
						mw.playing_format,
						mw.total_wickets AS wickets,
						mw.total_runs AS runs
					FROM
						match_wickets mw
					ORDER BY
						mw.bowler_id,
						mw.playing_format,
						mw.total_wickets DESC,
						mw.total_runs ASC
				),
				best_bowling_innings AS (
					SELECT
						bs.bowler_id,
						matches.playing_format,
						MAX(bs.wickets_taken) AS wickets
					FROM
						matches
						LEFT JOIN innings ON innings.match_id = matches.id
						AND innings.innings_number IS NOT NULL
						AND innings.is_super_over = FALSE
						LEFT JOIN bowling_scorecards bs ON bs.innings_id = innings.id
					WHERE
						matches.id = $1
					GROUP BY
						bs.bowler_id,
						matches.playing_format
				),
				best_bowling_figures AS (
					SELECT
						bbm.bowler_id,
						bbm.playing_format,
						twh.hauls_count AS ten_wicket_hauls,
						bbm.wickets AS best_match_wickets,
						bbm.runs AS best_match_runs,
						bbi.wickets AS best_innings_wickets
					FROM
						best_bowling_match bbm
						LEFT JOIN ten_wicket_hauls twh ON twh.bowler_id = bbm.bowler_id
						LEFT JOIN best_bowling_innings bbi ON bbi.bowler_id = bbm.bowler_id
						AND bbi.playing_format = bbm.playing_format
				),
				match_stats AS (
					-- Get batting and bowling stats per player per match
					SELECT
						m.playing_format,
						mse.player_id,
						-- stats
						ROW (
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
							bbf.ten_wicket_hauls,
							MIN(
								CASE
									WHEN bbf.best_innings_wickets = bws.wickets_taken THEN bws.runs_conceded
								END
							),
							bbf.best_innings_wickets,
							bbf.best_match_runs,
							bbf.best_match_wickets
						)::career_stats AS stats
					FROM
						matches m
						JOIN match_squad_entries mse ON mse.match_id = m.id
						JOIN innings i ON i.match_id = m.id
						LEFT JOIN batting_scorecards bs ON bs.innings_id = i.id
						AND mse.player_id = bs.batter_id
						LEFT JOIN bowling_scorecards bws ON bws.innings_id = i.id
						AND mse.player_id = bws.bowler_id
						LEFT JOIN best_bowling_figures bbf ON bbf.bowler_id = mse.player_id
						AND bbf.playing_format = m.playing_format
					WHERE
						m.id = $1
						AND i.is_super_over = FALSE
						AND i.innings_number IS NOT NULL
					GROUP BY
						m.playing_format,
						mse.player_id,
						bbf.ten_wicket_hauls,
						bbf.best_innings_wickets,
						bbf.best_match_wickets,
						bbf.best_match_runs
				)
			UPDATE players p
			SET -- Update individual format stats
				db_test_stats = (
					SELECT
						combine_career_stats (
							(
								SELECT
									ms.stats
								FROM
									match_stats ms
								WHERE
									ms.player_id = p.id
									AND ms.playing_format = 'Test'
							),
							db_test_stats
						)
				),
				db_odi_stats = (
					SELECT
						combine_career_stats (
							(
								SELECT
									ms.stats
								FROM
									match_stats ms
								WHERE
									ms.player_id = p.id
									AND ms.playing_format = 'ODI'
							),
							db_odi_stats
						)
				),
				db_t20i_stats = (
					SELECT
						combine_career_stats (
							(
								SELECT
									ms.stats
								FROM
									match_stats ms
								WHERE
									ms.player_id = p.id
									AND ms.playing_format = 'T20I'
							),
							db_t20i_stats
						)
				),
				db_fc_stats = (
					SELECT
						combine_career_stats (
							(
								SELECT
									ms.stats
								FROM
									match_stats ms
								WHERE
									ms.player_id = p.id
									AND (
										ms.playing_format = 'first_class'
										OR ms.playing_format = 'Test'
									)
							),
							db_fc_stats
						)
				),
				db_lista_stats = (
					SELECT
						combine_career_stats (
							(
								SELECT
									ms.stats
								FROM
									match_stats ms
								WHERE
									ms.player_id = p.id
									AND (
										ms.playing_format = 'list_a'
										OR ms.playing_format = 'ODI'
									)
							),
							db_lista_stats
						)
				),
				db_t20_stats = (
					SELECT
						combine_career_stats (
							(
								SELECT
									ms.stats
								FROM
									match_stats ms
								WHERE
									ms.player_id = p.id
									AND (
										ms.playing_format = 'T20'
										OR ms.playing_format = 'T20I'
									)
							),
							db_t20_stats
						)
				)
			WHERE
				EXISTS (
					SELECT
						1
					FROM
						match_stats ms
					WHERE
						ms.player_id = p.id
				);
		`

		_ = batch.Queue(statsUpdateQuery, input.MatchId)
	}

	return db.SendBatch(ctx, &batch).Close()
}

var matchInfoQuery = struct {
	selectFields  string
	joins         string
	groupByFields string
}{
	selectFields: `
		matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.event_match_number,
		matches.match_state, matches.match_state_description,
		
		matches.match_winner_team_id, matches.match_loser_team_id, matches.is_won_by_innings,
		matches.is_won_by_runs, matches.win_margin, matches.balls_remaining_after_win,
		matches.super_over_winner_id, matches.bowl_out_winner_id, matches.outcome_special_method,
		matches.toss_winner_team_id, matches.toss_loser_team_id, matches.is_toss_decision_bat,
		
		matches.season, matches.start_date, matches.end_date, matches.start_datetime_utc, matches.is_day_night, matches.ground_id, grounds.name, matches.main_series_id, main_series.name,

		matches.team1_id, team1.name, team1.image_url, matches.team2_id, team2.name, team2.image_url,

		(
			SELECT
				ARRAY_AGG (
					-- order is necessary for struct scanning
					ROW (
						innings.id, innings.innings_number, innings.batting_team_id, batting_team.name,
						
						innings.total_runs, innings.total_balls, innings.total_wickets, innings.innings_end, innings.target_runs, innings.max_overs
					)
				)
		) AS team_innings_short_info
	`,

	joins: `
		LEFT JOIN teams team1 ON matches.team1_id = team1.id 
		LEFT JOIN teams team2 ON matches.team2_id = team2.id

		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN series main_series ON matches.main_series_id = main_series.id

		LEFT JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE

		LEFT JOIN teams batting_team ON innings.batting_team_id = batting_team.id
	`,

	groupByFields: `
		matches.id, grounds.name,
		team1.id, team2.name, team1.image_url,
		team2.id, team2.name, team2.image_url,
		main_series.id, main_series.name
	`,
}

var matchHeaderQuery = struct {
	selectFields  string
	joins         string
	groupByFields string
}{
	selectFields: fmt.Sprintf(`
		%s,
		(
			SELECT
				ARRAY_AGG (
					-- order is necessary for struct scanning
					ROW (
						player_awards.player_id, players.name, player_awards.award_type
					)
				)
			FROM
				player_awards
				LEFT JOIN players ON player_awards.player_id = players.id
			WHERE
				player_awards.match_id = matches.id
		) AS match_awards
	`,
		matchInfoQuery.selectFields,
	),

	joins: matchInfoQuery.joins,

	groupByFields: matchInfoQuery.groupByFields,
}

func ReadMatches(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllMatchesResponse, error) {
	var response responses.AllMatchesResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "matches",
		DefaultLimit: 20,
		DefaultSort:  []string{"-start_date"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Match](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`
		SELECT %s
		FROM matches 
		%s
		%s
		GROUP BY %s
		%s
		%s`,
		matchInfoQuery.selectFields,
		matchInfoQuery.joins,
		queryInfoOutput.WhereClause,
		matchInfoQuery.groupByFields,
		queryInfoOutput.OrderByClause,
		queryInfoOutput.PaginationClause,
	)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.MatchInfo, error) {
		var match responses.MatchInfo

		err := rows.Scan(
			&match.MatchId, &match.PlayingLevel, &match.PlayingFormat, &match.MatchType, &match.EventMatchNumber,

			&match.MatchState, &match.MatchStateDescription,

			&match.MatchWinnerId, &match.MatchLoserId, &match.IsWonByInnings, &match.IsWonByRuns,
			&match.WinMargin, &match.BallsMargin, &match.SuperOverWinnerId, &match.BowlOutWinnerId, &match.OutcomeSpecialMethod, &match.TossWinnerId, &match.TossLoserId, &match.IsTossDecisionBat,

			&match.Season, &match.StartDate, &match.EndDate, &match.StartDateTimeUtc, &match.IsDayNight, &match.GroundId, &match.GroundName, &match.MainSeriesId, &match.MainSeriesName,

			&match.Team1Id, &match.Team1Name, &match.Team1ImageUrl, &match.Team2Id, &match.Team2Name, &match.Team2ImageUrl,

			&match.InningsScores,
		)

		return match, err
	})

	if len(matches) > queryInfoOutput.RecordsCount {
		response.Matches = matches[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Matches = matches
		response.Next = false
	}

	return response, err
}

func ReadMatchSummary(ctx context.Context, db DB_Exec, matchId int64) (responses.MatchSummary, error) {
	var response responses.MatchSummary
	matchHeader := &response.MatchHeader

	query := fmt.Sprintf(`
		WITH
			top_batters AS (
				SELECT
					bs.innings_id, bs.batter_id, batter.name AS batter_name,
					
					bs.runs_scored, bs.balls_faced, bs.fours_scored, bs.sixes_scored,

					ROW_NUMBER() OVER (
						PARTITION BY
							bs.innings_id
						ORDER BY
							(
								CASE
									WHEN bs.batter_id = innings.striker_id THEN 2
									WHEN bs.batter_id = innings.non_striker_id THEN 1
									ELSE 0
								END
							) DESC,
							bs.runs_scored DESC, bs.balls_faced ASC,
							bs.dismissal_type IS NULL, bs.sixes_scored, bs.fours_scored, bs.batting_position ASC
					) AS rn
				FROM
					batting_scorecards bs
					LEFT JOIN innings ON bs.innings_id = innings.id
					AND innings.innings_number IS NOT NULL
					AND innings.is_super_over = FALSE
					LEFT JOIN players batter ON bs.batter_id = batter.id
				WHERE
					innings.match_id = $1
			),
			aggregated_top_batters AS (
				SELECT
					top_batters.innings_id,
					ARRAY_AGG (
						ROW (
							top_batters.batter_id, top_batters.batter_name,
							
							top_batters.runs_scored, top_batters.balls_faced,

							top_batters.fours_scored, top_batters.sixes_scored
						)
					) AS batters
				FROM
					top_batters
				WHERE
					top_batters.rn <= 2
				GROUP BY
					top_batters.innings_id
			),
			top_bowlers AS (
				SELECT
					bs.innings_id, bs.bowler_id, bowler.name AS bowler_name,

					bs.balls_bowled / matches.balls_per_over + bs.balls_bowled %% matches.balls_per_over * 0.1 AS overs_bowled,
					bs.maiden_overs,
					bs.runs_conceded, bs.wickets_taken,

					ROW_NUMBER() OVER (
						PARTITION BY
							bs.innings_id
						ORDER BY
							(
								CASE
									WHEN bs.bowler_id = innings.bowler1_id THEN 2
									WHEN bs.bowler_id = innings.bowler2_id THEN 1
									ELSE 0
								END
							) DESC,
							bs.wickets_taken DESC, bs.runs_conceded ASC,
							bs.balls_bowled / matches.balls_per_over + bs.balls_bowled %% matches.balls_per_over * 0.1 DESC, bs.maiden_overs, bs.bowling_position ASC
					) AS rn
				FROM
					bowling_scorecards bs

					LEFT JOIN innings ON bs.innings_id = innings.id
					AND innings.innings_number IS NOT NULL
					AND innings.is_super_over = FALSE

					LEFT JOIN matches ON innings.match_id = matches.id
					
					LEFT JOIN players bowler ON bs.bowler_id = bowler.id
				WHERE
					innings.match_id = $1
			),
			aggregated_top_bowlers AS (
				SELECT
					top_bowlers.innings_id,
					ARRAY_AGG (
						ROW (
							top_bowlers.bowler_id, top_bowlers.bowler_name,
							
							top_bowlers.overs_bowled, top_bowlers.maiden_overs, top_bowlers.wickets_taken, top_bowlers.runs_conceded
						)
					) AS bowlers
				FROM
					top_bowlers
				WHERE
					top_bowlers.rn <= 2
				GROUP BY
					top_bowlers.innings_id
			),
			scorecard_summary AS (
				SELECT
					innings.id, innings.innings_number, innings.batting_team_id, batting_team.name,
					
					innings.total_runs, innings.total_wickets, innings.total_balls,
					
					ARRAY_AGG (aggregated_top_batters.batters),
					ARRAY_AGG (aggregated_top_bowlers.bowlers)
				FROM
					innings
					LEFT JOIN teams batting_team ON innings.batting_team_id = batting_team.id
					LEFT JOIN aggregated_top_batters ON aggregated_top_batters.innings_id = innings.id
					LEFT JOIN aggregated_top_bowlers ON aggregated_top_bowlers.innings_id = innings.id
				WHERE
					innings.match_id = $1
					AND innings.innings_number IS NOT NULL
					AND innings.is_super_over = FALSE
				GROUP BY
					innings.id, innings.innings_number, innings.batting_team_id, batting_team.name,
					
					innings.total_runs, innings.total_wickets, innings.total_balls
			),
			latest_commentary AS (
				SELECT
					deliveries.innings_id, deliveries.innings_delivery_number, deliveries.ball_number, deliveries.over_number,

					deliveries.batter_id, batter.name, deliveries.bowler_id, bowler.name, deliveries.fielder1_id, fielder1.name, deliveries.fielder2_Id, fielder2.name,
					
					deliveries.wides, deliveries.noballs, deliveries.legbyes, deliveries.byes, deliveries.total_runs, deliveries.is_four, deliveries.is_six,
					
					deliveries.player1_dismissed_id, player1_dismissed.name, deliveries.player1_dismissal_type, bs1.runs_scored, bs1.balls_faced, bs1.fours_scored, bs1.sixes_scored,

					deliveries.player2_dismissed_id, player2_dismissed.name, deliveries.player2_dismissal_type, bs2.runs_scored, bs2.balls_faced, bs2.fours_scored, bs2.sixes_scored,

					deliveries.commentary
				FROM
					deliveries
					
					LEFT JOIN innings ON deliveries.innings_id = innings.id
					AND innings.is_super_over = FALSE
					AND innings.innings_number IS NOT NULL
					
					LEFT JOIN matches ON innings.match_id = matches.id
					
					LEFT JOIN players batter ON deliveries.batter_id = batter.id
					LEFT JOIN players bowler ON deliveries.bowler_id = bowler.id
					LEFT JOIN players fielder1 ON deliveries.fielder1_id = fielder1.id
					LEFT JOIN players fielder2 ON deliveries.fielder2_id = fielder2.id
					LEFT JOIN players player1_dismissed ON deliveries.player1_dismissed_id = player1_dismissed.id
					LEFT JOIN players player2_dismissed ON deliveries.player2_dismissed_id = player2_dismissed.id

					LEFT JOIN batting_scorecards bs1 ON bs1.innings_id = deliveries.innings_id
					AND bs1.batter_id = deliveries.player1_dismissed_id
					LEFT JOIN batting_scorecards bs2 ON bs2.innings_id = deliveries.innings_id
					AND bs2.batter_id = deliveries.player2_dismissed_id
				WHERE
					matches.id = $1
				ORDER BY
					innings.innings_number DESC,
					deliveries.innings_delivery_number DESC
				FETCH FIRST
					50 ROWS ONLY
			)
		SELECT
			%s,
			(
				SELECT
					ARRAY_AGG (scorecard_summary.*)
				FROM
					scorecard_summary
			) AS scorecard_summary,
			(
				SELECT
					ARRAY_AGG (latest_commentary.*)
				FROM
					latest_commentary
			) AS latest_commentary
		FROM matches
		%s
		WHERE matches.id = $1
		GROUP BY %s;
	`, matchHeaderQuery.selectFields, matchHeaderQuery.joins, matchHeaderQuery.groupByFields)

	row := db.QueryRow(ctx, query, matchId)

	err := row.Scan(
		&matchHeader.MatchId, &matchHeader.PlayingLevel, &matchHeader.PlayingFormat, &matchHeader.MatchType, &matchHeader.EventMatchNumber,

		&matchHeader.MatchState, &matchHeader.MatchStateDescription,

		&matchHeader.MatchWinnerId, &matchHeader.MatchLoserId, &matchHeader.IsWonByInnings, &matchHeader.IsWonByRuns,
		&matchHeader.WinMargin, &matchHeader.BallsMargin, &matchHeader.SuperOverWinnerId, &matchHeader.BowlOutWinnerId, &matchHeader.OutcomeSpecialMethod, &matchHeader.TossWinnerId, &matchHeader.TossLoserId, &matchHeader.IsTossDecisionBat,

		&matchHeader.Season, &matchHeader.StartDate, &matchHeader.EndDate, &matchHeader.StartDateTimeUtc, &matchHeader.IsDayNight, &matchHeader.GroundId, &matchHeader.GroundName, &matchHeader.MainSeriesId, &matchHeader.MainSeriesName,

		&matchHeader.Team1Id, &matchHeader.Team1Name, &matchHeader.Team1ImageUrl, &matchHeader.Team2Id, &matchHeader.Team2Name, &matchHeader.Team2ImageUrl,

		&matchHeader.InningsScores,
		&matchHeader.PlayerAwards,

		&response.ScorecardSummary,
		&response.LatestCommentary,
	)

	return response, err
}

func ReadMatchFullScorecard(ctx context.Context, db DB_Exec, matchId int64) (responses.MatchScorecardResponse, error) {
	var response responses.MatchScorecardResponse
	matchHeader := &response.MatchHeader

	query := fmt.Sprintf(`
		SELECT
			%s,
			(
				SELECT
					ARRAY_AGG (
						-- order is necessary for struct scanning
						ROW (
							innings.id, innings.innings_number, innings.batting_team_id, batting_team.name,

							innings.total_runs, innings.total_balls, innings.total_wickets, innings.byes, innings.leg_byes, innings.wides, innings.noballs, innings.penalty,

							innings.innings_end, innings.target_runs, innings.max_overs,

							(
								SELECT
									ARRAY_AGG (
										-- order is necessary for struct scanning
										ROW (
											bs.batter_id, batter.name, bs.batting_position, bs.has_batted,

											bs.runs_scored, bs.balls_faced, bs.minutes_batted, bs.fours_scored, bs.sixes_scored,

											bs.dismissal_type, bs.dismissed_by_id, dismissed_by.name, bs.fielder1_id, fielder1.name, bs.fielder2_id, fielder2.name
										)
									)
								FROM
									batting_scorecards bs
									LEFT JOIN players batter ON bs.batter_id = batter.id
									LEFT JOIN players dismissed_by ON bs.dismissed_by_id = dismissed_by.id
									LEFT JOIN players fielder1 ON bs.fielder1_id = fielder1.id
									LEFT JOIN players fielder2 ON bs.fielder2_id = fielder2.id
								WHERE
									bs.innings_id = innings.id
							),

							(
								SELECT
									ARRAY_AGG (
										-- order is necessary for struct scanning
										ROW (
											bs.bowler_id, bowler.name, bs.bowling_position,

											bs.wickets_taken, bs.runs_conceded,
											
											bs.balls_bowled / matches.balls_per_over + bs.balls_bowled %% matches.balls_per_over * 0.1,
											
											bs.maiden_overs, bs.fours_conceded, bs.sixes_conceded, bs.wides_conceded, bs.noballs_conceded
										)
									)
								FROM
									bowling_scorecards bs

									LEFT JOIN players bowler ON bs.bowler_id = bowler.id

								WHERE
									bs.innings_id = innings.id AND
									bs.bowling_position IS NOT NULL
							),

							(
								SELECT
									ARRAY_AGG (
										-- order is necessary for struct scanning
										ROW (
											fow.batter_id, batter.name, fow.ball_number,
											fow.team_runs, fow.wicket_number, fow.dismissal_type
										)
									)
								FROM
									fall_of_wickets fow
									LEFT JOIN players batter ON fow.batter_id = batter.id
								WHERE
									fow.innings_id = innings.id
							
							)
						)
					)
			) AS match_innings
		
		FROM matches 
		%s
		WHERE matches.id = $1
		GROUP BY %s
	`,
		matchHeaderQuery.selectFields,
		matchHeaderQuery.joins,
		matchHeaderQuery.groupByFields,
	)

	row := db.QueryRow(ctx, query, matchId)

	err := row.Scan(
		&matchHeader.MatchId, &matchHeader.PlayingLevel, &matchHeader.PlayingFormat, &matchHeader.MatchType, &matchHeader.EventMatchNumber,

		&matchHeader.MatchState, &matchHeader.MatchStateDescription,

		&matchHeader.MatchWinnerId, &matchHeader.MatchLoserId, &matchHeader.IsWonByInnings, &matchHeader.IsWonByRuns,
		&matchHeader.WinMargin, &matchHeader.BallsMargin, &matchHeader.SuperOverWinnerId, &matchHeader.BowlOutWinnerId, &matchHeader.OutcomeSpecialMethod, &matchHeader.TossWinnerId, &matchHeader.TossLoserId, &matchHeader.IsTossDecisionBat,

		&matchHeader.Season, &matchHeader.StartDate, &matchHeader.EndDate, &matchHeader.StartDateTimeUtc, &matchHeader.IsDayNight, &matchHeader.GroundId, &matchHeader.GroundName, &matchHeader.MainSeriesId, &matchHeader.MainSeriesName,

		&matchHeader.Team1Id, &matchHeader.Team1Name, &matchHeader.Team1ImageUrl, &matchHeader.Team2Id, &matchHeader.Team2Name, &matchHeader.Team2ImageUrl,

		&matchHeader.InningsScores,
		&matchHeader.PlayerAwards,

		&response.InningsScorecards,
	)

	if err != nil {
		return response, err
	}

	return response, nil
}

func ReadMatchStats(ctx context.Context, db DB_Exec, matchId int64) (responses.MatchStatsResponse, error) {
	var response responses.MatchStatsResponse
	matchHeader := &response.MatchHeader

	query := fmt.Sprintf(`
		WITH overs_data AS (
			SELECT innings_id, over_number, SUM(deliveries.total_runs) AS runs,
				COUNT (
					DISTINCT CASE 
						WHEN deliveries.wides = 0 AND deliveries.noballs = 0 THEN deliveries.innings_delivery_number
					END
				) AS balls,
				COUNT(player1_dismissed_id) + COUNT(player2_dismissed_id) AS wickets
			FROM deliveries
			JOIN innings ON deliveries.innings_id = innings.id
				AND innings.innings_number IS NOT NULL
				AND innings.is_super_over = FALSE
			WHERE innings.match_id = $1
			GROUP BY innings_id, over_number
			ORDER BY innings_id ASC, over_number ASC
		)
		SELECT
			%s,
			(
				SELECT
					ARRAY_AGG (
						-- order is necessary for struct scanning
						ROW (
							innings.id, innings.innings_number, innings.batting_team_id, batting_team.name,

							(
								SELECT
									ARRAY_AGG (
										-- order is necessary for struct scanning
										ROW (
											bp.wicket_number, bp.is_unbeaten, bp.start_ball_number, bp.end_ball_number, bp.end_team_runs - bp.start_team_runs,

											bp.batter1_id, batter1.name, bp.batter1_runs, bp.batter1_balls,

											bp.batter2_id, batter2.name, bp.batter2_runs, bp.batter2_balls
										)
									)
								FROM
									batting_partnerships bp
								LEFT JOIN players batter1 ON bp.batter1_id = batter1.id
								LEFT JOIN players batter2 ON bp.batter2_id = batter2.id
									
								WHERE
									bp.innings_id = innings.id
							),

							(
								SELECT
									ARRAY_AGG (
										-- order is necessary for struct scanning
										ROW (
											over_number, runs, balls, wickets
										)
									)
								FROM
									overs_data
								WHERE
									overs_data.innings_id = innings.id
							)
						)
					)
			) AS match_innings
		
		FROM matches 
		%s
		WHERE matches.id = $1
		GROUP BY %s
	`,
		matchHeaderQuery.selectFields,
		matchHeaderQuery.joins,
		matchHeaderQuery.groupByFields,
	)

	row := db.QueryRow(ctx, query, matchId)

	err := row.Scan(
		&matchHeader.MatchId, &matchHeader.PlayingLevel, &matchHeader.PlayingFormat, &matchHeader.MatchType, &matchHeader.EventMatchNumber,

		&matchHeader.MatchState, &matchHeader.MatchStateDescription,

		&matchHeader.MatchWinnerId, &matchHeader.MatchLoserId, &matchHeader.IsWonByInnings, &matchHeader.IsWonByRuns,
		&matchHeader.WinMargin, &matchHeader.BallsMargin, &matchHeader.SuperOverWinnerId, &matchHeader.BowlOutWinnerId, &matchHeader.OutcomeSpecialMethod, &matchHeader.TossWinnerId, &matchHeader.TossLoserId, &matchHeader.IsTossDecisionBat,

		&matchHeader.Season, &matchHeader.StartDate, &matchHeader.EndDate, &matchHeader.StartDateTimeUtc, &matchHeader.IsDayNight, &matchHeader.GroundId, &matchHeader.GroundName, &matchHeader.MainSeriesId, &matchHeader.MainSeriesName,

		&matchHeader.Team1Id, &matchHeader.Team1Name, &matchHeader.Team1ImageUrl, &matchHeader.Team2Id, &matchHeader.Team2Name, &matchHeader.Team2ImageUrl,

		&matchHeader.InningsScores,
		&matchHeader.PlayerAwards,

		&response.Innings,
	)

	if err != nil {
		return response, err
	}

	return response, nil
}

func UpdateMatch(ctx context.Context, db *pgxpool.Pool, match *models.Match) error {
	query := `UPDATE matches SET start_date = $1, start_datetime_utc = $2, team1_id = $3, team2_id = $4, is_male = $5, main_series_id = $6, ground_id = $7, is_neutral_venue = $8, current_status = $9, final_result = $10, home_team_id = $11, away_team_id = $12, match_type = $13, playing_level = $14, playing_format = $15, season = $16, is_day_night = $17, outcome_special_method = $18, toss_winner_team_id = $19, toss_loser_team_id = $20, is_toss_decision_bat = $21, match_winner_team_id = $22, match_loser_team_id = $23, bowl_out_winner_id = $24, super_over_winner_id = $25, is_won_by_innings = $26, is_won_by_runs = $27, win_margin = $28, balls_remaining_after_win = $29, balls_per_over = $30, cricsheet_id = $31, is_bbb_done = $32, event_match_number = $33, end_date = $34, match_state = $35, match_state_description = $36 WHERE id = $37`

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	cmd, err := tx.Exec(ctx, query, match.StartDate, match.StartDateTimeUtc, match.Team1Id, match.Team2Id, match.IsMale, match.MainSeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate, match.MatchState, match.MatchStateDescription, match.Id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to update match")
	}

	if err = UpsertMatchSeriesEntries(ctx, tx, match.Id.Int64, match.SeriesListId); err != nil {
		return err
	}

	return nil
}

func UpsertCricsheetMatch(ctx context.Context, db *pgxpool.Pool, match *models.Match) (int64, error) {
	var matchId int64

	query := `
	INSERT INTO matches (start_date, start_datetime_utc, team1_id, team2_id, is_male, main_series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number, end_date, match_state, match_state_description) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36)
	ON CONFLICT (cricsheet_id)
	DO UPDATE SET start_date = $1, start_datetime_utc = $2, team1_id = $3, team2_id = $4, is_male = $5, main_series_id = $6, ground_id = $7, is_neutral_venue = $8, current_status = $9, final_result = $10, home_team_id = $11, away_team_id = $12, match_type = $13, playing_level = $14, playing_format = $15, season = $16, is_day_night = $17, outcome_special_method = $18, toss_winner_team_id = $19, toss_loser_team_id = $20, is_toss_decision_bat = $21, match_winner_team_id = $22, match_loser_team_id = $23, bowl_out_winner_id = $24, super_over_winner_id = $25, is_won_by_innings = $26, is_won_by_runs = $27, win_margin = $28, balls_remaining_after_win = $29, balls_per_over = $30, cricsheet_id = $31, is_bbb_done = $32, event_match_number = $33, end_date = $34, match_state = $35, match_state_description = $36
	RETURNING id
	`

	tx, err := db.Begin(ctx)
	if err != nil {
		return matchId, err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	err = tx.QueryRow(ctx, query, match.StartDate, match.StartDateTimeUtc, match.Team1Id, match.Team2Id, match.IsMale, match.MainSeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate, match.MatchState, match.MatchStateDescription).Scan(&matchId)

	if err != nil {
		return matchId, err
	}

	if err = UpsertMatchSeriesEntries(ctx, tx, matchId, match.SeriesListId); err != nil {
		return matchId, err
	}

	return matchId, err
}

func ReadMatchByCricsheetId(ctx context.Context, db DB_Exec, cricsheetId string) (models.Match, error) {
	var match models.Match

	query := `SELECT 
		id, cricsheet_id, event_match_number, start_date, start_datetime_utc, end_date, team1_id, team2_id, is_male,
	 
		ARRAY[mse.series_id], 
		
		main_series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, is_bbb_done, match_state, match_state_description 
		
		FROM matches 
		LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
		WHERE cricsheet_id = $1
	`

	row := db.QueryRow(ctx, query, cricsheetId)

	err := row.Scan(&match.Id, &match.CricsheetId, &match.EventMatchNumber, &match.StartDate, &match.StartDateTimeUtc, &match.EndDate, &match.Team1Id, &match.Team2Id, &match.IsMale, &match.SeriesListId, &match.MainSeriesId, &match.GroundId, &match.IsNeutralVenue, &match.CurrentStatus, &match.FinalResult, &match.HomeTeamId, &match.AwayTeamId, &match.MatchType, &match.PlayingLevel, &match.PlayingFormat, &match.Season, &match.IsDayNight, &match.OutcomeSpecialMethod, &match.TossWinnerId, &match.TossLoserId, &match.IsTossDecisionBat, &match.MatchWinnerId, &match.MatchLoserId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.BallsPerOver, &match.IsBBBDone, &match.MatchState, &match.MatchStateDescription)

	return match, err
}

func SetMatchBBBDone(ctx context.Context, db DB_Exec, matchId int64) error {
	query := `UPDATE matches SET is_bbb_done = true WHERE id = $1`
	cmd, err := db.Exec(ctx, query, matchId)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to set is_bbb_done")
	}

	return nil
}
