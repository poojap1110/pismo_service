package helper

import (
	"fmt"
	"strconv"
	"strings"
)

// ApplyQueryFilters ...
func ApplyQueryFilters(filters map[string]interface{}) string {
	var (
		condition []string
		sql       string
	)

	if len(filters) > 0 {
		for key, value := range filters {
			condition = append(condition, key+"='"+value.(string)+"'")
		}
	}
	if len(condition) > 0 {
		sql = sql + strings.Join(condition, " AND ")
	}
	return sql
}

// WithoutTrashed ...
func WithoutTrashed(query string) string {
	query = AppendWhereClause(query, "deleted_at is NULL")

	return query
}

// AppendWhereClause ...
func AppendWhereClause(query string, where string) string {
	if strings.Contains(query, "WHERE") {
		query = strings.Replace(query, "WHERE", fmt.Sprintf("WHERE %s AND", where), 1)
	} else if strings.Contains(query, "LIMIT") {
		query = strings.Replace(query, "LIMIT", fmt.Sprintf("WHERE %s LIMIT", where), 1)
	}

	return query
}

// ImplodeIntItems ...
func ImplodeIntItems(values []int, glue string) (rawString string) {
	valuesText := []string{}

	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	// Join our string slice.
	rawString = strings.Join(valuesText, glue)

	return
}
