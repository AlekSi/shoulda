package shoulda

import (
	"reflect"

	"github.com/AlekSi/shoulda/cmp"
)

// BeNilf checks that actual is untyped nil.
func BeNilf(tb TB, actual any, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	return assert(tb, actual == nil, m)
}

// BeNil checks that actual is untyped nil.
func BeNil(tb TB, actual any) bool {
	tb.Helper()

	m := dumpf(tb, "actual is not untyped nil, but %T:\n%s", actual)

	return assert(tb, actual == nil, m)
}

// NotBeNilf checks that actual is not (untyped) nil.
func NotBeNilf(tb TB, actual any, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	return assert(tb, actual != nil, m)
}

// NotBeNil checks that actual is not (untyped) nil.
func NotBeNil(tb TB, actual any) bool {
	tb.Helper()

	m := messagef("actual is untyped nil")

	return assert(tb, actual != nil, m)
}

// BeZerof checks that actual is the zero value of its type.
func BeZerof[T comparable](tb TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	var zero T
	return assert(tb, actual == zero, m)
}

// BeZero checks that actual is the zero value of its type.
func BeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := dumpf(tb, "actual is not zero, but %T:\n%s", actual)

	var zero T
	return assert(tb, actual == zero, m)
}

// NotBeZerof checks that actual is not the zero value of its type.
func NotBeZerof[T comparable](tb TB, actual T, msg string, args ...any) bool {
	tb.Helper()

	m := messagef(msg, args...)

	var zero T
	return assert(tb, actual != zero, m)
}

// NotBeZero checks that actual is not the zero value of its type.
func NotBeZero[T comparable](tb TB, actual T) bool {
	tb.Helper()

	m := messagef("actual is zero")

	var zero T
	return assert(tb, actual != zero, m)
}

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
