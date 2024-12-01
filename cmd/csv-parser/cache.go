package main

import (
	"sync"

	"github.com/mainlycricket/CricKendra/internal/responses"
)

// Generic Sync Cache

type CacheMap[K comparable, V any] struct {
	sync.Map
}

func (cache *CacheMap[K, V]) Set(key K, value V) {
	cache.Store(key, value)
}

func (cache *CacheMap[K, V]) Get(key K) (V, bool) {
	var value V

	data, ok := cache.Load(key)
	if !ok {
		return value, false
	}

	return data.(V), true
}

// Teams

type TeamKey struct {
	TeamName     string
	IsMale       bool
	PlayingLevel string
}

var TeamsCache CacheMap[TeamKey, int64]

// Players

type PlayerKey struct {
	CricsheetId string
}

var PlayersCache CacheMap[PlayerKey, responses.AllPlayers]

// Tours

type TourKey struct {
	Season        string
	TouringTeamId int64
	HostNationsId string // ids joined by a "_"
}

var ToursCache CacheMap[TourKey, int64]

// Tournament

type TournamentKey struct {
	Name          string
	IsMale        bool
	PlayingLevel  string
	PlayingFormat string
}

var TournamentsCache CacheMap[TournamentKey, int64]

// Series

type SeriesKey struct {
	Name          string
	Season        string
	IsMale        bool
	PlayingLevel  string
	PlayingFormat string
}

var SeriesCache CacheMap[SeriesKey, int64]

// Series Squad

type SeriesSquadKey struct {
	SeriesId int64
	TeamId   int64
}

var SeriesSquadCache CacheMap[SeriesSquadKey, int64]

// Seasons

var SeasonsCache CacheMap[string, bool]

// HostNations

var HostNationCache CacheMap[string, int64]

// Grounds

type GroundKey struct {
	Venue string
	City  string
}

var GroundsCache CacheMap[GroundKey, responses.AllGrounds]

// Cities

var CitiesCache CacheMap[string, int64]
