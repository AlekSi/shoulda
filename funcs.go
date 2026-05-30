package shoulda

import (
	"github.com/AlekSi/shoulda/cmp"
)

// Satisfy checks that predicate returns true for actual.
func Satisfy[A any](tb TB, actual A, predicate func(_ A) bool) bool {
	tb.Helper()

	m := msg("predicate is not satisfied for\nactual:   %v", actual)

	return assert(tb, predicate(actual), m)
}

// SatisfyWith checks that predicate returns true for actual and expected.
func SatisfyWith[A, E any](tb TB, actual A, expected E, predicate func(_ A, _ E) bool) bool {
	tb.Helper()

	m := msg("predicate is not satisfied with\nactual:   %v\nexpected: %v", actual, expected)

	return assert(tb, predicate(actual, expected), m)
}

// CompareWith checks that compare(actual, expected) returns order.
func CompareWith[A, E any](tb TB, actual A, expected E, order cmp.Order, compare func(_ A, _ E) int) bool {
	tb.Helper()

	m := msg("comparison result is not %d for\nactual:   %v\nexpected: %v", order, actual, expected)

	return assert(tb, compare(actual, expected) == int(order), m)
}

// CompareEqual checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqual[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msg(
		"comparison result is %s, not equal for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)

	return assert(tb, res == 0, m)
}

// CompareLess checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLess[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msg(
		"comparison result is %s, not less for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)

	return assert(tb, res == -1, m)
}

// CompareGreater checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreater[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msg(
		"comparison result is %s, not greater for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)

	return assert(tb, res == +1, m)
}
