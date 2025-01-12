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

	query := `INSERT INTO matches (start_date, start_time, team1_id, team2_id, is_male, main_series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34) RETURNING id`

	err = db.QueryRow(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.MainSeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate).Scan(&matchId)

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
	batch.Queue(`DELETE FROM match_series_entries WHERE match_id = $1`, matchId)

	for _, seriesId := range seriesListId {
		batch.Queue(query, matchId, seriesId)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
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
		SELECT
			matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.cricsheet_id, matches.balls_per_over, matches.event_match_number, matches.start_date, matches.end_date, matches.start_time, matches.is_day_night, matches.ground_id, grounds.name, matches.team1_id, t1.name, t1.image_url, matches.team2_id, t2.name, t2.image_url, matches.season, matches.main_series_id, ms.name, matches.current_status, matches.final_result, matches.outcome_special_method, matches.match_winner_team_id, matches.bowl_out_winner_id, matches.super_over_winner_id, matches.is_won_by_innings, matches.is_won_by_runs, matches.win_margin, matches.balls_remaining_after_win, matches.toss_winner_team_id, matches.is_toss_decision_bat,
			
			ARRAY_AGG ( ROW ( innings.innings_number, innings.batting_team_id, innings.total_runs, innings.total_balls, innings.total_wickets, innings.innings_end, innings.target_runs, innings.max_overs ) )
	
		FROM matches 
		
		LEFT JOIN teams t1 ON matches.team1_id = t1.id 
		LEFT JOIN teams t2 ON matches.team2_id = t2.id 
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN series ms ON matches.main_series_id = ms.id
		LEFT JOIN innings ON innings.match_id = matches.id AND innings.innings_number IS NOT NULL AND innings.is_super_over = FALSE
		
		GROUP BY matches.id, grounds.name, t1.name, t2.name, t1.image_url, t2.image_url, ms.name
		%s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllMatches, error) {
		var match responses.AllMatches

		err := rows.Scan(&match.Id, &match.PlayingLevel, &match.PlayingFormat, &match.MatchType, &match.CricsheetId, &match.BallsPerOver, &match.EventMatchNumber, &match.StartDate, &match.EndDate, &match.StartTime, &match.IsDayNight, &match.GroundId, &match.GroundName, &match.Team1Id, &match.Team1Name, &match.Team1ImageUrl, &match.Team2Id, &match.Team2Name, &match.Team2ImageUrl, &match.Season, &match.MainSeriesId, &match.MainSeriesName, &match.CurrentStatus, &match.FinalResult, &match.OutcomeSpecialMethod, &match.MatchWinnerId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.TossWinnerId, &match.IsTossDecisionBat, &match.Innings)

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

func ReadMatchById(ctx context.Context, db DB_Exec, matchId int) (responses.SingleMatchResponse, error) {
	var match responses.SingleMatchResponse

	query := `
	SELECT
		matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.cricsheet_id, matches.balls_per_over, matches.event_match_number, matches.start_date, matches.end_date, matches.start_time, matches.is_day_night, matches.season, matches.ground_id, grounds.name, matches.main_series_id, series.name, matches.team1_id, t1.name, t1.image_url, matches.team2_id, t2.name, t2.image_url, matches.current_status, matches.final_result, matches.outcome_special_method, matches.match_winner_team_id, matches.bowl_out_winner_id, matches.super_over_winner_id, matches.is_won_by_innings, matches.is_won_by_runs, matches.win_margin, matches.balls_remaining_after_win, matches.toss_winner_team_id, matches.is_toss_decision_bat,
	
		(
	        SELECT ARRAY_AGG ( ROW ( mse.team_id, mse.player_id, players.name, mse.is_captain, mse.is_wk, mse.is_debut, mse.is_vice_captain, mse.playing_status ) )
	        FROM
	            match_squad_entries mse
	            LEFT JOIN players ON mse.player_id = players.id
	        WHERE
	            mse.match_id = matches.id
	    ) AS squad_entries,
					
		(
	        SELECT ARRAY_AGG ( ROW (player_awards.player_id, player_awards.award_type) )
	        FROM
	            player_awards
	        WHERE
	            player_awards.match_id = matches.id
        ) AS match_awards,
	
        (
            SELECT
                ARRAY_AGG (
                    ROW ( innings.id, innings.innings_number, innings.batting_team_id, innings.bowling_team_id, innings.total_runs, innings.total_balls, innings.total_wickets, innings.byes, innings.leg_byes, innings.wides, innings.noballs, innings.penalty, innings.is_super_over, innings.innings_end, innings.target_runs, innings.max_overs ,
	                    (
	                        SELECT ARRAY_AGG ( ROW ( bs.batter_id, bs.batting_position, bs.has_batted, bs.runs_scored, bs.balls_faced, bs.minutes_batted, bs.fours_scored, bs.sixes_scored, bs.dismissal_type, bs.dismissed_by_id, bs.fielder1_id, bs.fielder2_id ) )
							FROM batting_scorecards bs
							WHERE bs.innings_id = innings.id
	                    ),
	                    (
	                        SELECT ARRAY_AGG ( ROW ( bs.bowler_id, bs.bowling_position, bs.wickets_taken, bs.runs_conceded, bs.balls_bowled, bs.maiden_overs, bs.fours_conceded, bs.sixes_conceded, bs.wides_conceded, bs.noballs_conceded ) ) 
							FROM bowling_scorecards bs
							WHERE bs.innings_id = innings.id
	                    )
                    )
                )
            FROM
                innings
            WHERE
                innings.match_id = matches.id
        ) AS innings
	
	FROM matches 

	LEFT JOIN teams t1 ON matches.team1_id = t1.id 
	LEFT JOIN teams t2 ON matches.team2_id = t2.id 
	LEFT JOIN grounds ON matches.ground_id = grounds.id
	LEFT JOIN series ON matches.main_series_id = series.id
	WHERE matches.id = $1
	`

	row := db.QueryRow(ctx, query, matchId)

	err := row.Scan(&match.Id, &match.PlayingLevel, &match.PlayingFormat, &match.MatchType, &match.CricsheetId, &match.BallsPerOver, &match.EventMatchNumber, &match.StartDate, &match.EndDate, &match.StartTime, &match.IsDayNight, &match.Season, &match.GroundId, &match.GroundName, &match.MainSeriesId, &match.MainSeriesName, &match.Team1Id, &match.Team1Name, &match.Team1ImageUrl, &match.Team2Id, &match.Team2Name, &match.Team2ImageUrl, &match.CurrentStatus, &match.FinalResult, &match.OutcomeSpecialMethod, &match.MatchWinnerId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.TossWinnerId, &match.IsTossDecisionBat, &match.SquadEntries, &match.MatchAwards, &match.Innings)

	return match, err
}

func UpdateMatch(ctx context.Context, db *pgxpool.Pool, match *models.Match) error {
	query := `UPDATE matches SET start_date = $1, start_time = $2, team1_id = $3, team2_id = $4, is_male = $5, main_series_id = $6, ground_id = $7, is_neutral_venue = $8, current_status = $9, final_result = $10, home_team_id = $11, away_team_id = $12, match_type = $13, playing_level = $14, playing_format = $15, season = $16, is_day_night = $17, outcome_special_method = $18, toss_winner_team_id = $19, toss_loser_team_id = $20, is_toss_decision_bat = $21, match_winner_team_id = $22, match_loser_team_id = $23, bowl_out_winner_id = $24, super_over_winner_id = $25, is_won_by_innings = $26, is_won_by_runs = $27, win_margin = $28, balls_remaining_after_win = $29, balls_per_over = $30, cricsheet_id = $31, is_bbb_done = $32, event_match_number = $33, end_date = $34 WHERE id = $35`

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

	cmd, err := tx.Exec(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.MainSeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate, match.Id)

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
	INSERT INTO matches (start_date, start_time, team1_id, team2_id, is_male, main_series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34)
	ON CONFLICT (cricsheet_id)
	DO UPDATE SET start_date = $1, start_time = $2, team1_id = $3, team2_id = $4, is_male = $5, main_series_id = $6, ground_id = $7, is_neutral_venue = $8, current_status = $9, final_result = $10, home_team_id = $11, away_team_id = $12, match_type = $13, playing_level = $14, playing_format = $15, season = $16, is_day_night = $17, outcome_special_method = $18, toss_winner_team_id = $19, toss_loser_team_id = $20, is_toss_decision_bat = $21, match_winner_team_id = $22, match_loser_team_id = $23, bowl_out_winner_id = $24, super_over_winner_id = $25, is_won_by_innings = $26, is_won_by_runs = $27, win_margin = $28, balls_remaining_after_win = $29, balls_per_over = $30, cricsheet_id = $31, is_bbb_done = $32, event_match_number = $33, end_date = $34
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

	err = tx.QueryRow(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.MainSeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate).Scan(&matchId)

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
		id, cricsheet_id, event_match_number, start_date, start_time, end_date, team1_id, team2_id, is_male,
	 
		ARRAY[mse.series_id], 
		
		main_series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, is_bbb_done 
		
		FROM matches 
		LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
		WHERE cricsheet_id = $1
	`

	row := db.QueryRow(ctx, query, cricsheetId)

	err := row.Scan(&match.Id, &match.CricsheetId, &match.EventMatchNumber, &match.StartDate, &match.StartTime, &match.EndDate, &match.Team1Id, &match.Team2Id, &match.IsMale, &match.SeriesListId, &match.MainSeriesId, &match.GroundId, &match.IsNeutralVenue, &match.CurrentStatus, &match.FinalResult, &match.HomeTeamId, &match.AwayTeamId, &match.MatchType, &match.PlayingLevel, &match.PlayingFormat, &match.Season, &match.IsDayNight, &match.OutcomeSpecialMethod, &match.TossWinnerId, &match.TossLoserId, &match.IsTossDecisionBat, &match.MatchWinnerId, &match.MatchLoserId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.BallsPerOver, &match.IsBBBDone)

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

func ReadMatchResultOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::match_final_result))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	matchResultOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var matchResult string
		err := row.Scan(&matchResult)
		return matchResult, err
	})

	return matchResultOptions, err
}

func ReadMatchTypeOptions(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::match_type))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	matchTypeOptions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var matchType string
		err := row.Scan(&matchType)
		return matchType, err
	})

	return matchTypeOptions, err
}

func ReadMatchFormats(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::playing_format))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	matchFormats, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var matchFormat string
		err := row.Scan(&matchFormat)
		return matchFormat, err
	})

	return matchFormats, err
}

func ReadMatchLevels(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	query := `SELECT unnest(enum_range(null::playing_level))`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	playingLevels, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var playingLevel string
		err := row.Scan(&playingLevel)
		return playingLevel, err
	})

	return playingLevels, err
}

func ReadSeriesMatches(ctx context.Context, db DB_Exec, seriesId int64) ([]responses.AllMatches, error) {
	query := `
	SELECT 
    	matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.cricsheet_id, matches.balls_per_over, matches.event_match_number, matches.start_date, matches.end_date, matches.start_time, matches.is_day_night, matches.ground_id, grounds.name, matches.team1_id, t1.name, t1.image_url, matches.team2_id, t2.name, t2.image_url, matches.season, matches.main_series_id, ms.name, matches.current_status, matches.final_result, matches.outcome_special_method, matches.match_winner_team_id, matches.bowl_out_winner_id, matches.super_over_winner_id, matches.is_won_by_innings, matches.is_won_by_runs, matches.win_margin, matches.balls_remaining_after_win, matches.toss_winner_team_id, matches.is_toss_decision_bat,
	    (
	        SELECT
	            ARRAY_AGG (
	                ROW ( innings.innings_number, innings.batting_team_id, innings.total_runs, innings.total_balls, innings.total_wickets, innings.innings_end, innings.target_runs, innings.max_overs )
	            )
	        FROM
	            innings
	        WHERE
	            innings.match_id = matches.id AND innings.is_super_over = FALSE AND innings.innings_number IS NOT NULL
	    )
	FROM
	    matches
	    LEFT JOIN match_series_entries mse ON mse.match_id = matches.id
	    LEFT JOIN series ON mse.series_id = series.id
	    LEFT JOIN series ms ON matches.main_series_id = ms.id
	    LEFT JOIN teams t1 ON matches.team1_id = t1.id
	    LEFT JOIN teams t2 ON matches.team2_id = t2.id
	    LEFT JOIN grounds ON matches.ground_id = grounds.id
	WHERE
	    series.id = $1
	`

	rows, err := db.Query(ctx, query, seriesId)
	if err != nil {
		return nil, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllMatches, error) {
		var match responses.AllMatches

		err := rows.Scan(&match.Id, &match.PlayingLevel, &match.PlayingFormat, &match.MatchType, &match.CricsheetId, &match.BallsPerOver, &match.EventMatchNumber, &match.StartDate, &match.EndDate, &match.StartTime, &match.IsDayNight, &match.GroundId, &match.GroundName, &match.Team1Id, &match.Team1Name, &match.Team1ImageUrl, &match.Team2Id, &match.Team2Name, &match.Team2ImageUrl, &match.Season, &match.MainSeriesId, &match.MainSeriesName, &match.CurrentStatus, &match.FinalResult, &match.OutcomeSpecialMethod, &match.MatchWinnerId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.TossWinnerId, &match.IsTossDecisionBat, &match.Innings)

		return match, err
	})

	return matches, err
}
