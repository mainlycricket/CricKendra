package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertTournament(ctx context.Context, db *pgxpool.Pool, tournament *models.Tournament) error {
	query := `INSERT INTO tournaments (name, is_male, playing_level, playing_format) VALUES($1, $2, $3, $4)`

	cmd, err := db.Exec(ctx, query, tournament.Name, tournament.IsMale, tournament.PlayingLevel, tournament.PlayingFormat)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert tournament")
	}

	return nil
}
