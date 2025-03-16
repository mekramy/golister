package golister

import (
	"slices"
	"strings"
)

// newConfig creates and returns a new config instance with default values.
func newConfig() *config {
	return &config{
		defaultLimit: 25,
		defaultSort:  "id",
		limits:       []uint{25, 50, 100, 250},
		sorts:        []string{},
		sorter:       postgreSQLSorter,
	}
}

// config holds configuration settings for the application.
type config struct {
	defaultLimit uint             // default limit for queries
	defaultSort  string           // default sort field for queries
	limits       []uint           // allowed limits for queries
	sorts        []string         // allowed sort fields for queries
	sorter       SQLSortGenerator // function to generate SQL sort clauses
}

// validateLimit checks if the provided limit is valid and returns it, otherwise returns the default limit.
func (c *config) validateLimit(v uint) uint {
	if v > 0 && (len(c.limits) == 0 || slices.Contains(c.limits, v)) {
		return v
	}
	return c.defaultLimit
}

// validateSort checks if the provided sort field is valid and returns it, otherwise returns the default sort field.
func (c *config) validateSort(s string) string {
	s = strings.TrimSpace(s)
	if s != "" && (len(c.sorts) == 0 || slices.Contains(c.sorts, s)) {
		return s
	}
	return c.defaultSort
}
