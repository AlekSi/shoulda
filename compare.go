package shoulda

import (
	"reflect"

	"github.com/AlekSi/shoulda/cmp"
)

// BeFalse checks that actual is false.
func BeFalse(tb TB, actual bool) bool {
	tb.Helper()

	m := messagef("is not false")

	return assert(tb, !actual, m)
}

// BeTrue checks that actual is true.
func BeTrue(tb TB, actual bool) bool {
	tb.Helper()

	m := messagef("is not true")

	return assert(tb, actual, m)
}

// BeDeepEqual checks that actual and expected are equal according to [reflect.DeepEqual].
func BeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	m := messagef("Values are not deep equal:\nactual:   %#v\nexpected: %#v", actual, expected)

	return assert(tb, reflect.DeepEqual(actual, expected), m)
}

// NotBeDeepEqual checks that actual and expected are not equal according to [reflect.DeepEqual].
func NotBeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	m := messagef("Values are deep equal:\nactual:   %#v\nexpected: %#v", actual, expected)

	return assert(tb, !reflect.DeepEqual(actual, expected), m)
}

// BeEqual checks that actual and expected are equal according to [cmp.Equal].
func BeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := messagef("Values are not equal:\nactual:   %v\nexpected: %v", actual, expected)

	return assert(tb, cmp.Equal(actual, expected), m)
}

// BeLess checks that actual is less than expected according to [cmp.Less].
func BeLess[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := messagef("actual:   %v\nis not less than\nexpected: %v", actual, expected)

	return assert(tb, cmp.Less(actual, expected), m)
}

// BeGreater checks that actual is greater than expected according to [cmp.Greater].
func BeGreater[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	m := messagef("actual:   %v\nis not greater than\nexpected: %v", actual, expected)

	return assert(tb, cmp.Greater(actual, expected), m)
}

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
