package cmp

import "fmt"

// Order represents the result of a comparison between two values.
type Order int

const (
	// OrderLess indicates that the first value is less than the second value.
	// It is the same value (-1) as returned by [cmp.Compare].
	OrderLess = Order(-1)

	// OrderEqual indicates that the first value is equal to the second value.
	// It is the same value (0) as returned by [cmp.Compare].
	OrderEqual = Order(0)

	// OrderGreater indicates that the first value is greater than the second value.
	// It is the same value (+1) as returned by [cmp.Compare].
	OrderGreater = Order(+1)
)

// String implements [fmt.Stringer].
func (o Order) String() string {
	switch o {
	case OrderLess:
		return "less"
	case OrderEqual:
		return "equal"
	case OrderGreater:
		return "greater"
	default:
		return fmt.Sprintf("Order(%d)", o)
	}
}

// check interfaces
var (
	_ fmt.Stringer = Order(0)
)
