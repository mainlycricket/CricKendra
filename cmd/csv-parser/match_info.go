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
	Id      int8
	Name    string
	Players map[string]int64
}

func extractMatchInfo(filePath, playingLevel, playingFormat string) (models.Match, error) {
	var match models.Match
	var team1, team2 TeamInfo
	var potmName, venue, city, eventName, alternateEventName string

	team1.Players, team2.Players = make(map[string]int64, 11), make(map[string]int64, 11)

	fp, err := os.Open(filePath)
	if err != nil {
		return match, err
	}
	defer fp.Close()

	matchId := strings.TrimSuffix(filepath.Base(filePath), "_info.csv")
	match.CricsheetId = pgtype.Text{String: matchId, Valid: true}

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
			return match, err
		}

		if row[0] != "info" {
			continue
		}

		if row[1] == "balls_per_over" {
			balls_per_over, err := strconv.ParseInt(row[2], 10, 64)
			if err != nil {
				return match, err
			}

			match.BallsPerOver = pgtype.Int8{Int64: balls_per_over, Valid: true}
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
				return match, err
			}

			match.EventMatchNumber = pgtype.Int8{Int64: match_number, Valid: true}
			continue
		}

		if row[1] == "gender" {
			if row[2] == "male" {
				match.IsMale = pgtype.Bool{Bool: true, Valid: true}
			} else {
				match.IsMale = pgtype.Bool{Bool: false, Valid: true}
			}
			continue
		}

		if row[1] == "season" {
			match.Season, err = handleSeason(row[2])
			if err != nil {
				return match, err
			}
			continue
		}

		if row[1] == "team" {
			teamName := row[2]

			teamField, err := handleTeam(teamName, playingLevel, match.IsMale.Bool)
			if err != nil {
				return match, err
			}

			if match.Team1Id.Valid {
				match.Team2Id = teamField
				team2.Id, team2.Name = int8(teamField.Int64), teamName
			} else {
				match.Team1Id = teamField
				team1.Id, team1.Name = int8(teamField.Int64), teamName
			}
			continue
		}

		if row[1] == "date" {
			date, err := time.Parse("2006/01/02", row[2])
			if err != nil {
				return match, err
			}

			match.StartDate = pgtype.Date{Time: date, Valid: true}
			continue
		}

		if row[1] == "bowl_out" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				return match, err
			}

			match.BowlOutWinnerId = teamField
			continue
		}

		if row[1] == "eliminator" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				return match, err
			}

			match.SuperOverWinnerId = teamField
			continue
		}

		if row[1] == "winner" {
			teamField, err := handleTeam(row[2], playingLevel, match.IsMale.Bool)
			if err != nil {
				return match, err
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
				return match, err
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
				return match, err
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

		if row[1] == "player" {
			playerName := row[3]
			if row[2] == team1.Name {
				team1.Players[playerName] = 0
			} else {
				team2.Players[playerName] = 0
			}
			continue
		}

		if row[1] == "registry" && row[2] == "people" {
			playerName, cricsheetId := row[3], row[4]

			if _, ok := team1.Players[playerName]; ok {
				player_id, err := handlePlayer(cricsheetId, match.IsMale.Bool, int64(team1.Id))
				if err != nil {
					return match, err
				}
				team1.Players[playerName] = player_id
			} else if _, ok := team2.Players[playerName]; ok {
				player_id, err := handlePlayer(cricsheetId, match.IsMale.Bool, int64(team2.Id))
				if err != nil {
					return match, err
				}
				team2.Players[playerName] = player_id
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
			potmName = row[2]
			continue
		}
	}

	if potmId, ok := team1.Players[potmName]; ok {
		match.PoTMsId = append(match.PoTMsId, pgtype.Int8{Int64: potmId, Valid: true})
	} else if potmId, ok := team2.Players[potmName]; ok {
		match.PoTMsId = append(match.PoTMsId, pgtype.Int8{Int64: potmId, Valid: true})
	}

	match.GroundId, err = handleVenue(venue, city)
	if err != nil {
		return match, err
	}

	if match.IsNeutralVenue.Bool {
		alternateEventName = fmt.Sprintf(`%s in %s %s Series`, team1.Name, team2.Name, playingFormat)
	} else {
		match.HomeTeamId = match.Team1Id
		match.AwayTeamId = match.Team2Id
		alternateEventName = fmt.Sprintf(`%s v %s %s Series`, team1.Name, team2.Name, playingFormat)
	}

	match.SeriesId, err = handleEvent(eventName, alternateEventName, &match)
	if err != nil {
		return match, err
	}

	return match, nil
}

func handlePlayer(cricsheetId string, isMale bool, teamId int64) (int64, error) {
	cacheKey := PlayerKey{CricsheetId: cricsheetId}
	player, exists := PlayersCache[cacheKey]
	if exists {
		return player.Id.Int64, nil
	}

	filters := url.Values{
		"cricsheet_id": []string{cricsheetId},
		"__limit":      []string{"1"},
	}

	dbResponse, err := dbutils.ReadPlayers(context.Background(), DB_POOL, filters)
	if err != nil {
		return 0, err
	}

	if len(dbResponse.Players) == 0 {
		cricsheetPeople, err := dbutils.ReadCricsheetPeopleById(context.Background(), DB_POOL, cricsheetId)

		if err != nil {
			return 0, err
		}

		newPlayer := models.Player{
			Name:               cricsheetPeople.Name,
			FullName:           cricsheetPeople.UniqueName,
			CricsheetId:        cricsheetPeople.Identifier,
			CricbuzzId:         cricsheetPeople.CricbuzzId,
			CricinfoId:         cricsheetPeople.CricinfoId,
			IsMale:             pgtype.Bool{Bool: isMale, Valid: true},
			TeamsRepresentedId: []pgtype.Int8{{Int64: teamId, Valid: true}},
		}

		err = dbutils.InsertPlayer(context.Background(), DB_POOL, &newPlayer)
		if err != nil {
			return 0, err
		}

		dbResponse, err = dbutils.ReadPlayers(context.Background(), DB_POOL, filters)
		if err != nil {
			return 0, err
		}
	}

	player = dbResponse.Players[0]
	PlayersCache[cacheKey] = player
	return player.Id.Int64, nil
}

func handleEvent(eventName, alternateEventName string, match *models.Match) (pgtype.Int8, error) {
	season, isMale := match.Season.String, match.IsMale.Bool
	playingLevel, playingFormat := match.PlayingLevel.String, match.PlayingFormat.String
	var err error

	touringTeam, hostNations, isTour := detectTour(eventName)

	if isTour {
		tourId, err := handleTour(touringTeam, season, playingLevel, hostNations, isMale)
		if err != nil {
			return pgtype.Int8{}, err
		}
		match.TourId = tourId
		eventName = alternateEventName
	} else {
		match.TournamentId, err = handleTournament(eventName, playingLevel, playingFormat, isMale)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	filters := url.Values{
		"name":           []string{eventName},
		"season":         []string{season},
		"playing_level":  []string{playingLevel},
		"playing_format": []string{playingFormat},
		"is_male":        []string{"true"},
		"__limit":        []string{"1"},
	}

	if !isMale {
		filters["is_male"][0] = "false"
	}

	dbResponse, err := dbutils.ReadSeries(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Series) == 0 {
		newSeries := models.Series{
			Name:          pgtype.Text{String: eventName, Valid: true},
			TeamsId:       []pgtype.Int8{match.Team1Id, match.Team2Id},
			TourId:        match.TourId,
			Season:        match.Season,
			IsMale:        match.IsMale,
			PlayingLevel:  match.PlayingLevel,
			PlayingFormat: match.PlayingFormat,
		}

		err = dbutils.InsertSeries(context.Background(), DB_POOL, &newSeries)
		if err != nil {
			return pgtype.Int8{}, err
		}

		dbResponse, err = dbutils.ReadSeries(context.Background(), DB_POOL, filters)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	match.SeriesId = dbResponse.Series[0].Id

	return pgtype.Int8{}, nil
}

func handleTournament(tournamentName, playingLevel, playingFormat string, isMale bool) (pgtype.Int8, error) {
	cacheKey := TournamentKey{
		Name:          tournamentName,
		IsMale:        isMale,
		PlayingLevel:  playingLevel,
		PlayingFormat: playingFormat,
	}

	tournament, exists := TournamentsCache[cacheKey]
	if exists {
		return tournament.Id, nil
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

	dbResponse, err := dbutils.ReadTournaments(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Tournaments) == 0 {
		return pgtype.Int8{}, nil
	}

	tournament = dbResponse.Tournaments[0]
	TournamentsCache[cacheKey] = tournament
	return tournament.Id, nil
}

func handleTour(touringTeam, season, playingLevel string, hostNations []string, isMale bool) (pgtype.Int8, error) {
	touringTeamId, err := handleTeam(touringTeam, playingLevel, isMale)
	if err != nil {
		return pgtype.Int8{}, err
	}

	var hostNationsId []pgtype.Int8
	for _, hostNationName := range hostNations {
		hostNationId, err := handleHostNation(hostNationName)
		if err != nil {
			return pgtype.Int8{}, err
		}
		hostNationsId = append(hostNationsId, hostNationId)
	}

	var hostNationsStr []string
	for _, id := range hostNationsId {
		hostNationsStr = append(hostNationsStr, strconv.FormatInt(id.Int64, 10))
	}

	cacheKey := TourKey{
		TouringTeamId: touringTeamId.Int64,
		HostNationsId: strings.Join(hostNationsStr, "_"),
		Season:        season,
	}

	tour, exists := ToursCache[cacheKey]
	if exists {
		return tour.Id, nil
	}

	filters := url.Values{
		"touring_team_id": []string{strconv.FormatInt(touringTeamId.Int64, 10)},
		"host_nations_id": hostNationsStr,
		"season":          []string{season},
		"__limit":         []string{"1"},
	}

	dbResponse, err := dbutils.ReadTours(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Tours) == 0 {
		newTour := models.Tour{
			TouringTeamId: touringTeamId,
			HostNationsId: hostNationsId,
			Season:        pgtype.Text{String: season, Valid: true},
		}

		err := dbutils.InsertTour(context.Background(), DB_POOL, &newTour)
		if err != nil {
			return pgtype.Int8{}, err
		}

		dbResponse, err = dbutils.ReadTours(context.Background(), DB_POOL, filters)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	tour = dbResponse.Tours[0]
	ToursCache[cacheKey] = tour

	return tour.Id, nil
}

func handleTeam(teamName, playingLevel string, isMale bool) (pgtype.Int8, error) {
	cacheKey := TeamKey{TeamName: teamName, IsMale: isMale, PlayingLevel: playingLevel}
	team, exists := TeamsCache[cacheKey]
	if exists {
		return pgtype.Int8{Int64: team.Id.Int64, Valid: true}, nil
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

	dbResponse, err := dbutils.ReadTeams(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Teams) == 0 {
		newTeam := models.Team{
			Name:         pgtype.Text{String: teamName, Valid: true},
			PlayingLevel: pgtype.Text{String: playingLevel, Valid: true},
			IsMale:       pgtype.Bool{Bool: isMale, Valid: true},
		}

		err := dbutils.InsertTeam(context.Background(), DB_POOL, &newTeam)
		if err != nil {
			return pgtype.Int8{}, err
		}

		dbResponse, err = dbutils.ReadTeams(context.Background(), DB_POOL, filters)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	team = dbResponse.Teams[0]
	TeamsCache[cacheKey] = team
	teamField := pgtype.Int8{Int64: team.Id.Int64, Valid: true}

	return teamField, nil
}

func handleSeason(season string) (pgtype.Text, error) {
	seasonExists := SeasonsCache[season]
	if seasonExists {
		return pgtype.Text{String: season, Valid: true}, nil
	}

	filters := url.Values{"season": []string{season}, "__limit": []string{"1"}}

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

	SeasonsCache[season] = true
	return pgtype.Text{String: season, Valid: true}, nil
}

func handleVenue(venue, city string) (pgtype.Int8, error) {
	cacheKey := GroundKey{Venue: venue, City: city}
	ground, exists := GroundsCache[cacheKey]
	if exists {
		return ground.Id, nil
	}

	cityId, err := handleCity(city)
	if err != nil {
		return pgtype.Int8{}, err
	}

	filters := url.Values{
		"name":    []string{venue},
		"city_id": []string{strconv.FormatInt(cityId.Int64, 10)},
		"__limit": []string{"1"},
	}

	dbResponse, err := dbutils.ReadGrounds(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Grounds) == 0 {
		newGround := models.Ground{
			Name:   pgtype.Text{String: venue, Valid: true},
			CityId: cityId,
		}

		err := dbutils.InsertGround(context.Background(), DB_POOL, &newGround)
		if err != nil {
			return pgtype.Int8{}, err
		}

		dbResponse, err = dbutils.ReadGrounds(context.Background(), DB_POOL, filters)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	ground = dbResponse.Grounds[0]
	GroundsCache[cacheKey] = ground

	return ground.Id, nil
}

func handleCity(cityName string) (pgtype.Int8, error) {
	city, exists := CitiesCache[cityName]
	if exists {
		return city.Id, nil
	}

	filters := url.Values{
		"name":    []string{cityName},
		"__limit": []string{"1"},
	}

	dbResponse, err := dbutils.ReadCities(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Cities) == 0 {
		newCity := models.City{
			Name: pgtype.Text{String: cityName, Valid: true},
		}

		err := dbutils.InsertCity(context.Background(), DB_POOL, &newCity)
		if err != nil {
			return pgtype.Int8{}, err
		}

		dbResponse, err = dbutils.ReadCities(context.Background(), DB_POOL, filters)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	city = dbResponse.Cities[0]
	CitiesCache[cityName] = city
	return city.Id, nil
}

func handleHostNation(hostNationName string) (pgtype.Int8, error) {
	hostNation, exists := HostNationCache[hostNationName]
	if exists {
		return pgtype.Int8{Int64: hostNation.Id.Int64, Valid: true}, nil
	}

	filters := url.Values{
		"name":    []string{hostNationName},
		"__limit": []string{"1"},
	}

	dbResponse, err := dbutils.ReadHostNations(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.HostNations) == 0 {
		newHostNation := models.HostNation{
			Name: pgtype.Text{String: hostNationName, Valid: true},
		}

		err := dbutils.InsertHostNation(context.Background(), DB_POOL, &newHostNation)
		if err != nil {
			return pgtype.Int8{}, err
		}

		dbResponse, err = dbutils.ReadHostNations(context.Background(), DB_POOL, filters)
		if err != nil {
			return pgtype.Int8{}, err
		}
	}

	hostNation = dbResponse.HostNations[0]
	HostNationCache[hostNationName] = hostNation
	hostNationField := pgtype.Int8{Int64: hostNation.Id.Int64, Valid: true}

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

func detectTour(eventName string) (string, []string, bool) {
	pattern := `^([A-Za-z\s]+) tour of ([A-Za-z\s]+(?: and [A-Za-z\s]+)*)$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(eventName)
	if len(matches) == 3 {
		teamName := matches[1]
		hostNations := matches[2]
		hostNationList := regexp.MustCompile(`\s+and\s+`).Split(hostNations, -1)
		return teamName, hostNationList, true
	}
	return "", nil, false
}
