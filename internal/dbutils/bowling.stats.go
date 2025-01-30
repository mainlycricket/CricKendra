package dbutils

import (
	"context"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/responses"
	statqueries "github.com/mainlycricket/CricKendra/internal/stat_queries"
)

// Function Names are in Read_Overall_Bowling_x_Stats format, x represents grouping

func Read_Overall_Bowling_Bowlers_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Bowler_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Bowler_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Bowlers(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	bowlers, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Bowler_Group, error) {
		var bowler responses.Overall_Bowling_Bowler_Group

		err := rows.Scan(&bowler.BowlerId, &bowler.BowlerName, &bowler.TeamsRepresented, &bowler.MinDate, &bowler.MaxDate, &bowler.MatchesPlayed, &bowler.InningsBowled, &bowler.OversBowled, &bowler.MaidenOvers, &bowler.RunsConceded, &bowler.WicketsTaken, &bowler.Average, &bowler.StrikeRate, &bowler.Economy, &bowler.FourWktHauls, &bowler.FiveWktHauls, &bowler.TenWktHauls, &bowler.BestMatchWkts, &bowler.BestMatchRuns, &bowler.BestInningsWkts, &bowler.BestInningsRuns, &bowler.FoursConceded, &bowler.SixesConceded)

		return bowler, err
	})

	if err != nil {
		return response, err
	}

	if len(bowlers) > limit {
		response.Stats = bowlers[:limit]
		response.Next = true
	} else {
		response.Stats = bowlers
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_TeamInnings_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_TeamInnings_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_TeamInnings_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_TeamInnings(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	inningsList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_TeamInnings_Group, error) {
		var innings responses.Overall_Bowling_TeamInnings_Group

		err := rows.Scan(&innings.MatchId, &innings.InningsNumber, &innings.BowlingTeamId, &innings.BowlingTeamName, &innings.BattingTeamId, &innings.BattingTeamName, &innings.Season, &innings.CityName, &innings.StartDate, &innings.PlayersCount, &innings.MatchesPlayed, &innings.InningsBowled, &innings.OversBowled, &innings.MaidenOvers, &innings.RunsConceded, &innings.WicketsTaken, &innings.Average, &innings.StrikeRate, &innings.Economy, &innings.FourWktHauls, &innings.FiveWktHauls, &innings.TenWktHauls, &innings.BestMatchWkts, &innings.BestMatchRuns, &innings.BestInningsWkts, &innings.BestInningsRuns, &innings.FoursConceded, &innings.SixesConceded)

		return innings, err
	})

	if err != nil {
		return response, err
	}

	if len(inningsList) > limit {
		response.Stats = inningsList[:limit]
		response.Next = true
	} else {
		response.Stats = inningsList
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Matches_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Match_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Match_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Matches(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Match_Group, error) {
		var match responses.Overall_Bowling_Match_Group

		err := rows.Scan(&match.MatchId, &match.Team1Id, &match.Team1Name, &match.Team2Id, &match.Team2Name, &match.Season, &match.CityName, &match.StartDate, &match.PlayersCount, &match.MatchesPlayed, &match.InningsBowled, &match.OversBowled, &match.MaidenOvers, &match.RunsConceded, &match.WicketsTaken, &match.Average, &match.StrikeRate, &match.Economy, &match.FourWktHauls, &match.FiveWktHauls, &match.TenWktHauls, &match.BestMatchWkts, &match.BestMatchRuns, &match.BestInningsWkts, &match.BestInningsRuns, &match.FoursConceded, &match.SixesConceded)

		return match, err
	})

	if err != nil {
		return response, err
	}

	if len(matches) > limit {
		response.Stats = matches[:limit]
		response.Next = true
	} else {
		response.Stats = matches
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Teams_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Team_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Team_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Teams(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	teams, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Team_Group, error) {
		var team responses.Overall_Bowling_Team_Group

		err := rows.Scan(&team.TeamId, &team.TeamName, &team.PlayersCount, &team.MinDate, &team.MaxDate, &team.MatchesPlayed, &team.InningsBowled, &team.OversBowled, &team.MaidenOvers, &team.RunsConceded, &team.WicketsTaken, &team.Average, &team.StrikeRate, &team.Economy, &team.FourWktHauls, &team.FiveWktHauls, &team.TenWktHauls, &team.BestMatchWkts, &team.BestMatchRuns, &team.BestInningsWkts, &team.BestInningsRuns, &team.FoursConceded, &team.SixesConceded)

		return team, err
	})

	if err != nil {
		return response, err
	}

	if len(teams) > limit {
		response.Stats = teams[:limit]
		response.Next = true
	} else {
		response.Stats = teams
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Oppositions_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Opposition_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Opposition_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Oppositions(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	oppositions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Opposition_Group, error) {
		var opposition responses.Overall_Bowling_Opposition_Group

		err := rows.Scan(&opposition.OppositionId, &opposition.OppositionName, &opposition.PlayersCount, &opposition.MinDate, &opposition.MaxDate, &opposition.MatchesPlayed, &opposition.InningsBowled, &opposition.OversBowled, &opposition.MaidenOvers, &opposition.RunsConceded, &opposition.WicketsTaken, &opposition.Average, &opposition.StrikeRate, &opposition.Economy, &opposition.FourWktHauls, &opposition.FiveWktHauls, &opposition.TenWktHauls, &opposition.BestMatchWkts, &opposition.BestMatchRuns, &opposition.BestInningsWkts, &opposition.BestInningsRuns, &opposition.FoursConceded, &opposition.SixesConceded)

		return opposition, err
	})

	if err != nil {
		return response, err
	}

	if len(oppositions) > limit {
		response.Stats = oppositions[:limit]
		response.Next = true
	} else {
		response.Stats = oppositions
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Ground_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Ground_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	grounds, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Ground_Group, error) {
		var ground responses.Overall_Bowling_Ground_Group

		err := rows.Scan(&ground.GroundId, &ground.GroundName, &ground.PlayersCount, &ground.MinDate, &ground.MaxDate, &ground.MatchesPlayed, &ground.InningsBowled, &ground.OversBowled, &ground.MaidenOvers, &ground.RunsConceded, &ground.WicketsTaken, &ground.Average, &ground.StrikeRate, &ground.Economy, &ground.FourWktHauls, &ground.FiveWktHauls, &ground.TenWktHauls, &ground.BestMatchWkts, &ground.BestMatchRuns, &ground.BestInningsWkts, &ground.BestInningsRuns, &ground.FoursConceded, &ground.SixesConceded)

		return ground, err
	})

	if err != nil {
		return response, err
	}

	if len(grounds) > limit {
		response.Stats = grounds[:limit]
		response.Next = true
	} else {
		response.Stats = grounds
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_HostNation_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_HostNation_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	hostNations, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_HostNation_Group, error) {
		var hostNation responses.Overall_Bowling_HostNation_Group

		err := rows.Scan(&hostNation.HostNationId, &hostNation.HostNationName, &hostNation.PlayersCount, &hostNation.MinDate, &hostNation.MaxDate, &hostNation.MatchesPlayed, &hostNation.InningsBowled, &hostNation.OversBowled, &hostNation.MaidenOvers, &hostNation.RunsConceded, &hostNation.WicketsTaken, &hostNation.Average, &hostNation.StrikeRate, &hostNation.Economy, &hostNation.FourWktHauls, &hostNation.FiveWktHauls, &hostNation.TenWktHauls, &hostNation.BestMatchWkts, &hostNation.BestMatchRuns, &hostNation.BestInningsWkts, &hostNation.BestInningsRuns, &hostNation.FoursConceded, &hostNation.SixesConceded)

		return hostNation, err
	})

	if err != nil {
		return response, err
	}

	if len(hostNations) > limit {
		response.Stats = hostNations[:limit]
		response.Next = true
	} else {
		response.Stats = hostNations
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Continents_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Continent_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Continent_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Continents(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	continents, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Continent_Group, error) {
		var continent responses.Overall_Bowling_Continent_Group

		err := rows.Scan(&continent.ContinentId, &continent.ContinentName, &continent.PlayersCount, &continent.MinDate, &continent.MaxDate, &continent.MatchesPlayed, &continent.InningsBowled, &continent.OversBowled, &continent.MaidenOvers, &continent.RunsConceded, &continent.WicketsTaken, &continent.Average, &continent.StrikeRate, &continent.Economy, &continent.FourWktHauls, &continent.FiveWktHauls, &continent.TenWktHauls, &continent.BestMatchWkts, &continent.BestMatchRuns, &continent.BestInningsWkts, &continent.BestInningsRuns, &continent.FoursConceded, &continent.SixesConceded)

		return continent, err
	})

	if err != nil {
		return response, err
	}

	if len(continents) > limit {
		response.Stats = continents[:limit]
		response.Next = true
	} else {
		response.Stats = continents
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Series_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Series_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Series_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Series(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	seriesList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Series_Group, error) {
		var series responses.Overall_Bowling_Series_Group

		err := rows.Scan(&series.SeriesId, &series.SeriesName, &series.SeriesSeason, &series.PlayersCount, &series.MinDate, &series.MaxDate, &series.MatchesPlayed, &series.InningsBowled, &series.OversBowled, &series.MaidenOvers, &series.RunsConceded, &series.WicketsTaken, &series.Average, &series.StrikeRate, &series.Economy, &series.FourWktHauls, &series.FiveWktHauls, &series.TenWktHauls, &series.BestMatchWkts, &series.BestMatchRuns, &series.BestInningsWkts, &series.BestInningsRuns, &series.FoursConceded, &series.SixesConceded)

		return series, err
	})

	if err != nil {
		return response, err
	}

	if len(seriesList) > limit {
		response.Stats = seriesList[:limit]
		response.Next = true
	} else {
		response.Stats = seriesList
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Tournaments_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Tournament_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Tournament_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Tournaments(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	tournaments, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Tournament_Group, error) {
		var tournament responses.Overall_Bowling_Tournament_Group

		err := rows.Scan(&tournament.TournamentId, &tournament.TournamentName, &tournament.PlayersCount, &tournament.MinDate, &tournament.MaxDate, &tournament.MatchesPlayed, &tournament.InningsBowled, &tournament.OversBowled, &tournament.MaidenOvers, &tournament.RunsConceded, &tournament.WicketsTaken, &tournament.Average, &tournament.StrikeRate, &tournament.Economy, &tournament.FourWktHauls, &tournament.FiveWktHauls, &tournament.TenWktHauls, &tournament.BestMatchWkts, &tournament.BestMatchRuns, &tournament.BestInningsWkts, &tournament.BestInningsRuns, &tournament.FoursConceded, &tournament.SixesConceded)

		return tournament, err
	})

	if err != nil {
		return response, err
	}

	if len(tournaments) > limit {
		response.Stats = tournaments[:limit]
		response.Next = true
	} else {
		response.Stats = tournaments
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Year_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Year_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	years, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Year_Group, error) {
		var year responses.Overall_Bowling_Year_Group

		err := rows.Scan(&year.Year, &year.PlayersCount, &year.MatchesPlayed, &year.InningsBowled, &year.OversBowled, &year.MaidenOvers, &year.RunsConceded, &year.WicketsTaken, &year.Average, &year.StrikeRate, &year.Economy, &year.FourWktHauls, &year.FiveWktHauls, &year.TenWktHauls, &year.BestMatchWkts, &year.BestMatchRuns, &year.BestInningsWkts, &year.BestInningsRuns, &year.FoursConceded, &year.SixesConceded)

		return year, err
	})

	if err != nil {
		return response, err
	}

	if len(years) > limit {
		response.Stats = years[:limit]
		response.Next = true
	} else {
		response.Stats = years
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Season_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Season_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	seasons, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Season_Group, error) {
		var season responses.Overall_Bowling_Season_Group

		err := rows.Scan(&season.Season, &season.PlayersCount, &season.MatchesPlayed, &season.InningsBowled, &season.OversBowled, &season.MaidenOvers, &season.RunsConceded, &season.WicketsTaken, &season.Average, &season.StrikeRate, &season.Economy, &season.FourWktHauls, &season.FiveWktHauls, &season.TenWktHauls, &season.BestMatchWkts, &season.BestMatchRuns, &season.BestInningsWkts, &season.BestInningsRuns, &season.FoursConceded, &season.SixesConceded)

		return season, err
	})

	if err != nil {
		return response, err
	}

	if len(seasons) > limit {
		response.Stats = seasons[:limit]
		response.Next = true
	} else {
		response.Stats = seasons
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Decades_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Decade_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Decade_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Decades(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	decades, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Decade_Group, error) {
		var decade responses.Overall_Bowling_Decade_Group

		err := rows.Scan(&decade.Decade, &decade.PlayersCount, &decade.MatchesPlayed, &decade.InningsBowled, &decade.OversBowled, &decade.MaidenOvers, &decade.RunsConceded, &decade.WicketsTaken, &decade.Average, &decade.StrikeRate, &decade.Economy, &decade.FourWktHauls, &decade.FiveWktHauls, &decade.TenWktHauls, &decade.BestMatchWkts, &decade.BestMatchRuns, &decade.BestInningsWkts, &decade.BestInningsRuns, &decade.FoursConceded, &decade.SixesConceded)

		return decade, err
	})

	if err != nil {
		return response, err
	}

	if len(decades) > limit {
		response.Stats = decades[:limit]
		response.Next = true
	} else {
		response.Stats = decades
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Bowling_Aggregate_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Aggregate_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Aggregate_Group]

	query, args, err := statqueries.Query_Overall_Bowling_Aggregate(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Aggregate_Group, error) {
		var record responses.Overall_Bowling_Aggregate_Group

		err = rows.Scan(&record.PlayersCount, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	response.Stats = records
	response.Next = false

	return response, nil
}

// Function Names are in Read_Individual_Bowling_x_Stats format, x represents grouping

func Read_Individual_Bowling_Innings_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Innings_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Innings_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Innings(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Innings_Group, error) {
		var record responses.Individual_Bowling_Innings_Group

		err := rows.Scan(&record.MatchId, &record.StartDate, &record.GroundId, &record.CityName, &record.InningsNumber, &record.BowlerId, &record.BowlerName, &record.BowlingTeamId, &record.BowlingTeamName, &record.BattingTeamId, &record.BattingTeamName, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Economy, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_MatchTotals_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_MatchTotals_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_MatchTotals_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_MatchTotals(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_MatchTotals_Group, error) {
		var record responses.Individual_Bowling_MatchTotals_Group

		err := rows.Scan(&record.MatchId, &record.StartDate, &record.GroundId, &record.CityName, &record.BowlerId, &record.BowlerName, &record.BowlingTeamId, &record.BowlingTeamName, &record.BattingTeamId, &record.BattingTeamName, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.Economy, &record.StrikeRate, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_Series_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Series_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Series_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Series(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Series_Group, error) {
		var record responses.Individual_Bowling_Series_Group

		err := rows.Scan(&record.SeriesId, &record.SeriesName, &record.SeriesSeason, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_Tournaments_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Tournament_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Tournament_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Tournaments(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Tournament_Group, error) {
		var record responses.Individual_Bowling_Tournament_Group

		err := rows.Scan(&record.TournamentId, &record.TournamentName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Ground_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Ground_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Ground_Group, error) {
		var record responses.Individual_Bowling_Ground_Group

		err := rows.Scan(&record.GroundId, &record.GroundName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_HostNation_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_HostNation_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_HostNation_Group, error) {
		var record responses.Individual_Bowling_HostNation_Group

		err := rows.Scan(&record.HostNationId, &record.HostNationName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_Oppositions_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Opposition_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Opposition_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Oppositions(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Opposition_Group, error) {
		var record responses.Individual_Bowling_Opposition_Group

		err := rows.Scan(&record.OppositionTeamId, &record.OppositionTeamName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Year_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Year_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Year_Group, error) {
		var record responses.Individual_Bowling_Year_Group

		err := rows.Scan(&record.Year, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)
		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}

func Read_Individual_Bowling_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Season_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Season_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Season_Group, error) {
		var record responses.Individual_Bowling_Season_Group

		err := rows.Scan(&record.Season, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.MaidenOvers, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.TenWktHauls, &record.BestMatchWkts, &record.BestMatchRuns, &record.BestInningsWkts, &record.BestInningsRuns, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if err != nil {
		return response, err
	}

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, nil
}
