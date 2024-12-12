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
