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

func ReadTeams(ctx context.Context, db *pgxpool.Pool, queryMap url.Values) (responses.AllTeamsResponse, error) {
	var response responses.AllTeamsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "teams",
		DefaultLimit: 20,
		DefaultSort:  []string{"id"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Team](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT id, name, is_male, image_url, playing_level, short_name FROM teams %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	teams, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllTeams, error) {
		var team responses.AllTeams

		err := rows.Scan(&team.Id, &team.Name, &team.IsMale, &team.ImageURL, &team.PlayingLevel, &team.ShortName)

		return team, err
	})

	if len(teams) > queryInfoOutput.RecordsCount {
		response.Teams = teams[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Teams = teams
		response.Next = false
	}

	return response, err
}
