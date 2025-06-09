package dbutils

import (
	"context"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	statqueries "github.com/mainlycricket/CricKendra/backend/internal/stat_queries"
)

// Function Names are in Read_Overall_Team_x_Stats format, x represents grouping

func Read_Overall_Team_Teams_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Teams_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Teams_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Teams(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	teams, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Teams_Group, error) {
		var team responses.Overall_Team_Teams_Group

		err := rows.Scan(&team.TeamId, &team.TeamName, &team.MinStartDate, &team.MaxStartDate, &team.MatchesPlayed, &team.MatchesWon, &team.MatchesLost, &team.WinLossRatio, &team.MatchesDrawn, &team.MatchesTied, &team.MatchesNoResult, &team.InningsCount, &team.TotalRuns, &team.TotalBalls, &team.TotalWickets, &team.Average, &team.ScoringRate, &team.HighestScore, &team.LowestScore)

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

func Read_Overall_Team_Players_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Players_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Players_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Players(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	players, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Players_Group, error) {
		var player responses.Overall_Team_Players_Group

		err := rows.Scan(&player.PlayerId, &player.PlayerName, &player.MinStartDate, &player.MaxStartDate, &player.TeamsCount, &player.MatchesPlayed, &player.MatchesWon, &player.MatchesLost, &player.WinLossRatio, &player.MatchesDrawn, &player.MatchesTied, &player.MatchesNoResult, &player.InningsCount, &player.TotalRuns, &player.TotalBalls, &player.TotalWickets, &player.Average, &player.ScoringRate, &player.HighestScore, &player.LowestScore)

		return player, err
	})

	if err != nil {
		return response, err
	}

	if len(players) > limit {
		response.Stats = players[:limit]
		response.Next = true
	} else {
		response.Stats = players
		response.Next = false
	}

	return response, nil
}

func Read_Overall_Team_Matches_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Matches_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Matches_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Matches(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Matches_Group, error) {
		var match responses.Overall_Team_Matches_Group

		err := rows.Scan(&match.MatchId, &match.Team1Id, &match.Team1Name, &match.Team2Id, &match.Team2Name, &match.GroundId, &match.CityName, &match.Season, &match.StartDate, &match.TeamsCount, &match.MatchesPlayed, &match.MatchesWon, &match.MatchesLost, &match.WinLossRatio, &match.MatchesDrawn, &match.MatchesTied, &match.MatchesNoResult, &match.InningsCount, &match.TotalRuns, &match.TotalBalls, &match.TotalWickets, &match.Average, &match.ScoringRate, &match.HighestScore, &match.LowestScore)

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

func Read_Overall_Team_Series_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Series_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Series_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Series(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	seriesList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Series_Group, error) {
		var series responses.Overall_Team_Series_Group

		err := rows.Scan(&series.SeriesId, &series.SeriesName, &series.SeriesSeason, &series.TeamsCount, &series.MatchesPlayed, &series.MatchesWon, &series.MatchesLost, &series.WinLossRatio, &series.MatchesDrawn, &series.MatchesTied, &series.MatchesNoResult, &series.InningsCount, &series.TotalRuns, &series.TotalBalls, &series.TotalWickets, &series.Average, &series.ScoringRate, &series.HighestScore, &series.LowestScore)

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

func Read_Overall_Team_Tournaments_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Tournament_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Tournament_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Tournaments(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	tournaments, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Tournament_Group, error) {
		var tournament responses.Overall_Team_Tournament_Group

		err := rows.Scan(&tournament.TournamentId, &tournament.TournamentName, &tournament.TeamsCount, &tournament.MatchesPlayed, &tournament.MatchesWon, &tournament.MatchesLost, &tournament.WinLossRatio, &tournament.MatchesDrawn, &tournament.MatchesTied, &tournament.MatchesNoResult, &tournament.InningsCount, &tournament.TotalRuns, &tournament.TotalBalls, &tournament.TotalWickets, &tournament.Average, &tournament.ScoringRate, &tournament.HighestScore, &tournament.LowestScore)

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

func Read_Overall_Team_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Grounds_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Grounds_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	grounds, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Grounds_Group, error) {
		var ground responses.Overall_Team_Grounds_Group

		err := rows.Scan(&ground.GroundId, &ground.GroundName, &ground.MinStartDate, &ground.MaxStartDate, &ground.TeamsCount, &ground.MatchesPlayed, &ground.MatchesWon, &ground.MatchesLost, &ground.WinLossRatio, &ground.MatchesDrawn, &ground.MatchesTied, &ground.MatchesNoResult, &ground.InningsCount, &ground.TotalRuns, &ground.TotalBalls, &ground.TotalWickets, &ground.Average, &ground.ScoringRate, &ground.HighestScore, &ground.LowestScore)

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

func Read_Overall_Team_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_HostNations_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_HostNations_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	hostNations, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_HostNations_Group, error) {
		var hostNation responses.Overall_Team_HostNations_Group

		err := rows.Scan(&hostNation.HostNationId, &hostNation.HostNationName, &hostNation.MinStartDate, &hostNation.MaxStartDate, &hostNation.TeamsCount, &hostNation.MatchesPlayed, &hostNation.MatchesWon, &hostNation.MatchesLost, &hostNation.WinLossRatio, &hostNation.MatchesDrawn, &hostNation.MatchesTied, &hostNation.MatchesNoResult, &hostNation.InningsCount, &hostNation.TotalRuns, &hostNation.TotalBalls, &hostNation.TotalWickets, &hostNation.Average, &hostNation.ScoringRate, &hostNation.HighestScore, &hostNation.LowestScore)

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

func Read_Overall_Team_Continents_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Continents_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Continents_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Continents(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	continents, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Continents_Group, error) {
		var continent responses.Overall_Team_Continents_Group

		err := rows.Scan(&continent.ContinentId, &continent.ContinentName, &continent.MinStartDate, &continent.MaxStartDate, &continent.TeamsCount, &continent.MatchesPlayed, &continent.MatchesWon, &continent.MatchesLost, &continent.WinLossRatio, &continent.MatchesDrawn, &continent.MatchesTied, &continent.MatchesNoResult, &continent.InningsCount, &continent.TotalRuns, &continent.TotalBalls, &continent.TotalWickets, &continent.Average, &continent.ScoringRate, &continent.HighestScore, &continent.LowestScore)

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

	return response, nil
}

func Read_Overall_Team_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Years_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Years_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	years, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Years_Group, error) {
		var year responses.Overall_Team_Years_Group

		err := rows.Scan(&year.Year, &year.TeamsCount, &year.MatchesPlayed, &year.MatchesWon, &year.MatchesLost, &year.WinLossRatio, &year.MatchesDrawn, &year.MatchesTied, &year.MatchesNoResult, &year.InningsCount, &year.TotalRuns, &year.TotalBalls, &year.TotalWickets, &year.Average, &year.ScoringRate, &year.HighestScore, &year.LowestScore)

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

func Read_Overall_Team_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Seasons_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Seasons_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	seasons, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Seasons_Group, error) {
		var season responses.Overall_Team_Seasons_Group

		err := rows.Scan(&season.Season, &season.TeamsCount, &season.MatchesPlayed, &season.MatchesWon, &season.MatchesLost, &season.WinLossRatio, &season.MatchesDrawn, &season.MatchesTied, &season.MatchesNoResult, &season.InningsCount, &season.TotalRuns, &season.TotalBalls, &season.TotalWickets, &season.Average, &season.ScoringRate, &season.HighestScore, &season.LowestScore)

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

func Read_Overall_Team_Decades_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Decades_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Decades_Group]

	query, args, limit, err := statqueries.Query_Overall_Team_Decades(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	decades, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Decades_Group, error) {
		var decade responses.Overall_Team_Decades_Group

		err := rows.Scan(&decade.Decade, &decade.TeamsCount, &decade.MatchesPlayed, &decade.MatchesWon, &decade.MatchesLost, &decade.WinLossRatio, &decade.MatchesDrawn, &decade.MatchesTied, &decade.MatchesNoResult, &decade.InningsCount, &decade.TotalRuns, &decade.TotalBalls, &decade.TotalWickets, &decade.Average, &decade.ScoringRate, &decade.HighestScore, &decade.LowestScore)

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

func Read_Overall_Team_Aggregate_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Team_Aggregate_Group], error) {
	var response responses.StatsResponse[responses.Overall_Team_Aggregate_Group]

	query, args, err := statqueries.Query_Overall_Team_Aggregate(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Team_Aggregate_Group, error) {
		var record responses.Overall_Team_Aggregate_Group

		err := rows.Scan(&record.TeamsCount, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

		return record, err
	})

	if err != nil {
		return response, err
	}

	response.Stats = records
	response.Next = false

	return response, nil
}

// Function Names are in Read_Individual_Team_x_Stats format, x represents grouping

func Read_Individual_Team_Innings_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_Innings_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_Innings_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_Innings(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_Innings_Group, error) {
		var record responses.Individual_Team_Innings_Group

		err := rows.Scan(&record.MatchId, &record.TeamId, &record.TeamName, &record.OppositionId, &record.OppositionName, &record.GroundId, &record.CityName, &record.StartDate, &record.FinalResult, &record.MatchWinnerId, &record.InningsId, &record.InningsNumber, &record.InningsEnd, &record.TotalRuns, &record.TotalWickets, &record.TotalOvers, &record.ScoringRate)

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

func Read_Individual_Team_MatchTotals_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_MatchTotals_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_MatchTotals_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_MatchTotals(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_MatchTotals_Group, error) {
		var record responses.Individual_Team_MatchTotals_Group

		err := rows.Scan(&record.MatchId, &record.TeamId, &record.TeamName, &record.OppositionId, &record.OppositionName, &record.GroundId, &record.CityName, &record.StartDate, &record.FinalResult, &record.MatchWinnerId, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate)

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

func Read_Individual_Team_MatchResults_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_MatchResults_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_MatchResults_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_MatchResults(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_MatchResults_Group, error) {
		var record responses.Individual_Team_MatchResults_Group

		err := rows.Scan(&record.MatchId, &record.TeamId, &record.TeamName, &record.OppositionId, &record.OppositionName, &record.GroundId, &record.CityName, &record.StartDate, &record.FinalResult, &record.MatchWinnerId, &record.TossWinnerId, &record.InningsNumber, &record.WinMargin, &record.BallsMargin, &record.IsWonByRuns, &record.IsWonByInnings)

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

func Read_Individual_Team_Series_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_Series_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_Series_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_Series(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_Series_Group, error) {
		var record responses.Individual_Team_Series_Group

		err := rows.Scan(&record.TeamId, &record.TeamName, &record.SeriesId, &record.SeriesName, &record.SeriesSeason, &record.MinStartDate, &record.MaxStartDate, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

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

func Read_Individual_Team_Tournaments_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_Tournaments_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_Tournaments_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_Tournaments(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_Tournaments_Group, error) {
		var record responses.Individual_Team_Tournaments_Group

		err := rows.Scan(&record.TeamId, &record.TeamName, &record.TournamentId, &record.TournamentName, &record.MinStartDate, &record.MaxStartDate, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

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

func Read_Individual_Team_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_Grounds_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_Grounds_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_Grounds_Group, error) {
		var record responses.Individual_Team_Grounds_Group

		err := rows.Scan(&record.TeamId, &record.TeamName, &record.GroundId, &record.GroundName, &record.MinStartDate, &record.MaxStartDate, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

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

func Read_Individual_Team_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_HostNations_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_HostNations_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_HostNations_Group, error) {
		var record responses.Individual_Team_HostNations_Group

		err := rows.Scan(&record.TeamId, &record.TeamName, &record.HostNationId, &record.HostNationName, &record.MinStartDate, &record.MaxStartDate, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

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

func Read_Individual_Team_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_Years_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_Years_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_Years_Group, error) {
		var record responses.Individual_Team_Years_Group

		err := rows.Scan(&record.TeamId, &record.TeamName, &record.Year, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

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

func Read_Individual_Team_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Team_Seasons_Group], error) {
	var response responses.StatsResponse[responses.Individual_Team_Seasons_Group]

	query, args, limit, err := statqueries.Query_Individual_Team_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		log.Println(query)
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Team_Seasons_Group, error) {
		var record responses.Individual_Team_Seasons_Group

		err := rows.Scan(&record.TeamId, &record.TeamName, &record.Season, &record.MatchesPlayed, &record.MatchesWon, &record.MatchesLost, &record.WinLossRatio, &record.MatchesDrawn, &record.MatchesTied, &record.MatchesNoResult, &record.InningsCount, &record.TotalRuns, &record.TotalBalls, &record.TotalWickets, &record.Average, &record.ScoringRate, &record.HighestScore, &record.LowestScore)

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
