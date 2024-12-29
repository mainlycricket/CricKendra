package dbutils

import (
	"context"
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
		return response, err
	}

	bowlers, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Bowler_Group, error) {
		var bowler responses.Overall_Bowling_Bowler_Group

		err := rows.Scan(&bowler.BowlerId, &bowler.BowlerName, &bowler.TeamsRepresented, &bowler.MinDate, &bowler.MaxDate, &bowler.MatchesPlayed, &bowler.InningsBowled, &bowler.OversBowled, &bowler.RunsConceded, &bowler.WicketsTaken, &bowler.Average, &bowler.StrikeRate, &bowler.Economy, &bowler.FourWktHauls, &bowler.FiveWktHauls, &bowler.BestInningsRuns, &bowler.BestInningsWkts, &bowler.FoursConceded, &bowler.SixesConceded)

		return bowler, err
	})

	if len(bowlers) > limit {
		response.Stats = bowlers[:limit]
		response.Next = true
	} else {
		response.Stats = bowlers
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_TeamInnings_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_TeamInnings_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_TeamInnings_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_TeamInnings(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	inningsList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_TeamInnings_Group, error) {
		var innings responses.Overall_Bowling_TeamInnings_Group

		err := rows.Scan(&innings.MatchId, &innings.InningsNumber, &innings.BowlingTeamId, &innings.BowlingTeamName, &innings.BattingTeamId, &innings.BattingTeamName, &innings.Season, &innings.CityName, &innings.StartDate, &innings.PlayersCount, &innings.MatchesPlayed, &innings.InningsBowled, &innings.OversBowled, &innings.RunsConceded, &innings.WicketsTaken, &innings.Average, &innings.StrikeRate, &innings.Economy, &innings.FourWktHauls, &innings.FiveWktHauls, &innings.BestInningsRuns, &innings.BestInningsWkts, &innings.FoursConceded, &innings.SixesConceded)

		return innings, err
	})

	if len(inningsList) > limit {
		response.Stats = inningsList[:limit]
		response.Next = true
	} else {
		response.Stats = inningsList
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Matches_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Match_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Match_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Matches(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Match_Group, error) {
		var match responses.Overall_Bowling_Match_Group

		err := rows.Scan(&match.MatchId, &match.Team1Id, &match.Team1Name, &match.Team2Id, &match.Team2Name, &match.Season, &match.CityName, &match.StartDate, &match.PlayersCount, &match.MatchesPlayed, &match.InningsBowled, &match.OversBowled, &match.RunsConceded, &match.WicketsTaken, &match.Average, &match.StrikeRate, &match.Economy, &match.FourWktHauls, &match.FiveWktHauls, &match.BestInningsRuns, &match.BestInningsWkts, &match.FoursConceded, &match.SixesConceded)

		return match, err
	})

	if len(matches) > limit {
		response.Stats = matches[:limit]
		response.Next = true
	} else {
		response.Stats = matches
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Teams_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Team_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Team_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Teams(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	teams, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Team_Group, error) {
		var team responses.Overall_Bowling_Team_Group

		err := rows.Scan(&team.TeamId, &team.TeamName, &team.PlayersCount, &team.MinDate, &team.MaxDate, &team.MatchesPlayed, &team.InningsBowled, &team.OversBowled, &team.RunsConceded, &team.WicketsTaken, &team.Average, &team.StrikeRate, &team.Economy, &team.FourWktHauls, &team.FiveWktHauls, &team.BestInningsRuns, &team.BestInningsWkts, &team.FoursConceded, &team.SixesConceded)

		return team, err
	})

	if len(teams) > limit {
		response.Stats = teams[:limit]
		response.Next = true
	} else {
		response.Stats = teams
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Oppositions_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Opposition_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Opposition_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Oppositions(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	oppositions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Opposition_Group, error) {
		var opposition responses.Overall_Bowling_Opposition_Group

		err := rows.Scan(&opposition.OppositionId, &opposition.OppositionName, &opposition.PlayersCount, &opposition.MinDate, &opposition.MaxDate, &opposition.MatchesPlayed, &opposition.InningsBowled, &opposition.OversBowled, &opposition.RunsConceded, &opposition.WicketsTaken, &opposition.Average, &opposition.StrikeRate, &opposition.Economy, &opposition.FourWktHauls, &opposition.FiveWktHauls, &opposition.BestInningsRuns, &opposition.BestInningsWkts, &opposition.FoursConceded, &opposition.SixesConceded)

		return opposition, err
	})

	if len(oppositions) > limit {
		response.Stats = oppositions[:limit]
		response.Next = true
	} else {
		response.Stats = oppositions
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Ground_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Ground_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	grounds, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Ground_Group, error) {
		var ground responses.Overall_Bowling_Ground_Group

		err := rows.Scan(&ground.GroundId, &ground.GroundName, &ground.PlayersCount, &ground.MinDate, &ground.MaxDate, &ground.MatchesPlayed, &ground.InningsBowled, &ground.OversBowled, &ground.RunsConceded, &ground.WicketsTaken, &ground.Average, &ground.StrikeRate, &ground.Economy, &ground.FourWktHauls, &ground.FiveWktHauls, &ground.BestInningsRuns, &ground.BestInningsWkts, &ground.FoursConceded, &ground.SixesConceded)

		return ground, err
	})

	if len(grounds) > limit {
		response.Stats = grounds[:limit]
		response.Next = true
	} else {
		response.Stats = grounds
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_HostNation_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_HostNation_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	hostNations, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_HostNation_Group, error) {
		var hostNation responses.Overall_Bowling_HostNation_Group

		err := rows.Scan(&hostNation.HostNationId, &hostNation.HostNationName, &hostNation.PlayersCount, &hostNation.MinDate, &hostNation.MaxDate, &hostNation.MatchesPlayed, &hostNation.InningsBowled, &hostNation.OversBowled, &hostNation.RunsConceded, &hostNation.WicketsTaken, &hostNation.Average, &hostNation.StrikeRate, &hostNation.Economy, &hostNation.FourWktHauls, &hostNation.FiveWktHauls, &hostNation.BestInningsRuns, &hostNation.BestInningsWkts, &hostNation.FoursConceded, &hostNation.SixesConceded)

		return hostNation, err
	})

	if len(hostNations) > limit {
		response.Stats = hostNations[:limit]
		response.Next = true
	} else {
		response.Stats = hostNations
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Continents_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Continent_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Continent_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Continents(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	continents, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Continent_Group, error) {
		var continent responses.Overall_Bowling_Continent_Group

		err := rows.Scan(&continent.ContinentId, &continent.ContinentName, &continent.PlayersCount, &continent.MinDate, &continent.MaxDate, &continent.MatchesPlayed, &continent.InningsBowled, &continent.OversBowled, &continent.RunsConceded, &continent.WicketsTaken, &continent.Average, &continent.StrikeRate, &continent.Economy, &continent.FourWktHauls, &continent.FiveWktHauls, &continent.BestInningsRuns, &continent.BestInningsWkts, &continent.FoursConceded, &continent.SixesConceded)

		return continent, err
	})

	if len(continents) > limit {
		response.Stats = continents[:limit]
		response.Next = true
	} else {
		response.Stats = continents
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Year_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Year_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	years, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Year_Group, error) {
		var year responses.Overall_Bowling_Year_Group

		err := rows.Scan(&year.Year, &year.PlayersCount, &year.MatchesPlayed, &year.InningsBowled, &year.OversBowled, &year.RunsConceded, &year.WicketsTaken, &year.Average, &year.StrikeRate, &year.Economy, &year.FourWktHauls, &year.FiveWktHauls, &year.BestInningsRuns, &year.BestInningsWkts, &year.FoursConceded, &year.SixesConceded)

		return year, err
	})

	if len(years) > limit {
		response.Stats = years[:limit]
		response.Next = true
	} else {
		response.Stats = years
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Bowling_Season_Group], error) {
	var response responses.StatsResponse[responses.Overall_Bowling_Season_Group]

	query, args, limit, err := statqueries.Query_Overall_Bowling_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	seasons, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Bowling_Season_Group, error) {
		var season responses.Overall_Bowling_Season_Group

		err := rows.Scan(&season.Season, &season.PlayersCount, &season.MatchesPlayed, &season.InningsBowled, &season.OversBowled, &season.RunsConceded, &season.WicketsTaken, &season.Average, &season.StrikeRate, &season.Economy, &season.FourWktHauls, &season.FiveWktHauls, &season.BestInningsRuns, &season.BestInningsWkts, &season.FoursConceded, &season.SixesConceded)

		return season, err
	})

	if len(seasons) > limit {
		response.Stats = seasons[:limit]
		response.Next = true
	} else {
		response.Stats = seasons
		response.Next = false
	}

	return response, err
}

func Read_Overall_Bowling_Aggregate_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.Overall_Bowling_Aggregate_Group, error) {
	var response responses.Overall_Bowling_Aggregate_Group

	query, args, err := statqueries.Query_Overall_Bowling_Aggregate(&queryMap)
	if err != nil {
		return response, err
	}

	err = db.QueryRow(ctx, query, args...).Scan(&response.PlayersCount, &response.MinDate, &response.MaxDate, &response.MatchesPlayed, &response.InningsBowled, &response.OversBowled, &response.RunsConceded, &response.WicketsTaken, &response.Average, &response.StrikeRate, &response.Economy, &response.FourWktHauls, &response.FiveWktHauls, &response.BestInningsRuns, &response.BestInningsWkts, &response.FoursConceded, &response.SixesConceded)

	return response, err
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
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Innings_Group, error) {
		var record responses.Individual_Bowling_Innings_Group

		err := rows.Scan(&record.MatchId, &record.StartDate, &record.GroundId, &record.CityName, &record.InningsNumber, &record.BowlerId, &record.BowlerName, &record.BowlingTeamId, &record.BowlingTeamName, &record.BattingTeamId, &record.BattingTeamName, &record.OversBowled, &record.RunsConceded, &record.WicketsTaken, &record.Economy, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, err
}

func Read_Individual_Bowling_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Ground_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Ground_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Ground_Group, error) {
		var record responses.Individual_Bowling_Ground_Group

		err := rows.Scan(&record.GroundId, &record.GroundName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.BestInningsRuns, &record.BestInningsWkts, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, err
}

func Read_Individual_Bowling_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_HostNation_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_HostNation_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_HostNation_Group, error) {
		var record responses.Individual_Bowling_HostNation_Group

		err := rows.Scan(&record.HostNationId, &record.HostNationName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.BestInningsRuns, &record.BestInningsWkts, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, err
}

func Read_Individual_Bowling_Oppositions_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Opposition_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Opposition_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Oppositions(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Opposition_Group, error) {
		var record responses.Individual_Bowling_Opposition_Group

		err := rows.Scan(&record.OppositionTeamId, &record.OppositionTeamName, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.BestInningsRuns, &record.BestInningsWkts, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, err
}

func Read_Individual_Bowling_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Year_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Year_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Year_Group, error) {
		var record responses.Individual_Bowling_Year_Group

		err := rows.Scan(&record.Year, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.BestInningsRuns, &record.BestInningsWkts, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, err
}

func Read_Individual_Bowling_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Bowling_Season_Group], error) {
	var response responses.StatsResponse[responses.Individual_Bowling_Season_Group]

	query, args, limit, err := statqueries.Query_Individual_Bowling_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Bowling_Season_Group, error) {
		var record responses.Individual_Bowling_Season_Group

		err := rows.Scan(&record.Season, &record.BowlerId, &record.BowlerName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBowled, &record.OversBowled, &record.RunsConceded, &record.WicketsTaken, &record.Average, &record.StrikeRate, &record.Economy, &record.FourWktHauls, &record.FiveWktHauls, &record.BestInningsRuns, &record.BestInningsWkts, &record.FoursConceded, &record.SixesConceded)

		return record, err
	})

	if len(records) > limit {
		response.Stats = records[:limit]
		response.Next = true
	} else {
		response.Stats = records
		response.Next = false
	}

	return response, err
}
