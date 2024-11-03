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

func InsertGround(ctx context.Context, db *pgxpool.Pool, ground *models.Ground) error {
	query := `INSERT INTO grounds (name, host_nation_id, city_id) VALUES($1, $2, $3)`

	cmd, err := db.Exec(ctx, query, ground.Name, ground.HostNationId, ground.CityId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert ground")
	}

	return nil
}

func ReadGrounds(ctx context.Context, db *pgxpool.Pool, queryMap url.Values) (responses.AllGroundsResponse, error) {
	var response responses.AllGroundsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "grounds",
		DefaultLimit: 50,
		DefaultSort:  []string{"name"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Ground](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT grounds.id, grounds.name, grounds.host_nation_id, host_nations.name, grounds.city_id, cities.name FROM grounds LEFT JOIN host_nations ON grounds.host_nation_id = host_nations.id LEFT JOIN cities ON grounds.city_id = cities.id %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	grounds, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllGrounds, error) {
		var ground responses.AllGrounds
		err := rows.Scan(&ground.Id, &ground.Name, &ground.HostNationId, &ground.HostNationName, &ground.CityId, &ground.CityName)
		return ground, err
	})

	if len(grounds) > queryInfoOutput.RecordsCount {
		response.Grounds = grounds[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Grounds = grounds
		response.Next = false
	}

	return response, err
}
