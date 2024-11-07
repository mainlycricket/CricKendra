package main

import "github.com/mainlycricket/CricKendra/internal/responses"

// Teams

type TeamKey struct {
	TeamName     string
	IsMale       bool
	PlayingLevel string
}

var TeamsCache = make(map[TeamKey]responses.AllTeams, 10)

// Players

type PlayerKey struct {
	CricsheetId string
}

var PlayersCache = make(map[PlayerKey]responses.AllPlayers, 10)

// Tours

type TourKey struct {
	Season        string
	TouringTeamId int64
	HostNationsId string // ids joined by a "_"
}

var ToursCache = make(map[TourKey]responses.AllTours, 10)

// Tournament

type TournamentKey struct {
	Name          string
	IsMale        bool
	PlayingLevel  string
	PlayingFormat string
}

var TournamentsCache = make(map[TournamentKey]responses.AllTournaments, 10)

// Series

type SeriesKey struct {
	Name          string
	Season        string
	IsMale        bool
	PlayingLevel  string
	PlayingFormat string
}

var SeriesCache = make(map[SeriesKey]responses.AllSeries, 10)

// Seasons

var SeasonsCache = make(map[string]bool, 10)

// HostNations

var HostNationCache = make(map[string]responses.AllHostNations, 10)

// Grounds

type GroundKey struct {
	Venue string
	City  string
}

var GroundsCache = make(map[GroundKey]responses.AllGrounds, 10)

// Cities

var CitiesCache = make(map[string]responses.AllCities, 10)
