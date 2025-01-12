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
