package shoulda

import (
	"github.com/AlekSi/shoulda/cmp"
)

// Satisfy checks that predicate returns true for actual.
func Satisfy[A any](tb TB, actual A, predicate func(_ A) bool) bool {
	tb.Helper()

	s := dumpf(tb, "actual is not satisfied by predicate:\nactual: %[2]s", actual, "")

	return assert(tb, predicate(actual), s)
}

// SatisfyWith checks that predicate returns true for actual and expected.
func SatisfyWith[A, E any](tb TB, actual A, expected E, predicate func(_ A, _ E) bool) bool {
	tb.Helper()

	m := msgDiff(
		tb,
		"actual and expected are not satisfied by predicate:\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, predicate(actual, expected), m)
}

// CompareWith checks that compare(actual, expected) returns order.
func CompareWith[A, E any](tb TB, actual A, expected E, order cmp.Order, compare func(_ A, _ E) int) bool {
	tb.Helper()

	switch order {
	case cmp.OrderEqual:
		return CompareEqual(tb, actual, expected, compare)
	case cmp.OrderLess:
		return CompareLess(tb, actual, expected, compare)
	case cmp.OrderGreater:
		return CompareGreater(tb, actual, expected, compare)
	default:
		return assert(tb, false, sprintf("invalid cmp.%s", order))
	}
}

// CompareEqual checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqual[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msgDiff(
		tb,
		"actual is not equal to expected, but "+cmp.Order(res).String()+":\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, res == 0, m)
}

// CompareLess checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLess[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msgDiff(
		tb,
		"actual is not less than expected, but "+cmp.Order(res).String()+":\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, res == -1, m)
}

// CompareGreater checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreater[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	m := msgDiff(
		tb,
		"actual is not greater than expected, but "+cmp.Order(res).String()+":\nactual: %[2]s\nexpected: %[4]s\n%[5]s",
		actual, expected,
	)

	return assert(tb, res == +1, m)
}
