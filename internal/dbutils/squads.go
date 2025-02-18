package dbutils

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func ReadSquadByMatchId(ctx context.Context, db DB_Exec, matchId int64) (responses.MatchSquadResponse, error) {
	var response responses.MatchSquadResponse
	matchHeader := &response.MatchHeader

	query := `
		WITH match_squads AS (
			SELECT 
				mse.team_id,
				teams.name,
				
				ARRAY_AGG(
					ROW(
						mse.player_id,
						players.name,
					
						mse.is_captain,
						mse.is_wk,
						mse.is_debut,
						mse.is_vice_captain,
						mse.playing_status
					)
				)

			FROM match_squad_entries mse

			LEFT JOIN teams ON mse.team_id = teams.id
			LEFT JOIN players ON mse.player_id = players.id
			
			WHERE mse.match_id = $1

			GROUP BY mse.team_id, teams.name
		)

		SELECT
			matches.id, matches.playing_level, matches.playing_format, matches.match_type, matches.event_match_number,

			-- Day 1, 2, etc - Test / FC
			-- Stumps, Innings Break, Tea/Lunch/Dinner, Stopped
			-- Need 50 runs, won by 5 wkts, trail/lead by 8 runs, won the toss and chose to bat, match starts in
	
			matches.season, matches.start_date, matches.end_date, matches.start_time, matches.is_day_night, matches.ground_id, grounds.name, matches.main_series_id, main_series.name,

			matches.team1_id, team1.name, team1.image_url, matches.team2_id, team2.name, team2.image_url,

			(
				SELECT
					ARRAY_AGG (
						-- order is necessary for struct scanning
						ROW (
							player_awards.player_id,
							players.name,
							player_awards.award_type
						)
					)
				FROM
					player_awards
					LEFT JOIN players ON player_awards.player_id = players.id
				WHERE
					player_awards.match_id = matches.id
			) AS match_awards,

			(
				SELECT
					ARRAY_AGG (
						-- order is necessary for struct scanning
						ROW (
							innings.innings_number, innings.batting_team_id, batting_team.name,

							innings.total_runs, innings.total_balls, innings.total_wickets, innings.innings_end, innings.target_runs, innings.max_overs
						)
					)
			) AS team_innings_short_info,

			ARRAY_AGG(match_squads.*) AS squads
		
		FROM matches

		LEFT JOIN teams team1 ON matches.team1_id = team1.id
		LEFT JOIN teams team2 ON matches.team2_id = team2.id
		LEFT JOIN grounds ON matches.ground_id = grounds.id
		LEFT JOIN series main_series ON matches.main_series_id = main_series.id

		LEFT JOIN innings ON innings.match_id = matches.id
			AND innings.innings_number IS NOT NULL
			AND innings.is_super_over = FALSE

		LEFT JOIN teams batting_team ON innings.batting_team_id = batting_team.id

		LEFT JOIN match_squads ON match_squads.team_id = innings.batting_team_id

		WHERE matches.id = $1

		GROUP BY
			matches.id, grounds.name, main_series.name, team1.name, team1.image_url, team2.name, team2.image_url
	`

	row := db.QueryRow(ctx, query, matchId)

	err := row.Scan(
		&matchHeader.MatchId, &matchHeader.PlayingLevel, &matchHeader.PlayingFormat, &matchHeader.MatchType, &matchHeader.EventMatchNumber,

		&matchHeader.Season, &matchHeader.StartDate, &matchHeader.EndDate, &matchHeader.StartTime, &matchHeader.IsDayNight, &matchHeader.GroundId, &matchHeader.GroundName, &matchHeader.MainSeriesId, &matchHeader.MainSeriesName,

		&matchHeader.Team1Id, &matchHeader.Team1Name, &matchHeader.Team1ImageUrl, &matchHeader.Team2Id, &matchHeader.Team2Name, &matchHeader.Team2ImageUrl,

		&matchHeader.PlayerAwards,
		&matchHeader.InningsScores,

		&response.TeamSquads,
	)

	if err != nil {
		return response, err
	}

	return response, err
}

func UpsertMatchSquadEntries(ctx context.Context, db DB_Exec, entries []models.MatchSquad) error {
	query := `
		INSERT INTO match_squad_entries 
			(player_id, match_id, team_id, is_captain, is_vice_captain, is_wk, is_debut, playing_status) 
			VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT(player_id, match_id)
		DO UPDATE SET 
			team_id = $3, is_captain = $4, is_vice_captain = $5, is_wk = $6, is_debut = $7, playing_status = $8
	`

	batch := &pgx.Batch{}
	for _, entry := range entries {
		batch.Queue(query, &entry.PlayerId, &entry.MatchId, &entry.TeamId, &entry.IsCaptain, &entry.IsViceCaptain, &entry.IsWk, &entry.IsDebut, &entry.PlayingStatus)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
}

func UpsertMatchSquadEntry(ctx context.Context, db DB_Exec, entry *models.MatchSquad) error {
	query := `
	INSERT INTO match_squad_entries (player_id, match_id, team_id, is_captain, is_vice_captain, is_wk, is_debut, playing_status) VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT(player_id, match_id)
	DO UPDATE SET team_id = $3, is_captain = $4, is_vice_captain = $5, is_wk = $6, is_debut = $7, playing_status = $8
	`

	cmd, err := db.Exec(ctx, query, entry.PlayerId, entry.MatchId, entry.TeamId, entry.IsCaptain, entry.IsViceCaptain, entry.IsWk, entry.IsDebut, entry.PlayingStatus)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert match squad entry")
	}

	return nil
}

func InsertSeriesSquad(ctx context.Context, db DB_Exec, entry *models.SeriesSquad) (int64, error) {
	var id int64

	query := `INSERT INTO series_squads (series_id, team_id, squad_label) VALUES($1, $2, $3) RETURNING id`

	err := db.QueryRow(ctx, query, entry.SeriesId, entry.TeamId, entry.SquadLabel).Scan(&id)

	return id, err
}

func ReadSeriesSquads(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllSeriesSquadResponse, error) {
	var response responses.AllSeriesSquadResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "series_squads",
		DefaultLimit: 20,
		DefaultSort:  []string{"id"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.SeriesSquad](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT id, series_id, team_id, squad_label FROM series_squads %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	squads, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllSeriesSquads, error) {
		var squad responses.AllSeriesSquads

		err := rows.Scan(&squad.Id, &squad.SeriesId, &squad.TeamId, &squad.SquadLabel)

		return squad, err
	})

	if len(squads) > queryInfoOutput.RecordsCount {
		response.Squads = squads[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Squads = squads
		response.Next = false
	}

	return response, err
}

func InsertSeriesSquadEntry(ctx context.Context, db DB_Exec, entry *models.SeriesSquadEntry) error {
	query := `INSERT INTO series_squad_entries (squad_id, player_id, is_captain, is_vice_captain, is_wk, playing_status) VALUES($1, $2, $3, $4, $5, $6)`

	cmd, err := db.Exec(ctx, query, entry.SquadId, entry.PlayerId, entry.IsCaptain, entry.IsViceCaptain, entry.IsWk, entry.PlayingStatus)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert series squad entry")
	}

	return nil
}

func UpsertSeriesSquadEntries(ctx context.Context, db DB_Exec, entries []models.SeriesSquadEntry) error {
	query := `
		INSERT INTO series_squad_entries 
			(squad_id, player_id, is_captain, is_vice_captain, is_wk, playing_status)
			VALUES($1, $2, $3, $4, $5, $6) 
		ON CONFLICT(squad_id, player_id) 
		DO UPDATE SET 
			is_captain = $3, is_vice_captain = $4, is_wk = $5, playing_status = $6
	`

	batch := &pgx.Batch{}
	for _, entry := range entries {
		batch.Queue(query, &entry.SquadId, &entry.PlayerId, &entry.IsCaptain, &entry.IsViceCaptain, &entry.IsWk, &entry.PlayingStatus)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
}

func UpsertSeriesSquadEntry(ctx context.Context, db DB_Exec, entry *models.SeriesSquadEntry) error {
	query := `
	INSERT INTO series_squad_entries (squad_id, player_id, is_captain, is_vice_captain, is_wk, playing_status) VALUES($1, $2, $3, $4, $5, $6) 
	ON CONFLICT(squad_id, player_id) 
	DO UPDATE SET is_captain = $3, is_vice_captain = $4, is_wk = $5, playing_status = $6`

	cmd, err := db.Exec(ctx, query, entry.SquadId, entry.PlayerId, entry.IsCaptain, entry.IsViceCaptain, entry.IsWk, entry.PlayingStatus)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to upsert series squad entry")
	}

	return nil
}
