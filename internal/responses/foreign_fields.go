package responses

import "github.com/jackc/pgx/v5/pgtype"

type TeamAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type MatchAsForeignField struct {
	Id    pgtype.Int8 `json:"id"`
	Team1 pgtype.Text `json:"team1"`
	Team2 pgtype.Text `json:"team2"`
	Date  pgtype.Date `json:"date"`
	City  pgtype.Text `json:"city"`
}

type HostNationAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type PlayerAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}
