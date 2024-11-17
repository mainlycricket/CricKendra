package dbutils

import (
	"context"

	"github.com/mainlycricket/CricKendra/internal/responses"
)

func ReadCricsheetPeopleById(ctx context.Context, db DB_Exec, identifier string) (responses.CricsheetPeople, error) {
	query := `SELECT identifier, name, unique_name, key_cricinfo, key_cricbuzz FROM cricsheet_people WHERE identifier = $1`

	row := db.QueryRow(ctx, query, identifier)

	var people responses.CricsheetPeople

	err := row.Scan(&people.Identifier, &people.Name, &people.UniqueName, &people.CricinfoId, &people.CricbuzzId)

	return people, err
}
