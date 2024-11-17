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

func InsertGround(ctx context.Context, db DB_Exec, ground *models.Ground) error {
	query := `INSERT INTO grounds (name, city_id) VALUES($1, $2)`

	cmd, err := db.Exec(ctx, query, ground.Name, ground.CityId)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert ground")
	}

	return nil
}

func ReadGrounds(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllGroundsResponse, error) {
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

	query := fmt.Sprintf(`SELECT grounds.id, grounds.name, grounds.city_id, cities.name, cities.host_nation_id, host_nations.name, host_nations.continent_id, continents.name FROM grounds LEFT JOIN cities ON grounds.city_id = cities.id LEFT JOIN host_nations ON cities.host_nation_id = host_nations.id LEFT JOIN continents ON host_nations.continent_id = continents.id %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	grounds, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllGrounds, error) {
		var ground responses.AllGrounds
		err := rows.Scan(&ground.Id, &ground.Name, &ground.CityId, &ground.CityName, &ground.HostNationId, &ground.HostNationName, &ground.ContinetId, &ground.ContinentName)
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
