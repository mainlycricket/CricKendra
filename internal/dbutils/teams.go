package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
)

func InsertTeam(ctx context.Context, db *pgxpool.Pool, team *models.Team) error {
	query := `INSERT INTO teams (name, is_male, image_url, playing_level) VALUES($1, $2, $3, $4)`

	cmd, err := db.Exec(ctx, query, team.Name, team.IsMale, team.ImageURL, team.PlayingLevel)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert team")
	}

	return nil
}

func ReadTeams(ctx context.Context, db *pgxpool.Pool) ([]models.Team, error) {
	query := `SELECT id, name, is_male, image_url, playing_level FROM teams`

	rows, err := db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Team, error) {
		var item models.Team

		err := rows.Scan(&item.Id, &item.Name, &item.IsMale, &item.ImageURL, &item.PlayingLevel)

		return item, err
	})

	return data, err
}
