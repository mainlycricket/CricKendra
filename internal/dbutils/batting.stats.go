package dbutils

import (
	"context"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/mainlycricket/CricKendra/internal/responses"
	statqueries "github.com/mainlycricket/CricKendra/internal/stat_queries"
)

// Function Names are in Read_Overall_Batting_x_Stats format, x represents grouping

func Read_Overall_Batting_Batters_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Batter_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Batter_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Batters(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	batters, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Batter_Group, error) {
		var batter responses.Overall_Batting_Batter_Group

		err := rows.Scan(&batter.BatterId, &batter.BatterName, &batter.TeamsRepresented, &batter.MinDate, &batter.MaxDate, &batter.MatchesPlayed, &batter.InningsBatted, &batter.RunsScored, &batter.BallsFaced, &batter.NotOuts, &batter.Average, &batter.StrikeRate, &batter.HighestScore, &batter.HighestNotOutScore, &batter.Centuries, &batter.HalfCenturies, &batter.FiftyPlusScores, &batter.Ducks, &batter.FoursScored, &batter.SixesScored)

		return batter, err
	})

	if len(batters) > limit {
		response.Stats = batters[:limit]
		response.Next = true
	} else {
		response.Stats = batters
		response.Next = false
	}

	return response, err
}

func Read_Overall_Batting_TeamInnings_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_TeamInnings_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_TeamInnings_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_TeamInnings(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	inningsList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_TeamInnings_Group, error) {
		var innings responses.Overall_Batting_TeamInnings_Group

		err := rows.Scan(&innings.MatchId, &innings.InningsNumber, &innings.BattingTeamId, &innings.BattingTeamName, &innings.BowlingTeamId, &innings.BowlingTeamName, &innings.Season, &innings.CityName, &innings.StartDate, &innings.PlayersCount, &innings.MatchesPlayed, &innings.InningsBatted, &innings.RunsScored, &innings.BallsFaced, &innings.NotOuts, &innings.Average, &innings.StrikeRate, &innings.HighestScore, &innings.HighestNotOutScore, &innings.Centuries, &innings.HalfCenturies, &innings.FiftyPlusScores, &innings.Ducks, &innings.FoursScored, &innings.SixesScored)

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

func Read_Overall_Batting_Matches_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Match_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Match_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Matches(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	matches, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Match_Group, error) {
		var match responses.Overall_Batting_Match_Group

		err := rows.Scan(&match.MatchId, &match.Team1Id, &match.Team1Name, &match.Team2Id, &match.Team2Name, &match.Season, &match.CityName, &match.StartDate, &match.PlayersCount, &match.MatchesPlayed, &match.InningsBatted, &match.RunsScored, &match.BallsFaced, &match.NotOuts, &match.Average, &match.StrikeRate, &match.HighestScore, &match.HighestNotOutScore, &match.Centuries, &match.HalfCenturies, &match.FiftyPlusScores, &match.Ducks, &match.FoursScored, &match.SixesScored)

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

func Read_Overall_Batting_Teams_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Team_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Team_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Teams(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	teams, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Team_Group, error) {
		var team responses.Overall_Batting_Team_Group

		err := rows.Scan(&team.TeamId, &team.TeamName, &team.PlayersCount, &team.MinDate, &team.MaxDate, &team.MatchesPlayed, &team.InningsBatted, &team.RunsScored, &team.BallsFaced, &team.NotOuts, &team.Average, &team.StrikeRate, &team.HighestScore, &team.HighestNotOutScore, &team.Centuries, &team.HalfCenturies, &team.FiftyPlusScores, &team.Ducks, &team.FoursScored, &team.SixesScored)

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

func Read_Overall_Batting_Oppositions_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Opposition_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Opposition_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Oppositions(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	teams, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Opposition_Group, error) {
		var team responses.Overall_Batting_Opposition_Group

		err := rows.Scan(&team.TeamId, &team.TeamName, &team.PlayersCount, &team.MinDate, &team.MaxDate, &team.MatchesPlayed, &team.InningsBatted, &team.RunsScored, &team.BallsFaced, &team.NotOuts, &team.Average, &team.StrikeRate, &team.HighestScore, &team.HighestNotOutScore, &team.Centuries, &team.HalfCenturies, &team.FiftyPlusScores, &team.Ducks, &team.FoursScored, &team.SixesScored)

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

func Read_Overall_Batting_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Ground_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Ground_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	grounds, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Ground_Group, error) {
		var ground responses.Overall_Batting_Ground_Group

		err := rows.Scan(&ground.GroundId, &ground.GroundName, &ground.PlayersCount, &ground.MinDate, &ground.MaxDate, &ground.MatchesPlayed, &ground.InningsBatted, &ground.RunsScored, &ground.BallsFaced, &ground.NotOuts, &ground.Average, &ground.StrikeRate, &ground.HighestScore, &ground.HighestNotOutScore, &ground.Centuries, &ground.HalfCenturies, &ground.FiftyPlusScores, &ground.Ducks, &ground.FoursScored, &ground.SixesScored)

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

func Read_Overall_Batting_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_HostNation_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_HostNation_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	hostNations, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_HostNation_Group, error) {
		var hostNation responses.Overall_Batting_HostNation_Group

		err := rows.Scan(&hostNation.HostNationId, &hostNation.HostNationName, &hostNation.PlayersCount, &hostNation.MinDate, &hostNation.MaxDate, &hostNation.MatchesPlayed, &hostNation.InningsBatted, &hostNation.RunsScored, &hostNation.BallsFaced, &hostNation.NotOuts, &hostNation.Average, &hostNation.StrikeRate, &hostNation.HighestScore, &hostNation.HighestNotOutScore, &hostNation.Centuries, &hostNation.HalfCenturies, &hostNation.FiftyPlusScores, &hostNation.Ducks, &hostNation.FoursScored, &hostNation.SixesScored)

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

func Read_Overall_Batting_Continents_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Continent_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Continent_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Continents(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	continents, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Continent_Group, error) {
		var continent responses.Overall_Batting_Continent_Group

		err := rows.Scan(&continent.ContinentId, &continent.ContinentName, &continent.PlayersCount, &continent.MinDate, &continent.MaxDate, &continent.MatchesPlayed, &continent.InningsBatted, &continent.RunsScored, &continent.BallsFaced, &continent.NotOuts, &continent.Average, &continent.StrikeRate, &continent.HighestScore, &continent.HighestNotOutScore, &continent.Centuries, &continent.HalfCenturies, &continent.FiftyPlusScores, &continent.Ducks, &continent.FoursScored, &continent.SixesScored)

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

func Read_Overall_Batting_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Year_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Year_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	years, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Year_Group, error) {
		var year responses.Overall_Batting_Year_Group

		err := rows.Scan(&year.Year, &year.PlayersCount, &year.MatchesPlayed, &year.InningsBatted, &year.RunsScored, &year.BallsFaced, &year.NotOuts, &year.Average, &year.StrikeRate, &year.HighestScore, &year.HighestNotOutScore, &year.Centuries, &year.HalfCenturies, &year.FiftyPlusScores, &year.Ducks, &year.FoursScored, &year.SixesScored)

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

func Read_Overall_Batting_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Overall_Batting_Season_Group], error) {
	var response responses.StatsResponse[responses.Overall_Batting_Season_Group]

	query, args, limit, err := statqueries.Query_Overall_Batting_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	seasons, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Overall_Batting_Season_Group, error) {
		var season responses.Overall_Batting_Season_Group

		err := rows.Scan(&season.Season, &season.PlayersCount, &season.MatchesPlayed, &season.InningsBatted, &season.RunsScored, &season.BallsFaced, &season.NotOuts, &season.Average, &season.StrikeRate, &season.HighestScore, &season.HighestNotOutScore, &season.Centuries, &season.HalfCenturies, &season.FiftyPlusScores, &season.Ducks, &season.FoursScored, &season.SixesScored)

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

func Read_Overall_Batting_Aggregate_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.Overall_Batting_Aggregate_Group, error) {
	var response responses.Overall_Batting_Aggregate_Group

	query, args, err := statqueries.Query_Overall_Batting_Aggregate(&queryMap)
	if err != nil {
		return response, err
	}

	err = db.QueryRow(ctx, query, args...).Scan(&response.PlayersCount, &response.MinDate, &response.MaxDate, &response.MatchesPlayed, &response.InningsBatted, &response.RunsScored, &response.BallsFaced, &response.NotOuts, &response.Average, &response.StrikeRate, &response.HighestScore, &response.HighestNotOutScore, &response.Centuries, &response.HalfCenturies, &response.FiftyPlusScores, &response.Ducks, &response.FoursScored, &response.SixesScored)

	return response, err
}

// Function Names are in Read_Individual_Batting_x_Stats format, x represents grouping

func Read_Individual_Batting_Innings_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Batting_Innings_Group], error) {
	var response responses.StatsResponse[responses.Individual_Batting_Innings_Group]

	query, args, limit, err := statqueries.Query_Individual_Batting_Innings(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Batting_Innings_Group, error) {
		var record responses.Individual_Batting_Innings_Group

		err := rows.Scan(&record.MatchId, &record.StartDate, &record.GroundId, &record.CityName, &record.InningsNumber, &record.BatterId, &record.BatterName, &record.BattingTeamId, &record.BattingTeamName, &record.BowlingTeamId, &record.BowlingTeamName, &record.RunsScored, &record.BallsFaced, &record.IsNotOut, &record.StrikeRate, &record.FoursScored, &record.SixesScored)

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

func Read_Individual_Batting_Grounds_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Batting_Ground_Group], error) {
	var response responses.StatsResponse[responses.Individual_Batting_Ground_Group]

	query, args, limit, err := statqueries.Query_Individual_Batting_Grounds(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Batting_Ground_Group, error) {
		var record responses.Individual_Batting_Ground_Group

		err := rows.Scan(&record.GroundId, &record.GroundName, &record.BatterId, &record.BatterName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBatted, &record.RunsScored, &record.BallsFaced, &record.NotOuts, &record.Average, &record.StrikeRate, &record.HighestScore, &record.HighestNotOutScore, &record.Centuries, &record.HalfCenturies, &record.FiftyPlusScores, &record.Ducks, &record.FoursScored, &record.SixesScored)

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

func Read_Individual_Batting_HostNations_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Batting_HostNation_Group], error) {
	var response responses.StatsResponse[responses.Individual_Batting_HostNation_Group]

	query, args, limit, err := statqueries.Query_Individual_Batting_HostNations(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Batting_HostNation_Group, error) {
		var record responses.Individual_Batting_HostNation_Group

		err := rows.Scan(&record.HostNationId, &record.HostNationName, &record.BatterId, &record.BatterName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBatted, &record.RunsScored, &record.BallsFaced, &record.NotOuts, &record.Average, &record.StrikeRate, &record.HighestScore, &record.HighestNotOutScore, &record.Centuries, &record.HalfCenturies, &record.FiftyPlusScores, &record.Ducks, &record.FoursScored, &record.SixesScored)

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

func Read_Individual_Batting_Oppositions_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Batting_Opposition_Group], error) {
	var response responses.StatsResponse[responses.Individual_Batting_Opposition_Group]

	query, args, limit, err := statqueries.Query_Individual_Batting_Oppositions(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Batting_Opposition_Group, error) {
		var record responses.Individual_Batting_Opposition_Group

		err := rows.Scan(&record.OppositionTeamId, &record.OppositionTeamName, &record.BatterId, &record.BatterName, &record.TeamsRepresented, &record.MinDate, &record.MaxDate, &record.MatchesPlayed, &record.InningsBatted, &record.RunsScored, &record.BallsFaced, &record.NotOuts, &record.Average, &record.StrikeRate, &record.HighestScore, &record.HighestNotOutScore, &record.Centuries, &record.HalfCenturies, &record.FiftyPlusScores, &record.Ducks, &record.FoursScored, &record.SixesScored)

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

func Read_Individual_Batting_Years_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Batting_Year_Group], error) {
	var response responses.StatsResponse[responses.Individual_Batting_Year_Group]

	query, args, limit, err := statqueries.Query_Individual_Batting_Years(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Batting_Year_Group, error) {
		var record responses.Individual_Batting_Year_Group

		err := rows.Scan(&record.Year, &record.BatterId, &record.BatterName, &record.TeamsRepresented, &record.MatchesPlayed, &record.InningsBatted, &record.RunsScored, &record.BallsFaced, &record.NotOuts, &record.Average, &record.StrikeRate, &record.HighestScore, &record.HighestNotOutScore, &record.Centuries, &record.HalfCenturies, &record.FiftyPlusScores, &record.Ducks, &record.FoursScored, &record.SixesScored)

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

func Read_Individual_Batting_Seasons_Stats(ctx context.Context, db DB_Exec, queryMap url.Values) (responses.StatsResponse[responses.Individual_Batting_Season_Group], error) {
	var response responses.StatsResponse[responses.Individual_Batting_Season_Group]

	query, args, limit, err := statqueries.Query_Individual_Batting_Seasons(&queryMap)
	if err != nil {
		return response, err
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return response, err
	}

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (responses.Individual_Batting_Season_Group, error) {
		var record responses.Individual_Batting_Season_Group

		err := rows.Scan(&record.Season, &record.BatterId, &record.BatterName, &record.TeamsRepresented, &record.MatchesPlayed, &record.InningsBatted, &record.RunsScored, &record.BallsFaced, &record.NotOuts, &record.Average, &record.StrikeRate, &record.HighestScore, &record.HighestNotOutScore, &record.Centuries, &record.HalfCenturies, &record.FiftyPlusScores, &record.Ducks, &record.FoursScored, &record.SixesScored)

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