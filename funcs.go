package shoulda

import (
	"github.com/AlekSi/shoulda/cmp"
)

// Satisfy checks that predicate returns true for actual.
func Satisfy[T any](tb TB, actual T, predicate func(_ T) bool) bool {
	tb.Helper()

	m := messagef("predicate is not satisfied for\nactual:   %v", actual)

	return assert(tb, predicate(actual), m)
}

// SatisfyWith checks that predicate returns true for actual and expected.
func SatisfyWith[T any](tb TB, actual, expected T, predicate func(_, _ T) bool) bool {
	tb.Helper()

	m := messagef("predicate is not satisfied with\nactual:   %v\nexpected: %v", actual, expected)

	return assert(tb, predicate(actual, expected), m)
}

// CompareWith checks that compare(actual, expected) returns order.
func CompareWith[T any](tb TB, actual, expected T, order cmp.Order, compare func(_, _ T) int) bool {
	tb.Helper()

	m := messagef("comparison result is not %d for\nactual:   %v\nexpected: %v", order, actual, expected)

	return assert(tb, compare(actual, expected) == int(order), m)
}

// CompareEqual checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqual[T any](tb TB, actual, expected T, compare func(_, _ T) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := messagef(
		"comparison result is %s, not equal for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)

	return assert(tb, res == 0, m)
}

// CompareLess checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLess[T any](tb TB, actual, expected T, compare func(_, _ T) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := messagef(
		"comparison result is %s, not less for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)

	return assert(tb, res == -1, m)
}

// CompareGreater checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreater[T any](tb TB, actual, expected T, compare func(_, _ T) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := messagef(
		"comparison result is %s, not greater for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)

	return assert(tb, res == +1, m)
}
