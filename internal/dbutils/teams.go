package dbutils

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func InsertTeam(ctx context.Context, db *pgxpool.Pool, team *models.Team) error {
	query := `INSERT INTO teams (name, is_male, image_url, playing_level, short_name) VALUES($1, $2, $3, $4, $5)`

	cmd, err := db.Exec(ctx, query, team.Name, team.IsMale, team.ImageURL, team.PlayingLevel, team.ShortName)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert team")
	}

	return nil
}

func ReadTeams(ctx context.Context, db *pgxpool.Pool) ([]responses.AllTeams, error) {
	query := `SELECT id, name, is_male, image_url, playing_level, short_name FROM teams`

	rows, err := db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllTeams, error) {
		var team responses.AllTeams

		err := rows.Scan(&team.Id, &team.Name, &team.IsMale, &team.ImageURL, &team.PlayingLevel, &team.ShortName)

		return team, err
	})

	return data, err
}
