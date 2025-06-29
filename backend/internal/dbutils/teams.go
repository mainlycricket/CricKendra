package dbutils

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	"github.com/mainlycricket/CricKendra/backend/pkg/pgxutils"
)

func InsertTeam(ctx context.Context, db DB_Exec, team *models.Team) (int64, error) {
	var id int64

	query := `INSERT INTO teams (name, is_male, image_url, playing_level, short_name) VALUES($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(ctx, query, team.Name, team.IsMale, team.ImageURL, team.PlayingLevel, team.ShortName).Scan(&id)

	return id, err
}

func ReadTeams(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllTeamsResponse, error) {
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
