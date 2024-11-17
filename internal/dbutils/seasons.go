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

func InsertSeason(ctx context.Context, db DB_Exec, season *models.Season) error {
	query := `INSERT INTO seasons (season) VALUES($1)`

	cmd, err := db.Exec(ctx, query, season.Season)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert season")
	}

	return nil
}

func ReadSeasons(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllSeasonsResponse, error) {
	var response responses.AllSeasonsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "seasons",
		DefaultLimit: 50,
		DefaultSort:  []string{"season"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Season](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT season FROM seasons %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	seasons, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var season string

		err := rows.Scan(&season)

		return season, err
	})

	if len(seasons) > queryInfoOutput.RecordsCount {
		response.Seasons = seasons[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Seasons = seasons
		response.Next = false
	}

	return response, err
}
