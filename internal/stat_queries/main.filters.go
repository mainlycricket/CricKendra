package statqueries

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	batting_stats = iota
	bowling_stats
	team_stats
)

type tableQuery struct {
	tableName  string
	fields     []string
	joins      map[string]int
	conditions []string
	args       []any
}

func (tq *tableQuery) prepareQuery() string {
	selectFields := prefixJoin(tq.fields, "SELECT", ", ")
	joins := stringifyOrderedMap(tq.joins)
	conditions := prefixJoin(tq.conditions, "WHERE", " AND ")

	query := fmt.Sprintf(`%s FROM %s %s %s`, selectFields, tq.tableName, joins, conditions)

	return query
}

type sqlWhere struct {
	args           []any
	matchQuery     *matchQuery
	inningsFilters *inningsFilters
}

func newSqlWhere() *sqlWhere {
	var sqlWhere = &sqlWhere{}
	sqlWhere.matchQuery = newMatchQuery()
	sqlWhere.inningsFilters = newInningsFilters()
	return sqlWhere
}

func (sqlWhere *sqlWhere) applyMatchFilters(params *url.Values, stats_type int) {
	sqlWhere.matchQuery.playingFormat(params)
	sqlWhere.matchQuery.isMale(params)

	sqlWhere.matchQuery.minStartDate(params)
	sqlWhere.matchQuery.maxStartDate(params)
	sqlWhere.matchQuery.season(params)

	sqlWhere.matchQuery.primaryTeam(params, stats_type, sqlWhere.inningsFilters)
	sqlWhere.matchQuery.oppositionTeam(params, stats_type, sqlWhere.inningsFilters)
	sqlWhere.matchQuery.homeAwayTeam(params, stats_type, sqlWhere.inningsFilters)
	sqlWhere.matchQuery.tossResult(params, stats_type, sqlWhere.inningsFilters)
	sqlWhere.matchQuery.batFieldFirst(params, stats_type, sqlWhere.inningsFilters)

	sqlWhere.matchQuery.continent(params)
	sqlWhere.matchQuery.hostNation(params)
	sqlWhere.matchQuery.ground(params)

	sqlWhere.matchQuery.series(params)
	sqlWhere.matchQuery.tournament(params)

	sqlWhere.args = append(sqlWhere.args, sqlWhere.matchQuery.args...)
}

func (sqlWhere *sqlWhere) applyFilters(params *url.Values, stats_type int) {
	if params == nil || len(*params) == 0 {
		return
	}

	sqlWhere.applyMatchFilters(params, stats_type)
}

func prefixJoin(words []string, prefix, sep string) string {
	if len(words) > 0 {
		return prefix + " " + strings.Join(words, sep)
	}

	return ""
}

func stringifyOrderedMap(m map[string]int) string {
	words := make([]string, len(m))

	for word, idx := range m {
		words[idx] = word
	}

	return strings.Join(words, " ")
}
