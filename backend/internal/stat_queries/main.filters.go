package statqueries

import (
	"fmt"
	"net/url"
	"strings"
)

type sqlWhere struct {
	stats_type     int
	group_type     int
	args           []any
	matchQuery     *matchQuery
	inningsFilters *inningsFilters
	qualifications *qualificationFilters
	sortingKeys    []sortingKey
	sortingClause  string
}

func newSqlWhere(stats_type, group_type int) *sqlWhere {
	var sqlWhere = &sqlWhere{stats_type: stats_type, group_type: group_type}
	sqlWhere.matchQuery = newMatchQuery()
	sqlWhere.inningsFilters = newInningsFilters()
	sqlWhere.qualifications = newQualificationFilters(stats_type, group_type)
	sqlWhere.sortingKeys = getSortingKeys(stats_type, group_type)
	return sqlWhere
}

func (sqlWhere *sqlWhere) applyFilters(params *url.Values) {
	if params == nil || len(*params) == 0 {
		sqlWhere.sortingClause = getSortingClause("", "", sqlWhere.sortingKeys)
		return
	}

	sqlWhere.matchQuery.placeholdersCount = len(sqlWhere.args)
	sqlWhere.matchQuery.applyMatchFilters(params)
	sqlWhere.args = append(sqlWhere.args, sqlWhere.matchQuery.args...)

	sqlWhere.inningsFilters.placeholdersCount = len(sqlWhere.args)
	sqlWhere.inningsFilters.applyInningsFilters(params, sqlWhere.stats_type)
	sqlWhere.args = append(sqlWhere.args, sqlWhere.inningsFilters.args...)

	sqlWhere.qualifications.placeholdersCount = len(sqlWhere.args)
	sqlWhere.qualifications.applyQualifications(params)
	sqlWhere.args = append(sqlWhere.args, sqlWhere.qualifications.args...)

	sort_by, sort_order := params.Get("sort_by"), params.Get("sort_order")
	sqlWhere.sortingClause = getSortingClause(sort_by, sort_order, sqlWhere.sortingKeys)
}

type tableQuery struct {
	tableName         string
	fields            []string
	joins             map[string]int
	conditions        []string
	args              []any
	placeholdersCount int
}

func (tq *tableQuery) prepareQuery() string {
	selectFields := prefixJoin(tq.fields, "SELECT", ", ")
	joins := stringifyOrderedMap(tq.joins)
	conditions := prefixJoin(tq.conditions, "WHERE", " AND ")

	query := fmt.Sprintf(`%s FROM %s %s %s`, selectFields, tq.tableName, joins, conditions)

	return query
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
