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

func InsertMatch(ctx context.Context, db DB_Exec, match *models.Match) (int64, error) {
	var id int64

	query := `INSERT INTO matches (start_date, start_time, team1_id, team2_id, is_male, series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34) RETURNING id`

	err := db.QueryRow(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.SeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate).Scan(&id)

	return id, err
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

	query := fmt.Sprintf(`SELECT 
	
	matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.cricsheet_id, matches.balls_per_over, matches.event_match_number, matches.start_date, matches.end_date, matches.start_time, matches.is_day_night, 
	
	matches.ground_id, grounds.name, 
	
	matches.team1_id, t1.name, t1.image_url, matches.team2_id, t2.name, t2.image_url,
	
	matches.current_status, matches.final_result, matches.outcome_special_method, matches.match_winner_team_id, matches.bowl_out_winner_id, matches.super_over_winner_id, matches.is_won_by_innings, matches.is_won_by_runs, matches.win_margin, matches.balls_remaining_after_win, 
	
	get_match_innings(matches.id)::team_innings_short_info[] AS innings,
	matches.series_id, series.name
	
	FROM matches 
	
	LEFT JOIN teams t1 ON matches.team1_id = t1.id 
	LEFT JOIN teams t2 ON matches.team2_id = t2.id 
	LEFT JOIN grounds ON matches.ground_id = grounds.id
	LEFT JOIN series ON matches.series_id = series.id
	
	%s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllMatches, error) {
		var match responses.AllMatches

		err := rows.Scan(&match.Id, &match.PlayingLevel, &match.PlayingFormat, &match.MatchType, &match.CricsheetId, &match.BallsPerOver, &match.EventMatchNumber, &match.StartDate, &match.EndDate, &match.StartTime, &match.IsDayNight, &match.GroundId, &match.GroundName, &match.Team1Id, &match.Team1Name, &match.Team1ImageUrl, &match.Team2Id, &match.Team2Name, &match.Team2ImageUrl, &match.CurrentStatus, &match.FinalResult, &match.OutcomeSpecialMethod, &match.MatchWinnerId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.Innings, &match.SeriesId, &match.SeriesName)

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

	query := `SELECT 
	
	matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.cricsheet_id, matches.balls_per_over, matches.event_match_number, matches.start_date, matches.end_date, matches.start_time, matches.is_day_night, matches.season,
	
	matches.ground_id, grounds.name,
	matches.series_id, series.name,
	
	matches.team1_id, t1.name, t1.image_url, matches.team2_id, t2.name, t2.image_url,
	
	matches.current_status, matches.final_result, matches.outcome_special_method, matches.match_winner_team_id, matches.bowl_out_winner_id, matches.super_over_winner_id, matches.is_won_by_innings, matches.is_won_by_runs, matches.win_margin, matches.balls_remaining_after_win, matches.toss_winner_team_id, matches.is_toss_decision_bat,
	
	get_innings_scorecard_entries($1)::innings_scorecard_record[] AS innings, 
	
	(SELECT ARRAY_AGG(ROW(team_id, player_id, players.name, is_captain, is_wk, is_debut, is_vice_captain)) FROM match_squad_entries LEFT JOIN players ON player_id = players.id WHERE match_id = matches.id AND team_id = matches.team1_id) AS squad_entries,
	
	(SELECT ARRAY_AGG(ROW(player_id, players.name, award_type)) FROM player_awards LEFT JOIN players ON player_awards.player_id = players.id WHERE player_awards.match_id = matches.id) AS players_of_the_series 
	
	FROM matches 
	
	LEFT JOIN teams t1 ON matches.team1_id = t1.id 
	LEFT JOIN teams t2 ON matches.team2_id = t2.id 
	LEFT JOIN grounds ON matches.ground_id = grounds.id
	LEFT JOIN series ON matches.series_id = series.id
	WHERE matches.id = $1`

	row := db.QueryRow(ctx, query, matchId)

	err := row.Scan(&match.Id, &match.PlayingLevel, &match.PlayingFormat, &match.MatchType, &match.CricsheetId, &match.BallsPerOver, &match.EventMatchNumber, &match.StartDate, &match.EndDate, &match.StartTime, &match.IsDayNight, &match.Season, &match.GroundId, &match.GroundName, &match.SeriesId, &match.SeriesName, &match.Team1Id, &match.Team1Name, &match.Team1ImageUrl, &match.Team2Id, &match.Team2Name, &match.Team2ImageUrl, &match.CurrentStatus, &match.FinalResult, &match.OutcomeSpecialMethod, &match.MatchWinnerId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.TossWinnerId, &match.IsTossDecisionBat, &match.Innings, &match.SquadEntries, &match.MatchAwards)

	return match, err
}

func UpdateMatch(ctx context.Context, db DB_Exec, match *models.Match) error {
	query := `UPDATE matches SET start_date = $1, start_time = $2, team1_id = $3, team2_id = $4, is_male = $5, series_id = $6, ground_id = $7, is_neutral_venue = $8, current_status = $9, final_result = $10, home_team_id = $11, away_team_id = $12, match_type = $13, playing_level = $14, playing_format = $15, season = $16, is_day_night = $17, outcome_special_method = $18, toss_winner_team_id = $19, toss_loser_team_id = $20, is_toss_decision_bat = $21, match_winner_team_id = $22, match_loser_team_id = $23, bowl_out_winner_id = $24, super_over_winner_id = $25, is_won_by_innings = $26, is_won_by_runs = $27, win_margin = $28, balls_remaining_after_win = $29, balls_per_over = $30, cricsheet_id = $31, is_bbb_done = $32, event_match_number = $33, end_date = $34 WHERE id = $35`

	cmd, err := db.Exec(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.SeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate, match.Id)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to update match")
	}

	return nil
}

func UpsertCricsheetMatch(ctx context.Context, db DB_Exec, match *models.Match) (int64, error) {
	var id int64

	query := `
	INSERT INTO matches (start_date, start_time, team1_id, team2_id, is_male, series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34)
	ON CONFLICT (cricsheet_id)
	DO UPDATE SET start_date = $1, start_time = $2, team1_id = $3, team2_id = $4, is_male = $5, series_id = $6, ground_id = $7, is_neutral_venue = $8, current_status = $9, final_result = $10, home_team_id = $11, away_team_id = $12, match_type = $13, playing_level = $14, playing_format = $15, season = $16, is_day_night = $17, outcome_special_method = $18, toss_winner_team_id = $19, toss_loser_team_id = $20, is_toss_decision_bat = $21, match_winner_team_id = $22, match_loser_team_id = $23, bowl_out_winner_id = $24, super_over_winner_id = $25, is_won_by_innings = $26, is_won_by_runs = $27, win_margin = $28, balls_remaining_after_win = $29, balls_per_over = $30, cricsheet_id = $31, is_bbb_done = $32, event_match_number = $33, end_date = $34
	RETURNING id
	`

	err := db.QueryRow(ctx, query, match.StartDate, match.StartTime, match.Team1Id, match.Team2Id, match.IsMale, match.SeriesId, match.GroundId, match.IsNeutralVenue, match.CurrentStatus, match.FinalResult, match.HomeTeamId, match.AwayTeamId, match.MatchType, match.PlayingLevel, match.PlayingFormat, match.Season, match.IsDayNight, match.OutcomeSpecialMethod, match.TossWinnerId, match.TossLoserId, match.IsTossDecisionBat, match.MatchWinnerId, match.MatchLoserId, match.BowlOutWinnerId, match.SuperOverWinnerId, match.IsWonByInnings, match.IsWonByRuns, match.WinMargin, match.BallsMargin, match.BallsPerOver, match.CricsheetId, match.IsBBBDone, match.EventMatchNumber, match.EndDate).Scan(&id)

	return id, err
}

func ReadMatchByCricsheetId(ctx context.Context, db DB_Exec, cricsheetId string) (models.Match, error) {
	var match models.Match

	query := `SELECT id, start_date, start_time, end_date, team1_id, team2_id, is_male, series_id, ground_id, is_neutral_venue, current_status, final_result, home_team_id, away_team_id, match_type, playing_level, playing_format, season, is_day_night, outcome_special_method, toss_winner_team_id, toss_loser_team_id, is_toss_decision_bat, match_winner_team_id, match_loser_team_id, bowl_out_winner_id, super_over_winner_id, is_won_by_innings, is_won_by_runs, win_margin, balls_remaining_after_win, balls_per_over, cricsheet_id, is_bbb_done, event_match_number FROM matches WHERE cricsheet_id = $1`

	row := db.QueryRow(ctx, query, cricsheetId)

	err := row.Scan(&match.Id, &match.StartDate, &match.StartTime, &match.EndDate, &match.Team1Id, &match.Team2Id, &match.IsMale, &match.SeriesId, &match.GroundId, &match.IsNeutralVenue, &match.CurrentStatus, &match.FinalResult, &match.HomeTeamId, &match.AwayTeamId, &match.MatchType, &match.PlayingLevel, &match.PlayingFormat, &match.Season, &match.IsDayNight, &match.OutcomeSpecialMethod, &match.TossWinnerId, &match.TossLoserId, &match.IsTossDecisionBat, &match.MatchWinnerId, &match.MatchLoserId, &match.BowlOutWinnerId, &match.SuperOverWinnerId, &match.IsWonByInnings, &match.IsWonByRuns, &match.WinMargin, &match.BallsMargin, &match.BallsPerOver, &match.CricsheetId, &match.IsBBBDone, &match.EventMatchNumber)

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
