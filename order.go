package golister

import (
	"fmt"
	"strings"
)

// Order represents the sorting order as a string.
type Order string

const (
	// Ascending represents ascending order.
	Ascending Order = "asc"

	// Descending represents descending order.
	Descending Order = "desc"
)

// String returns the string representation of the Order.
func (order Order) String() string {
	return strings.ToUpper(string(order))
}

// Numeric returns 1 for Ascending order and -1 for Descending order.
func (order Order) Numeric() int {
	if order == Ascending {
		return 1
	}
	return -1
}

// ParseOrder parses a value into an Order. It returns Descending for "-1" or "desc", and Ascending for any other value.
func ParseOrder(v any) Order {
	switch fmt.Sprint(v) {
	case "-1", "desc":
		return Descending
	default:
		return Ascending
	}
}
