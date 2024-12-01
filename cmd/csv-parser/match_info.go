package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

/*
balls_per_over: done
bowl_out: done
city: excluded
event: pending
date: done
eliminator: done
gender: done
match_number: done
match_referee: excluded
method: done
neutral_venue: done
outcome: done
player_of_match: done
registry: pending
reserve_umpire: excluded
season: done
team: done
toss_decision: done
toss_uncontested: done
toss_winner: done
tv_umpire: excluded
umpire: excluded
venue: done
winner: done
winner_innings: done
winner_runs: done
winner_wickets: done
*/

type TeamInfo struct {
	Id      int64
	Name    string
	Players map[string]string // map[playerName]cricsheetId
}

type MatchInfoResponse struct {
	Match     models.Match
	Team1Info TeamInfo
	Team2Info TeamInfo
	Err       error
}

func extractMatchInfo(filePath, playingLevel, playingFormat string, isMale bool, channel chan<- MatchInfoResponse) {
	var match models.Match
	var team1, team2 TeamInfo
	var venue, city, eventName string
	var potmNames []string
	var mainError error

	matchCricsheetId := strings.TrimSuffix(filepath.Base(filePath), "_info.csv")

	defer func() {
		if mainError != nil {
			channel <- MatchInfoResponse{
				Err: fmt.Errorf(" %s: %v", matchCricsheetId, mainError),
			}
		} else {
			channel <- MatchInfoResponse{
				Match:     match,
				Team1Info: team1,
				Team2Info: team2,
			}
		}
	}()

	team1.Players, team2.Players = make(map[string]string, 11), make(map[string]string, 11)

	fp, err := os.Open(filePath)
	if err != nil {
		mainError = err
		return
	}
	defer fp.Close()

	match.IsMale = pgtype.Bool{Bool: isMale, Valid: true}
	match.CricsheetId = pgtype.Text{String: matchCricsheetId, Valid: true}
	match.IsNeutralVenue = pgtype.Bool{Bool: true, Valid: true}
	match.PlayingFormat = pgtype.Text{String: playingFormat, Valid: true}
	match.PlayingLevel = pgtype.Text{String: playingLevel, Valid: true}
	match.FinalResult = pgtype.Text{String: "winner_decided", Valid: true}

	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			mainError = err
			return
		}

		if row[0] != "info" {
			continue
		}

		if row[1] == "balls_per_over" {
			balls_per_over, err := strconv.ParseInt(row[2], 10, 64)
			if err != nil {
				mainError = err
				return
			}

			match.BallsPerOver = pgtype.Int8{Int64: balls_per_over, Valid: true}
			continue
		}

		if row[1] == "method" {
			match.OutcomeSpecialMethod = pgtype.Text{String: row[2], Valid: true}
			continue
		}

		if row[1] == "outcome" {
			match.FinalResult = handleMatchOutcome(row[2])
			continue
		}

		if row[1] == "neutral_venue" && row[2] == "true" {
			match.IsNeutralVenue = pgtype.Bool{Bool: false, Valid: true}
			continue
		}

		if row[1] == "match_number" {
			match_number, err := strconv.ParseInt(row[2], 10, 64)
			if err != nil {
				mainError = err
				return
			}

			match.EventMatchNumber = pgtype.Int8{Int64: match_number, Valid: true}
			continue
		}

		if row[1] == "season" {
			match.Season, err = handleSeason(row[2])
			if err != nil {
				mainError = err
				return
			}
			continue
		}

		if row[1] == "team" {
			teamName := row[2]

			teamField, err := handleTeam(teamName, playingLevel, match.IsMale.Bool)
			if err != nil {
				mainError = err
				return
			}

			if match.Team1Id.Valid {
				match.Team2Id = teamField
				team2.Id, team2.Name = teamField.Int64, teamName
			} else {
				match.Team1Id = teamField
				team1.Id, team1.Name = teamField.Int64, teamName
			}
			continue
		}

		if row[1] == "date" {
			date, err := time.Parse("2006/01/02", row[2])
			if err != nil {
				mainError = err
				return
			}

			match.StartDate = pgtype.Date{Time: date, Valid: true}
			match.EndDate = pgtype.Date{Time: date, Valid: true}

			if playingFormat == "Test" || playingFormat == "first_class" {
				end_date := date.AddDate(0, 0, 5)
				match.EndDate = pgtype.Date{Time: end_date, Valid: true}
			}

			continue
		}

		if row[1] == "bowl_out" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				mainError = err
				return
			}

			match.BowlOutWinnerId = teamField
			continue
		}

		if row[1] == "eliminator" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				mainError = err
				return
			}

			match.SuperOverWinnerId = teamField
			continue
		}

		if row[1] == "winner" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				mainError = err
				return
			}

			match.MatchWinnerId = teamField

			if teamField.Int64 == match.Team1Id.Int64 {
				match.MatchLoserId = match.Team2Id
			} else {
				match.MatchLoserId = match.Team1Id
			}

			continue
		}

		if row[1] == "winner_innings" {
			if row[2] == "1" {
				match.IsWonByInnings = pgtype.Bool{Bool: true, Valid: true}
			} else {
				match.IsWonByInnings = pgtype.Bool{Bool: false, Valid: true}
			}
			continue
		}

		if row[1] == "winner_runs" || row[1] == "winner_wickets" {
			margin, err := strconv.ParseInt(row[2], 10, 64)
			if err != nil {
				mainError = err
				return
			}

			match.WinMargin = pgtype.Int8{Int64: margin, Valid: true}

			if row[1] == "winner_runs" {
				match.IsWonByRuns = pgtype.Bool{Bool: true, Valid: true}
			} else {
				match.IsWonByRuns = pgtype.Bool{Bool: false, Valid: true}
			}
			continue
		}

		if row[1] == "toss_winner" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				mainError = err
				return
			}

			match.TossWinnerId = teamField

			if teamField.Int64 == match.Team1Id.Int64 {
				match.TossLoserId = match.Team2Id
			} else {
				match.TossLoserId = match.Team1Id
			}

			continue
		}

		if row[1] == "toss_decision" {
			if row[2] == "bat" {
				match.IsTossDecisionBat = pgtype.Bool{Bool: true, Valid: true}
			} else if row[2] == "field" {
				match.IsTossDecisionBat = pgtype.Bool{Bool: false, Valid: true}
			}
			continue
		}

		if row[1] == "player" || row[1] == "players" {
			playerName := row[3]
			if row[2] == team1.Name {
				team1.Players[playerName] = ""
			} else {
				team2.Players[playerName] = ""
			}
			continue
		}

		if row[1] == "registry" && row[2] == "people" {
			playerName, cricsheetId := row[3], row[4]

			if _, ok := team1.Players[playerName]; ok {
				if err := handlePlayer(cricsheetId, match.IsMale.Bool); err != nil {
					mainError = err
					return
				}
				team1.Players[playerName] = cricsheetId
			} else if _, ok := team2.Players[playerName]; ok {
				if err := handlePlayer(cricsheetId, match.IsMale.Bool); err != nil {
					mainError = err
					return
				}
				team2.Players[playerName] = cricsheetId
			}
			continue
		}

		if row[1] == "event" {
			eventName = row[2]
			continue
		}

		if row[1] == "venue" {
			venue = row[2]
			continue
		}

		if row[1] == "city" {
			city = row[2]
			continue
		}

		if row[1] == "player_of_match" {
			potmNames = append(potmNames, row[2])
			continue
		}
	}

	ground, err := handleVenue(venue, city)
	if err != nil {
		mainError = fmt.Errorf(`error while handling venue: %v`, err)
		return
	}

	match.GroundId = ground.Id
	if !match.IsNeutralVenue.Bool {
		if ground.HostNationName.String == team1.Name {
			match.HomeTeamId = pgtype.Int8{Int64: team1.Id, Valid: true}
			match.AwayTeamId = pgtype.Int8{Int64: team2.Id, Valid: true}
		} else if ground.HostNationName.String == team2.Name {
			match.HomeTeamId = pgtype.Int8{Int64: team2.Id, Valid: true}
			match.AwayTeamId = pgtype.Int8{Int64: team1.Id, Valid: true}
		}
	}

	match.SeriesId, err = handleEvent(eventName, &match)
	if err != nil {
		mainError = fmt.Errorf(`error while handling event: %v`, err)
		return
	}

	matchId, err := dbutils.UpsertCricsheetMatch(context.Background(), DB_POOL, &match)
	if err != nil {
		mainError = err
		return
	}

	match.Id = pgtype.Int8{Int64: matchId, Valid: true}

	if err = insertSquadEntries(team1, match.Id.Int64, match.SeriesId.Int64); err != nil {
		mainError = fmt.Errorf(`error while handling squad entries of team1: %v`, err)
		return
	}

	if err = insertSquadEntries(team2, match.Id.Int64, match.SeriesId.Int64); err != nil {
		mainError = fmt.Errorf(`error while handling squad entries of team2: %v`, err)
		return
	}

	if err = insertPotmEntries(potmNames, team1, team2, match.Id.Int64); err != nil {
		mainError = fmt.Errorf(`error while inserting POTM entries: %v`, err)
		return
	}
}

func handlePlayer(cricsheetId string, isMale bool) error {
	cacheKey := PlayerKey{CricsheetId: cricsheetId}
	_, exists := PlayersCache.Get(cacheKey)
	if exists {
		return nil
	}

	filters := url.Values{
		"cricsheet_id": []string{cricsheetId},
		"__limit":      []string{"1"},
	}

	PlayerLock.Lock()
	defer PlayerLock.Unlock()

	dbResponse, err := dbutils.ReadPlayers(context.Background(), DB_POOL, filters)
	if err != nil {
		return fmt.Errorf(`error while reading reading players: %v`, err)
	}

	var player responses.AllPlayers

	if len(dbResponse.Players) == 0 {
		cricsheetPeople, err := dbutils.ReadCricsheetPeopleById(context.Background(), DB_POOL, cricsheetId)

		if err != nil {
			return fmt.Errorf(`failed to read player with id %s from cricsheet_people: %v`, cricsheetId, err)
		}

		newPlayer := models.Player{
			Name:        cricsheetPeople.Name,
			FullName:    cricsheetPeople.UniqueName,
			CricsheetId: cricsheetPeople.Identifier,
			CricbuzzId:  cricsheetPeople.CricbuzzId,
			CricinfoId:  cricsheetPeople.CricinfoId,
			IsMale:      pgtype.Bool{Bool: isMale, Valid: true},
		}

		playerId, err := dbutils.InsertPlayer(context.Background(), DB_POOL, &newPlayer)
		if err != nil {
			return fmt.Errorf(`error while inserting player: %v`, err)
		}

		player = responses.AllPlayers{
			Id:     pgtype.Int8{Int64: playerId, Valid: true},
			Name:   newPlayer.Name,
			IsMale: pgtype.Bool{Bool: isMale, Valid: true},
		}
	} else {
		player = dbResponse.Players[0]
	}

	PlayersCache.Set(cacheKey, player)
	return nil
}

func handleEvent(eventName string, match *models.Match) (pgtype.Int8, error) {
	hostNations := getTourHostNations(eventName)
	for _, hostNation := range hostNations {
		_, err := handleHostNation(hostNation)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	seriesId, err := handleSeries(eventName, match)
	if err != nil {
		return pgtype.Int8{}, err
	}

	return seriesId, nil
}

func handleSeries(name string, match *models.Match) (pgtype.Int8, error) {
	season, playingLevel, playingFormat, isMale := match.Season.String, match.PlayingLevel.String, match.PlayingFormat.String, match.IsMale.Bool

	cacheKey := SeriesKey{
		Name:          name,
		Season:        season,
		PlayingLevel:  playingLevel,
		PlayingFormat: playingFormat,
		IsMale:        isMale,
	}

	seriesId, ok := SeriesCache.Get(cacheKey)
	defer func() (pgtype.Int8, error) {
		if err := dbutils.UpsertSeriesTeamEntries(context.Background(), DB_POOL, seriesId, []pgtype.Int8{match.Team1Id, match.Team2Id}); err != nil {
			err = fmt.Errorf(`failed to upsert series team entries: %v`, err)
			return pgtype.Int8{}, err
		}
		return pgtype.Int8{Int64: seriesId, Valid: true}, nil
	}()

	if ok {
		return pgtype.Int8{Int64: seriesId, Valid: true}, nil
	}

	filters := url.Values{
		"name":           []string{name},
		"season":         []string{season},
		"playing_level":  []string{playingLevel},
		"playing_format": []string{playingFormat},
		"is_male":        []string{"true"},
		"__limit":        []string{"1"},
	}

	if !isMale {
		filters["is_male"][0] = "false"
	}

	SeriesLock.Lock()
	defer SeriesLock.Unlock()

	dbResponse, err := dbutils.ReadSeries(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Series) == 0 {
		newSeries := models.Series{
			Name:          pgtype.Text{String: name, Valid: true},
			TeamsId:       []pgtype.Int8{match.Team1Id, match.Team2Id},
			Season:        match.Season,
			IsMale:        match.IsMale,
			PlayingLevel:  match.PlayingLevel,
			PlayingFormat: match.PlayingFormat,
		}

		newSeries.TournamentId, err = handleTournament(name, playingLevel, playingFormat, isMale)
		if err != nil {
			return pgtype.Int8{}, err
		}

		seriesId, err = dbutils.InsertSeries(context.Background(), DB_POOL, &newSeries)
		if err != nil {
			return pgtype.Int8{}, err
		}
	} else {
		seriesId = dbResponse.Series[0].Id.Int64
		if err := dbutils.UpsertSeriesTeamEntries(context.Background(), DB_POOL, seriesId, []pgtype.Int8{match.Team1Id, match.Team2Id}); err != nil {
			err = fmt.Errorf(`failed to upsert series team entries: %v`, err)
			return pgtype.Int8{}, err
		}
	}

	SeriesCache.Set(cacheKey, seriesId)
	return pgtype.Int8{Int64: seriesId, Valid: true}, nil
}

func handleTournament(tournamentName, playingLevel, playingFormat string, isMale bool) (pgtype.Int8, error) {
	cacheKey := TournamentKey{
		Name:          tournamentName,
		IsMale:        isMale,
		PlayingLevel:  playingLevel,
		PlayingFormat: playingFormat,
	}

	tournamentId, exists := TournamentsCache.Get(cacheKey)
	if exists {
		return pgtype.Int8{Int64: tournamentId, Valid: true}, nil
	}

	filters := url.Values{
		"name":           []string{tournamentName},
		"playing_level":  []string{playingLevel},
		"playing_format": []string{playingFormat},
		"is_male":        []string{"true"},
		"__limit":        []string{"0"},
	}

	if !isMale {
		filters["is_male"][0] = "false"
	}

	TournamentLock.Lock()
	defer TournamentLock.Unlock()

	dbResponse, err := dbutils.ReadTournaments(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Tournaments) == 0 {
		return pgtype.Int8{}, nil
	}

	tournamentId = dbResponse.Tournaments[0].Id.Int64
	TournamentsCache.Set(cacheKey, tournamentId)
	return pgtype.Int8{Int64: tournamentId, Valid: true}, nil
}

func handleTeam(teamName, playingLevel string, isMale bool) (pgtype.Int8, error) {
	cacheKey := TeamKey{TeamName: teamName, IsMale: isMale, PlayingLevel: playingLevel}
	teamId, exists := TeamsCache.Get(cacheKey)
	if exists {
		return pgtype.Int8{Int64: teamId, Valid: true}, nil
	}

	filters := url.Values{
		"name":          []string{teamName},
		"playing_level": []string{playingLevel},
		"is_male":       []string{"true"},
		"__limit":       []string{"1"},
	}

	if !isMale {
		filters["is_male"][0] = "false"
	}

	TeamLock.Lock()
	defer TeamLock.Unlock()

	dbResponse, err := dbutils.ReadTeams(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Teams) == 0 {
		newTeam := models.Team{
			Name:         pgtype.Text{String: teamName, Valid: true},
			PlayingLevel: pgtype.Text{String: playingLevel, Valid: true},
			IsMale:       pgtype.Bool{Bool: isMale, Valid: true},
			ShortName:    pgtype.Text{String: defaultTeamShortName(teamName), Valid: true},
		}

		teamId, err = dbutils.InsertTeam(context.Background(), DB_POOL, &newTeam)
		if err != nil {
			return pgtype.Int8{}, err
		}
	} else {
		teamId = dbResponse.Teams[0].Id.Int64
	}

	TeamsCache.Set(cacheKey, teamId)
	teamField := pgtype.Int8{Int64: teamId, Valid: true}

	return teamField, nil
}

func handleSeason(season string) (pgtype.Text, error) {
	seasonExists, _ := SeasonsCache.Get(season)
	if seasonExists {
		return pgtype.Text{String: season, Valid: true}, nil
	}

	filters := url.Values{"season": []string{season}, "__limit": []string{"1"}}

	SeasonLock.Lock()
	defer SeasonLock.Unlock()

	dbResponse, err := dbutils.ReadSeasons(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Text{}, err
	}

	if len(dbResponse.Seasons) == 0 {
		newSeason := models.Season{Season: pgtype.Text{String: season, Valid: true}}
		err := dbutils.InsertSeason(context.Background(), DB_POOL, &newSeason)
		if err != nil {
			return pgtype.Text{}, err
		}
	}

	SeasonsCache.Set(season, true)
	return pgtype.Text{String: season, Valid: true}, nil
}

func handleVenue(venue, city string) (responses.AllGrounds, error) {
	cacheKey := GroundKey{Venue: venue, City: city}
	ground, exists := GroundsCache.Get(cacheKey)
	if exists {
		return ground, nil
	}

	cityId, err := handleCity(city)
	if err != nil {
		return responses.AllGrounds{}, fmt.Errorf(`error while handling city: %v`, err)
	}

	filters := url.Values{
		"name":    []string{venue},
		"__limit": []string{"1"},
	}

	if cityId.Valid {
		filters["city_id"] = []string{strconv.FormatInt(cityId.Int64, 10)}
	}

	GroundLock.Lock()
	defer GroundLock.Unlock()

	dbResponse, err := dbutils.ReadGrounds(context.Background(), DB_POOL, filters)
	if err != nil {
		return responses.AllGrounds{}, fmt.Errorf(`error while reading grounds: %v`, err)
	}

	if len(dbResponse.Grounds) == 0 {
		newGround := models.Ground{
			Name:   pgtype.Text{String: venue, Valid: true},
			CityId: cityId,
		}

		groundId, err := dbutils.InsertGround(context.Background(), DB_POOL, &newGround)
		if err != nil {
			return responses.AllGrounds{}, fmt.Errorf(`error while inserting ground: %v`, err)
		}

		ground = responses.AllGrounds{
			Id:     pgtype.Int8{Int64: groundId, Valid: true},
			Name:   newGround.Name,
			CityId: newGround.CityId,
		}
	} else {
		ground = dbResponse.Grounds[0]
	}

	GroundsCache.Set(cacheKey, ground)

	return ground, nil
}

func handleCity(cityName string) (pgtype.Int8, error) {
	if cityName == "" {
		return pgtype.Int8{}, nil
	}

	cityId, exists := CitiesCache.Get(cityName)
	if exists {
		return pgtype.Int8{Int64: cityId, Valid: true}, nil
	}

	filters := url.Values{
		"name":    []string{cityName},
		"__limit": []string{"1"},
	}

	CityLock.Lock()
	defer CityLock.Unlock()

	dbResponse, err := dbutils.ReadCities(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, fmt.Errorf(`error while reading cities: %v`, err)
	}

	if len(dbResponse.Cities) == 0 {
		newCity := models.City{
			Name: pgtype.Text{String: cityName, Valid: true},
		}

		cityId, err = dbutils.InsertCity(context.Background(), DB_POOL, &newCity)
		if err != nil {
			return pgtype.Int8{}, fmt.Errorf(`error while inserting city: %v`, err)
		}
	} else {
		cityId = dbResponse.Cities[0].Id.Int64
	}

	CitiesCache.Set(cityName, cityId)
	return pgtype.Int8{Int64: cityId, Valid: true}, nil
}

func handleHostNation(hostNationName string) (pgtype.Int8, error) {
	hostNationId, exists := HostNationCache.Get(hostNationName)
	if exists {
		return pgtype.Int8{Int64: hostNationId, Valid: true}, nil
	}

	filters := url.Values{
		"name":    []string{hostNationName},
		"__limit": []string{"1"},
	}

	HostNationLock.Lock()
	defer HostNationLock.Unlock()

	dbResponse, err := dbutils.ReadHostNations(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.HostNations) == 0 {
		newHostNation := models.HostNation{
			Name: pgtype.Text{String: hostNationName, Valid: true},
		}

		hostNationId, err = dbutils.InsertHostNation(context.Background(), DB_POOL, &newHostNation)
		if err != nil {
			return pgtype.Int8{}, err
		}
	} else {
		hostNationId = dbResponse.HostNations[0].Id.Int64
	}

	HostNationCache.Set(hostNationName, hostNationId)
	hostNationField := pgtype.Int8{Int64: hostNationId, Valid: true}

	return hostNationField, nil
}

func handleMatchOutcome(outcome string) pgtype.Text {
	switch outcome {
	case "draw":
		return pgtype.Text{String: "drawn", Valid: true}
	case "tie":
		return pgtype.Text{String: "tied", Valid: true}
	case "no result":
		return pgtype.Text{String: "no_result", Valid: true}
	default:
		return pgtype.Text{String: "winner_decided", Valid: true}
	}
}

func getTourHostNations(eventName string) []string {
	pattern := `^([A-Za-z\s]+) tour of ([A-Za-z\s]+(?: and [A-Za-z\s]+)*)$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(eventName)
	if len(matches) == 3 {
		hostNations := matches[2]
		hostNationList := regexp.MustCompile(`\s+and\s+`).Split(hostNations, -1)
		return hostNationList
	}
	return nil
}

func defaultTeamShortName(teamName string) string {
	words := strings.Split(teamName, " ")
	if len(words) == 1 {
		return strings.ToUpper(teamName[:3])
	}

	var shortName string

	for _, word := range words {
		shortName += strings.ToUpper(string(word[0]))
	}

	return shortName
}

func handleSeriesSquadId(teamInfo TeamInfo, seriesId int64) (pgtype.Int8, error) {
	cacheKey := SeriesSquadKey{TeamId: teamInfo.Id, SeriesId: seriesId}

	squadId, exists := SeriesSquadCache.Get(cacheKey)
	if exists {
		return pgtype.Int8{Int64: squadId, Valid: true}, nil
	}

	filters := url.Values{
		"team_id":   []string{strconv.FormatInt(teamInfo.Id, 10)},
		"series_id": []string{strconv.FormatInt(seriesId, 10)},
		"__limit":   []string{"1"},
	}

	SeriesSquadLock.Lock()
	defer SeriesSquadLock.Unlock()

	dbResponse, err := dbutils.ReadSeriesSquads(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Squads) == 0 {
		squadLabel := fmt.Sprintf(`%s squad`, teamInfo.Name)

		newSeriesSquad := models.SeriesSquad{
			SeriesId:   pgtype.Int8{Int64: seriesId, Valid: true},
			TeamId:     pgtype.Int8{Int64: teamInfo.Id, Valid: true},
			SquadLabel: pgtype.Text{String: squadLabel, Valid: true},
		}

		squadId, err = dbutils.InsertSeriesSquad(context.Background(), DB_POOL, &newSeriesSquad)
		if err != nil {
			return pgtype.Int8{}, err
		}
	} else {
		squadId = dbResponse.Squads[0].Id.Int64
	}

	SeriesSquadCache.Set(cacheKey, squadId)
	return pgtype.Int8{Int64: squadId, Valid: true}, nil
}

func insertSquadEntries(teamInfo TeamInfo, matchId, seriesId int64) error {
	squadId, err := handleSeriesSquadId(teamInfo, seriesId)
	if err != nil {
		return err
	}

	for playerName, cricsheetId := range teamInfo.Players {
		player, ok := PlayersCache.Get(PlayerKey{CricsheetId: cricsheetId})
		if !ok {
			return fmt.Errorf(`player %s not found in cache`, playerName)
		}

		playerTeamEntry := models.PlayerTeamEntry{
			PlayerId: player.Id,
			TeamId:   pgtype.Int8{Int64: teamInfo.Id, Valid: true},
		}

		if err := dbutils.UpsertPlayerTeamEntry(context.Background(), DB_POOL, &playerTeamEntry); err != nil {
			return fmt.Errorf(`failed to upsert team entry of player %s: %v`, playerName, err)
		}

		matchSquadEntry := models.MatchSquad{
			PlayerId:      player.Id,
			TeamId:        pgtype.Int8{Int64: teamInfo.Id, Valid: true},
			MatchId:       pgtype.Int8{Int64: matchId, Valid: true},
			PlayingStatus: pgtype.Text{String: "playing_xi", Valid: true},
		}

		if err := dbutils.UpsertMatchSquadEntry(context.Background(), DB_POOL, &matchSquadEntry); err != nil {
			return fmt.Errorf(`error while match squad upsertion of %s: %v`, playerName, err)
		}

		seriesSquadEntry := models.SeriesSquadEntry{
			PlayerId: player.Id,
			SquadId:  squadId,
		}

		if err := dbutils.UpsertSeriesSquadEntry(context.Background(), DB_POOL, &seriesSquadEntry); err != nil {
			return fmt.Errorf(`error while series squad upsertion of %s: %v`, playerName, err)
		}
	}

	return nil
}

func insertPotmEntries(potmNames []string, team1, team2 TeamInfo, matchId int64) error {
	for _, potmName := range potmNames {
		var cacheKey PlayerKey

		if cricsheetId, ok := team1.Players[potmName]; ok {
			cacheKey.CricsheetId = cricsheetId
		} else if cricsheetId, ok := team2.Players[potmName]; ok {
			cacheKey.CricsheetId = cricsheetId
		} else {
			return fmt.Errorf(`player %s not found in either team`, potmName)
		}

		player, ok := PlayersCache.Get(cacheKey)
		if !ok {
			return fmt.Errorf(`player %s not found in PlayersCache`, potmName)
		}

		awardEntry := models.PlayerAwardEntry{
			PlayerId:  player.Id,
			MatchId:   pgtype.Int8{Int64: matchId, Valid: true},
			AwardType: pgtype.Text{String: "player_of_the_match", Valid: true},
		}

		if err := dbutils.InsertPlayerAwardEntry(context.Background(), DB_POOL, &awardEntry); err != nil {
			return fmt.Errorf(`db error %v`, err)
		}
	}

	return nil
}
