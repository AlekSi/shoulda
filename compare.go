package shoulda

import (
	"reflect"

	"github.com/AlekSi/shoulda/cmp"
	"github.com/AlekSi/shoulda/internal"
)

// BeNil checks that actual is (untyped) nil.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	msg := internal.MsgFmt("actual: %[1]v (%[1]T)\nis not nil", actual)
	return assert(tb, actual == nil, msg)
}

// NotBeNil checks that actual is not (untyped) nil.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	return assert(tb, actual != nil, internal.MsgString("is nil"))
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T cmp.Ordered](tb TB, actual T) bool {
	tb.Helper()

	msg := internal.MsgFmt("actual: %v\nis not zero", actual)
	var zero T
	return assert(tb, actual == zero, msg)
}

// BeFalse checks that actual is false.
func BeFalse(tb TB, actual bool) bool {
	tb.Helper()

	return assert(tb, !actual, internal.MsgString("is not false"))
}

// BeTrue checks that actual is true.
func BeTrue(tb TB, actual bool) bool {
	tb.Helper()

	return assert(tb, actual, internal.MsgString("is not true"))
}

// BeDeepEqual checks that actual and expected are deeply equal.
func BeDeepEqual(tb TB, actual, expected any) bool {
	tb.Helper()

	msg := internal.MsgFmt("Values are not deep equal:\nactual:   %#v\nexpected: %#v", actual, expected)
	return assert(tb, reflect.DeepEqual(actual, expected), msg)
}

// BeEqual checks that actual and expected are equal.
func BeEqual[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	msg := internal.MsgFmt("Values are not equal:\nactual:   %v\nexpected: %v", actual, expected)
	return assert(tb, cmp.Equal(actual, expected), msg)
}

// BeLess checks that actual is less than expected.
func BeLess[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	msg := internal.MsgFmt("actual:   %v\nis not less than\nexpected: %v", actual, expected)
	return assert(tb, cmp.Less(actual, expected), msg)
}

// BeGreater checks that actual is greater than expected.
func BeGreater[T cmp.Ordered](tb TB, actual, expected T) bool {
	tb.Helper()

	msg := internal.MsgFmt("actual:   %v\nis not greater than\nexpected: %v", actual, expected)
	return assert(tb, cmp.Greater(actual, expected), msg)
}

// Satisfy checks that predicate returns true for actual.
func Satisfy[T any](tb TB, actual T, predicate func(_ T) bool) bool {
	tb.Helper()

	msg := internal.MsgFmt("predicate is not satisfied for\nactual:   %v", actual)
	return assert(tb, predicate(actual), msg)
}

// SatisfyWith checks that predicate returns true for actual and expected.
func SatisfyWith[T any](tb TB, actual, expected T, predicate func(_, _ T) bool) bool {
	tb.Helper()

	msg := internal.MsgFmt("predicate is not satisfied with\nactual:   %v\nexpected: %v", actual, expected)
	return assert(tb, predicate(actual, expected), msg)
}

// CompareWith checks that compare(actual, expected) returns order.
func CompareWith[T any](tb TB, actual, expected T, order cmp.Order, compare func(_, _ T) int) bool {
	tb.Helper()

	msg := internal.MsgFmt("comparison result is not %d for\nactual:   %v\nexpected: %v", order, actual, expected)
	return assert(tb, compare(actual, expected) == int(order), msg)
}

// CompareEqual checks that compare(actual, expected) returns 0 ([cmp.OrderEqual]).
func CompareEqual[T any](tb TB, actual, expected T, compare func(_, _ T) int) bool {
	tb.Helper()

	res := compare(actual, expected)
	msg := internal.MsgFmt(
		"comparison result is %s, not equal for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)
	return assert(tb, res == 0, msg)
}

// CompareLess checks that compare(actual, expected) returns -1 ([cmp.OrderLess]).
func CompareLess[T any](tb TB, actual, expected T, compare func(_, _ T) int) bool {
	tb.Helper()

	res := compare(actual, expected)
	msg := internal.MsgFmt(
		"comparison result is %s, not less for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)
	return assert(tb, res == -1, msg)
}

// CompareGreater checks that compare(actual, expected) returns 1 ([cmp.OrderGreater]).
func CompareGreater[T any](tb TB, actual, expected T, compare func(_, _ T) int) bool {
	tb.Helper()

	res := compare(actual, expected)
	msg := internal.MsgFmt(
		"comparison result is %s, not greater for\nactual:   %v\nexpected: %v",
		cmp.Order(res), actual, expected,
	)
	return assert(tb, res == +1, msg)
}
