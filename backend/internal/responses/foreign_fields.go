package responses

import "github.com/jackc/pgx/v5/pgtype"

type TeamAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type PlayerAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type HostNationAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type ContinentAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type GroundAsForeignField struct {
	Id             pgtype.Int8 `json:"id"`
	Name           pgtype.Text `json:"name"`
	CityName       pgtype.Text `json:"city_name"`
	HostNationName pgtype.Text `json:"host_nation_name"`
}

type SeriesAsForeignField struct {
	Id     pgtype.Int8 `json:"id"`
	Name   pgtype.Text `json:"name"`
	Season pgtype.Text `json:"season"`
}

type TournamentAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}
