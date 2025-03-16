package golister

import (
	"strconv"
	"strings"
)

// Sort represents a sorting field and order.
type Sort struct {
	Field string `json:"field"`
	Order Order  `json:"order"`
}

// SQLSortGenerator is a function type that generates SQL sort strings.
type SQLSortGenerator func(sorts []Sort, from uint64, limit uint) string

// mySQLSorter generates a MySQL-compatible ORDER BY clause with LIMIT.
func mySQLSorter(sorts []Sort, from uint64, limit uint) string {
	var sb strings.Builder
	sb.WriteString(" ORDER BY ")
	for i, sort := range sorts {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("`" + sort.Field + "` ")
		sb.WriteString(sort.Order.String())
	}
	sb.WriteString(" LIMIT " + strconv.FormatUint(from, 10))
	sb.WriteString(", " + strconv.FormatUint(uint64(limit), 10))
	return sb.String()
}

// postgreSQLSorter generates a PostgreSQL-compatible ORDER BY clause with LIMIT and OFFSET.
func postgreSQLSorter(sorts []Sort, from uint64, limit uint) string {
	var sb strings.Builder
	sb.WriteString(" ORDER BY ")
	for i, sort := range sorts {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(`"` + sort.Field + `" `)
		sb.WriteString(sort.Order.String())
	}
	sb.WriteString(" LIMIT " + strconv.FormatUint(uint64(limit), 10))
	sb.WriteString(" OFFSET " + strconv.FormatUint(from, 10))
	return sb.String()
}
