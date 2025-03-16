package golister

import (
	"slices"
	"strings"
)

// Options is a function type that modifies a lister instance.
type Options func(*config)

// WithLimits sets the default limit and valid limits for the lister.
func WithLimits(def uint, valids ...uint) Options {
	valids = slices.DeleteFunc(valids, func(v uint) bool {
		return v == 0
	})

	return func(c *config) {
		if def > 0 {
			c.defaultLimit = def
		}

		if len(valids) > 0 {
			c.limits = append([]uint{}, valids...)
		}
	}
}

// WithSorts sets the default sort and valid sorts for the lister.
func WithSorts(def string, valids ...string) Options {
	def = strings.TrimSpace(def)
	valids = slices.DeleteFunc(valids, func(v string) bool {
		return strings.TrimSpace(v) == ""
	})

	return func(c *config) {
		if def != "" {
			c.defaultSort = def
		}

		c.sorts = append([]string{}, valids...)
	}
}

// WithDefaultLimit sets the default limit for the lister.
func WithDefaultLimit(def uint) Options {
	return func(c *config) {
		if def > 0 {
			c.defaultLimit = def
		}
	}
}

// WithDefaultSort sets the default sort for the lister.
func WithDefaultSort(def string) Options {
	def = strings.TrimSpace(def)
	return func(c *config) {
		if def != "" {
			c.defaultSort = def
		}
	}
}

// WithMySQLSorter sets the SQL sorter to MySQL.
func WithMySQLSorter() Options {
	return func(c *config) {
		c.sorter = mySQLSorter
	}
}

// WithPostgreSQLSorter sets the SQL sorter to PostgreSQL.
func WithPostgreSQLSorter() Options {
	return func(c *config) {
		c.sorter = postgreSQLSorter
	}
}
