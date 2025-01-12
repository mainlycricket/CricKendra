package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func UpsertMatchAwardEntries(ctx context.Context, db DB_Exec, entries []models.PlayerAwardEntry) error {
	query := `
		INSERT INTO player_awards 
			(player_id, match_id, series_id, award_type) 
			VALUES($1, $2, $3, $4)
		ON CONFLICT(player_id, match_id, award_type)
		DO NOTHING
	`

	batch := &pgx.Batch{}
	for _, entry := range entries {
		batch.Queue(query, &entry.PlayerId, &entry.MatchId, &entry.SeriesId, &entry.AwardType)
	}

	batchResults := db.SendBatch(ctx, batch)
	return batchResults.Close()
}

func InsertPlayerAwardEntry(ctx context.Context, db DB_Exec, entry *models.PlayerAwardEntry) error {
	query := `INSERT INTO player_awards (player_id, match_id, series_id, award_type) VALUES($1, $2, $3, $4)`

	cmd, err := db.Exec(ctx, query, entry.PlayerId, entry.MatchId, entry.SeriesId, entry.AwardType)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert player award")
	}

	return nil
}
