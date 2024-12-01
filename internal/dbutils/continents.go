package dbutils

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertContinent(ctx context.Context, db *pgxpool.Pool, continent *models.Continent) (int64, error) {
	var id int64

	query := `INSERT INTO continents (name) VALUES($1) RETURNING id`

	err := db.QueryRow(ctx, query, continent.Name).Scan(&id)

	return id, err
}

func ReadContinents(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllContinentsResponse, error) {
	var response responses.AllContinentsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "continents",
		DefaultLimit: 50,
		DefaultSort:  []string{"name"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Continent](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT id, name FROM continents %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	continents, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllContinents, error) {
		var continent responses.AllContinents
		err := rows.Scan(&continent.Id, &continent.Name)
		return continent, err
	})

	if len(continents) > queryInfoOutput.RecordsCount {
		response.Continents = continents[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Continents = continents
		response.Next = false
	}

	return response, err
}
