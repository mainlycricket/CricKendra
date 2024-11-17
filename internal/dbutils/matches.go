package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertMatch(ctx context.Context, db DB_Exec, match *models.Match) error {
	query := `INSERT INTO matches (start_date, start_time, team1_id, team2_id, is_male, tournament_id, series_id, tour_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, players_of_the_match_id, balls_per_over, scorers_id, commentators_id, cricsheet_id, is_bbb_done) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37)`

	cmd, err := db.Exec(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.TournamentId, match.SeriesId, match.TourId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.PoTMsId, match.BallsPerOver, match.ScorersId, match.CommentatorsId, match.CricsheetId, match.IsBBBDone)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert match")
	}

	return nil
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

	query := fmt.Sprintf(`SELECT id, cricsheet_id FROM matches %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllMatches, error) {
		var match responses.AllMatches

		err := rows.Scan(&match.Id, &match.CricsheetId)

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

func UpdateMatch(ctx context.Context, db DB_Exec, match *models.Match) error {
	query := `UPDATE matches SET start_date = $1, start_time = $2, team1_id = $3, team2_id = $4, is_male = $5, tournament_id = $6, series_id = $7, tour_id = $8, ground_id = $9, is_neutral_venue = $10, current_status = $11, final_result = $12, home_team_id = $13, away_team_id = $14, match_type = $15, playing_level = $16, playing_format = $17, season = $18, is_day_night = $19, outcome_special_method = $20, toss_winner_team_id = $21, toss_loser_team_id = $22, is_toss_decision_bat = $23, match_winner_team_id = $24, match_loser_team_id = $25, bowl_out_winner_id = $26, super_over_winner_id = $27, is_won_by_innings = $28, is_won_by_runs = $29, win_margin = $30, balls_remaining_after_win = $31, players_of_the_match_id = $32, balls_per_over = $33, scorers_id = $34, commentators_id = $35, cricsheet_id = $36, is_bbb_done = $37 WHERE id = $38`

	cmd, err := db.Exec(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.TournamentId, match.SeriesId, match.TourId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.PoTMsId, match.BallsPerOver, match.ScorersId, match.CommentatorsId, match.CricsheetId, match.IsBBBDone, match.Id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to update match")
	}

	return nil
}

func UpsertMatch(ctx context.Context, db DB_Exec, match *models.Match) error {
	query := `
	INSERT INTO matches (start_date, start_time, team1_id, team2_id, is_male, tournament_id, series_id, tour_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, players_of_the_match_id, balls_per_over, scorers_id, commentators_id, cricsheet_id, is_bbb_done) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37)
	ON CONFLICT (cricsheet_id)
	DO UPDATE SET start_date = $1, start_time = $2, team1_id = $3, team2_id = $4, is_male = $5, tournament_id = $6, series_id = $7, tour_id = $8, ground_id = $9, is_neutral_venue = $10, current_status = $11, final_result = $12, home_team_id = $13, away_team_id = $14, match_type = $15, playing_level = $16, playing_format = $17, season = $18, is_day_night = $19, outcome_special_method = $20, toss_winner_team_id = $21, toss_loser_team_id = $22, is_toss_decision_bat = $23, match_winner_team_id = $24, match_loser_team_id = $25, bowl_out_winner_id = $26, super_over_winner_id = $27, is_won_by_innings = $28, is_won_by_runs = $29, win_margin = $30, balls_remaining_after_win = $31, players_of_the_match_id = $32, balls_per_over = $33, scorers_id = $34, commentators_id = $35, is_bbb_done = $37 
	`

	cmd, err := db.Exec(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.TournamentId, match.SeriesId, match.TourId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.PoTMsId, match.BallsPerOver, match.ScorersId, match.CommentatorsId, match.CricsheetId, match.IsBBBDone)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to upsert match")
	}

	return nil
}

func ReadMatchByCricsheetId(ctx context.Context, db DB_Exec, cricsheetId string) (models.Match, error) {
	var match models.Match

	query := `SELECT id, start_date, start_time, team1_id, team2_id, is_male, tournament_id, series_id, tour_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, players_of_the_match_id, balls_per_over, scorers_id, commentators_id, cricsheet_id, is_bbb_done FROM matches WHERE cricsheet_id = $1`

	row := db.QueryRow(ctx, query, cricsheetId)

	err := row.Scan(&match.Id, &match.StartDate, &match.StartTime, &match.Team1Id, &match.Team2Id, &match.IsMale, &match.TournamentId, &match.SeriesId, &match.TourId, &match.GroundId, &match.IsNeutralVenue, &match.CurrentStatus, &match.FinalResult, &match.HomeTeamId, &match.AwayTeamId, &match.MatchType, &match.PlayingLevel, &match.PlayingFormat, &match.Season, &match.IsDayNight, &match.OutcomeSpecialMethod, &match.TossWinnerId, &match.TossLoserId, &match.IsTossDecisionBat, &match.MatchWinnerId, &match.MatchLoserId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.PoTMsId, &match.BallsPerOver, &match.ScorersId, &match.CommentatorsId, &match.CricsheetId, &match.IsBBBDone)

	return match, err
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
