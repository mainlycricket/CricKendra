package main

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

var (
	cachedTeams           = teamCache{}
	cachedPlayers         = playerCache{}
	cachedTournaments     = tournamentCache{}
	cachedSeries          = seriesCache{}
	cachedSeriesSquads    = seriesSquadCache{}
	cachedSeasons         = seasonCache{}
	cachedHostNations     = hostNationCache{}
	cachedGrounds         = groundCache{}
	cachedCities          = cityCache{}
	cachedCricsheetPeople = cricsheetPeopleCache{}
)

// Generic Sync Cache
type cacheMap[K comparable, V any] struct {
	sync.Map
}

func (cache *cacheMap[K, V]) set(key K, value V) {
	cache.Store(key, value)
}

func (cache *cacheMap[K, V]) get(key K) (V, bool) {
	var value V

	data, ok := cache.Load(key)
	if !ok {
		return value, false
	}

	return data.(V), true
}

func (cache *cacheMap[K, V]) getFiltersFromKey(key K) url.Values {
	values := url.Values{"__page": []string{"1"}, "__limit": []string{"1"}}
	v := reflect.ValueOf(key)

	// If key is a pointer, get the underlying value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Only proceed if we have a struct
	if v.Kind() != reflect.Struct {
		return values
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Convert field value to string based on its type
		var stringValue string
		switch value.Kind() {
		case reflect.String:
			stringValue = value.String()
		case reflect.Bool:
			stringValue = strconv.FormatBool(value.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			stringValue = strconv.FormatInt(value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			stringValue = strconv.FormatUint(value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			stringValue = strconv.FormatFloat(value.Float(), 'f', -1, 64)
		default:
			continue // Skip unsupported types
		}

		if stringValue != "" {
			values.Set(field.Name, stringValue)
		}
	}

	return values
}

type cacheValue[T any] struct {
	data T
	lock sync.Mutex
}

// Teams

type teamKey struct {
	name          string
	is_male       bool
	playing_level string
}

type teamCache struct {
	cacheMap[teamKey, *cacheValue[pgtype.Int8]]
}

func (cache *teamCache) loadOrStore(key teamKey) (pgtype.Int8, error) {
	value, loaded := cache.LoadOrStore(key, &cacheValue[pgtype.Int8]{})
	teamValue := value.(*cacheValue[pgtype.Int8])

	if loaded && teamValue.data.Valid {
		return teamValue.data, nil
	}

	teamValue.lock.Lock()
	defer teamValue.lock.Unlock()

	if teamValue, ok := cache.get(key); ok && teamValue.data.Valid {
		return teamValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadTeams(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`error while reading teams: %v`, err)
	}

	if len(dbResponse.Teams) > 0 {
		teamValue.data = dbResponse.Teams[0].Id
		cache.set(key, teamValue)
		return teamValue.data, nil
	}

	newTeam := models.Team{
		Name:         pgtype.Text{String: key.name, Valid: true},
		IsMale:       pgtype.Bool{Bool: key.is_male, Valid: true},
		PlayingLevel: pgtype.Text{String: key.playing_level, Valid: true},
		ShortName:    pgtype.Text{String: defaultTeamShortName(key.name), Valid: true},
	}

	teamId, err := dbutils.InsertTeam(context.Background(), DB_POOL, &newTeam)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`error while inserting team: %v`, err)
	}

	teamValue.data = pgtype.Int8{Int64: teamId, Valid: true}
	cache.set(key, teamValue)
	return teamValue.data, nil
}

// Players

type playerKey struct {
	cricsheet_id string
}

type playerCache struct {
	cacheMap[playerKey, *cacheValue[responses.AllPlayers]]
}

func (cache *playerCache) loadOrStore(key playerKey, name string, is_male bool, teamId int64) (responses.AllPlayers, error) {
	playerTeamEntry := &models.PlayerTeamEntry{TeamId: pgtype.Int8{Int64: teamId, Valid: true}}

	value, loaded := cache.LoadOrStore(key, &cacheValue[responses.AllPlayers]{})
	playerValue := value.(*cacheValue[responses.AllPlayers])

	if loaded && playerValue.data.Id.Valid {
		playerTeamEntry.PlayerId = playerValue.data.Id
		if err := dbutils.UpsertPlayerTeamEntry(context.Background(), DB_POOL, playerTeamEntry); err != nil {
			return responses.AllPlayers{}, err
		}
		return playerValue.data, nil
	}

	playerValue.lock.Lock()
	defer playerValue.lock.Unlock()

	if playerValue, ok := cache.get(key); ok && playerValue.data.Id.Valid {
		playerTeamEntry.PlayerId = playerValue.data.Id
		if err := dbutils.UpsertPlayerTeamEntry(context.Background(), DB_POOL, playerTeamEntry); err != nil {
			return responses.AllPlayers{}, err
		}
		return playerValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadPlayers(context.Background(), DB_POOL, filters)
	if err != nil {
		return responses.AllPlayers{}, fmt.Errorf(`error while reading reading players: %v`, err)
	}

	if len(dbResponse.Players) > 0 {
		playerValue.data = dbResponse.Players[0]
		playerTeamEntry.PlayerId = playerValue.data.Id
		if err := dbutils.UpsertPlayerTeamEntry(context.Background(), DB_POOL, playerTeamEntry); err != nil {
			return responses.AllPlayers{}, err
		}
		cache.set(key, playerValue)
		return playerValue.data, nil
	}

	cricsheetPeople, err := cachedCricsheetPeople.loadOrStore(key.cricsheet_id, name)
	if err != nil {
		return responses.AllPlayers{}, fmt.Errorf(`failed to handle cricsheet people %s: %v`, key.cricsheet_id, err)
	}

	newPlayer := models.Player{
		Name:               cricsheetPeople.Name,
		FullName:           cricsheetPeople.UniqueName,
		CricsheetId:        cricsheetPeople.Identifier,
		CricbuzzId:         cricsheetPeople.CricbuzzId,
		CricinfoId:         cricsheetPeople.CricinfoId,
		IsMale:             pgtype.Bool{Bool: is_male, Valid: true},
		TeamsRepresentedId: []pgtype.Int8{{Int64: teamId, Valid: true}},
	}

	playerId, err := dbutils.InsertPlayer(context.Background(), DB_POOL, &newPlayer)
	if err != nil {
		return responses.AllPlayers{}, fmt.Errorf(`error while inserting player: %v`, err)
	}

	playerValue.data = responses.AllPlayers{
		Id:     pgtype.Int8{Int64: playerId, Valid: true},
		Name:   newPlayer.Name,
		IsMale: pgtype.Bool{Bool: is_male, Valid: true},
	}

	cache.set(key, playerValue)
	return playerValue.data, nil
}

// Tournament

type tournamentKey struct {
	name           string
	is_male        bool
	playing_level  string
	playing_format string
}

type tournamentCache struct {
	cacheMap[tournamentKey, *cacheValue[pgtype.Int8]]
}

func (cache *tournamentCache) loadOrStore(key tournamentKey) (pgtype.Int8, error) {
	value, loaded := cache.LoadOrStore(key, &cacheValue[pgtype.Int8]{})
	tournamentValue := value.(*cacheValue[pgtype.Int8])
	if loaded && tournamentValue.data.Valid {
		return tournamentValue.data, nil
	}

	tournamentValue.lock.Lock()
	defer tournamentValue.lock.Unlock()

	if tournamentValue, ok := cache.get(key); ok && tournamentValue.data.Valid {
		return tournamentValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadTournaments(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Tournaments) == 0 {
		return pgtype.Int8{}, nil
	}

	tournamentValue.data = dbResponse.Tournaments[0].Id
	cache.set(key, tournamentValue)
	return tournamentValue.data, nil
}

// Series

type seriesKey struct {
	name           string
	season         string
	is_male        bool
	playing_level  string
	playing_format string
}

type seriesValue struct {
	seriesId pgtype.Int8
	teamsId  []pgtype.Int8
}

type seriesCache struct {
	cacheMap[seriesKey, *cacheValue[seriesValue]]
}

func (cache *seriesCache) loadOrStore(key seriesKey, team1Id, team2Id int64, tourFlag string) (pgtype.Int8, error) {
	value, loaded := cache.LoadOrStore(key, &cacheValue[seriesValue]{})
	seriesValue := value.(*cacheValue[seriesValue])

	if loaded && seriesValue.data.seriesId.Valid {
		if err := handleSeriesTeamsEntries(seriesValue.data, team1Id, team2Id); err != nil {
			return pgtype.Int8{}, fmt.Errorf(`error while upserting series team entries: %v`, err)
		}

		return seriesValue.data.seriesId, nil
	}

	seriesValue.lock.Lock()
	defer seriesValue.lock.Unlock()

	if seriesValue, ok := cache.get(key); ok && seriesValue.data.seriesId.Valid {
		if err := handleSeriesTeamsEntries(seriesValue.data, team1Id, team2Id); err != nil {
			return pgtype.Int8{}, fmt.Errorf(`error while upserting series team entries: %v`, err)
		}

		return seriesValue.data.seriesId, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadSeries(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Series) > 0 {
		seriesValue.data.seriesId = dbResponse.Series[0].Id
		for _, team := range dbResponse.Series[0].Teams {
			seriesValue.data.teamsId = append(seriesValue.data.teamsId, team.Id)
		}

		if err := handleSeriesTeamsEntries(seriesValue.data, team1Id, team2Id); err != nil {
			return pgtype.Int8{}, fmt.Errorf(`error while upserting series team entries: %v`, err)
		}

		cache.set(key, seriesValue)
		return seriesValue.data.seriesId, nil
	}

	// check if series is a part of a tournament

	tournamentKey := tournamentKey{
		name:           key.name,
		is_male:        key.is_male,
		playing_level:  key.playing_level,
		playing_format: key.playing_format,
	}

	tournamentId, err := cachedTournaments.loadOrStore(tournamentKey)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if err := cachedSeasons.loadOrStore(seasonKey{season: key.season}); err != nil {
		return pgtype.Int8{}, fmt.Errorf(`error while ensuring series season: %v`, err)
	}

	newSeries := models.Series{
		Name:          pgtype.Text{String: key.name, Valid: true},
		TeamsId:       []pgtype.Int8{{Int64: team1Id, Valid: true}, {Int64: team2Id, Valid: true}},
		TourFlag:      pgtype.Text{String: tourFlag, Valid: tourFlag != ""},
		Season:        pgtype.Text{String: key.season, Valid: true},
		IsMale:        pgtype.Bool{Bool: key.is_male, Valid: true},
		PlayingLevel:  pgtype.Text{String: key.playing_level, Valid: true},
		PlayingFormat: pgtype.Text{String: key.playing_format, Valid: true},
		TournamentId:  tournamentId,
	}

	seriesId, err := dbutils.InsertSeries(context.Background(), DB_POOL, &newSeries)
	if err != nil {
		return pgtype.Int8{}, err
	}

	seriesValue.data.seriesId = pgtype.Int8{Int64: seriesId, Valid: true}
	seriesValue.data.teamsId = newSeries.TeamsId
	cache.set(key, seriesValue)
	return seriesValue.data.seriesId, nil
}

// Series Squad

type seriesSquadKey struct {
	series_id int64
	team_id   int64
}

type seriesSquadCache struct {
	cacheMap[seriesSquadKey, *cacheValue[pgtype.Int8]]
}

func (cache *seriesSquadCache) loadOrStore(key seriesSquadKey, teamName string) (pgtype.Int8, error) {
	value, loaded := cache.LoadOrStore(key, &cacheValue[pgtype.Int8]{})
	squadValue := value.(*cacheValue[pgtype.Int8])

	if loaded && squadValue.data.Valid {
		return squadValue.data, nil
	}

	squadValue.lock.Lock()
	defer squadValue.lock.Unlock()

	if squadValue, ok := cache.get(key); ok && squadValue.data.Valid {
		return squadValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadAllSeriesSquads(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Squads) > 0 {
		squadValue.data = dbResponse.Squads[0].Id
		cache.set(key, squadValue)
		return squadValue.data, nil
	}

	squadLabel := fmt.Sprintf(`%s squad`, teamName)
	newSeriesSquad := models.SeriesSquad{
		SeriesId:   pgtype.Int8{Int64: key.series_id, Valid: true},
		TeamId:     pgtype.Int8{Int64: key.team_id, Valid: true},
		SquadLabel: pgtype.Text{String: squadLabel, Valid: true},
	}

	squadId, err := dbutils.InsertSeriesSquad(context.Background(), DB_POOL, &newSeriesSquad)
	if err != nil {
		return pgtype.Int8{}, err
	}

	squadValue.data = pgtype.Int8{Int64: squadId, Valid: true}
	cache.set(key, squadValue)
	return squadValue.data, nil
}

// Seasons

type seasonKey struct {
	season string
}

type seasonCache struct {
	cacheMap[seasonKey, *cacheValue[bool]]
}

func (cache *seasonCache) loadOrStore(key seasonKey) error {
	value, loaded := cache.LoadOrStore(key, &cacheValue[bool]{})
	seasonValue := value.(*cacheValue[bool])

	if loaded && seasonValue.data {
		return nil
	}

	seasonValue.lock.Lock()
	defer seasonValue.lock.Unlock()

	if seasonValue, ok := cache.get(key); ok && seasonValue.data {
		return nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadSeasons(context.Background(), DB_POOL, filters)
	if err != nil {
		return fmt.Errorf(`error while reading seasons: %v`, err)
	}

	if len(dbResponse.Seasons) > 0 {
		seasonValue.data = true
		cache.set(key, seasonValue)
		return nil
	}

	newSeason := models.Season{
		Season: pgtype.Text{String: key.season, Valid: true},
	}

	if err := dbutils.InsertSeason(context.Background(), DB_POOL, &newSeason); err != nil {
		return fmt.Errorf(`error while inserting season: %v`, err)
	}

	seasonValue.data = true
	cache.set(key, seasonValue)
	return nil
}

// HostNations

type hostNationKey struct {
	name string
}

type hostNationCache struct {
	cacheMap[hostNationKey, *cacheValue[pgtype.Int8]]
}

func (cache *hostNationCache) loadOrStore(key hostNationKey) (pgtype.Int8, error) {
	value, loaded := cache.LoadOrStore(key, &cacheValue[pgtype.Int8]{})
	hostNationValue := value.(*cacheValue[pgtype.Int8])

	if loaded && hostNationValue.data.Valid {
		return hostNationValue.data, nil
	}

	hostNationValue.lock.Lock()
	defer hostNationValue.lock.Unlock()

	if hostNationValue, ok := cache.get(key); ok && hostNationValue.data.Valid {
		return hostNationValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadHostNations(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`failed to read host nations: %v`, err)
	}

	if len(dbResponse.HostNations) > 0 {
		hostNationValue.data = dbResponse.HostNations[0].Id
		cache.set(key, hostNationValue)
		return hostNationValue.data, nil
	}

	newHostNation := models.HostNation{
		Name: pgtype.Text{String: key.name, Valid: true},
	}

	hostNationId, err := dbutils.InsertHostNation(context.Background(), DB_POOL, &newHostNation)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`failed to insert host nation: %v`, err)
	}

	hostNationValue.data = pgtype.Int8{Int64: hostNationId, Valid: true}
	cache.set(key, hostNationValue)
	return hostNationValue.data, nil
}

// Grounds

type groundKey struct {
	name    string
	city_id int64
}

type groundCache struct {
	cacheMap[groundKey, *cacheValue[responses.AllGrounds]]
}

func (cache *groundCache) loadOrStore(key groundKey) (responses.AllGrounds, error) {
	value, loaded := cache.LoadOrStore(key, &cacheValue[responses.AllGrounds]{})
	groundValue := value.(*cacheValue[responses.AllGrounds])

	if loaded && groundValue.data.Id.Valid {
		return groundValue.data, nil
	}

	groundValue.lock.Lock()
	defer groundValue.lock.Unlock()

	if groundValue, ok := cache.get(key); ok && groundValue.data.Id.Valid {
		return groundValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	if key.city_id == 0 {
		filters.Del("city_id")
	}

	dbResponse, err := dbutils.ReadGrounds(context.Background(), DB_POOL, filters)
	if err != nil {
		return responses.AllGrounds{}, fmt.Errorf(`error while reading grounds: %v`, err)
	}

	if len(dbResponse.Grounds) > 0 {
		groundValue.data = dbResponse.Grounds[0]
		cache.set(key, groundValue)
		return groundValue.data, nil
	}

	newGround := models.Ground{
		Name:   pgtype.Text{String: key.name, Valid: true},
		CityId: pgtype.Int8{Int64: key.city_id, Valid: key.city_id > 0},
	}

	groundId, err := dbutils.InsertGround(context.Background(), DB_POOL, &newGround)
	if err != nil {
		return responses.AllGrounds{}, fmt.Errorf(`error while inserting ground: %v`, err)
	}

	groundValue.data = responses.AllGrounds{
		Id:     pgtype.Int8{Int64: groundId, Valid: true},
		Name:   newGround.Name,
		CityId: newGround.CityId,
	}

	cache.set(key, groundValue)
	return groundValue.data, nil
}

// Cities

type cityKey struct {
	name string
}

type cityCache struct {
	cacheMap[cityKey, *cacheValue[pgtype.Int8]]
}

func (cache *cityCache) loadOrStore(key cityKey) (pgtype.Int8, error) {
	if key.name == "" {
		return pgtype.Int8{}, nil
	}

	value, loaded := cache.LoadOrStore(key, &cacheValue[pgtype.Int8]{})
	cityValue := value.(*cacheValue[pgtype.Int8])

	if loaded && cityValue.data.Valid {
		return cityValue.data, nil
	}

	cityValue.lock.Lock()
	defer cityValue.lock.Unlock()

	if cityValue, ok := cache.get(key); ok && cityValue.data.Valid {
		return cityValue.data, nil
	}

	filters := cache.getFiltersFromKey(key)
	dbResponse, err := dbutils.ReadCities(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`error while reading cities: %v`, err)
	}

	if len(dbResponse.Cities) > 0 {
		cityValue.data = dbResponse.Cities[0].Id
		cache.set(key, cityValue)
		return cityValue.data, nil
	}

	newCity := models.City{
		Name: pgtype.Text{String: key.name, Valid: true},
	}

	cityId, err := dbutils.InsertCity(context.Background(), DB_POOL, &newCity)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`error while inserting city: %v`, err)
	}

	cityValue.data = pgtype.Int8{Int64: cityId, Valid: true}
	cache.set(key, cityValue)
	return cityValue.data, nil
}

type cricsheetPeopleCache struct {
	cacheMap[string, *cacheValue[responses.CricsheetPeople]]
}

func (cache *cricsheetPeopleCache) loadOrStore(identifier, name string) (responses.CricsheetPeople, error) {
	value, loaded := cache.LoadOrStore(identifier, &cacheValue[responses.CricsheetPeople]{})
	peopleValue := value.(*cacheValue[responses.CricsheetPeople])

	if loaded && peopleValue.data.Identifier.Valid {
		return peopleValue.data, nil
	}

	peopleValue.lock.Lock()
	defer peopleValue.lock.Unlock()

	if peopleValue, ok := cache.get(identifier); ok && peopleValue.data.Identifier.Valid {
		return peopleValue.data, nil
	}

	cricsheetPeople, err := dbutils.ReadCricsheetPeopleById(context.Background(), DB_POOL, identifier)
	if err == nil {
		return cricsheetPeople, nil
	}

	if err.Error() != "no rows in result set" {
		return responses.CricsheetPeople{}, err
	}

	if err := dbutils.InsertCricsheetPeople(context.Background(), DB_POOL, identifier, name); err != nil {
		return responses.CricsheetPeople{}, err
	}

	cricsheetPeople.Identifier = pgtype.Text{String: identifier, Valid: true}
	cricsheetPeople.Name = pgtype.Text{String: name, Valid: true}
	return cricsheetPeople, nil
}

func defaultTeamShortName(teamName string) string {
	words := strings.Split(teamName, " ")
	if len(words) == 1 {
		return strings.ToUpper(teamName[:3])
	}

	var shortName string

	for _, word := range words {
		if unicode.IsUpper(rune(word[0])) {
			shortName += string(word[0])
		}
	}

	return shortName
}

func handleSeriesTeamsEntries(sv seriesValue, team1Id, team2Id int64) error {
	var teamsId []pgtype.Int8

	if !slices.Contains(sv.teamsId, pgtype.Int8{Int64: team1Id, Valid: true}) {
		teamsId = append(teamsId, pgtype.Int8{Int64: team1Id, Valid: true})
	}

	if !slices.Contains(sv.teamsId, pgtype.Int8{Int64: team2Id, Valid: true}) {
		teamsId = append(teamsId, pgtype.Int8{Int64: team2Id, Valid: true})
	}

	if len(teamsId) == 0 {
		return nil
	}

	err := dbutils.UpsertSeriesTeamEntries(context.Background(), DB_POOL, sv.seriesId, teamsId)
	return err
}
