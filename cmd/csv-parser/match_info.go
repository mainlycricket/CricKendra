package main

import (
	"context"
	"encoding/csv"
	"io"
	"net/url"
	"os"
	"path/filepath"
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
	var potmName, eventName string

	fp, err := os.Open(filePath)
	if err != nil {
		return match, err
	}

	matchId := strings.TrimSuffix(filepath.Base(filePath), "_info.csv")
	match.CricsheetId = pgtype.Text{String: matchId, Valid: true}

	match.PlayingFormat = pgtype.Text{String: playingFormat, Valid: true}
	match.PlayingLevel = pgtype.Text{String: playingLevel, Valid: true}
	match.FinalResult = pgtype.Text{String: "winner_decided", Valid: true}

	reader := csv.NewReader(fp)

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

		if row[1] == "event" {
			eventName = row[2]
		}

		if row[1] == "outcome" {
			match.FinalResult = handleMatchOutcome(row[2])
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

		if row[1] == "venue" {
			match.GroundId, match.HostNationId, err = handleVenueAndHostNation(row[2])
			if err != nil {
				return match, err
			}
			continue
		}

		if row[1] == "team" {
			teamName := row[2]

			teamField, err := handleTeam(teamName, match.IsMale.Bool)
			if err != nil {
				return match, nil
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
			teamField, err := handleTeam(row[2], match.IsMale.Bool)
			if err != nil {
				return match, err
			}

			match.BowlOutWinnerId = teamField
			continue
		}

		if row[1] == "eliminator" {
			teamField, err := handleTeam(row[2], match.IsMale.Bool)
			if err != nil {
				return match, err
			}

			match.SuperOverWinnerId = teamField
			continue
		}

		if row[1] == "winner" {
			teamField, err := handleTeam(row[2], match.IsMale.Bool)
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
			teamField, err := handleTeam(row[2], match.IsMale.Bool)
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

		if row[1] == "player_of_match" {
			potmName = row[3]
		}
	}

	if potmId, ok := team1.Players[potmName]; ok {
		match.PoTMsId = append(match.PoTMsId, pgtype.Int8{Int64: potmId, Valid: true})
	} else if potmId, ok := team2.Players[potmName]; ok {
		match.PoTMsId = append(match.PoTMsId, pgtype.Int8{Int64: potmId, Valid: true})
	}

	handleEvent(eventName, match.Season.String, match.IsMale.Bool, &match)

	return match, nil
}

func handlePlayer(cricsheetId string, isMale bool, teamId int64) (int64, error) {
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

	return dbResponse.Players[0].Id.Int64, nil
}

func handleEvent(eventName, season string, isMale bool, match *models.Match) error {
	filters := url.Values{
		"name":    []string{eventName},
		"season":  []string{season},
		"is_male": []string{"true"},
		"__limit": []string{"1"},
	}

	if !isMale {
		filters["is_male"] = []string{"false"}
	}

	dbResponse, err := dbutils.ReadSeries(context.Background(), DB_POOL, filters)
	if err != nil {
		return err
	}

	if len(dbResponse.Series) == 0 {
		newSeries := models.Series{
			Name:          pgtype.Text{String: eventName, Valid: true},
			Season:        pgtype.Text{String: season, Valid: true},
			IsMale:        pgtype.Bool{Bool: isMale, Valid: true},
			PlayingLevel:  match.PlayingLevel,
			PlayingFormat: match.PlayingFormat,
			TeamsId:       []pgtype.Int8{match.Team1Id, match.Team2Id},
			HostNationsId: []pgtype.Int8{match.HostNationId},
		}

		err = dbutils.InsertSeries(context.Background(), DB_POOL, &newSeries)
		if err != nil {
			return err
		}

		dbResponse, err = dbutils.ReadSeries(context.Background(), DB_POOL, filters)
		if err != nil {
			return err
		}
	}

	match.SeriesId = dbResponse.Series[0].Id

	return nil
}

func handleTeam(teamName string, isMale bool) (pgtype.Int8, error) {
	filters := url.Values{
		"name":    []string{teamName},
		"is_male": []string{"true"},
		"__limit": []string{"1"},
	}

	if !isMale {
		filters["male"][0] = "false"
	}

	dbResponse, err := dbutils.ReadTeams(context.Background(), DB_POOL, filters)
	if err != nil {
		return pgtype.Int8{}, err
	}

	if len(dbResponse.Teams) == 0 {
		newTeam := models.Team{
			Name:   pgtype.Text{String: teamName, Valid: true},
			IsMale: pgtype.Bool{Bool: isMale, Valid: true},
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

	team := dbResponse.Teams[0]
	teamField := pgtype.Int8{Int64: team.Id.Int64, Valid: true}

	return teamField, nil
}

func handleSeason(season string) (pgtype.Text, error) {
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

	return pgtype.Text{String: season, Valid: true}, nil
}

func handleVenueAndHostNation(venue string) (groundId, hostNationId pgtype.Int8, err error) {
	filters := url.Values{"name": []string{venue}, "__limit": []string{"1"}}

	dbResponse, err := dbutils.ReadGrounds(context.Background(), DB_POOL, filters)
	if err != nil {
		return groundId, hostNationId, err
	}

	if len(dbResponse.Grounds) == 0 {
		newGround := models.Ground{Name: pgtype.Text{String: venue, Valid: true}}

		err := dbutils.InsertGround(context.Background(), DB_POOL, &newGround)
		if err != nil {
			return groundId, hostNationId, err
		}

		dbResponse, err = dbutils.ReadGrounds(context.Background(), DB_POOL, filters)
		if err != nil {
			return groundId, hostNationId, err
		}
	}

	groundId = dbResponse.Grounds[0].Id
	hostNationId = dbResponse.Grounds[0].HostNationId

	return groundId, hostNationId, nil
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
