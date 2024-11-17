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

func InsertTournament(ctx context.Context, db *pgxpool.Pool, tournament *models.Tournament) error {
	query := `INSERT INTO tournaments (name, is_male, playing_level, playing_format) VALUES($1, $2, $3, $4)`

	cmd, err := db.Exec(ctx, query, tournament.Name, tournament.IsMale, tournament.PlayingLevel, tournament.PlayingFormat)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() < 1 {
		return errors.New("failed to insert tournament")
	}

	return nil
}

func ReadTournaments(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.AllTournamentsResponse, error) {
	var response responses.AllTournamentsResponse

	queryInfoInput := pgxutils.QueryInfoInput{
		UrlQuery:     queryMap,
		TableName:    "tournaments",
		DefaultLimit: 20,
		DefaultSort:  []string{"id"},
	}

	queryInfoOutput, err := pgxutils.ParseQuery[models.Tournament](queryInfoInput)
	if err != nil {
		return response, err
	}

	query := fmt.Sprintf(`SELECT id, name, is_male, playing_level, playing_format FROM tournaments %s %s %s`, queryInfoOutput.WhereClause, queryInfoOutput.OrderByClause, queryInfoOutput.PaginationClause)

	rows, err := db.Query(ctx, query, queryInfoOutput.Args...)
	if err != nil {
		return response, err
	}

	tournaments, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.AllTournaments, error) {
		var tournament responses.AllTournaments

		err := rows.Scan(&tournament.Id, &tournament.Name, &tournament.IsMale, &tournament.PlayingLevel, &tournament.PlayingFormat)

		return tournament, err
	})

	if len(tournaments) > queryInfoOutput.RecordsCount {
		response.Tournaments = tournaments[:queryInfoOutput.RecordsCount]
		response.Next = true
	} else {
		response.Tournaments = tournaments
		response.Next = false
	}

	return response, err
}
