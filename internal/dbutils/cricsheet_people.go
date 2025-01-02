package dbutils

import (
	"context"
	"errors"

	"github.com/mainlycricket/CricKendra/internal/responses"
)

func InsertCricsheetPeople(ctx context.Context, db DB_Exec, identifier, name string) error {
	query := `INSERT INTO cricsheet_people (identifier, name) VALUES($1, $2)`

	cmd, err := db.Exec(ctx, query, identifier, name)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert people")
	}

	return nil
}

func ReadCricsheetPeopleById(ctx context.Context, db DB_Exec, identifier string) (responses.CricsheetPeople, error) {
	query := `SELECT identifier, name, unique_name, key_cricinfo, key_cricbuzz FROM cricsheet_people WHERE identifier = $1`

	row := db.QueryRow(ctx, query, identifier)

	var people responses.CricsheetPeople

	err := row.Scan(&people.Identifier, &people.Name, &people.UniqueName, &people.CricinfoId, &people.CricbuzzId)

	return people, err
}
