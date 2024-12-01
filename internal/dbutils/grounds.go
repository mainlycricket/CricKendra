package dbutils

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/pkg/pgxutils"
)

func InsertGround(ctx context.Context, db DB_Exec, ground *models.Ground) (int64, error) {
	var id int64

	query := `INSERT INTO grounds (name, city_id) VALUES($1, $2) RETURNING id`

	err := db.QueryRow(ctx, query, ground.Name, ground.CityId).Scan(&id)

	return id, err
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
