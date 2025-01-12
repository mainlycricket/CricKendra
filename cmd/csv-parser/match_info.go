package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
)

type teamInfo struct {
	id      int64
	name    string
	players map[string]string // map[playerName]cricsheetId
}

func newTeamInfo() *teamInfo {
	teamInfo := teamInfo{}
	teamInfo.players = make(map[string]string, 11)
	return &teamInfo
}

type matchInfo struct {
	infoFilePath string
	match        *models.Match
	team1Info    *teamInfo
	team2Info    *teamInfo
}

type matchInfoInput struct {
	filePath, cricsheetId              string
	playingFormat, playingLevel        string
	balls_per_over                     int64
	team1Name, team2Name               string
	is_male                            bool
	season                             string
	dates                              []time.Time
	event_match_number                 int64
	eventName                          string
	groundName, cityName               string
	is_neutral_venue                   bool
	toss_winner_name                   string
	is_toss_decision_bat               bool
	player_of_match_names              []string
	match_winner_name                  string
	is_won_by_innings, is_won_by_runs  bool
	outcome                            string // no result, draw, tie (default: winner decided)
	special_method                     string // D/L, VJD, Awarded, 1st innings score, Lost fewer wickets
	bowl_out_winner, super_over_winner string
	win_margin                         int64
	team1Players, team2Players         map[string]string // name -> cricsheet_id
}

func newInfoInput(filePath, playingFormat string) *matchInfoInput {
	var infoInput matchInfoInput

	infoInput.dates = make([]time.Time, 0, 1)
	infoInput.team1Players = make(map[string]string, 11)
	infoInput.team2Players = make(map[string]string, 11)
	infoInput.player_of_match_names = make([]string, 0, 1)

	infoInput.filePath = filePath
	infoInput.cricsheetId = strings.TrimSuffix(filepath.Base(filePath), "_info.csv")
	infoInput.playingFormat = playingFormat
	infoInput.playingLevel = "international"
	if !slices.Contains([]string{"Test", "ODI", "T20I"}, playingFormat) {
		infoInput.playingFormat = "domestic"
	}

	infoInput.is_neutral_venue = false
	infoInput.is_won_by_innings = false
	infoInput.outcome = "winner decided"

	return &infoInput
}

func parseMatchInfoFile(filePath, playingFormat string, channel chan<- info_parse_response) {
	var mainErr error
	infoInput := newInfoInput(filePath, playingFormat)

	defer func() {
		response := info_parse_response{}
		if mainErr != nil {
			response.err = fmt.Errorf(`%s: %v`, infoInput.cricsheetId, mainErr)
		} else {
			response.infoInput = infoInput
		}
		channel <- response
	}()

	fp, err := os.Open(filePath)
	if err != nil {
		mainErr = fmt.Errorf(`error while opening file: %v`, err)
		return
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1 // to have variable number of records per row

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			mainErr = fmt.Errorf(`error while reading file: %v`, err)
			return
		}

		if row[0] != "info" {
			continue
		}

		if row[1] == "balls_per_over" {
			infoInput.balls_per_over, _ = strconv.ParseInt(row[2], 10, 64)
			continue
		}

		if row[1] == "team" {
			if infoInput.team1Name == "" {
				infoInput.team1Name = row[2]
			} else {
				infoInput.team2Name = row[2]
			}
			continue
		}

		if row[1] == "gender" {
			infoInput.is_male = row[2] == "male"
			continue
		}

		if row[1] == "season" {
			infoInput.season = row[2]
			continue
		}

		if row[1] == "date" {
			date, _ := time.Parse("2006/01/02", row[2])
			infoInput.dates = append(infoInput.dates, date)
			continue
		}

		if row[1] == "event" {
			infoInput.eventName = row[2]
			continue
		}

		if row[1] == "match_number" {
			infoInput.event_match_number, _ = strconv.ParseInt(row[2], 10, 64)
			continue
		}

		if row[1] == "venue" {
			infoInput.groundName = row[2]
			continue
		}

		if row[1] == "city" {
			infoInput.cityName = row[2]
			continue
		}

		if row[1] == "neutral_venue" {
			infoInput.is_neutral_venue = row[2] == "true"
			continue
		}

		if row[1] == "toss_winner" {
			infoInput.toss_winner_name = row[2]
			continue
		}

		if row[1] == "toss_decison" {
			infoInput.is_toss_decision_bat = row[2] == "bat"
			continue
		}

		if row[1] == "player_of_match" {
			infoInput.player_of_match_names = append(infoInput.player_of_match_names, row[2])
			continue
		}

		if row[1] == "winner" {
			infoInput.match_winner_name = row[2]
			continue
		}

		if row[1] == "winner_runs" {
			infoInput.is_won_by_runs = true
			infoInput.win_margin, _ = strconv.ParseInt(row[2], 10, 64)
			continue
		}

		if row[1] == "winner_wickets" {
			infoInput.is_won_by_runs = false
			infoInput.win_margin, _ = strconv.ParseInt(row[2], 10, 64)
			continue
		}

		if row[1] == "winner_innings" && row[2] == "1" {
			infoInput.is_won_by_innings = false
			continue
		}

		if row[1] == "method" {
			infoInput.special_method = row[2]
			continue
		}

		if row[1] == "outcome" {
			infoInput.outcome = row[2]
			continue
		}

		if row[1] == "bowl_out" {
			infoInput.bowl_out_winner = row[2]
			continue
		}

		if row[1] == "eliminator" {
			infoInput.super_over_winner = row[2]
			continue
		}

		if row[1] == "player" || row[1] == "players" {
			if row[2] == infoInput.team1Name {
				infoInput.team1Players[row[3]] = ""
			} else {
				infoInput.team2Players[row[3]] = ""
			}
			continue
		}

		if row[1] == "registry" && row[2] == "people" {
			playerName, playerId := row[3], row[4]
			if _, ok := infoInput.team1Players[playerName]; ok {
				infoInput.team1Players[playerName] = playerId
			} else if _, ok := infoInput.team2Players[playerName]; ok {
				infoInput.team2Players[playerName] = playerId
			}
			continue
		}
	}
}

func (input *matchInfoInput) initalizeMatch(channel chan<- match_init_response) {
	var mainError error
	output := &matchInfo{
		infoFilePath: input.filePath,
		match:        &models.Match{},
		team1Info:    &teamInfo{},
		team2Info:    &teamInfo{},
	}

	defer func() {
		var response match_init_response
		if mainError != nil {
			response.err = fmt.Errorf("%s: %v", input.cricsheetId, mainError)
		} else {
			response.matchInfo = output
		}
		channel <- response
	}()

	// NOTE: ORDER IS CRUCIAL

	// season
	if err := cachedSeasons.loadOrStore(seasonKey{season: input.season}); err != nil {
		mainError = fmt.Errorf(`error while handling season: %v`, err)
		return
	}

	// teams
	if err := output.setMatchTeams(input, true); err != nil {
		mainError = fmt.Errorf(`error while setting team1: %d`, err)
		return
	}
	if err := output.setMatchTeams(input, false); err != nil {
		mainError = fmt.Errorf(`error while setting team2: %d`, err)
		return
	}

	// venue
	if err := output.setVenueDetails(input); err != nil {
		mainError = fmt.Errorf(`error while setting venue details: %v`, err)
		return
	}

	// event
	if err := output.setMatchSeries(input); err != nil {
		mainError = fmt.Errorf(`error while setting match series: %v`, err)
		return
	}

	// players
	if err := output.ensurePlayers(input, true); err != nil {
		mainError = fmt.Errorf(`failed to ensure team 1 players: %v`, err)
		return
	}
	if err := output.ensurePlayers(input, false); err != nil {
		mainError = fmt.Errorf(`failed to ensure team 2 players: %v`, err)
		return
	}

	match := output.match

	// match winner & loser
	if input.match_winner_name == output.team1Info.name {
		match.MatchWinnerId, match.MatchLoserId = match.Team1Id, match.Team2Id
	} else if input.match_winner_name == output.team2Info.name {
		match.MatchWinnerId, match.MatchLoserId = match.Team2Id, match.Team1Id
	}

	// toss winner & loser
	if input.toss_winner_name == output.team1Info.name {
		match.TossWinnerId, match.TossLoserId = match.Team1Id, match.Team2Id
	} else if input.toss_winner_name == output.team2Info.name {
		match.TossWinnerId, match.TossLoserId = match.Team2Id, match.Team1Id
	}

	// bowl out winner
	if input.bowl_out_winner == output.team1Info.name {
		match.BowlOutWinnerId = match.Team1Id
	} else if input.bowl_out_winner == output.team2Info.name {
		match.BowlOutWinnerId = match.Team2Id
	}

	// super over winner
	if input.super_over_winner == output.team1Info.name {
		match.SuperOverWinnerId = match.Team1Id
	} else if input.super_over_winner == output.team2Info.name {
		match.SuperOverWinnerId = match.Team2Id
	}

	match.CricsheetId = pgtype.Text{String: input.cricsheetId, Valid: true}
	match.EventMatchNumber = pgtype.Int8{Int64: input.event_match_number, Valid: input.event_match_number > 0}
	match.StartDate = pgtype.Date{Time: input.dates[0], Valid: true}
	match.EndDate = pgtype.Date{Time: input.dates[len(input.dates)-1], Valid: true}
	match.IsMale = pgtype.Bool{Bool: input.is_male, Valid: true}
	match.IsNeutralVenue = pgtype.Bool{Bool: input.is_neutral_venue, Valid: true}
	match.FinalResult = pgtype.Text{String: input.outcome, Valid: input.outcome != ""}
	match.PlayingLevel = pgtype.Text{String: input.playingLevel, Valid: true}
	match.PlayingFormat = pgtype.Text{String: input.playingFormat, Valid: true}
	match.Season = pgtype.Text{String: input.season, Valid: true}
	match.OutcomeSpecialMethod = pgtype.Text{String: input.special_method, Valid: input.special_method != ""}
	match.IsTossDecisionBat = pgtype.Bool{Bool: input.is_toss_decision_bat, Valid: true}
	match.IsWonByInnings = pgtype.Bool{Bool: input.is_won_by_innings, Valid: input.match_winner_name != ""}
	match.IsWonByRuns = pgtype.Bool{Bool: input.is_won_by_runs, Valid: input.match_winner_name != ""}
	match.WinMargin = pgtype.Int8{Int64: input.win_margin, Valid: input.match_winner_name != ""}
	match.BallsPerOver = pgtype.Int8{Int64: input.balls_per_over, Valid: true}
	match.IsBBBDone = pgtype.Bool{Bool: false, Valid: true}

	matchId, err := dbutils.UpsertCricsheetMatch(context.Background(), DB_POOL, match)
	if err != nil {
		mainError = fmt.Errorf(`failed to upsert cricsheet match: %v`, err)
		return
	}

	match.Id = pgtype.Int8{Int64: matchId, Valid: true}

	if err := output.upsertSquadEntries(true); err != nil {
		mainError = fmt.Errorf(`failed to upsert team1 squads: %v`, err)
		return
	}

	if err := output.upsertSquadEntries(false); err != nil {
		mainError = fmt.Errorf(`failed to upsert team2 squads: %v`, err)
		return
	}

	output.match = match
}

func (matchInfo *matchInfo) setMatchTeams(input *matchInfoInput, isTeam1 bool) error {
	teamKey := teamKey{name: input.team2Name, is_male: input.is_male, playing_level: input.playingLevel}
	if isTeam1 {
		teamKey.name = input.team1Name
	}

	teamId, err := cachedTeams.loadOrStore(teamKey)
	if err != nil {
		return err
	}

	teamInfo := &teamInfo{id: teamId.Int64, name: teamKey.name}

	if isTeam1 {
		matchInfo.match.Team1Id = teamId
		teamInfo.players = input.team1Players
		matchInfo.team1Info = teamInfo
	} else {
		matchInfo.match.Team2Id = teamId
		teamInfo.players = input.team2Players
		matchInfo.team2Info = teamInfo
	}

	return nil
}

func (matchInfo *matchInfo) setVenueDetails(input *matchInfoInput) error {
	renameKey := venue{groundName: input.groundName, cityName: input.cityName}
	if renameData, ok := renamedVenues[renameKey]; ok {
		input.groundName = renameData.groundName
		input.cityName = renameData.cityName
	}

	cityId, err := cachedCities.loadOrStore(cityKey{name: input.cityName})
	if err != nil {
		return fmt.Errorf(`error while loading city %s: %v`, input.cityName, err)
	}

	ground, err := cachedGrounds.loadOrStore(groundKey{name: input.groundName, city_id: cityId.Int64})
	if err != nil {
		return fmt.Errorf(`error while loading ground %s: %v`, input.groundName, err)
	}

	matchInfo.match.GroundId = ground.Id

	if !input.is_neutral_venue {
		if ground.HostNationName.String == matchInfo.team1Info.name {
			matchInfo.match.HomeTeamId = matchInfo.match.Team1Id
			matchInfo.match.AwayTeamId = matchInfo.match.Team2Id
		} else if ground.HostNationName.String == matchInfo.team2Info.name {
			matchInfo.match.HomeTeamId = matchInfo.match.Team2Id
			matchInfo.match.AwayTeamId = matchInfo.match.Team1Id
		}
	}

	return nil
}

func (matchInfo *matchInfo) setMatchSeries(input *matchInfoInput) error {
	seriesKey := seriesKey{
		name:           input.eventName,
		is_male:        input.is_male,
		playing_format: input.playingFormat,
		playing_level:  input.playingLevel,
	}

	if renamedSeries, ok := renamedSeriesNames[seriesKey]; ok {
		seriesKey.name = renamedSeries
	}

	seriesKey.season = input.season
	if renamedSeason, ok := renamedSeriesSeaons[seriesKey]; ok {
		seriesKey.season = renamedSeason
	}

	// check if tour
	touringTeam, hostNations := getTourInfo(seriesKey.name)
	if touringTeam != "" {
		for _, hostNation := range hostNations {
			if _, err := cachedHostNations.loadOrStore(hostNationKey{name: hostNation}); err != nil {
				return fmt.Errorf(`failed to load host nation %s: %v`, hostNation, err)
			}
		}

		hostTeam := input.team1Name
		if touringTeam == input.team1Name {
			hostTeam = input.team2Name
		}

		subSeriesKey := seriesKey
		subSeriesKey.name = fmt.Sprintf("%s in %s %s series", touringTeam, hostTeam, input.playingFormat)

		subSeriesId, err := cachedSeries.loadOrStore(subSeriesKey, 1, 2, "tour_series")
		if err != nil {
			return fmt.Errorf(`failed to load tour series: %v`, err)
		}

		matchInfo.match.SeriesListId = append(matchInfo.match.SeriesListId, subSeriesId)
	}

	var tourFlag string
	if len(matchInfo.match.SeriesListId) > 0 {
		tourFlag = "tour_sub_series"
	}

	mainSeriesId, err := cachedSeries.loadOrStore(seriesKey, 1, 2, tourFlag)
	if err != nil {
		return fmt.Errorf(`failed to load main series: %v`, err)
	}

	matchInfo.match.MainSeriesId = mainSeriesId
	matchInfo.match.SeriesListId = append(matchInfo.match.SeriesListId, matchInfo.match.MainSeriesId)

	return nil
}

func (matchInfo *matchInfo) ensurePlayers(input *matchInfoInput, isTeam1 bool) error {
	teamId, players := matchInfo.team2Info.id, matchInfo.team2Info.players
	if isTeam1 {
		teamId, players = matchInfo.team1Info.id, matchInfo.team1Info.players
	}

	for name, cricsheetId := range players {
		if cricsheetId == "" {
			return fmt.Errorf(`no cricsheet id found for player %s: `, name)
		}

		key := playerKey{cricsheet_id: cricsheetId}
		if _, err := cachedPlayers.loadOrStore(key, name, input.is_male, teamId); err != nil {
			return fmt.Errorf(`failed to load player %s: %v`, name, err)
		}
	}

	if isTeam1 {
		matchInfo.team1Info.players = players
	} else {
		matchInfo.team2Info.players = players
	}

	return nil
}

func (matchInfo *matchInfo) upsertSquadEntries(isTeam1 bool) error {
	teamInfo := matchInfo.team2Info
	if isTeam1 {
		teamInfo = matchInfo.team1Info
	}

	playersId, err := teamInfo.getPlayersId()
	if err != nil {
		return err
	}

	seriesSquadKey := seriesSquadKey{team_id: teamInfo.id}

	for _, playerId := range playersId {
		matchSquadEntry := &models.MatchSquad{
			MatchId:       matchInfo.match.Id,
			PlayerId:      pgtype.Int8{Int64: playerId, Valid: true},
			TeamId:        pgtype.Int8{Int64: teamInfo.id, Valid: true},
			PlayingStatus: pgtype.Text{String: "playing_xi", Valid: true},
		}

		if err := dbutils.UpsertMatchSquadEntry(context.Background(), DB_POOL, matchSquadEntry); err != nil {
			return fmt.Errorf(`failed to upsert match squad entry of player %d: %v`, playerId, err)
		}

		for _, seriesId := range matchInfo.match.SeriesListId {
			seriesSquadKey.series_id = seriesId.Int64
			seriesSquadId, err := cachedSeriesSquads.loadOrStore(seriesSquadKey, teamInfo.name)
			if err != nil {
				return fmt.Errorf(`error while loading series %d squad: %v`, seriesId.Int64, err)
			}

			seriesSquadEntry := &models.SeriesSquadEntry{
				PlayerId: pgtype.Int8{Int64: playerId, Valid: true},
				SquadId:  seriesSquadId,
			}

			if err := dbutils.UpsertSeriesSquadEntry(context.Background(), DB_POOL, seriesSquadEntry); err != nil {
				return fmt.Errorf(`failed to upsert series %d squad entry for player %d: %v`, seriesId.Int64, playerId, err)
			}
		}
	}

	return nil
}

func (matchInfo *matchInfo) insertPotmEntries(input *matchInfoInput) error {
	potmEntries := make([]models.PlayerAwardEntry, 0, 1)

	for _, potmName := range input.player_of_match_names {
		cricsheetId, ok := matchInfo.team1Info.players[potmName]
		if !ok {
			if cricsheetId, ok = matchInfo.team2Info.players[potmName]; !ok {
				return fmt.Errorf(`player of the match %s not found in either team info`, potmName)
			}
		}

		playerValue, _ := cachedPlayers.get(playerKey{cricsheet_id: cricsheetId})
		player := playerValue.data

		potmEntry := models.PlayerAwardEntry{
			MatchId:   matchInfo.match.Id,
			PlayerId:  player.Id,
			AwardType: pgtype.Text{String: "player_of_the_match", Valid: true},
		}

		potmEntries = append(potmEntries, potmEntry)
	}

	if err := dbutils.UpsertMatchAwardEntries(context.Background(), DB_POOL, potmEntries); err != nil {
		return fmt.Errorf(`failed to upsert potm entries: %v`, err)
	}

	return nil
}

func (teamInfo *teamInfo) getPlayersId() ([]int64, error) {
	playersId := make([]int64, 0, 11)

	for playerName, cricsheetId := range teamInfo.players {
		playerValue, ok := cachedPlayers.get(playerKey{cricsheet_id: cricsheetId})
		if !ok {
			return nil, fmt.Errorf("player %s not found in cache", playerName)
		}
		playersId = append(playersId, playerValue.data.Id.Int64)
	}

	return playersId, nil
}

func getTourInfo(eventName string) (string, []string) {
	pattern := `^([A-Za-z\s]+) tour of ([A-Za-z\s]+(?: and [A-Za-z\s]+)*)$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(eventName)
	if len(matches) == 3 {
		teamName := matches[1]
		hostNations := matches[2]
		hostNationList := regexp.MustCompile(`\s+and\s+`).Split(hostNations, -1)
		return teamName, hostNationList
	}
	return "", nil
}
