package shoulda

import (
	"github.com/AlekSi/shoulda/cmp"
)

// Satisfy checks that predicate returns true for actual.
func Satisfy[A any](tb TB, actual A, predicate func(_ A) bool) bool {
	tb.Helper()

	args := []any{Dump(tb, actual)}

	m := msgf("actual is not satisfied by predicate:\nactual: %s", args...)

	return assert(tb, predicate(actual), m)
}

// Satisfyf checks that predicate returns true for actual.
func Satisfyf[A any](tb TB, actual A, predicate func(_ A) bool, format string, args ...any) bool {
	tb.Helper()

	args = append([]any{Dump(tb, actual)}, args...)

	m := msgf("actual is not satisfied by predicate:\nactual: %s\n"+format, args...)

	return assert(tb, predicate(actual), m)
}

// SatisfyWith checks that predicate returns true for actual and expected.
func SatisfyWith[A, E any](tb TB, actual A, expected E, predicate func(_ A, _ E) bool) bool {
	tb.Helper()

	args := []any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual and expected are not satisfied by predicate:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, predicate(actual, expected), m)
}

// SatisfyWithf checks that predicate returns true for actual and expected.
func SatisfyWithf[A, E any](tb TB, actual A, expected E, predicate func(_ A, _ E) bool, format string, args ...any) bool {
	tb.Helper()

	args = append([]any{
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}, args...)

	m := msgf("actual and expected are not satisfied by predicate:\nactual: %s\nexpected: %s\n%s"+format, args...)

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
		return assert(tb, false, msgf("invalid cmp.%s", order))
	}
}

// CompareWithf checks that compare(actual, expected) returns order.
func CompareWithf[A, E any](tb TB, actual A, expected E, order cmp.Order, compare func(_ A, _ E) int, format string, args ...any) bool {
	tb.Helper()

	switch order {
	case cmp.OrderEqual:
		return CompareEqualf(tb, actual, expected, compare, format, args...)
	case cmp.OrderLess:
		return CompareLessf(tb, actual, expected, compare, format, args...)
	case cmp.OrderGreater:
		return CompareGreaterf(tb, actual, expected, compare, format, args...)
	default:
		m := msgf("invalid cmp.%s\n"+format, append([]any{order}, args...)...)
		return assert(tb, false, m)
	}
}

// CompareEqual checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqual[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	args := []any{
		cmp.Order(res),
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not equal to expected, but %s:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, res == 0, m)
}

// CompareEqualf checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqualf[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int, format string, args ...any) bool {
	tb.Helper()

	res := compare(actual, expected)

	args = append([]any{
		cmp.Order(res),
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}, args...)

	m := msgf("actual is not equal to expected, but %s:\nactual: %s\nexpected: %s\n%s"+format, args...)

	return assert(tb, res == 0, m)
}

// CompareLess checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLess[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	args := []any{
		cmp.Order(res),
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not less than expected, but %s:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, res == -1, m)
}

// CompareLessf checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLessf[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int, format string, args ...any) bool {
	tb.Helper()

	res := compare(actual, expected)

	args = append([]any{
		cmp.Order(res),
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}, args...)

	m := msgf("actual is not less than expected, but %s:\nactual: %s\nexpected: %s\n%s"+format, args...)

	return assert(tb, res == -1, m)
}

// CompareGreater checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreater[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int) bool {
	tb.Helper()

	res := compare(actual, expected)

	args := []any{
		cmp.Order(res),
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}

	m := msgf("actual is not greater than expected, but %s:\nactual: %s\nexpected: %s\n%s", args...)

	return assert(tb, res == +1, m)
}

// CompareGreaterf checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreaterf[A, E any](tb TB, actual A, expected E, compare func(_ A, _ E) int, format string, args ...any) bool {
	tb.Helper()

	res := compare(actual, expected)

	args = append([]any{
		cmp.Order(res),
		Dump(tb, actual),
		Dump(tb, expected),
		Diff(tb, "actual", actual, "expected", expected),
	}, args...)

	m := msgf("actual is not greater than expected, but %s:\nactual: %s\nexpected: %s\n%s"+format, args...)

	return assert(tb, res == +1, m)
}
